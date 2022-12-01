package main

import (
	"fmt"
	"os"
)

// go run gen.go dayN
func main() {
	if os.Args[1] == "" {
		panic("not enough args")
	}

	pkg := os.Args[1]
	if err := os.Mkdir(pkg, 0755); err != nil {
		panic(err)
	}

	createFile := func(name, content string) {
		d := []byte(content)
		if err := os.WriteFile(name, d, 0644); err != nil {
			panic(err)
		}
	}

	createFile(fmt.Sprintf("%s/partial_input.txt", pkg), "")
	createFile(fmt.Sprintf("%s/full_input.txt", pkg), "")

	// could probably use templates if this got more complex
	pkgContent := fmt.Sprintf(`package %s

func Part1(path string) int {
	return 0
}

func Part2(path string) int {
	return 0
}
`, pkg)
	createFile(fmt.Sprintf("%s/%s.go", pkg, pkg), pkgContent)

	testContent := fmt.Sprintf(`package %s

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestPart1(t *testing.T) {
	helpers.TestsForFunction(
		t,
		Part1,
		"Part1",
		0,
		0,
	)
}

func TestPart2(t *testing.T) {
	helpers.TestsForFunction(
		t,
		Part2,
		"Part2",
		0,
		0,
	)
}
`, pkg)
	createFile(fmt.Sprintf("%s/%s_test.go", pkg, pkg), testContent)
}
