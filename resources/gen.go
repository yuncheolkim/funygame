package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"squrefight/utils"
	"strconv"
	"strings"
)

func main() {

	r, e := os.Open(`/Users/jinyunzhe/go/src/squrefight/resources/proto/message.proto`)
	if e != nil {
		panic(e)
	}
	reader := bufio.NewReader(r)
	var lines []string
	start := false
	for {

		line, _, e := reader.ReadLine()

		if len(line) > 0 {
			l := string(line)

			if start && strings.Contains(l, "}") {
				start = false
			}
			if start {
				lines = append(lines, l)
			}

			if !start && strings.Contains(l, "oneof") {
				start = true
			}

			fmt.Println(l)
		}

		if e != nil {
			break
		}
	}

	var reqLines []Line
	for _, v := range lines {
		line := ParseLine(v)
		if strings.HasSuffix(line.Type, "Request") {
			reqLines = append(reqLines, line)
		}
	}

	if len(reqLines) > 0 {
		GenTemplate(reqLines)
	}
}

type Line struct {
	Type string
	Name string
	Id   int
}

func ParseLine(line string) Line {
	line = strings.TrimSpace(line)
	r := regexp.MustCompile(`\s|\t`)
	split := r.Split(line, -1)

	i, _ := strconv.Atoi(strings.TrimRight(split[3], ";"))

	return Line{Type: split[0], Name: utils.Capitalize(split[1]), Id: i}
}

func WriteFile() {

}

func GenTemplate(l []Line) {
	t := template.New("gen")
	b, _ := ioutil.ReadFile("resources/gen.tmpl")

	t, e := t.Parse(string(b))
	if e != nil {
		panic(e)
	}

	file, _ := os.Create("pbmsg/msg_gen.go")

	t.Execute(file, l)

}
