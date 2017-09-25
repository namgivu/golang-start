package main

import "testing"

func TestSomeHelper(t *testing.T) {
  cases := []struct {
    in, want string
  }{
    {"world", "Hello world"},
    {"there", "Hello there"},
  }

  for _, c := range cases {
    got := SayHello(c.in)
    if got != c.want {
      t.Errorf("SayHello(%q) == %q, want %q", c.in, got, c.want)
    }
  }

}
