package dao

import (
	"database/sql"
	"fmt"

	"github.com/rcliao/e2etest"
)

// Service defines the persistance layer
type Service struct {
	db *sql.DB
}

// New is constructor for creating MySQL service
func New(db *sql.DB) *Service {
	return &Service{db}
}

// UpdateStatus update status to MySQL DB
func (s *Service) UpdateStatus(status e2etest.Status) error {
	return nil
}

// CreateStatus update an existing status to MySQL DB
func (s *Service) CreateStatus(status e2etest.Status) error {
	return nil
}

// GetStatus create a new status to MySQL DB
func (s *Service) GetStatus(status e2etest.Status) e2etest.Status {
	return e2etest.Status{}
}

// StoreToken stores the access token to DB so that it can be used for API calls
func (s *Service) StoreToken(token string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO token (`access_token`) VALUES (?)")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM token")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(token)
	err = tx.Commit()
	return err
}

// GetToken returns the token stored from #StoreToken earlier
func (s *Service) GetToken() string {
	var token string
	err := s.db.QueryRow("SELECT access_token FROM token").Scan(&token)
	if err != nil {
		fmt.Println("Failed to scan token from DB because", err)
		return ""
	}
	return token
}
