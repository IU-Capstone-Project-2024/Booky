package validator

import (
	models2 "booky-back/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

func (v *Validator) ValidateCourse(c *models2.Course) (ValidationErrors, error) {
	// Title
	v.Check(len(c.Title) != 0, "title", "must be provided")

	// Tracks
	v.Check(len(c.Tracks) != 0, "tracks", "must be provided")
	v.Check(len(lo.Uniq(c.Tracks)) == len(c.Tracks), "tracks", "must have unique values")

	// Semester
	v.Check(c.Semester != 0, "semester", "must be provided")

	// Year
	v.Check(c.Year != 0, "year", "must be provided")
	v.Check(c.Year > 2000, "year", "must be greater than 2000")
	v.Check(c.Year < 2100, "year", "must be less than 2100")

	return v.Errors, nil
}

func (v *Validator) ValidateNote(c *models2.Note) (ValidationErrors, error) {
	// Title
	v.Check(len(c.Title) != 0, "title", "must be provided")

	// Body
	v.Check(len(c.Body) != 0, "body", "must be provided")

	// CourseID
	v.Check(len(c.CourseID) != 0, "course_id", "must be provided")

	// Publisher
	v.Check(len(c.Publisher.ID) != 0, "publisher_id", "must be provided")

	return v.Errors, nil
}

func (v *Validator) ValidateFile(f *models2.File) (ValidationErrors, error) {
	// Filename
	v.Check(len(f.Filename) != 0, "filename", "must be provided")

	// Content
	v.Check(len(f.Content) != 0, "content", "must be provided")

	// CourseID
	v.Check(len(f.CourseID) != 0, "course_id", "must be provided")

	// Publisher
	v.Check(len(f.Publisher.ID) != 0, "publisher_id", "must be provided")

	return v.Errors, nil
}

func (v *Validator) ValidateID(id string) (ValidationErrors, error) {
	_, err := uuid.Parse(id)
	v.Check(err == nil, "id", "must be a valid UUID")

	return v.Errors, nil
}
