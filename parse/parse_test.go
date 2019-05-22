package parse

import (
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
			if !entities.Equals(got, tt.expected) {
				t.Errorf("ParseTags(%q) == %q, want %q", tt.in, got, tt.expected)
			}
		})
	}
}
