package web

import (
	"github.com/labstack/echo/v4"
	"github.com/zzzz465/portal/sd/internal/store"
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

func (s *HTTPServer) registerMiddlewares() {
	storeGroup := s.echo.Group("/store")
	RegisterStoreHandlers(storeGroup, s.store)
}

func (s *HTTPServer) Start() error {
	return s.echo.Start(":3000")
}
