package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	NewProblemAttempt()
}

func NewProblemAttempt() {
	r := flag.String("s", "", "The name of the source where the problem comes from (e.g. grokking)")
	p := flag.String("p", "", "The name of the problem (use-names-with-dashes)")
	f := flag.String("f", "", "The name of a file if different from problem (useCamelCase)")

	flag.Parse()

	dirs := map[string]string{
		"grokking": "grokking-the-coding-interview",
		"epi":      "elements-of-programming-interviews",
	}

	if _, ok := dirs[*r]; !ok {
		fmt.Printf("Invalid resource name: %s", *r)
		return
	}

	var base string
	time := time.Now().Format("2006-01-02-03-04-05")
	name := *p + "-" + time[:19]
	dir := filepath.Join("src", dirs[*r], name)

	if *f != "" {
		base = *f
	} else {
		base = *p
	}

	test := filepath.Join(dir, base+".test.ts")
	file := filepath.Join(dir, base+".ts")

	err := os.MkdirAll(dir, 0777)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(
		test,
		[]byte(
			fmt.Sprintf(
				`import { %s } from './%s';

describe('%s', () => {
	it('', () => {
		%s()
	});
});
`, base, base, base, base)),
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(
		file,
		[]byte(
			fmt.Sprintf(
				`export function %s() {
	//
}
`, base)),
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("code", test)

	cmd.Dir = wd

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("code", file)

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
