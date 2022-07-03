package usecase

import (
	"github.com/google/uuid"
	"github.com/matheus-dr/go-test/entity"
)

type CreateCourse struct {
	Repository entity.CourseRepository
}

func (c CreateCourse) Execute(input CreateCourseInputDTO) (CreateCourseOutputDTO, error) {
	course := entity.Course{}

	course.Id = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Status = input.Status

	err := c.Repository.Insert(course)
	if err != nil {
		return CreateCourseOutputDTO{}, err
	}

	output := CreateCourseOutputDTO{}

	output.Id = course.Id
	output.Name = course.Name
	output.Description = course.Description
	output.Status = input.Status

	return output, nil
}
