package roasters

import (
	"net/http"
	"github.com/gin-gonic/gin"
        t "lyseggen.xyz/brewd/types"
)

func getRoastersActual() {
    
}
func getRoasters(c *gin.Context) {
    roasters := []t.Roaster{}
    roasters[0] = t.Roaster{Id: 1, Name: "Tim Wendelboe", CountryId: 1, Address: "Ss"}

    c.JSON(http.StatusOK, gin.H{
        "message": roasters,
})
}
