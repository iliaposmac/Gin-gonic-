package service

import (
	"github.com/iliaposmac/todo-app"
	"github.com/iliaposmac/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo, listRepo}
}

func (s *TodoItemService) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)

	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAllItems(userId int, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAllItems(userId, listId)
}

func (s *TodoItemService) GetById(userId int, listId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, item todo.UpdateTodoItem) error {
	return s.repo.Update(userId, itemId, item)
}
