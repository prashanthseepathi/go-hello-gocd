package main


import "testing"


func TestAlwaysPass(t *testing.T) {
if 2+2 != 4 {
t.Fatal("math broke :(")
}
}
