package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/arnestaphorsius/liveop-go/models"
    "github.com/arnestaphorsius/liveop-go/db"
    "fmt"
)

func GetObjecten(c *gin.Context) {

    conn := db.Connection()

    var objecten []models.Object

    postcode := c.Query("postcode")
    straat := c.Query("straat")
    huisnummer := c.Query("huisnummer")
    x_coord := c.Query("x")
    y_coord := c.Query("y")

    query := conn.Model(&objecten);

    if postcode != "" { query = query.Where("adres_postcode = ?", postcode) }
    if straat != "" { query = query.Where("adres_straatnaam = ?", straat) }
    if huisnummer != "" { query = query.Where("adres_huisnummer = ?", huisnummer) }

    if x_coord != "" && y_coord != "" {
        query = query.Where(
            "ST_DWithin(geometry_center_geo, ST_SETSRID(" +
            "ST_MakePoint(?::double precision, ?::double precision), 28992), 200)", x_coord, y_coord).
        Order("ST_Distance(geometry_center_geo, ST_SetSRID(" +
            "ST_MakePoint(?::double precision, ?::double precision), 28992))", x_coord, y_coord)

    }

    if err := query.Limit(20).Select(); err != nil || objecten == nil {

        NotFound(c)
        fmt.Println(err)

    } else {
        c.IndentedJSON(200, gin.H{"objecten": objecten})
    }

}