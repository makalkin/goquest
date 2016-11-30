package tests

import "github.com/revel/revel/testing"

type APIV1Test struct {
	testing.TestSuite
}

func (t *APIV1Test) Before() {
	println("Set up")
}

func (t *APIV1Test) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *APIV1Test) After() {
	println("Tear down")
}