package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	// "github.com/teerasaknrt/go-unit-testing/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Repository represent the todo repository contranct
// type Repository interface {
// 	Create(ctx context.Context, data *models.ToDo) (int, error)
// 	Read(ctx context.Context, id int) (*models.ToDo, error)
// }

type todoRepository struct {
	DB *sql.DB
}

// NewToDoRepository will create an object that represent the todo.Repository interface
func NewToDoRepository(db *sql.DB) Repository {
	createTable(db)
	return &todoRepository{
		DB: db,
	}
}

func createTable(db *sql.DB) (err error) {

	const usersQry = `CREATE TABLE IF NOT EXISTS todo (
		id serial PRIMARY KEY,
		title varchar(200) NOT NULL,
		description text NOT NULL,
		reminder timestamp NOT NULL
	)`

	if _, err = db.Exec(usersQry); err != nil {
		logrus.Fatalf("cannot create todo table : %s", err.Error())
		return status.Error(codes.Unknown, "failed to connect to database - --> "+err.Error())
	}

	return nil
}

func (r *todoRepository) Create(ctx context.Context, t *models.ToDo) (id int, err error) {

	query := `
	INSERT INTO todo(
		title,
		description,
		reminder
	) VALUES(
		$1, $2, $3
	) RETURNING
		id;
	`
	err = r.DB.QueryRowContext(ctx, query,
		t.Title,
		t.Description,
		t.Reminder,
	).Scan(&id)

	if err != nil {
		logrus.Errorf("cannot create a new todo: %s", err)
		return
	}
	logrus.Debug("id : %d", id)

	return
}

func (r *todoRepository) Read(ctx context.Context, id int) (*models.ToDo, error) {
	task := &models.ToDo{}

	query := `
		SELECT
			id,
			title,
			description,
			reminder
		FROM todo
		WHERE id = $1
	`
	err := r.DB.QueryRowContext(ctx, query,
		id,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Reminder,
	)

	if err != nil {
		logrus.Errorf("cannot get task by id: %s", err.Error())
		return nil, err
	}
	return task, nil
}
