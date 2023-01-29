package service

import (
	"test-api/entity"
	"test-api/model"
)

type TodoService interface {
	Save(entity.Todo) entity.Todo
	GetAll() []entity.Todo
	Delete(int32) entity.Todo
	Update(entity.Todo) entity.Todo
}
type todoService struct {
}

func New() *todoService {
	return &todoService{}
}
func (t *todoService) Update(todo entity.Todo) entity.Todo {
	return model.Update(todo)
}
func (t *todoService) Delete(id int32) entity.Todo {
	return model.Delete(id)
}
func (t *todoService) Save(todo entity.Todo) entity.Todo {
	if err := model.Add(todo); err != nil {
		return entity.Todo{}
	} else {
		return todo
	}
}
func (t *todoService) GetAll() []entity.Todo {
	if result, err := model.GetAll(); err != nil {
		return []entity.Todo{}
	} else {
		return result
	}
}
