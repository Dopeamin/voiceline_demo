package mappers

import (
	"github.com/voiceline/backend/internal/domain/entities"
	"github.com/voiceline/backend/internal/interface/dto"
)

// UserMapper handles mapping between User entity and DTOs
type UserMapper struct{}

// NewUserMapper creates a new UserMapper
func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

// ToDTO converts a User entity to a UserDTO
func (m *UserMapper) ToDTO(user *entities.User) *dto.UserDTO {
	if user == nil {
		return nil
	}

	return &dto.UserDTO{
		ID:        user.ID.String(),
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}

// ToAuthResponseDTO converts authentication output to AuthResponseDTO
func (m *UserMapper) ToAuthResponseDTO(token string, user *entities.User) *dto.AuthResponseDTO {
	return &dto.AuthResponseDTO{
		Token: token,
		User:  m.ToDTO(user),
	}
}

