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
	var creationDates []time.Time
	var results []*model.GroupedPlayer

	// Subquery to get distinct creation dates for the given team ID, ordered and paginated
	subquery := r.DB.Model(&model.Player{}).
		Select("DISTINCT created_at").
		Where("team_id = ?", teamID).
		Order("created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize)

	// Get the distinct creation dates
	err := r.DB.Table("(?) as sub", subquery).
		Select("sub.created_at").
		Scan(&creationDates).Error
	if err != nil {
		return nil, err
	}

	if len(creationDates) == 0 {
		return results, nil
	}

	// Get players for the paginated creation dates
	var players []model.Player
	err = r.DB.Where("team_id = ? AND created_at IN ?", teamID, creationDates).
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
	currentTime := time.Now()
	for _, creationDate := range creationDates {
		week := calculateWeekDifference(currentTime, creationDate)
		results = append(results, &model.GroupedPlayer{
			CreationDate: creationDate,
			TeamID:       teamID,
			Week:         week,
			Players:      groupedPlayersMap[creationDate],
		})
	}

	return results, nil
}

func calculateWeekDifference(current, creation time.Time) uint16 {
	currentYear, currentWeek := current.ISOWeek()
	creationYear, creationWeek := creation.ISOWeek()
	yearDiff := currentYear - creationYear
	weekDiff := int(currentWeek - creationWeek)
	return uint16(yearDiff*52 + weekDiff)
}
