package main
import (
    "fmt"

"github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/db_service")


func main() {
    db, err := db_service.NewPostgresService(db_service.PostgresConfig{})
    if err != nil {
        panic(err)
    }
    defer db.Close()

    fmt.Println("Connected successfully")
}