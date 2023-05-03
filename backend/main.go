package main

import (
	"github.com/coma64/bahn-alarm-backend/handlers"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	bahnAlarmApi := &handlers.BahnAlarmApi{}
	apiHandler := server.NewStrictHandler(bahnAlarmApi, nil)
	server.RegisterHandlers(e, apiHandler)

	e.File("/docs", "static/docs.html")
	e.Static("/static/swagger", "swagger-ui/dist")
	e.File("/openapi.yml", "openapi.yml")

	e.Logger.Fatal(e.Start(":8090"))
}
