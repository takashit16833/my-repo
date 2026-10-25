package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
    actual := EvenOrOdd(10)

    if actual != "even" {
        t.Errorf("expected: even, actual: %s", actual)
    }
}
