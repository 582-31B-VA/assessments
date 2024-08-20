package main

import "testing"

func TestWithdraw(t *testing.T) {
	t.Run("Sufficiant funds", func(t *testing.T) {
		b := newBank()
		a := b.newAccount("", 10)
		got, _ := a.withdraw(5)
		want := cents(5)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("Insufficiant funds", func(t *testing.T) {
		b := newBank()
		a := b.newAccount("", 5)
		got, err := a.withdraw(10)
		want := cents(5)
		if got != want && err != nil {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}
