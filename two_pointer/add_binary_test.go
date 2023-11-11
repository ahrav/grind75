package two_pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		str1 string
		str2 string
		want string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
		{"1", "0", "1"},
		{"0", "1", "1"},
		{"1", "1", "10"},
		{"", "1", "1"},
		{"1", "", "1"},
		{"", "", ""},
	}

	for _, tc := range testCases {
		got := addBinary(tc.str1, tc.str2)
		assert.Equal(t, tc.want, got)
	}
}

func BenchmarkAddBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addBinary("1010", "1011")
	}
}
