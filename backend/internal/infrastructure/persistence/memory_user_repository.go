package persistence

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/voiceline/backend/internal/domain/entities"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type MemoryUserRepository struct {
	users      map[uuid.UUID]*entities.User
	emailIndex map[string]uuid.UUID
	mu         sync.RWMutex
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:      make(map[uuid.UUID]*entities.User),
		emailIndex: make(map[string]uuid.UUID),
	}
}

func (r *MemoryUserRepository) Create(ctx context.Context, user *entities.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	r.emailIndex[user.Email] = user.ID
	return nil
}

func (r *MemoryUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (r *MemoryUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id, exists := r.emailIndex[email]
	if !exists {
		return nil, ErrUserNotFound
	}

	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (r *MemoryUserRepository) Update(ctx context.Context, user *entities.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return ErrUserNotFound
	}

	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return ErrUserNotFound
	}

	delete(r.emailIndex, user.Email)
	delete(r.users, id)
	return nil
}
