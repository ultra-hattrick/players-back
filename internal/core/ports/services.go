package ports

import "github.com/ultra-hattrick/players-back/internal/core/domain"

type PlayerService interface {
	GetGroupedPlayersByTeamID(teamID uint, pageNumber, pageSize int) ([]*domain.GroupedPlayer, error)
	CreatePlayersHistory(teamID uint, loadedByJob bool) ([]*domain.Player, error)
}
