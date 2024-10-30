package main

import (
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

func generatePendingTransactions(numOfTransaction int) []*PendingTransaction {
	result := []*PendingTransaction{}
	totalExpectedDeposit := 0

	fmt.Printf("=== Generating %v Pending Deposit Transactions ===\n", numOfTransaction)
	for i := range numOfTransaction {
		pendingTransaction := &PendingTransaction{
			Id:            i + 1,
			DepositAmount: rand.Intn(10) + 1,
		}
		totalExpectedDeposit += pendingTransaction.DepositAmount
		result = append(result, pendingTransaction)

		// fmt.Printf("Transaction Id.%v, Amount To Deposit: %v \n",
		// 	pendingTransaction.Id, pendingTransaction.DepositAmount)
	}

	fmt.Printf("---> It is expected that a total amount of %v is added to the balance!\n", totalExpectedDeposit)
	return result
}

func (b *Balance) processPendingTransaction(
	mutex *sync.Mutex,
	wg *sync.WaitGroup,
	pendingTransaction *PendingTransaction,
) {
	defer wg.Done()

	mutex.Lock()
	b.Total += pendingTransaction.DepositAmount
	mutex.Unlock()

	// fmt.Printf("Transaction Id. %v Completed: %v added to the balance!\n",
	// 	pendingTransaction.Id, pendingTransaction.DepositAmount)
}

func (b *Balance) PrinceTotalBalance() {
	fmt.Printf("---> The current total balance is %v \n", b.Total)
}

func main() {
	var (
		wg       sync.WaitGroup
		mutex    sync.Mutex
		useMutex bool

		numOfTransaction = 1000
	)
	wg.Add(numOfTransaction)

	myBalance := &Balance{Total: 0}
	myBalance.PrinceTotalBalance()

	fmt.Println("=== Virtual Banking Deposit System ===")
	if useMutex {
		fmt.Println("NOTICE: MUTEX IS USED")
	} else {
		fmt.Println("NOTICE: MUTEX IS NOT USED!")
	}

	pendingDepositTransactions := generatePendingTransactions(numOfTransaction)
	fmt.Println("=== Start Processing Pending Transactions ===", "\n...\n...")

	for _, e := range pendingDepositTransactions {
		go myBalance.processPendingTransaction(&mutex, &wg, e)
	}
	wg.Wait()
	fmt.Printf("=== All Pending %v Transactions Are Processed ===\n",
		len(pendingDepositTransactions))

	myBalance.PrinceTotalBalance()
}
