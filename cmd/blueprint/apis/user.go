package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/trung85-go/go-project-blueprint/cmd/blueprint/daos"
	"github.com/trung85-go/go-project-blueprint/cmd/blueprint/services"
	"log"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
