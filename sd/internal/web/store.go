package web

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	errors2 "github.com/zzzz465/portal/sd/internal/errors"
	"github.com/zzzz465/portal/sd/internal/store"
	"path"

	_ "github.com/zzzz465/portal/sd/docs"
)

func RegisterStoreHandlers(g *echo.Group, store store.Store) {
	g.GET("/", func(c echo.Context) error {
		return c.String(200, fmt.Sprintf("usage: %s", path.Join(c.Path(), "<datasource>")))
	})

	g.GET("/:datasource", getRecordsHandler(store))
}

// getRecordsHandler godoc
// @Summary      get all records of given data source
// @Tags         records
// @Produce      json
// @Success      200  {object}  []types.Record
// @Failure      400
// @Failure      500
// @Router       /store [get]
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
