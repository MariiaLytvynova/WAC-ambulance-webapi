package ambulance_wl

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type implExaminationsAPI struct {
}

func NewExaminationsApi() ExaminationsAPI {
	return &implExaminationsAPI{}
}

func (o implExaminationsAPI) CreateExamination(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implExaminationsAPI) DeleteExamination(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implExaminationsAPI) UpdateExamination(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}	
