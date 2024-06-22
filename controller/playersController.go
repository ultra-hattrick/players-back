package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultra-hattrick/players-back/service"
	"github.com/ultra-hattrick/players-back/utils"
	"gitlab.com/uchile1/helper/helperHttp"
	"gitlab.com/uchile1/helper/helperLog"
)

type PlayersController struct {
	PlayersService *service.PlayersService
}

func NewPlayersController(service *service.PlayersService) *PlayersController {
	return &PlayersController{PlayersService: service}
}

func (controller *PlayersController) GetPlayersByIDAndTeamID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}

func (controller *PlayersController) GetPlayersByTeamID(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("team_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse("Invalid team id"))
		return
	}
	pageSize, err := utils.GetQueryInt(c, "pageSize", 1, func(arg int) bool { return arg <= 20 })
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	page, err := utils.GetQueryInt(c, "page", 1, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	playersGroup, err := controller.PlayersService.GetGroupedPlayersByTeamID(uint(teamID), page, pageSize)
	if err != nil {
		helperLog.Logger.Error().Err(err).Msg("Failed to fetch")
		c.JSON(http.StatusInternalServerError, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, helperHttp.NewSuccessResponse(playersGroup))
}

func (controller *PlayersController) CreatePlayersByTeamID(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("team_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse("Invalid team id"))
		return
	}
	players, err := controller.PlayersService.CreatePlayersHistory(uint(teamID))
	if err != nil {
		helperLog.Logger.Error().Err(err).Msg("Failed to fetch")
		c.JSON(http.StatusInternalServerError, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, helperHttp.NewSuccessResponse(players))
}

func (controller *PlayersController) FetchPlayersByTeamIDAndMatchID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}

func (controller *PlayersController) GetPlayersByTeamIDAndMatchID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}
