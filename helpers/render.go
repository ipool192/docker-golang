package helpers

import (
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var render *renderer.Render

func init() {
	options := renderer.Options{
		ParseGlobPattern: "./templates/*.html",
	}
	render = renderer.New(options)
}

// RenderHTML - Rendering html
func RenderHTML(res http.ResponseWriter, httpStatus int, template string) {
	render.HTML(res, httpStatus, template, nil)
}
