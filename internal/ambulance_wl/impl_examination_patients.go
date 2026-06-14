package ambulance_wl

import (
    "net/http"

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
 //posli vsetky vysetrenia pacienta ako json
}

func (o implExaminationPatientAPI) DeleteExamination(c *gin.Context) {
//	patientId := c.Param("patientId")
//	examinationId := c.Param("id")
	
	// c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implExaminationPatientAPI) UpdateExamination(c *gin.Context) {
	//patientId := c.Param("patientId")
//	examinationId := c.Param("id")
	
	//c.AbortWithStatus(http.StatusNotImplemented)
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
