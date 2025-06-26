package main

import "fmt"

func main(){
	var chethanAccount =CreateAccount("chethan","chethan@gmail.com",23,"52638764481",100)
	chethanAccount.CheckBalance()
	chethanAccount.CreditAmount(1000)
	chethanAccount.CheckBalance()
	chethanAccount.WithdrawAmount(500)
	
}


type Account struct{
	Name string
	Email  string
	Age int
	AccountNo string
	Balance float32

}
func CreateAccount(name string,email string ,age int ,accountno string ,balance float32) Account {
	return Account{
		Name: name,
		Email: email,
		Age: age,
		AccountNo: accountno,
		Balance: balance,

	}
}
func (a *Account) CheckBalance(){
	fmt.Println(a.Balance)
}
func (a *Account) CreditAmount(amount float32){
a.Balance=a.Balance + amount 
}

func(a *Account) WithdrawAmount(amount float32){
	if a.Balance >= amount{
		a.Balance = a.Balance - amount
		fmt.Println(a.Balance)

	} else {
		fmt.Println("error")
	}

}