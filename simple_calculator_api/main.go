package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "calculator", "Calc-O-Matic 3000")
	})

	e.GET("/calcResult", func(c echo.Context) error {
		operation := c.FormValue("operation")
		op1, err := strconv.ParseFloat(c.FormValue("op1"), 64)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "operand 1 is not a number",
			})
		}

		op2, err := strconv.ParseFloat(c.FormValue("op2"), 64)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "operand 2 is not a number",
			})
		}

		var result float64

		switch operation {
		case "+":
			result = op1 + op2
		case "-":
			result = op1 - op2
		case "*":
			result = op1 * op2
		case "/":
			if op2 == 0 {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "cannot perform division by zero",
				})
			}

			result = op1 / op2

		default:
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid operation",
			})
		}

		return c.JSON(http.StatusOK, map[string]float64{
			"result": result,
		})
	})

	e.Logger.Fatal(e.Start(":3306"))
}
