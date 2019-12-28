package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hi", index)
	http.HandleFunc("/box", box)
	http.Handle("/favicon", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("received a request")
	tpl.ExecuteTemplate(w, "index.gohtml", "reza")
}

// type row struct {
// 	Cells []int
// }

// func (r row) stringer() string {
// 	var sb strings.Builder
// 	sb.WriteString("<tr>")
// 	for _, c := range r.Cells {
// 		sb.WriteString(fmt.Sprintf("<tc>%d</tc>", c))
// 	}
// 	sb.WriteString("<tr/>")
// 	return sb.String()
// }

type table struct {
	// Rows []row
	Rows [][]int
}

func (t table) String() string {
	var sb strings.Builder
	sb.WriteString("<table><tbody>\n")
	for _, r := range t.Rows {
		var sbInner strings.Builder
		sbInner.WriteString("\t<tr>\n\t\t")
		for _, c := range r {
			sbInner.WriteString(fmt.Sprintf("<td>%d</td>", c))
		}
		sbInner.WriteString("\n\t</tr>\n")
		sb.WriteString(sbInner.String())
	}
	sb.WriteString("</tbody></table>")
	return sb.String()
}

func NewTable(x, y int) table {
	values := [][]int{}
	for i := 0; i < y; i++ {
		values = append(values, make([]int, x))
	}
	return table{values}
}

func box(w http.ResponseWriter, r *http.Request) {
	// rows = r.URL.Query().Get("rows")
	// cols = r.URL.Query().Get("cols")
	table := NewTable(2, 4)
	log.Println("received a request for a box")
	log.Printf("\n%s\n", table)
	tpl.ExecuteTemplate(w, "box.gohtml", fmt.Sprintf("%s", table))
}
