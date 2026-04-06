package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	id     int
	status string
}

var (
	r  = rand.New(rand.NewSource(time.Now().UnixNano()))
	mu = sync.Mutex{}
)

func Run() {
	values := []int{0, 2, 4, 6, 8}
	ch := make(chan int, len(values))
	for _, v := range values {
		go func() {
			ch <- v * 2
		}()
	}
	for range len(values) {
		fmt.Println(<-ch)
	}
	///
	var wg sync.WaitGroup
	wg.Add(2)
	orders := generateOrders(20)
	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrders(orders)
	}()

	wg.Wait()
	printOrders(orders)

}

func (or Order) String() string {
	return fmt.Sprintf("Order[Id: %d, Status: %s]", or.id, or.status)
}

func generateOrders(count int) []Order {
	orders := make([]Order, count)
	for ind := range count {
		orders[ind] = Order{id: ind + 1, status: "pending"}
	}
	return orders
}

func processOrders(orders []Order) {
	mu.Lock()
	defer mu.Unlock()
	for _, order := range orders {
		fmt.Println("Processing...order ", order.id)
	}
}

func updateOrders(orders []Order) {
	mu.Lock()
	defer mu.Unlock()
	status := []string{"processing", "delievered", "shipped"}
	for idx := range len(orders) {
		orders[idx].status = status[r.Intn(len(status))]
	}
}

func printOrders(ors []Order) {
	for _, order := range ors {
		fmt.Println(order)
	}
}
