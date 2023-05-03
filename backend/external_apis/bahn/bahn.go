package bahn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

func doRequest[T interface{}](ctx context.Context, query string, responseBody *T) error {
	request, err := createBahnRequest(ctx, query)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	log.Debug().Str("url", request.URL.String()).Msg("requesting bahn api")

	var response *http.Response
	if response, err = http.DefaultClient.Do(request); err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Err(err).Msg("failed to close response body")
		}
	}()

	if response.StatusCode != 200 {
		if config.Conf.Debug {
			logResponse(response)
		}

		return BadResponse
	}

	if err = json.NewDecoder(response.Body).Decode(responseBody); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

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
	stations := &PlacesResponse{}
	return stations, doRequest(ctx, "/spam/v1/places?query="+query, stations)
}

func FetchConnections(
	ctx context.Context,
	departure time.Time,
	originStationId string,
	destinationStationId string,
) (*ConnectionsResponse, error) {
	query := url.Values{}

	query.Set("departureTime", departure.Format(time.RFC3339))
	query.Set("publicTransportModes", "regional_rail,urban_rail,metro,tram,bus")

	// These coordinates are required and must be different but have no effect on the results
	query.Set("origin", "0,0")
	query.Set("destination", "0,1")

	query.Set("originStationId", originStationId)
	query.Set("destinationStationId", destinationStationId)

	connections := &ConnectionsResponse{}
	return connections, doRequest(ctx, "/navigation/v2/search/trips?"+query.Encode(), connections)
}

func FetchConnectionsByPlace(ctx context.Context, departure time.Time, origin, destination *Place) (*ConnectionsResponse, error) {
	return FetchConnections(ctx, departure, origin.StationID, destination.StationID)
}
