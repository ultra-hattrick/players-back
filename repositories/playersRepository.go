package repositories

import (
	"time"

	"github.com/ultra-hattrick/players-back/model"
	"gorm.io/gorm"
)

type PlayersRepository struct {
	DB *gorm.DB
}

func NewPlayersRepository(db *gorm.DB) *PlayersRepository {
	return &PlayersRepository{DB: db}
}

func (r *PlayersRepository) InsertPlayers(players []*model.Player) error {
	return r.DB.Create(players).Error
}

func (r *PlayersRepository) GetGroupedPlayersByTeamID(teamID uint, page, pageSize int) ([]*model.GroupedPlayer, error) {

	var results []*model.GroupedPlayer

	// Subquery to get distinct creation dates for the given team ID, ordered and paginated
	subquery := r.DB.Model(&model.Player{}).
		Select("DISTINCT created_at").
		Where(&model.Player{TeamID: teamID}).
		Order("created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize)

		// Get players for the paginated creation dates
	var players []model.Player
	err := r.DB.Where("team_id = ? AND created_at IN (?)", teamID, subquery).
		Order("created_at DESC").
		Find(&players).Error
	if err != nil {
		return nil, err
	}

	// Map to store players grouped by creation date
	groupedPlayersMap := make(map[time.Time][]model.Player)

	for _, player := range players {
		creationDate := player.CreatedAt
		groupedPlayersMap[creationDate] = append(groupedPlayersMap[creationDate], player)
	}

	// Populate results with grouped players

	for key := range groupedPlayersMap {
		results = append(results, &model.GroupedPlayer{
			CreationDate: key,
			TeamID:       teamID,
			Players:      groupedPlayersMap[key],
		})
	}

	return results, nil
}
