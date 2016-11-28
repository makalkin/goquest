package controllers

import "github.com/revel/revel"

type Circle struct {
	*revel.Controller
}

func (c Circle) Add() revel.Result  {
	return c.RenderText("")
}


func (c Circle) GetOne(id string) revel.Result  {
	return c.RenderText("")
}

func (c Circle) GetMany() revel.Result  {
	return c.RenderText("")
}

func (c Circle) Update() revel.Result  {
	return c.RenderText("")
}

func (c Circle) Delete() revel.Result  {
	return c.RenderText("")
}