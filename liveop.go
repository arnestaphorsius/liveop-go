package main

import (
    "github.com/gin-gonic/gin"
    . "github.com/arnestaphorsius/liveop-go/controllers"
)

func main() {

    router := gin.Default()

    router.GET("api/objecten", GetObjecten)

    router.Run(":8040")
}