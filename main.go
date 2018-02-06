package main

import "fmt"

const prompt = `
Please enter number of operation:
1. Create an account
2. Show accout detail
3. Withdraw
4. Deposit
5. Make transfer
6. Show all account detail
7. Delete an account
8. Exit
`
func main() {
	fmt.Println("Welcome to Account System")
Exit:
	for {
		fmt.Println(prompt)
		var num int
		fmt.Scanf("%d\n", &num)
		switch num {
		case 1:
			fmt.Println("please input your <name> <balance>:")
			var name string
			var balance float64
			fmt.Scanf("%s %f\n", &name, &balance)
			err := NewAccount(name, balance)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Println("create success")
		case 2:
			fmt.Println("please input account id:")
			var id int64
			fmt.Scanf("%d\n", &id)
			a, err := GetAccount(id)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Printf("%#v\n", a)
		case 3:
			fmt.Println("please input account <id> <withdraw>:")
			var id int64
			var withdraw float64
			fmt.Scanf("%d %f\n", &id, &withdraw)
			a, err := MakeWithdraw(id, withdraw)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Printf("%#v\n", a)
		case 4:
			fmt.Println("please input your account <id> <balance>:")
			var id int64
			var balance float64
			fmt.Scanf("%d %f\n", &id, &balance)
			a, err := MakeDeposit(id, balance)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			fmt.Printf("%#v\n", a)
		case 5:
			fmt.Println("please input <outid> <money> <inid>:")
			var outId, inId int64
			var money float64
			fmt.Scanf("%d %f %d\n", &outId, &money, &inId)
			err := MakeTransfer(outId, money, inId)
			if err != nil {
				fmt.Printf("transfer error: %v\n", err)
			}
			fmt.Println("transfer success")
		case 6:
			as, err := ListAccounts()
			if err != nil {
				fmt.Printf("query error: %v\n", err)
			}
			for _, a := range as {
				fmt.Printf("account: %#v\n", a)
			}
		case 7:
			fmt.Println("please input delete account <id>:")
			var id int64 
			fmt.Scanf("%d\n", &id)
			err := DeleteAccount(id)
			if err != nil {
				fmt.Printf("delete error: %v\n", err)
			}
			fmt.Println("delete success")
		case 8:
			break Exit
		}
	}
}