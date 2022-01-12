package mflag

import (
	"flag"

	"github.com/muesli/mango"
)

// FlagVisitor is used to visit all flags and track them in a mango.ManPage.
func FlagVisitor(m *mango.ManPage) func(*flag.Flag) {
	return func(f *flag.Flag) {
		_ = m.Root.AddFlag(mango.Flag{
			Name:  f.Name,
			Usage: f.Usage,
		})
	}
}
