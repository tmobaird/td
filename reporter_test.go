package main

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestReportTodos(t *testing.T) {
	t.Run("turns todos into string output", func(t *testing.T) {
		todos := []Todo{
			{Name: "Hello", CreatedAt: "dummy", Id: uuid.New()},
			{Name: "World", CreatedAt: "dummy", Id: uuid.New()},
		}

		got := ReportTodos(todos, false)
		want := "- Hello\n- World\n"

		assertCorrectMessage(t, got, want)
	})

	t.Run("shows todo id when verbose", func(t *testing.T) {
		todoId := uuid.New()
		todos := []Todo{
			{Name: "Hello", CreatedAt: "dummy", Id: todoId},
		}

		got := ReportTodos(todos, true)
		want := "- Hello (" + todos[0].Id.String() + ")\n"

		assertCorrectMessage(t, got, want)
	})
}

func TestReportError(t *testing.T) {
	t.Run("turns error into string output", func(t *testing.T) {
		err := errors.New("Failed to write to file")

		got := ReportError(err, "add")
		want := "ERROR - Failed to execute add\n- Failed to write to file\n"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}