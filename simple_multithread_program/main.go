package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
)

type Balance struct {
	Total int
}

type PendingTransaction struct {
	Id            int
	DepositAmount int
}

// generate data of transactions to be processed
func generatePendingTransactions(numOfTransaction int) []*PendingTransaction {
	result := []*PendingTransaction{}
	totalExpectedDeposit := 0

	for range numOfTransaction {
		pendingTransaction := &PendingTransaction{
			DepositAmount: rand.Intn(10) + 1,
		}
		totalExpectedDeposit += pendingTransaction.DepositAmount
		result = append(result, pendingTransaction)
	}

	fmt.Printf("---> It is expected that a total amount of %v is added to the balance!\n\n", totalExpectedDeposit)
	return result
}

// process transaction data, this function will be passed into concurrency
func (b *Balance) processPendingTransaction(
	mutex *sync.Mutex,
	wg *sync.WaitGroup,
	useMutex bool,
	pendingTransaction *PendingTransaction,
) {
	defer wg.Done() // decrement work group count when the function is done

	if useMutex {
		// lock b.Total resource
		// other process cannot access unless it is unlocked
		mutex.Lock()
	}
	b.Total += pendingTransaction.DepositAmount
	if useMutex {
		mutex.Unlock()
	}
}

func (b *Balance) PrinceTotalBalance() {
	fmt.Printf("---> The current total balance is %v \n", b.Total)
}

func main() {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	useMutex := flag.Bool("mutex", false, "enable mutex")
	flag.Parse()

	myBalance := &Balance{Total: 0}

	fmt.Println("=== Virtual Banking Deposit System ===")
	if *useMutex {
		fmt.Println("NOTICE: MUTEX IS USED!")
	} else {
		fmt.Println("NOTICE: MUTEX IS NOT USED!")
	}

	fmt.Println()
	myBalance.PrinceTotalBalance()

	numOfTransaction := 1000
	fmt.Printf("=== Generating %v Pending Deposit Transactions ===\n", numOfTransaction)
	pendingDepositTransactions := generatePendingTransactions(numOfTransaction)

	fmt.Println("=== Start Processing All Pending Transactions With Go Concurrency ===", "\n...\n...")

	wg.Add(numOfTransaction) // assign wait group to equal amount of the transaction
	for _, e := range pendingDepositTransactions {
		// start a new GoRoutine
		go myBalance.processPendingTransaction(&mutex, &wg, *useMutex, e)
	}

	wg.Wait() // wait until the number of work group is decremented to 0
	fmt.Printf("=== All Pending %v Transactions Are Processed ===\n",
		len(pendingDepositTransactions))

	myBalance.PrinceTotalBalance()
}
