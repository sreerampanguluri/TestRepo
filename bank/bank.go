package main

import "fmt"

type Person struct {
	bal int
	dep int
	wit int
}

func Deposit() {
	var ram Person
	ram.bal = 1000
	fmt.Println("Enter the amount you want to deposit")
	fmt.Scan(&ram.dep)
	ram.bal = ram.bal + ram.dep
	fmt.Println(ram.bal)
	if ram.dep > 50000 {
		panic("The maximum amount you can deposited is $5000 in a single trasaction")
	}
}

func Withdraw() {
	var ram Person
	ram.bal = 1000
	fmt.Println("Enter how much amount you would like to draw")
	fmt.Scan(&ram.wit)
	if ram.bal > 0 {
		ram.bal = ram.bal - ram.wit
		fmt.Println(ram.bal)
	}
	if ram.wit > ram.bal {
		panic("Amount entered is more then the your balance")
	}
}

func Defer() {

	defer fmt.Println("Thank you for Banking with us")

	fmt.Println("Enter 1 for deposit or Enter 2 for withdrawal")
	var option int
	fmt.Scan(&option)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	switch option {
	case 1:
		Deposit()
	case 2:
		Withdraw()

	}

}

func main() {
	Defer()
}
