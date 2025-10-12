package mango

import (
	"strings"
	"testing"
	"time"
)

// DummyBuilder is an implementation of [Builder] to be used in mango's tests.
type DummyBuilder struct {
	strings.Builder
}

func (b *DummyBuilder) Heading(section uint, title, description string, ts time.Time) {
	// TODO: implement this when needed
}
func (b *DummyBuilder) Paragraph() {
	// TODO: implement this when needed
}
func (b *DummyBuilder) Indent(n int) {
	// TODO: implement this when needed
}
func (b *DummyBuilder) IndentEnd() {
	// TODO: implement this when needed
}
func (b *DummyBuilder) TaggedParagraph(indentation int) {
	// TODO: implement this when needed
}
func (b *DummyBuilder) List(text string) {
	// TODO: implement this when needed
}
func (b *DummyBuilder) Section(text string) {
	// TODO: implement this when needed
}
func (b *DummyBuilder) EndSection() {
	// TODO: implement this when needed
}
func (b *DummyBuilder) Text(text string) {
	b.WriteString(text)
}
func (b *DummyBuilder) TextBold(text string) {
	b.WriteString("<b>")
	b.WriteString(text)
	b.WriteString("</b>")
}
func (b *DummyBuilder) TextItalic(text string) {
	b.WriteString("<i>")
	b.WriteString(text)
	b.WriteString("</i>")
}
func (b *DummyBuilder) String() string { return b.Builder.String() }

func TestBuildSynopsis(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic", "test [OPTION]... [ARGUMENT]...", "<b>test</b> [<i>OPTION</i>]... [<i>ARGUMENT</i>]..."},
		{"non-closed brackets", "test [argument", "<b>test</b> [argument"},
		{"end with lbracket", "test argument[", "<b>test</b> argument["},
		{"end with rbracket", "test argument]", "<b>test</b> argument]"},
		{"single bracketed arg", "test [argument]", "<b>test</b> [<i>argument</i>]"},
		{"no args", "test", "<b>test</b>"},
		{"no args with space", "test ", "<b>test</b>"},
		{"empty", "", "<b></b>"},
		{"single space", "", "<b></b>"},
		{"starts with space", " test args", "<b>test</b> args"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := NewManPage(1, "", "")
			w := new(DummyBuilder)
			m.buildSynopsis(w, c.input)
			actual := w.String()
			if actual != c.expected {
				t.Fatalf("expected %q, got %q", c.expected, actual)
			}
		})
	}
}
