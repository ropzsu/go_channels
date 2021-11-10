package hello

import "testing"

func TestHelloWorld(t *testing.T) {
  t.Errorf("HelloWorld (%q) == (%q), want %q ", "hello", "world",  "hello")
}


/*
ZhiwendeMacBook-Pro:39-testing brianguo$ go test
--- FAIL: TestHelloWorld (0.00s)
    helloWorld_test.go:6: HelloWorld ("hello") == ("world"), want "hello" 
FAIL
exit status 1
FAIL	hello	0.687s
ZhiwendeMacBook-Pro:39-testing brianguo$



*/


/*
package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

*/
