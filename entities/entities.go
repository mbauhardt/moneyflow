package entities

import (
	"strconv"
)

// ************
// tag section
// ************
type TagModifier string

const (
	add    TagModifier = "+"
	remove TagModifier = "-"
)

type Tag struct {
	Modifier TagModifier
	Name     string
}

func TagToAdd(name string) Tag {
	return Tag{Modifier: add, Name: name}
}

func TagToRemove(name string) Tag {
	return Tag{Modifier: remove, Name: name}
}

func TagEquals(t1, t2 []Tag) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i, v := range t1 {
		if v.Modifier != t2[i].Modifier {
			return false
		}
		if v.Name != t2[i].Name {
			return false
		}
	}
	return true
}

func ToString(t *Tag) string {
	return string(t.Modifier) + t.Name
}

// **************
// money section
// **************

type Money struct {
	Value int64
}

func MoneyEquals(m1, m2 *Money) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}
	return m1.Value == m2.Value
}

func MoneyToString(m *Money) string {
	if m.Value < 0 {
		return "-€" + strconv.FormatInt(m.Value*-1, 10)
	} else {
		return "€" + strconv.FormatInt(m.Value, 10)
	}
}

// **************
// description section
// **************

type Description struct {
	Value string
}

func DescriptionEquals(m1, m2 *Description) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}
	return m1.Value == m2.Value
}
