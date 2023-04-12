package usecase

import (
	"context"

	"CBI/internal/domain/entity"
	"CBI/internal/domain/repository"
)

type StudentsUseCase interface {
	FindAll(ctx context.Context) ([]entity.Students, error)
	FindById(ctx context.Context, id int) (entity.Students, error)
	Register(ctx context.Context, newStudent entity.Students) error
	EditById(ctx context.Context, id int, newStudent entity.Students) error
	Unregister(ctx context.Context, id int) error
}

type studentsUseCase struct {
	StudentsRepository repository.StudentsRepository
}

func NewStudentsUseCase(studentsRepository repository.StudentsRepository) StudentsUseCase {
	return &studentsUseCase{
		StudentsRepository: studentsRepository,
	}
}
func (s *studentsUseCase) FindAll(ctx context.Context) ([]entity.Students, error) {

	students, err := s.StudentsRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (uc *studentsUseCase) FindById(ctx context.Context, id int) (entity.Students, error) {
	student, err := uc.StudentsRepository.GetById(ctx, id)
	if err != nil {
		return entity.Students{}, err
	}
	return student, nil
}

func (uc *studentsUseCase) Register(ctx context.Context, newStudent entity.Students) error {
	err := uc.StudentsRepository.Create(ctx, newStudent)
	if err != nil {
		return err
	}
	return nil
}

func (uc *studentsUseCase) EditById(ctx context.Context, id int, newStudent entity.Students) error {
	err := uc.StudentsRepository.UpdateById(ctx, id, newStudent)
	if err != nil {
		return err
	}
	return nil
}

func (uc *studentsUseCase) Unregister(ctx context.Context, id int) error {
	err := uc.StudentsRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
