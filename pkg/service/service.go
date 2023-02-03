package service

import (
	"github.com/AlibekDalgat/todo-app"
	"github.com/AlibekDalgat/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, id int) (list todo.TodoList, err error)
	UpdateById(userId, id int, input todo.UpdateListInput) error
	DeleteById(userId, id int) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, listId, id int) (list todo.TodoItem, err error)
	UpdateById(userId, listId, id int, input todo.UpdateItemInput) error
	DeleteById(userId, listId, id int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
