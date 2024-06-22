package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/dghubble/oauth1"
	"github.com/ultra-hattrick/players-back/model"
	"github.com/ultra-hattrick/players-back/repositories"
	"gitlab.com/uchile1/helper/helperCommon"
	"gitlab.com/uchile1/helper/helperLog"
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

func (s *PlayersService) CreatePlayersHistory(teamID uint) ([]*model.Player, error) {
	hd, err := getResultsFromHattrick(os.Getenv("PLAYERS_TEAM") + strconv.Itoa(int(teamID)))
	if err != nil {
		return nil, fmt.Errorf("error al obtener los jugadores del team ID: %d, desde hattrick: %v", teamID, err)
	}
	if hd.Players == nil {
		return nil, fmt.Errorf("invalid value players, must be distinct to null for team ID: %d", teamID)
	}
	for _, v := range hd.Players {
		v.TeamID = teamID
	}
	return hd.Players, s.PlayersRepository.InsertPlayers(hd.Players)
}

func getResultsFromHattrick(pathHattrick string) (*model.HattrickData, error) {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	httpClient := config.Client(oauth1.NoContext, oauth1.NewToken(os.Getenv("OAUTH1_TOKEN"), os.Getenv("OAUTH1_TOKEN_SECRET")))
	path := fmt.Sprintf("%s%s", os.Getenv("BASE_RESOURCE_URL"), pathHattrick)
	resp, err := httpClient.Get(path)
	helperLog.Logger.Warn().Str(
		"function", helperCommon.GetFrame(1).Function,
	).Msgf("Se ocupa API Hattrick url: %s", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Imprimir el status code
	helperLog.Logger.Debug().Msgf("HTTP Status Code: %d para la url: %s", resp.StatusCode, path)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Status Code: %d para la url: %s", resp.StatusCode, path)
	}

	// Leer el contenido del body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	helperLog.Logger.Debug().Msgf("Response Body: %s", string(body))

	// Reiniciar el cuerpo de la respuesta para que pueda ser decodificado
	resp.Body = io.NopCloser(bytes.NewBuffer(body))

	var hattrickData model.HattrickData
	err = xml.NewDecoder(resp.Body).Decode(&hattrickData)
	if err != nil {
		return nil, err
	}
	if hattrickData.Error != nil {
		return nil, fmt.Errorf("%s", *hattrickData.Error)
	}
	// helperLog.Logger.Debug().Msgf("--->Arena: %v", hattrickData.Match.DetailsMatch.Arena)
	return &hattrickData, nil
}
