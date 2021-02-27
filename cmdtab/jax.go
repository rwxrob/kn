package cmdtab

import (
	"fmt"

	"github.com/rwxrob/cmdtab"
)

const jax = `
<script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
<script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"></script>
`

func init() {
	x := cmdtab.New("jax")
	x.Summary = `outpus the Mathjax HTML snippet`

	x.Method = func(args []string) (err error) {
		fmt.Println(jax)
		return nil
	}
}
