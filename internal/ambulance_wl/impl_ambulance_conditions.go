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
        //var cnd ambulance_wl.Condition
	ambulanceId := c.Param("ambulanceId")

	if ambulanceId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing ambulanceId",
		})
		return
	}

	ctx := c.Request.Context()

	// SQL query
	rows, err := o.db.DB.QueryContext(ctx, `
		SELECT value, code, reference, typical_duration_minutes
		FROM conditions
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()


    dbCnd, _ := o.db.GetCondition(ctx, code)
    var cnd ambulance_wl.Condition
    result := ambulance_wl.Condition{
	Value: dbCnd.Value,
	Code: dbCnd.Code,
	Reference: dbCnd.Reference,
	TypicalDurationMinutes: dbCnd.TypicalDurationMinutes,
}

	// výsledný zoznam
	//var result []ambulance_wl.Condition

	for rows.Next() { //nacitam riadky z db
		var cnd ambulance_wl.Condition

		err := rows.Scan(
			&cnd.Value,
			&cnd.Code,
			&cnd.Reference,
			&cnd.TypicalDurationMinutes,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		result = append(result, cnd) //pridam do vysledneho zoznamu
	}

	// kontrola chyby iterácie
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result) //vratim vysledok vo forme jsonu
}