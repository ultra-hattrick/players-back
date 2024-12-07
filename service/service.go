package service

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ultra-hattrick/helper/utils"
	"github.com/ultra-hattrick/players-back/model"
	"github.com/ultra-hattrick/players-back/repositories"
)

type PlayersService struct {
	PlayersRepository *repositories.PlayersRepository
}

func NewPlayersService(repo *repositories.PlayersRepository) *PlayersService {
	return &PlayersService{PlayersRepository: repo}
}

func (s *PlayersService) GetGroupedPlayersByTeamID(teamID uint, page, pageSize int) ([]*model.GroupedPlayer, error) {
	return s.PlayersRepository.GetGroupedPlayersByTeamID(teamID, page, pageSize)
}

func (s *PlayersService) CreatePlayersHistory(teamID uint, loadedByJob bool) ([]*model.Player, error) {
	var hd model.HattrickData
	err := utils.GetResultsFromHattrick(os.Getenv("PLAYERS_TEAM")+strconv.Itoa(int(teamID)), &hd)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los jugadores del team ID: %d, desde hattrick: %v", teamID, err)
	}
	if hd.Error != nil {
		return nil, fmt.Errorf("error al obtener los jugadores del team ID: %d, hattrick error: %s", teamID, *hd.Error)
	}
	if hd.Players == nil {
		return nil, fmt.Errorf("invalid value players, must be distinct to null for team ID: %d", teamID)
	}
	//verificar si hay cambios en la forma y condicion de los jugadores,
	// si hubo cambios, es que se cargó en entrenamiento y se puede cargar un registro histórico
	GroupLastTraining, err := s.GetGroupedPlayersByTeamID(teamID, 1, 1)
	if err != nil {
		return nil, fmt.Errorf("error al obtener último registro de entrenamiento para el team ID: %d: %v", teamID, err)
	}
	if len(GroupLastTraining) != 0 {
		loadedTraining := false
		playersLastTraining := GroupLastTraining[0].Players
		for _, p := range playersLastTraining {
			for _, ht := range hd.Players {
				if ht.PlayerID == p.PlayerID && (ht.PlayerForm != p.PlayerForm || ht.StaminaSkill != p.StaminaSkill || ht.TSI != p.TSI) {
					loadedTraining = true
					break
				}
			}
			if loadedTraining {
				break
			}
		}
		if !loadedTraining {
			return nil, fmt.Errorf("no training (yet!) for team id: %d, last training loaded: %v", teamID, playersLastTraining[0].CreatedAt)
		}
	}
	for _, v := range hd.Players {
		v.TeamID = teamID
		v.LoadedByJob = loadedByJob
	}
	return hd.Players, s.PlayersRepository.InsertPlayers(hd.Players)
}
