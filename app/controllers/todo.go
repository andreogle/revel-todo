package controllers

import "github.com/revel/revel"

type Todo struct {
	*revel.Controller
}

func (c Todo) Index() revel.Result {
	return c.Render()
}
