package double

import (
	"testing"
)

func TestInt(t *testing.T) {
	type test struct {
		in   int
		want int
	}

	var (
		tests = []test{
			test{in: 1, want: 2},
			test{in: 2, want: 4},
			test{in: 3, want: 6},
			test{in: 10, want: 20},
			test{in: 100, want: 200},
			test{in: -1, want: -2},
			test{in: -2, want: -4},
			test{in: -3, want: -6},
			test{in: -10, want: -20},
			test{in: -100, want: -200},
		}
	)

	for _, tt := range tests {
		got := Int(tt.in)
		if got != tt.want {
			t.Errorf("Int(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}

func TestInt32(t *testing.T) {
	type test struct {
		in   int32
		want int32
	}

	var (
		tests = []test{
			test{in: 1, want: 2},
			test{in: 2, want: 4},
			test{in: 3, want: 6},
			test{in: 10, want: 20},
			test{in: 100, want: 200},
			test{in: -1, want: -2},
			test{in: -2, want: -4},
			test{in: -3, want: -6},
			test{in: -10, want: -20},
			test{in: -100, want: -200},
		}
	)

	for _, tt := range tests {
		got := Int32(tt.in)
		if got != tt.want {
			t.Errorf("Int32(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}

func TestUInt(t *testing.T) {
	type test struct {
		in   uint
		want uint
	}

	var (
		tests = []test{
			test{in: 1, want: 2},
			test{in: 2, want: 4},
			test{in: 3, want: 6},
			test{in: 10, want: 20},
			test{in: 100, want: 200},
		}
	)

	for _, tt := range tests {
		got := UInt(tt.in)
		if got != tt.want {
			t.Errorf("UInt(%d) = %d; want %d", tt.in, got, tt.want)
		}
	}
}
