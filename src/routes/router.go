package routes

import (
	
	"net/http"
	"github.com/gin-gonic/gin"
	"domain/model"

)


type Router struct {
	 Ginrouter *gin.Engine
	 Routergroup *gin.RouterGroup
}


func NewRouter() *Router {
	
	r := new(Router)
	r.Ginrouter = gin.Default()
	r.Routergroup = r.Ginrouter.Group("api/v1")
	{
		r.Routergroup.GET("/:id", GetCity)
		
		r.Routergroup.DELETE("/", notImplemented)
		r.Routergroup.POST("/", notImplemented)
		r.Routergroup.PUT("/", notImplemented)
	}
	
	r.Ginrouter.Run(":8080")
	return r;
}


func notImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error" : "method not Implemented"})
}

func  GetCity(c *gin.Context){
	 
	 city := c.Param("id")
	 dbconn := model.NewDatabase("eoin", "pass","weather_app") 
	 citydao := model.NewCityDAO(dbconn)
	 citydata := citydao.GetByCity(city)
	 
	 if(len(citydata) < 1){
	 	c.JSON(http.StatusNotFound, gin.H{"error" : "not found"})
	 } else {
	 	 c.JSON(http.StatusOK,   gin.H{"cities" : citydata})
	 }
}




