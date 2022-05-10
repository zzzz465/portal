package web

import (
    "github.com/labstack/echo/v4"
    echoSwagger "github.com/swaggo/echo-swagger"
    "github.com/zzzz465/portal/sd/internal/store"

    _ "github.com/zzzz465/portal/sd/docs"
)

type HTTPServer struct {
    store store.Store
    echo  *echo.Echo
}

func NewHTTPServer(store store.Store) *HTTPServer {
    e := echo.New()

    s := &HTTPServer{
        store: store,
        echo:  e,
    }

    s.registerMiddlewares()

    return s
}

// @title Swagger API
// @version 1.0
func (s *HTTPServer) registerMiddlewares() {
    storeGroup := s.echo.Group("/records")
    RegisterStoreHandlers(storeGroup, s.store)

    s.echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (s *HTTPServer) Start() error {
    return s.echo.Start(":3000")
}
