package ambulance_wl

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/db_service"

)

type implExaminationsAPI struct {
    db *db_service.PostgresService
}

func NewExaminationsApi(db *db_service.PostgresService) ExaminationsAPI {
	return &implExaminationsAPI{
		db: db,
	}
}

func (o implExaminationsAPI) CreateExamination(c *gin.Context) {
	// patientId := c.Param("patientId")
	//examinationId := c.Param("id")
	
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
//	patientId := c.Param("patientId")
//	examinationId := c.Param("id")
	
	// c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implExaminationsAPI) UpdateExamination(c *gin.Context) {
	//patientId := c.Param("patientId")
//	examinationId := c.Param("id")
	
	//c.AbortWithStatus(http.StatusNotImplemented)
}	
