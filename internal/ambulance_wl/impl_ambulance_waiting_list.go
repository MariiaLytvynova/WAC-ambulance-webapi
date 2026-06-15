package ambulance_wl

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/MariiaLytvynova/WAC-ambulance-webapi/internal/db_service"

)

type implAmbulanceWaitingListAPI struct {
    db *db_service.PostgresService

}

func NewAmbulanceWaitingListApi(db *db_service.PostgresService) AmbulanceWaitingListAPI {
return &implAmbulanceWaitingListAPI{
		db: db,
	}
}

func (o implAmbulanceWaitingListAPI) CreateWaitingListEntry(c *gin.Context) {
    c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceWaitingListAPI) DeleteWaitingListEntry(c *gin.Context) {
    c.AbortWithStatus(http.StatusNotImplemented)
}

func (o implAmbulanceWaitingListAPI) GetWaitingListEntries(c *gin.Context) {

    rows, err := o.db.DB.Query(`
        SELECT id_patient, name, fullname FROM patients
    `) //ziskaj vsetky vysetrenia pacienta z db

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal server error",
        })
        return
    }
    defer rows.Close()

    var waiting [] WaitingListEntry

    for rows.Next() {
        var e WaitingListEntry

        err := rows.Scan(
            &e.IdPatient,
            &e.Name,
            &e.Fullname,
        )

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Scan error",
            })
            return
        }

        waiting = append(waiting, e)
    }

    c.JSON(http.StatusOK, waiting)


}

func (o implAmbulanceWaitingListAPI) GetWaitingListEntry(c *gin.Context) {
    rows, err := o.db.DB.Query(`
        SELECT id_patient, name, fullname FROM patients
    `) //ziskaj vsetky vysetrenia pacienta z db

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal server error",
        })
        return
    }
    defer rows.Close()

    var waiting [] WaitingListEntry

    for rows.Next() {
        var e WaitingListEntry

        err := rows.Scan(
            &e.IdPatient,
            &e.Name,
            &e.Fullname,
        )

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Scan error",
            })
            return
        }

        waiting = append(waiting, e)
    }

    c.JSON(http.StatusOK, waiting)

}

func (o implAmbulanceWaitingListAPI) UpdateWaitingListEntry(c *gin.Context) {
    c.AbortWithStatus(http.StatusNotImplemented)
}