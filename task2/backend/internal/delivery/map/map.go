package _map

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"software-engineering-2/internal/delivery/map/common"
	_map "software-engineering-2/internal/usecase/map"
)

type Delivery struct {
	usecase _map.UseCase
}

func NewMapDelivery(usecase _map.UseCase) *Delivery {
	return &Delivery{
		usecase: usecase,
	}
}

func (d *Delivery) RegisterRoutes(g *echo.Group) {
	mapGroup := g.Group("/maps")

	mapGroup.GET("/names", d.getAllMapNames)
	mapGroup.POST("", d.addMap)
	mapGroup.DELETE("/:mapName", d.deleteMap)

	mapGroup.GET("/:mapName/cities", d.getAllCities)
	mapGroup.POST("/:mapName/cities", d.addCity)
	mapGroup.PATCH("/:mapName/cities/:cityName", d.updateCityName)
	mapGroup.DELETE("/:mapName/cities/:cityName", d.deleteCity)

	mapGroup.GET("/:mapName/roads", d.getAllRoads)
	mapGroup.POST("/:mapName/roads", d.addRoad)
	mapGroup.PATCH("/:mapName/roads", d.updateRoadCost)
	mapGroup.DELETE("/:mapName/roads", d.deleteRoad)

	mapGroup.POST("/:mapName/undo", d.undo)
	mapGroup.POST("/:mapName/redo", d.redo)
	mapGroup.POST("/:mapName/download", d.download)
	mapGroup.POST("/upload", d.upload)
}

func (d *Delivery) getAllMapNames(c echo.Context) error {
	maps, err := d.usecase.GetAllMapNames()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to get all map names: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Data: maps,
	})
}

func (d *Delivery) getAllCities(c echo.Context) error {
	mapName := c.Param("mapName")

	cities, err := d.usecase.GetAllCities(mapName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to get all cities: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Data: cities,
	})
}

func (d *Delivery) getAllRoads(c echo.Context) error {
	mapName := c.Param("mapName")

	roads, err := d.usecase.GetAllRoads(mapName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to get all roads: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Data: roads,
	})
}

func (d *Delivery) addMap(c echo.Context) error {
	mapName := c.QueryParam("name")

	if err := d.usecase.AddMap(mapName); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to add map: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusCreated, common.Response{
		Message: "Map added successfully",
	})
}

func (d *Delivery) addCity(c echo.Context) error {
	mapName := c.Param("mapName")

	var req common.CityCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Error: fmt.Sprintf("Invalid reqeust body: %s", err.Error()),
		})
	}

	if err := d.usecase.AddCity(mapName, req); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to add city: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusCreated, common.Response{
		Message: "City added successfully",
	})
}

func (d *Delivery) addRoad(c echo.Context) error {
	mapName := c.Param("mapName")

	var req common.RoadCreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Error: fmt.Sprintf("Invalid reqeust body: %s", err.Error()),
		})
	}

	if err := d.usecase.AddRoad(mapName, req); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to add road: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Message: "Road added successfully",
	})
}

func (d *Delivery) updateCityName(c echo.Context) error {
	mapName := c.Param("mapName")
	oldCityName := c.Param("cityName")
	newCityName := c.QueryParam("name")

	if err := d.usecase.UpdateCityName(mapName, oldCityName, newCityName); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to update city name: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Message: "City updated successfully",
	})
}

func (d *Delivery) updateRoadCost(c echo.Context) error {
	mapName := c.Param("mapName")

	var req common.RoadUpdateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Error: fmt.Sprintf("Invalid reqeust body: %s", err.Error()),
		})
	}

	if err := d.usecase.UpdateRoadCost(mapName, req); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to update road cost: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Message: "Road cost updated successfully",
	})
}

func (d *Delivery) deleteMap(c echo.Context) error {
	mapName := c.Param("mapName")

	if err := d.usecase.DeleteMap(mapName); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to delete map: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Message: "Map deleted successfully",
	})
}

func (d *Delivery) deleteCity(c echo.Context) error {
	mapName := c.Param("mapName")
	cityName := c.Param("cityName")

	if err := d.usecase.DeleteCity(mapName, cityName); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to delete city: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Message: "City deleted successfully",
	})
}

func (d *Delivery) deleteRoad(c echo.Context) error {
	mapName := c.Param("mapName")

	var req common.RoadDeleteRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Error: fmt.Sprintf("Invalid reqeust body: %s", err.Error()),
		})
	}

	if err := d.usecase.DeleteRoad(mapName, req); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to delete road: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Message: "Road deleted successfully",
	})
}

func (d *Delivery) undo(c echo.Context) error {
	mapName := c.Param("mapName")

	if err := d.usecase.Undo(mapName); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to undo action: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (d *Delivery) redo(c echo.Context) error {
	mapName := c.Param("mapName")

	if err := d.usecase.Redo(mapName); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to redo action: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (d *Delivery) download(c echo.Context) error {
	mapName := c.Param("mapName")

	mapData, err := d.usecase.Download(mapName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to download map: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Data: mapData,
	})
}

func (d *Delivery) upload(c echo.Context) error {
	var mapData _map.MapData

	if err := c.Bind(&mapData); err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Error: fmt.Sprintf("Invalid request body: %s", err.Error()),
		})
	}

	if err := d.usecase.Upload(&mapData); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Error: fmt.Sprintf("Failed to upload map"),
		})
	}

	return c.JSON(http.StatusOK, common.Response{
		Message: "Map uploaded successfully",
	})
}
