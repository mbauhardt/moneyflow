package parse

import (
	"errors"
	"strings"
	"testing"

	"github.com/mbauhardt/moneyflow/entities"
)

func TestParseTags(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected []entities.Tag
	}{

		{"this is not a tag", "this is not a tag to add", []entities.Tag{}},
		{"this is a tag to add", "this is a +tag to add", []entities.Tag{entities.TagToAdd("tag")}},
		{"this is a tag to remove", "this is a -tag to remove", []entities.Tag{entities.TagToRemove("tag")}},
		{"two tags", "this is a +tag and this is also a +tag2 to add", []entities.Tag{entities.TagToAdd("tag"), entities.TagToAdd("tag2")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseTags(tt.in)
			if !entities.TagEquals(got, tt.expected) {
				t.Errorf("ParseTags(%q) == %q, want %q", tt.in, got, tt.expected)
			}
		})
	}
}

func ErrorEquals(e1, e2 error) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if e1 == nil || e2 == nil {
		return false
	}
	return strings.Compare(e1.Error(), e2.Error()) == 0
}

func TestParseMoney(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected *entities.Money
		err      error
	}{
		{"AbsentMoney", "There is no money", nil, nil},
		{"TwoTimesMoney", "foo $300 and $600", nil, errors.New("More than one money is detected: $300 $600")},
		{"One Expense", "buy food -$800", &entities.Money{Value: -800}, nil},
		{"One Expense in the middle of command", "buy food for -$800 abc", &entities.Money{Value: -800}, nil},
		{"One Income", "rent $300", &entities.Money{Value: 300}, nil},
		{"One Income in the middle", "rent $300 my flat", &entities.Money{Value: 300}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMoney(tt.in)
			if !entities.MoneyEquals(got, tt.expected) {
				t.Errorf("ParseMoney(%q) == %q, want %q", tt.in, got, tt.expected)
			}
			if !ErrorEquals(err, tt.err) {
				t.Errorf("ParseMoney(%q) == %q, want error %q", tt.in, err, tt.err)
			}
		})
	}
}
