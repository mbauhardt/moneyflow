package entities

import "testing"

func TestTagEquals(t *testing.T) {
	type args struct {
		t1 []Tag
		t2 []Tag
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"different len",
			args{
				[]Tag{Tag{Modifier: add, Name: "tag1"}, Tag{Modifier: add, Name: "tag2"}},
				[]Tag{Tag{Modifier: add, Name: "tag1"}},
			},
			false,
		},
		{"different modifier",
			args{
				[]Tag{Tag{Modifier: remove, Name: "tag"}},
				[]Tag{Tag{Modifier: add, Name: "tag"}},
			},
			false,
		},
		{"different name",
			args{
				[]Tag{Tag{Modifier: add, Name: "tag"}},
				[]Tag{Tag{Modifier: add, Name: "tag2"}},
			},
			false,
		},
		{"same modifier and name",
			args{
				[]Tag{Tag{Modifier: add, Name: "tag"}},
				[]Tag{Tag{Modifier: add, Name: "tag"}},
			},
			true,
		},
		{"same modifier and name but in different order",
			args{
				[]Tag{Tag{Modifier: add, Name: "tag"}, Tag{Modifier: add, Name: "tag2"}},
				[]Tag{Tag{Modifier: add, Name: "tag2"}, Tag{Modifier: add, Name: "tag"}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TagEquals(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoneyEquals(t *testing.T) {
	type args struct {
		m1 *Money
		m2 *Money
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same value",
			args{
				&Money{Value: 23},
				&Money{Value: 23},
			},
			true,
		},
		{"both nil",
			args{
				nil,
				nil,
			},
			true,
		},
		{"first is nil",
			args{
				nil,
				&Money{Value: 1},
			},
			false,
		},
		{"second is nil",
			args{
				&Money{Value: 1},
				nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoneyEquals(tt.args.m1, tt.args.m2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
