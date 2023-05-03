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

// Defines values for ConnectionAlarmType.
const (
	ConnectionAlarmTypeConnectionAlarm ConnectionAlarmType = "connection-alarm"
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

// BahnStation defines model for BahnStation.
type BahnStation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// ConnectionAlarm defines model for ConnectionAlarm.
type ConnectionAlarm struct {
	Connection SimpleConnection    `json:"connection"`
	Message    string              `json:"message"`
	Type       ConnectionAlarmType `json:"type"`
}

// ConnectionAlarmType defines model for ConnectionAlarm.Type.
type ConnectionAlarmType string

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

// PushNotificationSubscriptionCreate defines model for PushNotificationSubscriptionCreate.
type PushNotificationSubscriptionCreate struct {
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
	Endpoint *string            `json:"endpoint,omitempty"`
	Keys     *map[string]string `json:"keys,omitempty"`
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
	From       BahnStation        `json:"from"`
	Id         *int               `json:"id,omitempty"`
	To         BahnStation        `json:"to"`
}

// TrackedConnectionList defines model for TrackedConnectionList.
type TrackedConnectionList struct {
	Connections []TrackedConnection `json:"connections"`
	Pagination  Pagination          `json:"pagination"`
}

// TrackedConnectionUpdate defines model for TrackedConnectionUpdate.
type TrackedConnectionUpdate struct {
	Departures []TrackedDepartureWrite `json:"departures"`
}

// TrackedConnectionWrite defines model for TrackedConnectionWrite.
type TrackedConnectionWrite struct {
	Departures []TrackedDepartureWrite `json:"departures"`
	From       BahnStation             `json:"from"`
	To         BahnStation             `json:"to"`
}

// TrackedDeparture defines model for TrackedDeparture.
type TrackedDeparture struct {
	Delay     int                    `json:"delay"`
	Departure time.Time              `json:"departure"`
	Status    TrackedDepartureStatus `json:"status"`
}

// TrackedDepartureStatus defines model for TrackedDeparture.Status.
type TrackedDepartureStatus string

// TrackedDepartureWrite defines model for TrackedDepartureWrite.
type TrackedDepartureWrite struct {
	Departure time.Time `json:"departure"`
}

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

// VapidKeys defines model for VapidKeys.
type VapidKeys struct {
	PublicKey string `json:"publicKey"`
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
type PostNotificationsPushSubscriptionsJSONRequestBody = PushNotificationSubscriptionCreate

// PutNotificationsPushSubscriptionsIdJSONRequestBody defines body for PutNotificationsPushSubscriptionsId for application/json ContentType.
type PutNotificationsPushSubscriptionsIdJSONRequestBody = PushNotificationSubscription

// PostTrackingConnectionsJSONRequestBody defines body for PostTrackingConnections for application/json ContentType.
type PostTrackingConnectionsJSONRequestBody = TrackedConnectionWrite

// PutTrackingConnectionsIdJSONRequestBody defines body for PutTrackingConnectionsId for application/json ContentType.
type PutTrackingConnectionsIdJSONRequestBody = TrackedConnectionUpdate

// AsConnectionAlarm returns the union data inside the Alarm_Content as a ConnectionAlarm
func (t Alarm_Content) AsConnectionAlarm() (ConnectionAlarm, error) {
	var body ConnectionAlarm
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromConnectionAlarm overwrites any union data inside the Alarm_Content as the provided ConnectionAlarm
func (t *Alarm_Content) FromConnectionAlarm(v ConnectionAlarm) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeConnectionAlarm performs a merge with any union data inside the Alarm_Content, using the provided ConnectionAlarm
func (t *Alarm_Content) MergeConnectionAlarm(v ConnectionAlarm) error {
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

	// (GET /notifications/vapid-keys)
	GetNotificationsVapidKeys(ctx echo.Context) error

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

// GetNotificationsVapidKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetNotificationsVapidKeys(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetNotificationsVapidKeys(ctx)
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
	router.GET(baseURL+"/notifications/vapid-keys", wrapper.GetNotificationsVapidKeys)
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
