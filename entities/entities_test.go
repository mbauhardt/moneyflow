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
				[]Tag{Tag{modifier: add, name: "tag1"}, Tag{modifier: add, name: "tag2"}},
				[]Tag{Tag{modifier: add, name: "tag1"}},
			},
			false,
		},
		{"different modifier",
			args{
				[]Tag{Tag{modifier: remove, name: "tag"}},
				[]Tag{Tag{modifier: add, name: "tag"}},
			},
			false,
		},
		{"different name",
			args{
				[]Tag{Tag{modifier: add, name: "tag"}},
				[]Tag{Tag{modifier: add, name: "tag2"}},
			},
			false,
		},
		{"same modifier and name",
			args{
				[]Tag{Tag{modifier: add, name: "tag"}},
				[]Tag{Tag{modifier: add, name: "tag"}},
			},
			true,
		},
		{"same modifier and name but in different order",
			args{
				[]Tag{Tag{modifier: add, name: "tag"}, Tag{modifier: add, name: "tag2"}},
				[]Tag{Tag{modifier: add, name: "tag2"}, Tag{modifier: add, name: "tag"}},
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
