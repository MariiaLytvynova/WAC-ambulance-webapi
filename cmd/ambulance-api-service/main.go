package main

import (
    "log"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/MariiaLytvynova/WAC-ambulance-webapi/api"
    "github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/ambulance_wl"
    "github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/db_service"
    "context"
    "time"
    "github.com/gin-contrib/cors"
)

func main() {
    log.Printf("Server started")

    //connection to database
    db, err := db_service.NewPostgresService(
        db_service.PostgresConfig{},
    )
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    port := os.Getenv("AMBULANCE_API_PORT")
    if port == "" {
        port = "8080"
    }
    environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
    if !strings.EqualFold(environment, "production") { // case insensitive comparison
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.New()
    engine.Use(gin.Recovery())

    corsMiddleware := cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{""},
        AllowCredentials: false,
        MaxAge: 12 * time.Hour,
    })
    engine.Use(corsMiddleware)

    // request routings
    handleFunctions := &ambulance_wl.ApiHandleFunctions{ //vytvaram objekty, ktore implementuju rozhrania
    AmbulanceConditionsAPI:  ambulance_wl.NewAmbulanceConditionsApi(db),
    AmbulanceWaitingListAPI: ambulance_wl.NewAmbulanceWaitingListApi(),
    ExaminationsAPI:         ambulance_wl.NewExaminationsApi(),
  }
  ambulance_wl.NewRouterWithGinEngine(engine, *handleFunctions) //spracuje url poziadavku

    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}