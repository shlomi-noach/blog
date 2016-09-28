package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	fileName := os.Args[1]
	b, _ := ioutil.ReadFile(fileName)
	re := regexp.MustCompile(`(?s)<blockquote>[\s]*<pre>(.*?)</pre>[\s]*</blockquote>`)
	f := func(input []byte) []byte {
		s := string(input)
		s = strings.Replace(s, "<br />", "", -1)
		s = strings.Replace(s, "<br/>", "", -1)
		s = strings.Replace(s, "<p>", "", -1)
		s = strings.Replace(s, "</p>", "", -1)
		return []byte(s)
	}
	replaced := re.ReplaceAllFunc(b, f)
	fmt.Print(strings.TrimSpace(string(replaced)))
}
