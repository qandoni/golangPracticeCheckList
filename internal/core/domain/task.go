package domain

import (
	core_errors "checklist/internal/core/errors"
	"fmt"
)

type Task struct {
	ID          int
	Title       string
	Description *string
	Completed   bool
}

func NewTask(
	id int,
	title string,
	description *string,
	completed bool,
) Task {
	return Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func NewTaskUnitialized(
	title string,
	description *string,
) Task {
	return NewTask(
		UninitializedID,
		title,
		description,
		false,
	)
}

func (t *Task) Validate() error {
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 50 {
		return fmt.Errorf(
			"invalid 'Title' len: %d: %w",
			titleLen,
			core_errors.ErrInvalidArgument,
		)
	}
	if t.Description != nil {
		descriptionLen := len([]rune(*t.Description))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf(
				"invalid 'Description' len: %d, %w",
				descriptionLen,
				core_errors.ErrInvalidArgument,
			)
		}
	}
	return nil
}

func (t *Task) SetCompleted(completed bool) {
	t.Completed = completed
}
