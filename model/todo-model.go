package model

import (
	"database/sql"
	"fmt"
	"test-api/entity"

	_ "github.com/go-sql-driver/mysql"
)

func Add(todo entity.Todo) error {
	conn, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec("INSERT INTO todo.todo(title, completed) values(?,?)", todo.Title, todo.Completed)
	if err != nil {
		return err
	}
	return nil
}
func Update(todo entity.Todo) entity.Todo {
	conn, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		return entity.Todo{}
	}
	defer conn.Close()
	// mystr := "UPDATE todo.todo SET title=\"" + todo.Title + "\", completed=" + strconv.FormatBool(todo.Completed) + " WHERE id=" + strconv.Itoa(int(todo.Id))
	// _, err = conn.Query(mystr)
	_, err = conn.Query("UPDATE todo.todo SET title=?, completed=? where id=?", todo.Title, todo.Completed, todo.Id)

	if err != nil {
		fmt.Println(err)
		return entity.Todo{}
	}
	return todo
}
func GetAll() ([]entity.Todo, error) {
	conn, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rows, err := conn.Query("SELECT * FROM todo.todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []entity.Todo{}
	for rows.Next() {
		t := entity.Todo{}
		err := rows.Scan(&t.Id, &t.Title, &t.Completed)
		if err != nil {
			continue
			// return entity.Todo{}, nil
		}
		todos = append(todos, t)
	}
	return todos, nil
}
func Delete(id int32) entity.Todo {
	conn, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		return entity.Todo{}
	}
	todo := findById(id)
	defer conn.Close()
	result, err := conn.Query("DELETE FROM todo.todo WHERE id=?", id)
	if err != nil {
		return entity.Todo{}
	}
	defer result.Close()
	return todo
}

func findById(id int32) entity.Todo {
	conn, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		return entity.Todo{}
	}
	defer conn.Close()
	rows, err := conn.Query("SELECT * FROM todo.todo WHERE id=?", id)
	if err != nil {
		return entity.Todo{}
	}
	defer rows.Close()
	todos := []entity.Todo{}
	for rows.Next() {
		t := entity.Todo{}
		err := rows.Scan(&t.Id, &t.Title, &t.Completed)
		if err != nil {
			continue
			// return entity.Todo{}, nil
		}
		todos = append(todos, t)
	}
	return todos[0]
}
