package double

import (
	"testing"
)

func TestInt(t *testing.T) {
	type test struct {
		in   int
		want int
		got  int
	}

	var (
		tests = []test{
			test{in: 1, want: 2, got: Int(1)},
			test{in: 2, want: 4, got: Int(2)},
			test{in: 3, want: 6, got: Int(3)},
			test{in: 10, want: 20, got: Int(10)},
			test{in: 100, want: 200, got: Int(100)},
			test{in: -1, want: -2, got: Int(-1)},
			test{in: -2, want: -4, got: Int(-2)},
			test{in: -3, want: -6, got: Int(-3)},
			test{in: -10, want: -20, got: Int(-10)},
			test{in: -100, want: -200, got: Int(-100)},
		}
	)

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("Int(%d) = %d; want %d", tt.in, tt.got, tt.want)
		}
	}
}
