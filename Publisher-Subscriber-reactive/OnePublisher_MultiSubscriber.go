package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
	"time"
)
type Employee1 struct {
	Name string
	Age  string
	EmployeeId string
}

var waitOneToManyPubSub sync.WaitGroup

func main(){
	publisherChannel := make(chan rxgo.Item)
	waitOneToManyPubSub.Add(2)
	ctx := context.Background()
	// Making publisher and Subscriber run on different threads
	go OnePublisher(publisherChannel)
	go MultiSubscriber(publisherChannel,ctx)
	waitOneToManyPubSub.Wait()
}

func MultiSubscriber(publishedChannel chan rxgo.Item ,ctx context.Context){
	// Connectable Observer is used to listen
	observe := rxgo.FromChannel(publishedChannel,rxgo.WithPublishStrategy())
	// Subscriber 1
    observe.DoOnNext(func(i interface{}) {
		fmt.Printf("Response Subscriber 2 : %d\n ",i)
	})

	// Subscriber 2
	observe.DoOnNext(func(i interface{}) {
		fmt.Printf("Response Subscriber 1 : %d\n ",i)
	})
	// Waiting for all the data to consume
	observe.Connect(ctx)

	waitOneToManyPubSub.Done()
}
func OnePublisher(Employees chan rxgo.Item){
	for {
		Employees <- rxgo.Of(Employee1{
			Name:       "Pat",
			Age:        "35",
			EmployeeId: "32145",
		})
		Employees <- rxgo.Of(Employee1{
			Name:       "Sam",
			Age:        "45",
			EmployeeId: "32135",
		})
		Employees <- rxgo.Of(Employee1{
			Name:       "Mike",
			Age:        "55",
			EmployeeId: "32165",
		})
		time.Sleep(100 * time.Millisecond)
	}
	// Using infinite loop is used for demonstration a blocking call can be added
	// in a real time scenario where data can come from external source

	//close(Employees)
	waitOneToManyPubSub.Done()
}
