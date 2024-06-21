package service

import "github.com/ultra-hattrick/players-back/repositories"

type PlayersService struct {
	PlayersRepository *repositories.PlayersRepository
}

func NewPlayersService(repo *repositories.PlayersRepository) *PlayersService {
	return &PlayersService{PlayersRepository: repo}
}
