package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultra-hattrick/helper/utils"
	_ "github.com/ultra-hattrick/players-back/internal/core/domain"
	"github.com/ultra-hattrick/players-back/internal/core/ports"
	"gitlab.com/uchile1/helper/helperHttp"
	"gitlab.com/uchile1/helper/helperLog"
)

type PlayersHandler struct {
	PlayersService ports.PlayerService
}

func NewPlayersHandler(service ports.PlayerService) *PlayersHandler {
	return &PlayersHandler{PlayersService: service}
}

// GetPlayersByIDAndTeamID godoc
// @Summary      Get players by ID and Team ID
// @Description  Get players by ID and Team ID - Not Implemented
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        team_id   path      int  true  "Team ID"
// @Param        player_id path      int  true  "Player ID"
// @Success      501       {object}  helperHttp.Response
// @Router       /{player_id} [get]
func (h *PlayersHandler) GetPlayersByIDAndTeamID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}

// GetPlayersByTeamID godoc
// @Summary      Get grouped players by Team ID
// @Description  Get a list of players for a specific team
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        team_id   path      int  true  "Team ID"
// @Param        pageNumber      query     int  false "Page number"
// @Param        pageSize  query     int  false "Page size"
// @Success      200       {object}  helperHttp.Response{data=[]domain.GroupedPlayer}
// @Failure      400       {object}  helperHttp.Response
// @Failure      500       {object}  helperHttp.Response
// @Router       / [get]
func (h *PlayersHandler) GetPlayersByTeamID(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("team_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse("Invalid team id"))
		return
	}
	pageSize, err := utils.GetQueryInt(c, "pageSize", 10, func(arg int) bool { return arg <= 20 })
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	pageNumber, err := utils.GetQueryInt(c, "pageNumber", 1, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	playersGroup, err := h.PlayersService.GetGroupedPlayersByTeamID(uint(teamID), pageNumber, pageSize)
	if err != nil {
		helperLog.Logger.Error().Err(err).Msg("Failed to fetch")
		c.JSON(http.StatusInternalServerError, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, helperHttp.NewSuccessResponse(playersGroup))
}

// CreatePlayersByTeamID godoc
// @Summary      Create players history
// @Description  Fetch players from Hattrick and store them in the database
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        team_id       path      int  true  "Team ID"
// @Param        loaded_by_job query     bool false "Loaded by job"
// @Success      200           {object}  helperHttp.Response{data=[]domain.Player}
// @Failure      400           {object}  helperHttp.Response
// @Failure      500           {object}  helperHttp.Response
// @Router       / [post]
func (h *PlayersHandler) CreatePlayersByTeamID(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("team_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse("Invalid team id"))
		return
	}
	loadedByJob, err := utils.GetQueryBool(c, "loaded_by_job", false)
	if err != nil {
		c.JSON(http.StatusBadRequest, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	players, err := h.PlayersService.CreatePlayersHistory(uint(teamID), loadedByJob)
	if err != nil {
		helperLog.Logger.Error().Err(err).Msg("Failed to fetch")
		c.JSON(http.StatusInternalServerError, helperHttp.NewErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, helperHttp.NewSuccessResponse(players))
}

// FetchPlayersByTeamIDAndMatchID godoc
// @Summary      Fetch players by Team ID and Match ID
// @Description  Fetch players by Team ID and Match ID - Not Implemented
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        team_id  path      int  true  "Team ID"
// @Param        match_id path      int  true  "Match ID"
// @Success      501      {object}  helperHttp.Response
// @Router       /{team_id}/player/match/{match_id} [get]
func (h *PlayersHandler) FetchPlayersByTeamIDAndMatchID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}

// GetPlayersByTeamIDAndMatchID godoc
// @Summary      Get players by Team ID and Match ID
// @Description  Get players by Team ID and Match ID - Not Implemented
// @Tags         players
// @Accept       json
// @Produce      json
// @Param        team_id  path      int  true  "Team ID"
// @Param        match_id path      int  true  "Match ID"
// @Success      501      {object}  helperHttp.Response
// @Router       /{team_id}/player/match/{match_id} [post]
func (h *PlayersHandler) GetPlayersByTeamIDAndMatchID(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, helperHttp.NewErrorResponse("Not Implemented"))
}
