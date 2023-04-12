package repository

import (
	"CBI/configs"
	"CBI/internal/domain/entity"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type StudentsRepository interface {
	GetAll(ctx context.Context) ([]entity.Students, error)
	GetById(ctx context.Context, id int) (entity.Students, error)
	Create(ctx context.Context, newStudent entity.Students) error
	UpdateById(ctx context.Context, id int, newStudent entity.Students) error
	Delete(ctx context.Context, id int) error
}

type studentsRepository struct {
	DB     *sql.DB
	config configs.Config
}

func NewStudentsRepository(db *sql.DB, config configs.Config) StudentsRepository {
	return &studentsRepository{
		DB:     db,
		config: config,
	}
}

func (repo *studentsRepository) GetAll(ctx context.Context) ([]entity.Students, error) {
	query := "SELECT * FROM students"
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []entity.Students{}
	for rows.Next() {
		var s entity.Students
		err := rows.Scan(&s.ID, &s.Name, &s.Class, &s.Email, &s.Password)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

func (repo *studentsRepository) GetById(ctx context.Context, id int) (entity.Students, error) {
	query := "SELECT * FROM students WHERE id=$1"
	row := repo.DB.QueryRowContext(ctx, query, id)

	var s entity.Students
	err := row.Scan(&s.ID, &s.Name, &s.Class, &s.Email, &s.Password)
	if err != nil {
		return entity.Students{}, err
	}
	return s, nil
}

func (repo *studentsRepository) Create(ctx context.Context, newStudent entity.Students) error {
	query := "INSERT INTO students(name, class, email, password) VALUES($1, $2, $3, $4)"
	_, err := repo.DB.ExecContext(ctx, query, newStudent.Name, newStudent.Class, newStudent.Email, newStudent.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *studentsRepository) UpdateById(ctx context.Context, id int, newStudent entity.Students) error {
	query := "UPDATE students SET name=$1, class=$2, email=$3, password=$4 WHERE id=$5"
	_, err := repo.DB.ExecContext(ctx, query, newStudent.Name, newStudent.Class, newStudent.Email, newStudent.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *studentsRepository) Delete(ctx context.Context, id int) error {
	_, err := repo.DB.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
