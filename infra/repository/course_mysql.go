package repository

import (
	"database/sql"

	"github.com/matheus-dr/go-test/entity"
)

type CourseMysqlRepository struct {
	Db *sql.DB
}

// Implementation of Course Repository interface
func (c CourseMysqlRepository) Insert(course entity.Course) error {
	stmt, err := c.Db.Prepare("INSERT INTO courses(id_course, name, description, status) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		course.Id,
		course.Name,
		course.Description,
		course.Status,
	)
	if err != nil {
		return err
	}

	return nil
}
