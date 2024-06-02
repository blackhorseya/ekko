package templates

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

//go:embed *.html
var f embed.FS

var templates = template.Must(template.ParseFS(f, "*.html"))

// SetHTMLTemplate set html template.
func SetHTMLTemplate(r *gin.Engine) {
	r.SetHTMLTemplate(templates)
}
