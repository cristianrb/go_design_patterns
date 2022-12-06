package main

import "fmt"

type BankAccount struct {
	balance int
}

func Deposit(b *BankAccount, amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
}

func Withdraw(b *BankAccount, amount int) {
	if b.balance >= amount {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
	}
}

func main() {
	ba := &BankAccount{0}
	var commands []func()
	commands = append(commands, func() {
		Deposit(ba, 100)
	})
	commands = append(commands, func() {
		Withdraw(ba, 25)
	})

	for _, cmd := range commands {
		cmd()
	}

	fmt.Println(*ba)
}
