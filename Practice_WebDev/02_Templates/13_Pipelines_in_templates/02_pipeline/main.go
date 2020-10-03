package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl":  fdouble,
	"fsq":   fsquare,
	"fsqrt": fsquareroot,
}

func init() {
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func fdouble(x int) int {
	return x + x
}

func fsquare(x int) float64 {
	return math.Pow(float64(x), 2)
}

func fsquareroot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
