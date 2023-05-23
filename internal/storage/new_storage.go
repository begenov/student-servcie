package storage

import (
	"context"
	"database/sql"

	"github.com/begenov/student-servcie/internal/models"
	"github.com/begenov/student-servcie/internal/storage/postgresql"
)

type Students interface {
	CreateStudent(ctx context.Context, student models.Student) error
	GetStudentByID(ctx context.Context, id int) (*models.Student, error)
	Update(ctx context.Context, student models.Student) error
	Delete(ctx context.Context, studentID int) error
	ByIDCourses(ctx context.Context, studentID int) ([]models.Student, error)
}

type Admins interface {
	CreateAdmin(ctx context.Context, admin models.Admin) error
	AdminByEmail(ctx context.Context, email string) (*models.Admin, error)
	RefreshTokenUpdate(ctx context.Context, session models.Session, adminID int) error
}

type Storage struct {
	Students Students
	Admins   Admins
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Students: postgresql.NewStudentsStorage(db),
		Admins:   postgresql.NewAdminsStorage(db),
	}
}