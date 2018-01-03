package chapter1

import (
    "testing"
    )

func TestMultiplication(t *testing.T) {

    var bucks = New_Dollar(5)
    var result int = bucks.Times(2)

    if result != 10 {
        t.Errorf("Dollar 5 times(2) == %v, want 10", result)
    }
}

