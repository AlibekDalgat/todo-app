package service

import (
	"github.com/AlibekDalgat/todo-app"
	"github.com/AlibekDalgat/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo}
}

func (listService *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return listService.repo.Create(userId, list)
}
