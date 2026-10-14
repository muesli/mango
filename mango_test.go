package mango

import (
	"os"
	"testing"
	"time"
)

type mockBuilder struct {
	date time.Time
}

func (m *mockBuilder) Heading(section uint, title, description string, ts time.Time) {
	m.date = ts
}

func (m *mockBuilder) Paragraph()                     {}
func (m *mockBuilder) Indent(n int)                  {}
func (m *mockBuilder) IndentEnd()                    {}
func (m *mockBuilder) TaggedParagraph(indentation int) {}
func (m *mockBuilder) List(text string)              {}
func (m *mockBuilder) Section(text string)           {}
func (m *mockBuilder) EndSection()                   {}
func (m *mockBuilder) Text(text string)              {}
func (m *mockBuilder) TextBold(text string)          {}
func (m *mockBuilder) TextItalic(text string)        {}
func (m *mockBuilder) String() string                { return "" }

func TestManPageDateDefault(t *testing.T) {
	os.Unsetenv("SOURCE_DATE_EPOCH")
	mp := NewManPage(1, "test", "a test command")
	mb := &mockBuilder{}
	before := time.Now().Add(-time.Second)
	mp.Build(mb)
	after := time.Now().Add(time.Second)

	if mb.date.Before(before) || mb.date.After(after) {
		t.Errorf("expected default date to be current time, got %v", mb.date)
	}
}

func TestManPageSourceDateEpoch(t *testing.T) {
	expectedTime := time.Unix(1600000000, 0).UTC()
	os.Setenv("SOURCE_DATE_EPOCH", "1600000000")
	defer os.Unsetenv("SOURCE_DATE_EPOCH")

	mp := NewManPage(1, "test", "a test command")
	mb := &mockBuilder{}
	mp.Build(mb)

	if !mb.date.Equal(expectedTime) {
		t.Errorf("expected date from SOURCE_DATE_EPOCH to be %v, got %v", expectedTime, mb.date)
	}
}

func TestManPageWithDate(t *testing.T) {
	os.Setenv("SOURCE_DATE_EPOCH", "1600000000")
	defer os.Unsetenv("SOURCE_DATE_EPOCH")

	customTime := time.Unix(1700000000, 0).UTC()
	mp := NewManPage(1, "test", "a test command").WithDate(customTime)
	mb := &mockBuilder{}
	mp.Build(mb)

	if !mb.date.Equal(customTime) {
		t.Errorf("expected custom date from WithDate to override SOURCE_DATE_EPOCH, got %v", mb.date)
	}
}
