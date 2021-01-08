package bank

var deposits = make(chan int)
var withdraws = make(chan int)
var available = make(chan bool)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	if amount < 0 {
		return false
	}
	withdraws <- amount
	return <-available
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount

		case amount := <-withdraws:
			if balance < amount {
				available <- false
			} else {
				balance -= amount
				available <- true
			}

		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
