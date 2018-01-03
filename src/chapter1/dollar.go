package chapter1

import "fmt"

type Dollar struct {
    amount int
}

func New_Dollar(a int) (Dollar) {
    return Dollar{amount: a}
}

func (d Dollar) GetAmount() {
    fmt.Println(d.amount)
}

func (d *Dollar) SetAmount(n int) {
    d.amount = n
}

func (d Dollar) Times(t int) (int) {
    return d.amount * t
}

