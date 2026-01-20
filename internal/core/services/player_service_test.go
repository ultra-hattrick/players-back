package services

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ultra-hattrick/players-back/internal/core/domain"
)

type MockPlayerRepository struct {
	mock.Mock
}

func (m *MockPlayerRepository) InsertPlayers(players []*domain.Player) error {
	args := m.Called(players)
	return args.Error(0)
}

func (m *MockPlayerRepository) GetGroupedPlayersByTeamID(teamID uint, page, pageSize int) ([]*domain.GroupedPlayer, error) {
	args := m.Called(teamID, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.GroupedPlayer), args.Error(1)
}

func TestGetGroupedPlayersByTeamID(t *testing.T) {
	mockRepo := new(MockPlayerRepository)
	service := NewPlayersService(mockRepo)

	t.Run("success", func(t *testing.T) {
		expectedPlayers := []*domain.GroupedPlayer{
			{
				TeamID:       1,
				CreationDate: time.Now(),
				Players:      []domain.Player{},
			},
		}
		mockRepo.On("GetGroupedPlayersByTeamID", uint(1), 1, 10).Return(expectedPlayers, nil)

		result, err := service.GetGroupedPlayersByTeamID(1, 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, expectedPlayers, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetGroupedPlayersByTeamID", uint(2), 1, 10).Return(nil, errors.New("db error"))

		result, err := service.GetGroupedPlayersByTeamID(2, 1, 10)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
