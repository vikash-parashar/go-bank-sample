package db

import (
	"database/sql"
	"errors"
	"go-server/models"
	"go-server/utils"
	"log"
)

func (db *DB) GetUserByEmailID(email string) (*models.User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, role
        FROM users
        WHERE email = $1
    `
	user := &models.User{}
	err := db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		log.Printf("Error fetching user by email: %v", err)
		return nil, err
	}
	return user, nil
}

func (db *DB) RegisterUser(user *models.User) error {
	query := `
        INSERT INTO users (first_name, last_name, phone, email, password, role)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	err := db.QueryRow(query, user.FirstName, user.LastName, user.Phone, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}
	return nil
}

func (db *DB) UpdateUserPassword(userID int, newPassword string) error {
	query := `
        UPDATE users
        SET password = $2
        WHERE id = $1
    `
	_, err := db.Exec(query, userID, newPassword)
	if err != nil {
		log.Printf("Error updating user password: %v", err)
		return err
	}
	return nil
}

// GetAllUsers retrieves all active user records.
func (db *DB) GetAllUsers() ([]*models.User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, role
        FROM users
    `
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
		if err != nil {
			log.Printf("Error scanning user rows: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v", err)
		return nil, err
	}
	return users, nil
}

// GetUserByResetToken retrieves a user by their reset token.
func (db *DB) GetUserByResetToken(resetToken string) (*models.User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, role
        FROM users
        WHERE reset_token = $1
    `
	user := &models.User{}
	err := db.QueryRow(query, resetToken).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found by reset token")
		}
		log.Printf("Error fetching user by reset token: %v", err)
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user's information in the database.
func (db *DB) UpdateUser(user *models.User) error {
	query := `
        UPDATE users
        SET first_name = $1, last_name = $2, email = $3, password = $4, role = $5
        WHERE id = $6
    `
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

// SetResetToken sets the reset token for a user in the database.
func (db *DB) SetResetToken(userID int, resetToken string) error {
	query := `
        UPDATE users
        SET reset_token = $1
        WHERE id = $2
    `
	_, err := db.Exec(query, resetToken, userID)
	if err != nil {
		log.Printf("Error setting reset token: %v", err)
		return err
	}
	return nil
}

// ClearResetToken clears the reset token for a user in the database.
func (db *DB) ClearResetToken(userID int) error {
	query := `
        UPDATE users
        SET reset_token = NULL
        WHERE id = $1
    `
	_, err := db.Exec(query, userID)
	if err != nil {
		log.Printf("Error clearing reset token: %v", err)
		return err
	}
	return nil
}

// VerifyResetToken verifies the reset token for a user.
func (db *DB) VerifyResetToken(resetToken string) (*models.User, error) {
	query := `
        SELECT id, email, reset_token
        FROM users
        WHERE reset_token = $1
    `
	user := &models.User{}
	err := db.QueryRow(query, resetToken).Scan(&user.ID, &user.Email, &user.ResetToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("reset token not found")
		}
		log.Printf("Error fetching user by reset token: %v", err)
		return nil, err
	}

	// Check if the reset token has expired (optional)
	if utils.IsTokenExpired(user.ResetTokenExpiry) {
		return nil, errors.New("reset token has expired")
	}

	return user, nil
}
