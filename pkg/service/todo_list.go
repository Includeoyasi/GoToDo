package service

import (
	"github.com/Includeoyasi/todo-app"
	"github.com/Includeoyasi/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (as *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return as.repo.Create(userId, list)
}

func (as *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return as.repo.GetAll(userId)
}

func (as *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return as.repo.GetById(userId, listId)
}

func (as *TodoListService) Delete(userId, listId int) error {
	return as.repo.Delete(userId, listId)
}

func (as *TodoListService) Update(userId, listId int, input todo.UpdateTodoListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return as.repo.Update(userId, listId, input)
}
