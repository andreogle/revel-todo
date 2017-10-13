package controllers

import (
	"strconv"
	"todo/app/models"
	"todo/app/routes"

	"github.com/revel/revel"
)

type Todo struct {
	*revel.Controller
}

var todoList = []models.Todo{}

func (c Todo) CheckLogin() revel.Result {
	if _, ok := c.Session["username"]; ok {
		return nil
	}
	c.Flash.Error("Please log in first!")
	return c.Redirect(routes.App.Index())
}

func (c Todo) Index() revel.Result {
	return c.Render(todoList)
}

func (c Todo) AddTodo(todo models.Todo) revel.Result {
	todo.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Todo.Index)
	}

	todo.Number = len(todoList)
	todo.Number++
	todo.Creator = c.Session["username"]
	todoList = append(todoList, todo)

	return c.Redirect(Todo.Index)
}

func (c Todo) CompleteTodo() revel.Result {
	c.Request.ParseForm()
	for i, _ := range c.Request.Form {
		t, _ := strconv.Atoi(i)
		t--
		todoList[t].Completed = true
	}
	return c.Redirect(Todo.Index)
}

func (c Todo) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}
