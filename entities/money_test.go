package entities

import "testing"

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
