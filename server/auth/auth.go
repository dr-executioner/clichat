package auth

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    string
	Username  string
	Password  string
	CreatedAt time.Time
	IsOnline  bool
}
type AuthManager struct {
	users map[string]User
	mu    sync.Mutex
}

func NewAuthManager() *AuthManager {
	return &AuthManager{
		users: make(map[string]User),
	}
}

var users = make(map[string]User)

func (am *AuthManager) Register(username, password string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	if _, exists := am.users[username]; exists {
		return errors.New("Username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	am.users[username] = User{
		UserId:   uuid.NewString(),
		Username: username,
		Password: string(hashed),
	}

	return nil
}

func (am *AuthManager) Login(username, password string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	user, exists := am.users[username]

	if !exists {
		return errors.New("User Does not exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("Wrong Password")
	}

	return nil
}
