package controllers

import "github.com/revel/revel"

type Quest struct {
	*revel.Controller
}


func (c Quest) Add(title string, experience int, currency int) revel.Result {
	//c.Validation.MinSize(title, 3).Message("Cannot understand shit")
	//
	//if c.Validation.HasErrors() {
	//	errorData := make(map[string]interface{})
	//	errorData["errors"] = c.Validation.Errors
	//	c.Validation.Keep()
	//	c.FlashParams()
	//	c.Response.Status = 400
	//	return c.RenderJson(errorData)
	//}
	data := make(map[string]interface{})
	data["success"] = true
	return c.RenderJson(data)
}

