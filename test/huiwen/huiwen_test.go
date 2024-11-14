package huiwen

import (
	"reflect"
	"testing"
)

func TestIspalindrome(t *testing.T) {
	type test struct {
		input string
		want  bool
	}
	tests := []test{
		{input: "asffsa", want: true},
		{input: "sdfdfs", want: false},
	}
	for _, tc := range tests {
		got := IsPalindrome(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected:%v, got:%v", tc.want, got)
		}
	}
}

func BenchmarkIspalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("agga")
	}
}
