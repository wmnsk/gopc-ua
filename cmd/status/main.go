// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/liuxgo/opcua/cmd/service/goname"
)

func main() {
	log.SetFlags(0)

	in := flag.String("in", "schema/StatusCode.csv", "path to StatusCodes.csv")
	out := flag.String("out", "ua/status_gen.go", "path to generated file")
	flag.Parse()

	if *in == "" {
		log.Fatal("-in is required")
	}
	if *out == "" {
		log.Fatal("-out is required")
	}

	f, err := os.Open(*in)
	if err != nil {
		log.Fatalf("Error reading %s: %v", *in, err)
	}
	defer f.Close()

	// rows have the format of "name,id,description"
	// but the description is not quoted and can contain commas
	// Therefore, do not use the csv.Reader here.

	var rows [][]string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		rows = append(rows, strings.SplitN(sc.Text(), ",", 3))
	}
	if sc.Err() != nil {
		log.Fatalf("Error parsing csv: %v", err)
	}

	// goify the name and prefix with Status
	for i := range rows {
		rows[i][0] = "Status" + goname.Format(rows[i][0])
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, rows); err != nil {
		log.Fatalf("Error generating code: %v", err)
	}

	bfmt, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatalf("Error formatting source: %v", err)
	}

	if err := ioutil.WriteFile(*out, bfmt, 0644); err != nil {
		log.Fatalf("Error writing %s: %v", *out, err)
	}
	log.Printf("Wrote %s", *out)
}

var tmpl = template.Must(template.New("").Parse(`
// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Code generated by cmd/status. DO NOT EDIT!

package ua

import "fmt"

// StatusCode is an error type for a status code.
type StatusCode uint32

func (n StatusCode) Error() string {
	if d, ok := StatusCodes[n]; ok {
		return fmt.Sprintf("%s %s (0x%X)", d.Text, d.Name, uint32(n))
	}
	return fmt.Sprintf("0x%X", uint32(n))
}

var (
	StatusOK StatusCode = 0x0
    StatusUncertain StatusCode = 0x40000000
    StatusBad StatusCode = 0x80000000

	{{range .}}{{index . 0}} StatusCode = {{index . 1}}
	{{end}}
)

type StatusCodeDesc struct {
	Name string
	Text string
}

// StatusCodes maps status codes to the status code error types.
var StatusCodes = map[StatusCode]StatusCodeDesc{
	StatusOK: StatusCodeDesc{Name: "OK", Text: ""},
	StatusUncertain: StatusCodeDesc{Name: "Uncertain", Text: ""},
	StatusBad: StatusCodeDesc{Name: "Bad", Text: ""},
	{{range .}}{{index . 0}}: StatusCodeDesc{ Name: "{{index . 0}}", Text: "{{index . 2}}" },
	{{end}}
}
`))
