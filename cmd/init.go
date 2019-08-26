
package main

import (
	"gitlab.com/fusemedia/wgen/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"text/template"
	"unicode"
)

func main() {
	_, callerFile, _, _ := runtime.Caller(0)
	currDir := filepath.Dir(callerFile)

	filename:= filepath.Join(currDir, "workflow.gotpl")

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"type": func(obj interface{}) string {
			return reflect.TypeOf(obj).Name()
		},
		"firstLetterToLower": func(str string) string {
			a := []rune(str)
			a[0] = unicode.ToLower(a[0])
			return string(a)
		},
		"title": func(str string) string {
			return strings.Title(str)
		},
	}

	tmpl := template.New("").Funcs(funcMap)

	b ,err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	tmpl, err = tmpl.New(filepath.Base(filename)).Parse(string(b))
	if err != nil {
		panic(err)
	}

	dir, err := os.Getwd()

	if err != nil { panic(err) }

	outFile, err := os.Create(filepath.Join(dir, "workflow.go"))

	if err != nil { panic(err) }

	err = tmpl.Execute(outFile,  config.NewConfig("flow.json"))

	if err != nil { panic(err) }
}
