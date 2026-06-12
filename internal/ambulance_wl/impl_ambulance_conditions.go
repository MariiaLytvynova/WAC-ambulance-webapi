package ambulance_wl

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type implAmbulanceConditionsAPI struct { // struktura implementacie 
}

func NewAmbulanceConditionsApi() AmbulanceConditionsAPI { //vrati object
    return &implAmbulanceConditionsAPI{}
}

func (o implAmbulanceConditionsAPI) GetConditions(c *gin.Context) { //mame getconditions v interface AmbulanvceConditionsAPI
    c.AbortWithStatus(http.StatusNotImplemented) //vrati clientovy
}