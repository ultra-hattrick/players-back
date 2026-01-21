package postgres

import (
	"time"

	"github.com/ultra-hattrick/players-back/internal/core/domain"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	DB *gorm.DB
}

func NewPlayersRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{DB: db}
}

func (r *PlayerRepository) InsertPlayers(players []*domain.Player) error {
	return r.DB.Create(players).Error
}

func (r *PlayerRepository) GetGroupedPlayersByTeamID(teamID uint, page, pageSize int) ([]*domain.GroupedPlayer, error) {

	var results []*domain.GroupedPlayer

	// Subquery to find the latest ID for each player_id of the given team
	subquery := r.DB.Model(&domain.Player{}).
		Select("MAX(id)").
		Where("team_id = ?", teamID).
		Group("player_id")

	// Get the latest records for each player with pagination
	var players []domain.Player
	err := r.DB.Where("id IN (?)", subquery).
		Order("created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&players).Error
	if err != nil {
		return nil, err
	}

	// Map to store players grouped by creation date
	groupedPlayersMap := make(map[time.Time][]domain.Player)

	for _, player := range players {
		creationDate := player.CreatedAt
		groupedPlayersMap[creationDate] = append(groupedPlayersMap[creationDate], player)
	}

	// Populate results with grouped players

	for key := range groupedPlayersMap {
		results = append(results, &domain.GroupedPlayer{
			CreationDate: key,
			TeamID:       teamID,
			Players:      groupedPlayersMap[key],
		})
	}

	return results, nil
}
