// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Defines values for TrackedDepartureStatus.
const (
	Canceled   TrackedDepartureStatus = "canceled"
	Delayed    TrackedDepartureStatus = "delayed"
	NotChecked TrackedDepartureStatus = "not-checked"
	OnTime     TrackedDepartureStatus = "on-time"
)

// Defines values for Urgency.
const (
	Error Urgency = "error"
	Info  Urgency = "info"
	Warn  Urgency = "warn"
)

// Alarm defines model for Alarm.
type Alarm struct {
	Content   Alarm_Content `json:"content"`
	CreatedAt time.Time     `json:"createdAt"`
	Id        int           `json:"id"`
	Urgency   Urgency       `json:"urgency"`
}

// Alarm_Content defines model for Alarm.Content.
type Alarm_Content struct {
	union json.RawMessage
}

// AlarmsList defines model for AlarmsList.
type AlarmsList struct {
	Alarms     []Alarm    `json:"alarms"`
	Pagination Pagination `json:"pagination"`
}

// BahnConnection defines model for BahnConnection.
type BahnConnection struct {
	Departure struct {
		ScheduledTime time.Time `json:"scheduledTime"`
	} `json:"departure"`
}

// BahnConnectionsList defines model for BahnConnectionsList.
type BahnConnectionsList struct {
	Connections []BahnConnection `json:"connections"`
}

// BahnPlace defines model for BahnPlace.
type BahnPlace struct {
	Id        string `json:"id"`
	Label     string `json:"label"`
	Name      string `json:"name"`
	StationId string `json:"stationId"`
}

// BahnPlacesList defines model for BahnPlacesList.
type BahnPlacesList struct {
	Places []BahnPlace `json:"places"`
}

// CancelledAlarm defines model for CancelledAlarm.
type CancelledAlarm struct {
	Connection SimpleConnection `json:"connection"`
	IsCanceled bool             `json:"isCanceled"`
}

// DelayChangeAlarm defines model for DelayChangeAlarm.
type DelayChangeAlarm struct {
	Connection      SimpleConnection `json:"connection"`
	NewDelayMinutes int              `json:"newDelayMinutes"`
}

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	// TotalItems The total amount of found items, not just on this page
	TotalItems int `json:"totalItems"`
}

// PushNotificationSubscription defines model for PushNotificationSubscription.
type PushNotificationSubscription struct {
	CreatedAt    *time.Time      `json:"createdAt,omitempty"`
	Id           *int            `json:"id,omitempty"`
	IsEnabled    bool            `json:"isEnabled"`
	Name         string          `json:"name"`
	Subscription RawSubscription `json:"subscription"`
}

// PushNotificationSubscriptionList defines model for PushNotificationSubscriptionList.
type PushNotificationSubscriptionList struct {
	Pagination    Pagination                     `json:"pagination"`
	Subscriptions []PushNotificationSubscription `json:"subscriptions"`
}

// RawSubscription defines model for RawSubscription.
type RawSubscription struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		Auth   string `json:"auth"`
		P256dh string `json:"p256dh"`
	} `json:"keys"`
}

// RegisterRequest defines model for RegisterRequest.
type RegisterRequest struct {
	InviteToken string `json:"inviteToken"`
	Password    string `json:"password"`
	Username    string `json:"username"`
}

// SimpleConnection defines model for SimpleConnection.
type SimpleConnection struct {
	Departure time.Time `json:"departure"`
	FromName  string    `json:"fromName"`
	ToName    string    `json:"toName"`
}

// TrackedConnection defines model for TrackedConnection.
type TrackedConnection struct {
	Departures []TrackedDeparture `json:"departures"`
	FromId     string             `json:"fromId"`
	FromName   string             `json:"fromName"`
	Id         *int               `json:"id,omitempty"`
	ToId       string             `json:"toId"`
	ToName     string             `json:"toName"`
}

// TrackedConnectionList defines model for TrackedConnectionList.
type TrackedConnectionList struct {
	Connections []TrackedConnection `json:"connections"`
	Pagination  Pagination          `json:"pagination"`
}

// TrackedConnectionUpdate defines model for TrackedConnectionUpdate.
type TrackedConnectionUpdate struct {
	Departures []TrackedDeparture `json:"departures"`
}

// TrackedDeparture defines model for TrackedDeparture.
type TrackedDeparture struct {
	Departure time.Time               `json:"departure"`
	Status    *TrackedDepartureStatus `json:"status,omitempty"`
}

// TrackedDepartureStatus defines model for TrackedDeparture.Status.
type TrackedDepartureStatus string

// TrackingStats defines model for TrackingStats.
type TrackingStats struct {
	CanceledConnectionCount int `json:"canceledConnectionCount"`
	DelayedConnectionCount  int `json:"delayedConnectionCount"`
	NextDeparture           struct {
		ConnectionId int       `json:"connectionId"`
		Departure    time.Time `json:"departure"`
	} `json:"nextDeparture"`
	TotalConnectionCount int `json:"totalConnectionCount"`
}

// Urgency defines model for Urgency.
type Urgency string

// User defines model for User.
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        int       `json:"id"`
	IsAdmin   bool      `json:"isAdmin"`
	Name      string    `json:"name"`
}

// ValidationFailed defines model for ValidationFailed.
type ValidationFailed struct {
	Message *string `json:"message,omitempty"`
}

// GetAlarmsParams defines parameters for GetAlarms.
type GetAlarmsParams struct {
	Page    *int     `form:"page,omitempty" json:"page,omitempty"`
	Size    *int     `form:"size,omitempty" json:"size,omitempty"`
	Urgency *Urgency `form:"urgency,omitempty" json:"urgency,omitempty"`
}

// GetBahnConnectionsParams defines parameters for GetBahnConnections.
type GetBahnConnectionsParams struct {
	Departure time.Time `form:"departure" json:"departure"`
	FromId    string    `form:"fromId" json:"fromId"`
	ToId      string    `form:"toId" json:"toId"`
}

// GetBahnPlacesParams defines parameters for GetBahnPlaces.
type GetBahnPlacesParams struct {
	Name string `form:"name" json:"name"`
}

// GetNotificationsPushSubscriptionsParams defines parameters for GetNotificationsPushSubscriptions.
type GetNotificationsPushSubscriptionsParams struct {
	Page *int `form:"page,omitempty" json:"page,omitempty"`
	Size *int `form:"size,omitempty" json:"size,omitempty"`
}

// GetTrackingConnectionsParams defines parameters for GetTrackingConnections.
type GetTrackingConnectionsParams struct {
	Page *int `form:"page,omitempty" json:"page,omitempty"`
	Size *int `form:"size,omitempty" json:"size,omitempty"`
}

// PostAuthLoginJSONRequestBody defines body for PostAuthLogin for application/json ContentType.
type PostAuthLoginJSONRequestBody = LoginRequest

// PostAuthRegisterJSONRequestBody defines body for PostAuthRegister for application/json ContentType.
type PostAuthRegisterJSONRequestBody = RegisterRequest

// PostNotificationsPushSubscriptionsJSONRequestBody defines body for PostNotificationsPushSubscriptions for application/json ContentType.
type PostNotificationsPushSubscriptionsJSONRequestBody = PushNotificationSubscription

// PutNotificationsPushSubscriptionsIdJSONRequestBody defines body for PutNotificationsPushSubscriptionsId for application/json ContentType.
type PutNotificationsPushSubscriptionsIdJSONRequestBody = PushNotificationSubscription

// PostTrackingConnectionsJSONRequestBody defines body for PostTrackingConnections for application/json ContentType.
type PostTrackingConnectionsJSONRequestBody = TrackedConnection

// PutTrackingConnectionsIdJSONRequestBody defines body for PutTrackingConnectionsId for application/json ContentType.
type PutTrackingConnectionsIdJSONRequestBody = TrackedConnectionUpdate

// AsDelayChangeAlarm returns the union data inside the Alarm_Content as a DelayChangeAlarm
func (t Alarm_Content) AsDelayChangeAlarm() (DelayChangeAlarm, error) {
	var body DelayChangeAlarm
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDelayChangeAlarm overwrites any union data inside the Alarm_Content as the provided DelayChangeAlarm
func (t *Alarm_Content) FromDelayChangeAlarm(v DelayChangeAlarm) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDelayChangeAlarm performs a merge with any union data inside the Alarm_Content, using the provided DelayChangeAlarm
func (t *Alarm_Content) MergeDelayChangeAlarm(v DelayChangeAlarm) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsCancelledAlarm returns the union data inside the Alarm_Content as a CancelledAlarm
func (t Alarm_Content) AsCancelledAlarm() (CancelledAlarm, error) {
	var body CancelledAlarm
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCancelledAlarm overwrites any union data inside the Alarm_Content as the provided CancelledAlarm
func (t *Alarm_Content) FromCancelledAlarm(v CancelledAlarm) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCancelledAlarm performs a merge with any union data inside the Alarm_Content, using the provided CancelledAlarm
func (t *Alarm_Content) MergeCancelledAlarm(v CancelledAlarm) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t Alarm_Content) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Alarm_Content) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /alarms)
	GetAlarms(ctx echo.Context, params GetAlarmsParams) error

	// (DELETE /alarms/{id})
	DeleteAlarmsId(ctx echo.Context, id int) error

	// (POST /auth/login)
	PostAuthLogin(ctx echo.Context) error

	// (POST /auth/logout)
	PostAuthLogout(ctx echo.Context) error

	// (GET /auth/me)
	GetAuthMe(ctx echo.Context) error

	// (POST /auth/register)
	PostAuthRegister(ctx echo.Context) error

	// (GET /bahn/connections)
	GetBahnConnections(ctx echo.Context, params GetBahnConnectionsParams) error

	// (GET /bahn/places)
	GetBahnPlaces(ctx echo.Context, params GetBahnPlacesParams) error

	// (GET /notifications/push-subscriptions)
	GetNotificationsPushSubscriptions(ctx echo.Context, params GetNotificationsPushSubscriptionsParams) error

	// (POST /notifications/push-subscriptions)
	PostNotificationsPushSubscriptions(ctx echo.Context) error

	// (DELETE /notifications/push-subscriptions/{id})
	DeleteNotificationsPushSubscriptionsId(ctx echo.Context, id int) error

	// (PUT /notifications/push-subscriptions/{id})
	PutNotificationsPushSubscriptionsId(ctx echo.Context, id int) error

	// (GET /tracking/connections)
	GetTrackingConnections(ctx echo.Context, params GetTrackingConnectionsParams) error

	// (POST /tracking/connections)
	PostTrackingConnections(ctx echo.Context) error

	// (DELETE /tracking/connections/{id})
	DeleteTrackingConnectionsId(ctx echo.Context, id int) error

	// (PUT /tracking/connections/{id})
	PutTrackingConnectionsId(ctx echo.Context, id int) error

	// (GET /tracking/stats)
	GetTrackingStats(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAlarms converts echo context to params.
func (w *ServerInterfaceWrapper) GetAlarms(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAlarmsParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// ------------- Optional query parameter "urgency" -------------

	err = runtime.BindQueryParameter("form", true, false, "urgency", ctx.QueryParams(), &params.Urgency)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter urgency: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAlarms(ctx, params)
	return err
}

// DeleteAlarmsId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteAlarmsId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteAlarmsId(ctx, id)
	return err
}

// PostAuthLogin converts echo context to params.
func (w *ServerInterfaceWrapper) PostAuthLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostAuthLogin(ctx)
	return err
}

// PostAuthLogout converts echo context to params.
func (w *ServerInterfaceWrapper) PostAuthLogout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostAuthLogout(ctx)
	return err
}

// GetAuthMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetAuthMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAuthMe(ctx)
	return err
}

// PostAuthRegister converts echo context to params.
func (w *ServerInterfaceWrapper) PostAuthRegister(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostAuthRegister(ctx)
	return err
}

// GetBahnConnections converts echo context to params.
func (w *ServerInterfaceWrapper) GetBahnConnections(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBahnConnectionsParams
	// ------------- Required query parameter "departure" -------------

	err = runtime.BindQueryParameter("form", true, true, "departure", ctx.QueryParams(), &params.Departure)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter departure: %s", err))
	}

	// ------------- Required query parameter "fromId" -------------

	err = runtime.BindQueryParameter("form", true, true, "fromId", ctx.QueryParams(), &params.FromId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fromId: %s", err))
	}

	// ------------- Required query parameter "toId" -------------

	err = runtime.BindQueryParameter("form", true, true, "toId", ctx.QueryParams(), &params.ToId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter toId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBahnConnections(ctx, params)
	return err
}

// GetBahnPlaces converts echo context to params.
func (w *ServerInterfaceWrapper) GetBahnPlaces(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBahnPlacesParams
	// ------------- Required query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, true, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBahnPlaces(ctx, params)
	return err
}

// GetNotificationsPushSubscriptions converts echo context to params.
func (w *ServerInterfaceWrapper) GetNotificationsPushSubscriptions(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNotificationsPushSubscriptionsParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNotificationsPushSubscriptions(ctx, params)
	return err
}

// PostNotificationsPushSubscriptions converts echo context to params.
func (w *ServerInterfaceWrapper) PostNotificationsPushSubscriptions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostNotificationsPushSubscriptions(ctx)
	return err
}

// DeleteNotificationsPushSubscriptionsId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteNotificationsPushSubscriptionsId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteNotificationsPushSubscriptionsId(ctx, id)
	return err
}

// PutNotificationsPushSubscriptionsId converts echo context to params.
func (w *ServerInterfaceWrapper) PutNotificationsPushSubscriptionsId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutNotificationsPushSubscriptionsId(ctx, id)
	return err
}

// GetTrackingConnections converts echo context to params.
func (w *ServerInterfaceWrapper) GetTrackingConnections(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTrackingConnectionsParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTrackingConnections(ctx, params)
	return err
}

// PostTrackingConnections converts echo context to params.
func (w *ServerInterfaceWrapper) PostTrackingConnections(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostTrackingConnections(ctx)
	return err
}

// DeleteTrackingConnectionsId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTrackingConnectionsId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteTrackingConnectionsId(ctx, id)
	return err
}

// PutTrackingConnectionsId converts echo context to params.
func (w *ServerInterfaceWrapper) PutTrackingConnectionsId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutTrackingConnectionsId(ctx, id)
	return err
}

// GetTrackingStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetTrackingStats(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTrackingStats(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/alarms", wrapper.GetAlarms)
	router.DELETE(baseURL+"/alarms/:id", wrapper.DeleteAlarmsId)
	router.POST(baseURL+"/auth/login", wrapper.PostAuthLogin)
	router.POST(baseURL+"/auth/logout", wrapper.PostAuthLogout)
	router.GET(baseURL+"/auth/me", wrapper.GetAuthMe)
	router.POST(baseURL+"/auth/register", wrapper.PostAuthRegister)
	router.GET(baseURL+"/bahn/connections", wrapper.GetBahnConnections)
	router.GET(baseURL+"/bahn/places", wrapper.GetBahnPlaces)
	router.GET(baseURL+"/notifications/push-subscriptions", wrapper.GetNotificationsPushSubscriptions)
	router.POST(baseURL+"/notifications/push-subscriptions", wrapper.PostNotificationsPushSubscriptions)
	router.DELETE(baseURL+"/notifications/push-subscriptions/:id", wrapper.DeleteNotificationsPushSubscriptionsId)
	router.PUT(baseURL+"/notifications/push-subscriptions/:id", wrapper.PutNotificationsPushSubscriptionsId)
	router.GET(baseURL+"/tracking/connections", wrapper.GetTrackingConnections)
	router.POST(baseURL+"/tracking/connections", wrapper.PostTrackingConnections)
	router.DELETE(baseURL+"/tracking/connections/:id", wrapper.DeleteTrackingConnectionsId)
	router.PUT(baseURL+"/tracking/connections/:id", wrapper.PutTrackingConnectionsId)
	router.GET(baseURL+"/tracking/stats", wrapper.GetTrackingStats)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
