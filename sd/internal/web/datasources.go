package web

import (
    "github.com/labstack/echo/v4"
    "github.com/zzzz465/portal/sd/internal/datasource"
)

func RegisterDatasourcesHandlers(g *echo.Group, dsMap map[string]datasource.Datasource) {
    g.GET("", getDatasourcesHandler(dsMap))
    g.GET("/", getDatasourcesHandler(dsMap))
}

type getDatasourcesResponse struct {
    DataSources []string `json:"dataSources"`
}

// getDatasourcesHandler godoc
// @Summary      get all datasources
// @Tags         datasources
// @Produce      json
// @Success      200  {object}  getDatasourcesResponse
// @Failure      400
// @Failure      500
// @Router       /datasources [get]
func getDatasourcesHandler(storeMap map[string]datasource.Datasource) func(c echo.Context) error {
    res := getDatasourcesResponse{
        DataSources: []string{},
    }

    for k, _ := range storeMap {
        res.DataSources = append(res.DataSources, k)
    }

    return func(c echo.Context) error {
        return c.JSON(200, res)
    }
}
