package routes

import (
	"net/http"
	"strconv"

	"github.com/eoinahern/go_service/domain/model"
	"github.com/eoinahern/go_service/utils"
	"github.com/gin-gonic/gin"
)

/*
* bit verbose for my liking also dont like this repeating dc connection.
**/

type Router struct {
	Ginrouter   *gin.Engine
	Routergroup *gin.RouterGroup
}

func NewRouter() *Router {

	r := new(Router)
	r.Ginrouter = gin.Default()
	r.Routergroup = r.Ginrouter.Group("api/v1")
	{
		r.Routergroup.GET("/:id", GetCity)
		r.Routergroup.GET("/:lat,:long", GetWeatherData)
		r.Routergroup.DELETE("/", notImplemented)
		r.Routergroup.POST("/", notImplemented)
		r.Routergroup.PUT("/", notImplemented)
	}

	r.Ginrouter.Run(":8080")
	return r
}

func notImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "method not Implemented"})
}

func GetCity(c *gin.Context) {

	city := c.Param("id")

	dbconn := model.NewDatabase("eoin", "pass", "weather_app")
	citydao := model.NewCityDAO(dbconn)
	citydata := citydao.GetByCity(city)

	if len(citydata) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"cities": citydata})
	}
}

func GetWeatherData(c *gin.Context) {

	lat, err1 := strconv.ParseFloat(c.Param("lat"), 64)
	longit, err2 := strconv.ParseFloat(c.Param("long"), 64)

	if err1 != nil || err2 != nil {
		println("conv error")
	}

	println(lat)
	println(longit)

	//need seperate dbconnection here.
	dbconn := model.NewDatabase("eoin", "pass", "weather_app")
	citydao := model.NewCityDAO(dbconn)
	allcities := citydao.GetAllCities()

	city := utils.FindClosest(allcities, lat, longit)
	weatherdao := model.NewDailyWeatherDAO(dbconn)
	cityslice := weatherdao.Get(city.Name)

	if len(cityslice) > 0 {
		c.JSON(http.StatusOK, gin.H{"data": cityslice})
	}

}
