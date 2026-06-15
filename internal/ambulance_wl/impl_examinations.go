package ambulance_wl

import (
    "net/http"
	"fmt"
    "github.com/gin-gonic/gin"
	"github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/db_service"

)

type implExaminationsAPI struct {
    db *db_service.PostgresService
}
//response from frontens
type ExaminationResponse struct {
    Id         int    `json:"id"`
    PatientId  int    `json:"patientId"`
    Day        string `json:"day"`
    StartTime  string `json:"startTime"`
    EndTime    string `json:"endTime"`
    Name       string `json:"name"`
}
func NewExaminationsApi(db *db_service.PostgresService) ExaminationsAPI {
	return &implExaminationsAPI{
		db: db,
	}
}

func (o implExaminationsAPI) CreateExamination(c *gin.Context) {
	
	// c.AbortWithStatus(http.StatusNotImplemented)
	_, err := o.db.GetCondition(c.Request.Context(), "some_code") //ziskame condition z db
	if err != nil {
		if err == db_service.ErrNotFound { //return json error
			c.JSON(http.StatusNotFound, gin.H{"error": "Condition not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
}

func (o implExaminationsAPI) DeleteExamination(c *gin.Context) {
	
	// c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implExaminationsAPI) UpdateExamination(c *gin.Context) {
fmt.Println("UpdateExamination HIT")
    var examination ExaminationResponse
    if err := c.BindJSON(&examination); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
fmt.Printf("UPDATE payload: %+v\n", examination)
    _, err := o.db.DB.Exec(`
        UPDATE examinations
        SET
            day = $1,
            start_time = $2,
            end_time = $3,
            name_examination = $4
        WHERE id = $5 and id_patient = $6
    `,
        examination.Day,
        examination.StartTime,
        examination.EndTime,
        examination.Name,
        examination.Id,
		examination.PatientId,
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "updated",
    })
}