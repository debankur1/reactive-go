package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
)

var wait sync.WaitGroup
func ProduceImages(images chan rxgo.Item){
   for i:=0;i<10;i++ {
   	    images <- rxgo.Of("a"+string(i)+".jpg")
   }
   //Optionally these channel can be kept open if these process is never ending
   close(images)
	wait.Done()
}

func ClearImage(images chan rxgo.Item,outputChannel chan rxgo.Item) {
	observe := rxgo.FromChannel(images)
	for data := range observe.
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			// Can call an external api to clear the images
			newImage := fmt.Sprint(item) + " Colored_Image"
			return newImage, nil
		},
			rxgo.WithBufferedChannel(1)).Observe() {
		    outputChannel <- data
	}
	wait.Done()
}

func UploadFiles(Output chan rxgo.Item){
	if Output!=nil {
		for data := range Output {
            fmt.Println("Uploading files")
			fmt.Println(data.V)
		}
	}
	close(Output)
	fmt.Println("Output is closed")
	wait.Done()
}

//func main()  {
//	images := make(chan rxgo.Item)
//	OutputChannel := make(chan rxgo.Item)
//	wait.Add(3)
//	go ProduceImages(images)
//	go ClearImage(images,OutputChannel)
//    go UploadFiles(OutputChannel)
//	wait.Wait()
//}