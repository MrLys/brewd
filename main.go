package main

import (
        roast "lyseggen.xyz/brewd/roasters"
	"net/http"
	"github.com/gin-gonic/gin"
)



func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(
            http.StatusOK, gin.H{
                "message": "pong",
            },
    )}) 
    r.GET("/roasters", roast.getRoasters)
    r.Run()
}

