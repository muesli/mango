# mango

[![Build Status](https://github.com/muesli/mango/workflows/build/badge.svg)](https://github.com/muesli/mango/actions)
[![Go ReportCard](https://goreportcard.com/badge/muesli/mango)](https://goreportcard.com/report/muesli/mango)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/muesli/mango)

mango is a man-page generator for the Go flag, pflag, and cobra packages. It
extracts commands, flags, and arguments from your program and enables it to
self-document.

## Usage with flag:

```go
import (
    "flag"
    "fmt"

    "github.com/muesli/mango"
    "github.com/muesli/mango/mflag"
    "github.com/muesli/roff"
)

var (
    one = flag.String("one", "", "first value")
    two = flag.String("two", "", "second value")
)

func main() {
    flag.Parse()

    manPage := mango.NewManPage(1, "mango", "mango - a man-page generator").
        WithLongDescription("mango is a man-page generator for Go.\n"+
            "Features:\n"+
            "* User-friendly\n"+
            "* Plugable").
        WithSection("Copyright", "(C) 2022 Christian Muehlhaeuser.\n"+
            "Released under MIT license.")

    flag.VisitAll(mflag.FlagVisitor(manPage))
    fmt.Println(manPage.Build(roff.NewDocument()))
}
```

Mango will extract all the flags from your app and generate a man-page similar
to this example:

![mango](/mango.png)

## Usage with pflag:

```go
import (
    "fmt"

    "github.com/muesli/mango"
    "github.com/muesli/mango/mpflag"
    "github.com/muesli/roff"
    flag "github.com/spf13/pflag"
)

func main() {
    flag.Parse()

    manPage := mango.NewManPage(1, "mango", "mango - a man-page generator").
        WithLongDescription("mango is a man-page generator for Go.").
        WithSection("Copyright", "(C) 2022 Christian Muehlhaeuser.\n"+
            "Released under MIT license.")

    flag.VisitAll(mpflag.PFlagVisitor(manPage))
    fmt.Println(manPage.Build(roff.NewDocument()))
}
```

## Usage with cobra:

```go
import (
	"fmt"

	mcobra "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
)

var (
    rootCmd = &cobra.Command{
        Use:   "mango",
        Short: "A man-page generator",
    }
)

func main() {
    manPage, err := mcobra.NewManPageFromCobra(1, rootCmd)
    if err != nil {
        panic(err)
    }

    manPage = manPage.WithSection("Copyright", "(C) 2022 Christian Muehlhaeuser.\n"+
        "Released under MIT license.")

    fmt.Println(manPage.Build(roff.NewDocument()))
}
```

## Feedback

Got some feedback or suggestions? Please open an issue or drop me a note!

* [Twitter](https://twitter.com/mueslix)
* [The Fediverse](https://mastodon.social/@fribbledom)
