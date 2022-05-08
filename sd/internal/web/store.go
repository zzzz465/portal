package web

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	errors2 "github.com/zzzz465/portal/sd/internal/errors"
	"github.com/zzzz465/portal/sd/internal/store"
)

func RegisterStoreHandlers(g *echo.Group, store store.Store) {
	g.GET("/", func(c echo.Context) error {
		return c.String(200, fmt.Sprintf("usage: %s/<datasource>", c.Path()))
	})

	g.GET("/:datasource", getRecordsHandler(store))
}

func getRecordsHandler(store store.Store) func(c echo.Context) error {
	return func(c echo.Context) error {
		datasource := c.Param("datasource")
		records, err := store.GetRecord(datasource)
		if err != nil {
			if errors.Is(err, errors2.ErrNotExist) {
				return c.String(400, fmt.Sprintf("datasource %s not exists.", datasource))
			} else {
				return c.String(500, fmt.Sprintf("internal error: %+v", err))
			}
		}

		return c.JSON(200, records)
	}
}
