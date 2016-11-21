// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"

type tQuest struct{}

var Quest tQuest

func (_ tQuest) Add(
	title string,
	experience int,
	currency int,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "title", title)
	revel.Unbind(args, "experience", experience)
	revel.Unbind(args, "currency", currency)
	return revel.MainRouter.Reverse("Quest.Add", args).Url
}

type tApp struct{}

var App tApp

func (_ tApp) Index() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Quests() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("App.Quests", args).Url
}

func (_ tApp) Auth(
	code string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.Auth", args).Url
}

func (_ tApp) Logout() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("App.Logout", args).Url
}

type tTestRunner struct{}

var TestRunner tTestRunner

func (_ tTestRunner) Index() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
	suite string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
	suite string,
	test string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List() string {
	args := make(map[string]string)

	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}

type tStatic struct{}

var Static tStatic

func (_ tStatic) Serve(
	prefix string,
	filepath string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
	moduleName string,
	prefix string,
	filepath string,
) string {
	args := make(map[string]string)

	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}
