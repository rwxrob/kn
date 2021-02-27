package cmdtab

import (
	"bytes"
	"fmt"
	"io"
	"os"

	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/rwxrob/cmdtab"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func ToHTML(md []byte, buf io.Writer) error {
	// TODO detect math in the input
	/*
			mjax := `
		`
			fmt.Fprintln(buf, mjax)
	*/

	gold := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithExtensions(mathjax.MathJax),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
	if err := gold.Convert(md, buf); err != nil {
		return err
	}
	return nil
}
func init() {
	x := cmdtab.New("html")
	x.Summary = `converts KEG Mark to HTML`
	x.Usage = `[<file>]`

	x.Description = `
		Converts KEG Mark (Markdown) into HTML. Reads from stdin if no *file*
		argument passed.`

	x.Method = func(args []string) (err error) {
		var mark []byte
		var buf bytes.Buffer

		switch {
		case len(args) >= 1:
			mark, err = os.ReadFile(args[0])
		default:
			mark, err = io.ReadAll(os.Stdin)
		}
		if err != nil {
			return err
		}

		err = ToHTML(mark, &buf)
		if err != nil {
			return err
		}
		fmt.Print(buf.String())

		return nil
	}
}
