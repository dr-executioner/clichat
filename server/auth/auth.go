package auth

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   string
	Username string
	Password string
}
type AuthManager struct {
	db *sql.DB
}

func NewAuthManager(dsn string) (*AuthManager, error) {

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping Db: %w", err)
	}

	return &AuthManager{db: db}, nil
}

func (am *AuthManager) Register(username, password string) error {
	var exists bool

	err := am.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)

	if err != nil {
		return fmt.Errorf("Failed to check existing user: %w", err)
	}

	if exists {
		return fmt.Errorf("Username '%s' already exists", username)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("Failed to hash password: %w", err)
	}

	_, err = am.db.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", username, string(hashedPassword))

	if err != nil {
		return fmt.Errorf("Failed to insert user : %w", err)
	}
	return nil
}

func (am *AuthManager) Login(username, password string) error {
	var hashedPassword string

	err := am.db.QueryRow("SELECT password_hash FROM users WHERE username =? ", username).Scan(&hashedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("User does not exist")
		}
		return fmt.Errorf("Failed to fetch user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("Invalid Password")

	}
	return nil
}
