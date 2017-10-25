package Router

import (
	handler "github.com/NirmalVatsyayan/FuzzyWuzzyGo/Handlers"
	"gopkg.in/gin-gonic/gin.v1"
)

func Routes(port string) *gin.Engine{

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apis := router.Group("/v1.0/")
	apis.POST("/fuzzyscore", handler.GetFuzzyScore)

	return router
}

