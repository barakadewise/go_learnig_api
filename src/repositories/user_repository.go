package repositories

import (
	"database/sql"
	"fmt"
	"sample-api/src/entities"
)

type UserRepositoryInterface interface {
	SaveUser(user entities.User) (entities.User, error)
	GetUserById(id int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(id int, user entities.User) error
	DeleteUser(id int) error
}

type UserRepository struct {
	DB *sql.DB
}

// Constructor
func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

// SaveUser inserts a new user into the database
func (r *UserRepository) SaveUser(user entities.User) (entities.User, error) {

	var existingID int
	err := r.DB.QueryRow("SELECT id FROM users WHERE email = ?", user.Email).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		return entities.User{}, err
	}
	if existingID != 0 {
		return entities.User{}, fmt.Errorf("email already exists")
	}

	query := `INSERT INTO users(email, name,address,age, dob,company_id) VALUES (?,?,?,?,?,?)`
	result, err := r.DB.Exec(query, user.Email, user.Name, user.Address, user.Age, user.DOB, user.CompanyID)
	if err != nil {
		return entities.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.User{}, err
	}
	user.ID = int(id)
	return user, nil
}

// GetUserById fetches a user by ID
func (r *UserRepository) GetUserById(id int) (entities.User, error) {
	var user entities.User
	query := `SELECT id,name,email,dob,age,address,company_id FROM users WHERE id=?`
	row := r.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.DOB, &user.Age, &user.Address, &user.CompanyID)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.User{}, fmt.Errorf("user not found")
		}
		return entities.User{}, err
	}
	return user, nil
}

// GetAllUsers fetches all users
func (r *UserRepository) GetAllUsers() ([]entities.User, error) {
	query := `SELECT id,name,email,dob,age,address,company_id FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var u entities.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.DOB, &u.Age, &u.Address, &u.CompanyID)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// UpdateUser updates an existing user
func (r *UserRepository) UpdateUser(id int, user entities.User) error {
	query := `UPDATE users SET name=?, email=?, dob=?, age=?, address=?, company_id=? WHERE id=?`
	result, err := r.DB.Exec(query, user.Name, user.Email, user.DOB, user.Age, user.Address, user.CompanyID, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

// DeleteUser removes a user by ID
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=?`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
