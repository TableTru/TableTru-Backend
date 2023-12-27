package infrastructure

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

//GinRouter -> Gin Router
type GinRouter struct {
    Gin *gin.Engine
}

//NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

    httpRouter := gin.Default()
    
    config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // ระบุโดเมนที่ต้องการให้รับ CORS
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	httpRouter.Use(cors.New(config))

    // httpRouter.Use(cors.New(cors.Config{
    //     AllowOrigins: []string{"http://localhost:3000"},
    //     AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
    //     AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
    // }))


    httpRouter.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
    })
    return GinRouter{
        Gin: httpRouter,
    }

}