package utils

import (
	"github.com/revel/revel"
	"strconv"
	"fmt"
)

type APIError struct {
	Field string `json:"field,omitempty"`
	Msg   string `json:"msg"`
}

func (e APIError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s", e.Field, e.Msg)
	} else {
		return e.Msg
	}
}

func GetPaging(c *revel.Controller) (int, int) {
	var page, perPage int
	if _, ok := c.Params.Query["page"]; ok {
		page, _ = strconv.Atoi(c.Params.Query.Get("page"))
	} else {
		page = 1
	}
	if _, ok := c.Params.Query["perPage"]; ok {
		perPage, _ = strconv.Atoi(c.Params.Query.Get("perPage"))
	} else {
		perPage = 20
	}
	return page, perPage
}

func RenderJsonError(c *revel.Controller, code int, err error) revel.Result {
	c.Response.Status = code
	return c.RenderJson(&map[string]interface{}{"error": err})
}
