package mappers

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/voiceline/backend/internal/domain/entities"
	"github.com/voiceline/backend/internal/interface/mappers"
)

var NewUserMapper = mappers.NewUserMapper

func TestUserMapper_ToDTO(t *testing.T) {
	mapper := NewUserMapper()

	t.Run("Convert valid user", func(t *testing.T) {
		user := &entities.User{
			ID:        uuid.New(),
			Email:     "test@example.com",
			Name:      "Test User",
			CreatedAt: time.Now(),
		}

		dto := mapper.ToDTO(user)

		assert.NotNil(t, dto)
		assert.Equal(t, user.ID.String(), dto.ID)
		assert.Equal(t, user.Email, dto.Email)
		assert.Equal(t, user.Name, dto.Name)
		assert.Equal(t, user.CreatedAt, dto.CreatedAt)
	})

	t.Run("Convert nil user", func(t *testing.T) {
		dto := mapper.ToDTO(nil)
		assert.Nil(t, dto)
	})
}

func TestUserMapper_ToAuthResponseDTO(t *testing.T) {
	mapper := NewUserMapper()

	user := &entities.User{
		ID:        uuid.New(),
		Email:     "test@example.com",
		Name:      "Test User",
		CreatedAt: time.Now(),
	}
	token := "test-token"

	dto := mapper.ToAuthResponseDTO(token, user)

	assert.NotNil(t, dto)
	assert.Equal(t, token, dto.Token)
	assert.NotNil(t, dto.User)
	assert.Equal(t, user.ID.String(), dto.User.ID)
	assert.Equal(t, user.Email, dto.User.Email)
}
