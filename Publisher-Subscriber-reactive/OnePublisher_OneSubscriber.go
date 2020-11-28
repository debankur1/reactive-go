package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
	"time"
)

type Employee struct {
	Name string
	Age  string
	EmployeeId string
}
var wait sync.WaitGroup

//func main(){
//	publisherChannel := make(chan rxgo.Item)
//	wait.Add(2)
//	// Making publisher and Subscriber run on different threads
//	go Publisher(publisherChannel)
//    go Subscriber(publisherChannel)
//	wait.Wait()
//
//}
func Subscriber(publishedChannel chan rxgo.Item ){
	// Basic Observable Listening to the Publisher
	observe := rxgo.FromChannel(publishedChannel)
	// Response from Observable
	response := observe.DoOnNext(func(i interface{}) {
		fmt.Printf("Response: %d\n ",i)
	})
	// Waiting for all the data to consume
	<-response
	wait.Done()
}
func Publisher(Employees chan rxgo.Item){
	for {
		Employees <- rxgo.Of(Employee{
			Name:       "Pat",
			Age:        "35",
			EmployeeId: "32145",
		})
		Employees <- rxgo.Of(Employee{
			Name:       "Sam",
			Age:        "45",
			EmployeeId: "32135",
		})
		Employees <- rxgo.Of(Employee{
			Name:       "Mike",
			Age:        "55",
			EmployeeId: "32165",
		})
		time.Sleep(100 * time.Millisecond)
	}
	// Using infinite loop is used for demonstration a blocking call can be added
	// in a real time scenario where data can come from external source

	//close(Employees)
	wait.Done()
}
