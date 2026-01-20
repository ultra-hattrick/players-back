package ports

import "github.com/ultra-hattrick/players-back/internal/core/domain"

type PlayerRepository interface {
	InsertPlayers(players []*domain.Player) error
	GetGroupedPlayersByTeamID(teamID uint, page, pageSize int) ([]*domain.GroupedPlayer, error)
}
