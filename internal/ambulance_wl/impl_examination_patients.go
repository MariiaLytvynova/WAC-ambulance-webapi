package ambulance_wl

import (
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
	"github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/db_service"

)

type implExaminationPatientAPI struct {
    db *db_service.PostgresService
}

func NewExaminationsPatientApi(db *db_service.PostgresService) ExaminationpatientAPI {
	return &implExaminationPatientAPI{
		db: db,
	}
}

func (o implExaminationPatientAPI) CreateExamination(c *gin.Context) {
	fmt.Println("Add new examination")
 var examination ExaminationResponse
    if err := c.BindJSON(&examination); err != nil {
            fmt.Println("----- BindJSON error:", err)

        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    fmt.Println("id_patient =", examination.PatientId)

     _, err := o.db.DB.Exec(`
        INSERT INTO examinations (
            id_patient,
            day,
            start_time,
            end_time,
            name_examination
        )
        VALUES ($1, $2, $3, $4, $5)
    `,
        examination.PatientId,
        examination.Day,
        examination.StartTime,
        examination.EndTime,
        examination.Name,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Examination created successfully",
    })
}

func (o implExaminationPatientAPI) DeleteExamination(c *gin.Context) {
//	patientId := c.Param("patientId")
//	examinationId := c.Param("id")
	
	// c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implExaminationPatientAPI) UpdateExamination(c *gin.Context) {
	fmt.Println("UpdateExamination HIT")
    var examination ExaminationResponse
    if err := c.BindJSON(&examination); err != nil {
            fmt.Println("----- BindJSON error:", err)

        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
fmt.Printf(" AAAAAAAAAAAAAAAAAAAAAA: %+v\n", examination)
//     res, err := o.db.DB.Exec(`
//         UPDATE examinations
//         SET
//             day = $1,
//             start_time = $2,
//             end_time = $3,
//             name_examination = $4
//         WHERE id = $5
//     `,
//         examination.Day,
//         examination.StartTime,
//         examination.EndTime,
//         examination.Name,
//         examination.Id,
//     )

//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{
//             "error": err.Error(),
//         })
//         return
//     }
// rows, _ := res.RowsAffected()
// fmt.Println("ROWS AFFECTED:", rows)
//     c.JSON(http.StatusOK, gin.H{
//         "message": "updated",
//     })
res, err := o.db.DB.Exec(`
    UPDATE examinations
    SET day = $1,
        start_time = $2,
        end_time = $3,
        name_examination = $4
    WHERE id = $5
`,
    examination.Day,
    examination.StartTime,
    examination.EndTime,
    examination.Name,
    examination.Id,
)

if err != nil {
    fmt.Println("SQL ERROR:", err)
    return
}

rows, _ := res.RowsAffected()
fmt.Println("ROWS AFFECTED:", rows)
}	

func (o implExaminationPatientAPI) GetExaminations(c *gin.Context) {
    patientId := c.Param("patientId") //ziskaj s frontend api callu

    rows, err := o.db.DB.Query(`
        SELECT id, id_patient, day, start_time, end_time, name_examination
        FROM examinations
        WHERE id_patient = $1
    `, patientId) //ziskaj vsetky vysetrenia pacienta z db

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal server error",
        })
        return
    }
    defer rows.Close()

    var examinations []Examination

    for rows.Next() {
        var e Examination

        err := rows.Scan(
            &e.Id,
            &e.PatientId,
            &e.Day,
            &e.StartTime,
            &e.EndTime,
            &e.Name,
        )

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Scan error",
            })
            return
        }

        examinations = append(examinations, e)
    }

    c.JSON(http.StatusOK, examinations)
}
