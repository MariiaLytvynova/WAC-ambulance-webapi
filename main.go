package main //is used to create db tables 
import (
    "fmt"

"github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/db_service"
)


func main() {
    db, err := db_service.NewPostgresService(db_service.PostgresConfig{})
    if err != nil {
        panic(err)
    }
    createTables(db)
    defer db.Close()

    fmt.Println("Connected successfully")

}

func createTables(db *db_service.PostgresService) {
    query := `
        CREATE TABLE IF NOT EXISTS doctors (
            id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            name VARCHAR(255) NOT NULL
        )
    `

    _, err := db.DB.Exec(query)
    if err != nil {
        panic(err)
    }

    query2 := `
        CREATE TABLE IF NOT EXISTS patients (
            id_patient INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            id_doctor INT NOT NULL,
            name VARCHAR(255) NOT NULL,
            fullname VARCHAR(255) NOT NULL,

            CONSTRAINT fk_patient_doctor
            FOREIGN KEY (id_doctor)
            REFERENCES doctors(id)
)
    `

    _, err = db.DB.Exec(query2)
    if err != nil {
        panic(err)
    }

    query3 := `
       CREATE TABLE IF NOT EXISTS examinations (
            id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            id_patient INT NOT NULL,
            day VARCHAR(50) NOT NULL,
            start_time TIME NOT NULL,
            end_time TIME NOT NULL,
            name_examination VARCHAR(255) NOT NULL,

            CONSTRAINT fk_examination_patient
            FOREIGN KEY (id_patient)
            REFERENCES patients(id_patient)
)
    `
    _, err = db.DB.Exec(query3)
    if err != nil {
        panic(err)
    }
}
