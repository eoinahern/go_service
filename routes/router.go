package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/eoinahern/go_service/domain/model"
	"github.com/eoinahern/go_service/utils"
	"github.com/gin-gonic/gin"
)

/*
* bit verbose for my liking also dont like this repeating dc connection.
**/

//Username: 	bd145d3b601f2e
//Password: 	532d35c9
//heroku_1587748f259385b

//var dbconn = model.NewDatabase("bd145d3b601f2e", "532d35c9", "heroku_1587748f259385b")
var dbconn = model.NewDatabase("eoin", "pass", "weather_app")

type Router struct {
	Ginrouter   *gin.Engine
	Routergroup *gin.RouterGroup
}

func NewRouter() *Router {

	r := new(Router)
	r.Ginrouter = gin.Default()
	r.Routergroup = r.Ginrouter.Group("api/v1")
	{
		//r.Routergroup.GET("/:id", GetCity)
		r.Routergroup.GET("/:lat/:long", GetWeatherData)
		r.Routergroup.DELETE("/", notImplemented)
		r.Routergroup.POST("/", notImplemented)
		r.Routergroup.PUT("/", notImplemented)
	}

	//run on herokus port no!!!
	//for testing use 8080

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//r.Ginrouter.Run(":8080")

	r.Ginrouter.Run(":" + port)
	return r
}

func notImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "method not Implemented"})
}

func GetCity(c *gin.Context) {

	city := c.Param("id")

	//dbconn := model.NewDatabase("bd145d3b601f2e", "532d35c9", "heroku_1587748f259385b")
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

	//dbconn := model.NewDatabase("bd145d3b601f2e", "532d35c9", "heroku_1587748f259385b")
	citydao := model.NewCityDAO(dbconn)
	allcities := citydao.GetAllCities()

	city := utils.FindClosest(allcities, lat, longit)
	fmt.Printf("closest city name iss = %s \n", city.Name)

	weatherdao := model.NewDailyWeatherDAO(dbconn)
	//problem exists with get call

	cityslice := weatherdao.Get(string(city.Name))
	fmt.Println(cityslice)

	var status int

	if len(cityslice) == 0 {
		status = 204
	} else {
		status = 200
	}

	c.Abort()
	c.JSON(status, gin.H{"data": cityslice})
}
