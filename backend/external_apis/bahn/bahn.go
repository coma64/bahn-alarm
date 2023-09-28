package bahn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/metrics"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	BadResponse = fmt.Errorf("bad bahn api response")
)

var randomUuid = uuid.NewString()

func createBahnRequest(ctx context.Context, path string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, config.Conf.Bahn.BaseUrl+path, nil)
	if err != nil {
		return nil, err
	}

	request.WithContext(ctx)
	request.Header.Set("Time-Zone", "UTC")
	request.Header.Set("X-Api-Key", config.Conf.Bahn.ApiKey)
	// Simply using some uuid here somehow causes the stations endpoint to return more relevant results
	request.Header.Set("X-Installation-Id", randomUuid)

	return request, err
}

func doRequest[T interface{}](ctx context.Context, path, query string, responseBody *T) error {
	request, err := createBahnRequest(ctx, path+query)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	start := time.Now()
	var response *http.Response
	if response, err = http.DefaultClient.Do(request); err != nil {
		metrics.BahnApiRequestDuration.WithLabelValues("-1", path).Observe(0)
		return fmt.Errorf("error sending request: %w", err)
	}

	elapsed := metrics.AccurateSecondsSince(start)

	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Err(err).Msg("failed to close response body")
		}
	}()

	metrics.BahnApiRequestDuration.
		WithLabelValues(strconv.Itoa(response.StatusCode), path).
		Observe(elapsed)

	if response.StatusCode != 200 {
		if config.Conf.Debug {
			logResponse(response)
		}

		return BadResponse
	}

	buf := strings.Builder{}
	if _, err = io.Copy(&buf, response.Body); err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	if err = json.Unmarshal([]byte(buf.String()), responseBody); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if _, err = db.Db.ExecContext(ctx, "insert into bahnapisearchresponses (data) values ($1)", buf.String()); err != nil {
		return fmt.Errorf("error writing bahn api response to db: %w", err)
	}

	log.Debug().
		Str("path", path).
		Str("query", query).
		Int("statusCode", response.StatusCode).
		Float64("durationSeconds", elapsed).
		Msg("Requested bahn API")

	return nil
}

func logResponse(response *http.Response) {
	if dump, err := httputil.DumpResponse(response, true); err != nil {
		log.Err(err).Msg("Failed to log response")
	} else {
		log.Printf("%q", dump)
	}
}

func FetchPlaces(ctx context.Context, query string) (*PlacesResponse, error) {
	log.Debug().Str("package", "bahn").Str("query", query).Msg("Fetching places")

	stations := &PlacesResponse{}
	return stations, doRequest(ctx, "/spam/v1/places", "?query="+query, stations)
}

func FetchConnections(
	ctx context.Context,
	departure time.Time,
	originStationId string,
	destinationStationId string,
) (*ConnectionsResponse, error) {
	log.Debug().
		Str("package", "bahn").
		Str("originStationId", originStationId).
		Str("destinationStationid", destinationStationId).
		Time("departure", departure).
		Msg("Fetching connections")

	query := url.Values{}

	query.Set("departureTime", departure.Format(time.RFC3339))
	query.Set("publicTransportModes", "regional_rail,urban_rail,metro,tram,bus")

	// These coordinates are required and must be different but have no effect on the results
	query.Set("origin", "0,0")
	query.Set("destination", "0,1")

	query.Set("originStationId", originStationId)
	query.Set("destinationStationId", destinationStationId)

	connections := &ConnectionsResponse{}
	return connections, doRequest(ctx, "/navigation/v2/search/trips", "?"+query.Encode(), connections)
}
