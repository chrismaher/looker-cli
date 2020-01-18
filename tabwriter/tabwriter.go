package tabwriter

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/chrismaher/looker-cli/client"
)

func TabPrinter(inputs []client.Response, keys []string) {
	var str string
	var vs []interface{}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)

	for _, val := range inputs {
		str = ""
		vs = nil
		for _, v := range keys {
			str += "%v\t"
			vs = append(vs, val[v])
		}
		str += "\n"
		fmt.Fprintf(w, str, vs...)
	}
	w.Flush()
}
