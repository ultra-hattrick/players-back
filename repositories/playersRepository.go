package repositories

import "gorm.io/gorm"

type PlayersRepository struct {
	DB *gorm.DB
}

func NewPlayersRepository(db *gorm.DB) *PlayersRepository {
	return &PlayersRepository{DB: db}
}
