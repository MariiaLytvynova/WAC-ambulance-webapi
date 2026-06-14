package ambulance_wl
// package db_service
import (
	"net/http"

	// "github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/ambulance_wl"
	"github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
)

type implAmbulanceConditionsAPI struct { // struktura implementacie 
    db *db_service.PostgresService

}

func NewAmbulanceConditionsApi(db *db_service.PostgresService) AmbulanceConditionsAPI { //vrati object
return &implAmbulanceConditionsAPI{
        db: db,
    }
}


// func (o implAmbulanceConditionsAPI) GetConditions(c *gin.Context) { //mame getconditions v interface AmbulanvceConditionsAPI
//     c.AbortWithStatus(http.StatusNotImplemented) //vrati clientovy
// }

func (o implAmbulanceConditionsAPI) GetConditions(c *gin.Context) {
    _, err := o.db.GetCondition(c.Request.Context(), "some_code") //ziskame condition z db
	if err != nil {
		if err == db_service.ErrNotFound { //return json error
			c.JSON(http.StatusNotFound, gin.H{"error": "Condition not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	// c.AbortWithStatus(http.StatusNotImplemented) //vrati clientovy

}