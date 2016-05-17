package main

import (
	"fmt"
  "github.com/foo/bar"
  "github.com/foo/baz"
)

func main() {
	fmt.Println("Hello world!")
  bar.Bar()
  baz.Baz()
}
