package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ultra-hattrick/players-back/service"
	"gitlab.com/uchile1/helper/helperHttp"
)

type PlayersController struct {
	PlayersService *service.PlayersService
}

func NewPlayersController(service *service.PlayersService) *PlayersController {
	return &PlayersController{PlayersService: service}
}

func (controller *PlayersController) GetPlayersByTeamID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}

func (controller *PlayersController) GetPlayersByIDAndTeamID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}

func (controller *PlayersController) FetchPlayersByTeamID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}

func (controller *PlayersController) FetchPlayersByTeamIDAndMatchID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}

func (controller *PlayersController) GetPlayersByTeamIDAndMatchID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}
