package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// Student declaring a struct
type Student struct {

	// declaring fields which are
	// exported and accessible
	// outside of package as they
	// begin with a capital letter
	Name  string
	Marks int64
}

// main function
func main() {

	// defining an object of struct
	std1 := Student{"Bob", 94}

	// "New" creates a new template
	// with name passed as argument
	tpl := template.New("template_1")

	// "Parse" parses a string into a template
	tpl, _ = tpl.Parse("Hello {{.Name}}, your marks are {{.Marks}}%!")

	// standard output to print merged data
	builder := new(strings.Builder)
	err := tpl.Execute(builder, std1)
	// if there is no error,
	// prints the output
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(builder.String())
	ExampleTemplate()
}

func ExampleTemplate() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

	// Output:
	// Dear Aunt Mildred,
	//
	// It was a pleasure to see you at the wedding.
	// Thank you for the lovely bone china tea set.
	//
	// Best wishes,
	// Josie
	//
	// Dear Uncle John,
	//
	// It is a shame you couldn't make it to the wedding.
	// Thank you for the lovely moleskin pants.
	//
	// Best wishes,
	// Josie
	//
	// Dear Cousin Rodney,
	//
	// It is a shame you couldn't make it to the wedding.
	//
	// Best wishes,
	// Josie
}
