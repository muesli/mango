package main

import (
	"fmt"

	"github.com/muesli/mango"
	"github.com/muesli/mango/mpflag"
	"github.com/muesli/roff"
	flag "github.com/spf13/pflag"
)

var (
	one = flag.String("one", "", "first value")
	two = flag.String("two", "", "second value")
)

func main() {
	flag.Parse()

	manPage := mango.NewManPage(1, "mango", "mango - a man-page generator").
		WithLongDescription("mango is a man-page generator for the Go flag, pflag, and cobra packages.\n"+
			"Features:\n"+
			"* User-friendly\n"+
			"* Plugable").
		WithSection("Authors", "mango was written by Christian Muehlhaeuser <https://github.com/muesli/mango>").
		WithSection("Copyright", "Copyright (C) 2022 Christian Muehlhaeuser.\n"+
			"Released under MIT license.")

	flag.VisitAll(mpflag.PFlagVisitor(manPage))

	fmt.Println(manPage.Build(roff.NewDocument()))
}
