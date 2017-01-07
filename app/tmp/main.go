// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	_ "github.com/makalkin/goquest/app"
	controllers "github.com/makalkin/goquest/app/api/v1/app/controllers"
	controllers0 "github.com/makalkin/goquest/app/controllers"
	tests "github.com/makalkin/goquest/tests"
	controllers1 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers2 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
	"reflect"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")

	revel.RegisterController((*controllers.Circle)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name:           "Add",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "GetOne",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "GetMany",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "Update",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "Delete",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Join",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers.Quest)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "title", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "experience", Type: reflect.TypeOf((*int)(nil))},
					&revel.MethodArg{Name: "currency", Type: reflect.TypeOf((*int)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GetOne",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "GetMany",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "GetMe",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "token", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers0.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					20: []string{},
				},
			},
			&revel.MethodType{
				Name: "Quests",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					24: []string{},
				},
			},
			&revel.MethodType{
				Name: "Auth",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "code", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name:           "Logout",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers1.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.RegisterController((*controllers2.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{},
				RenderArgNames: map[int][]string{
					72: []string{
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil))},
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil))},
				},
				RenderArgNames: map[int][]string{
					125: []string{},
				},
			},
			&revel.MethodType{
				Name:           "List",
				Args:           []*revel.MethodArg{},
				RenderArgNames: map[int][]string{},
			},
		})

	revel.DefaultValidationKeys = map[string]map[int]string{}
	testing.TestSuites = []interface{}{
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
