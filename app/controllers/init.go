package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod(Todo.CheckLogin, revel.BEFORE)
}
