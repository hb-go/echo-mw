// Base on github.com/robvdl/pongo2gin
package pongo2echo

import (
	"errors"
	"io"
	"path"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

// RenderOptions is used to configure the renderer.
type RenderOptions struct {
	TemplateDir string
	ContentType string
	Debug       bool
}

// Pongo2Render is a custom Gin template renderer using Pongo2.
type Pongo2Render struct {
	Options  *RenderOptions
	Template *pongo2.Template
	Context  pongo2.Context
}

// New creates a new Pongo2Render instance with custom Options.
func New(options RenderOptions) Pongo2Render {
	return Pongo2Render{
		Options: &options,
	}
}

// Default creates a Pongo2Render instance with default options.
func Default() Pongo2Render {
	return New(RenderOptions{
		TemplateDir: "templates",
		ContentType: "text/html; charset=utf-8",
		Debug:       true,
	})
}

func (p Pongo2Render) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var template *pongo2.Template
	filename := path.Join(p.Options.TemplateDir, name)

	// always read template files from disk if in debug mode, use cache otherwise.
	if p.Options.Debug {
		template = pongo2.Must(pongo2.FromFile(filename))
	} else {
		template = pongo2.Must(pongo2.FromCache(filename))
	}

	context, exist := data.(map[string]interface{})
	if !exist {
		panic(errors.New("Pongo2 context error!").Error())
	}

	err := template.ExecuteWriter(context, w)
	return err
}
