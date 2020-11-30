package main

import (
"context"
"fmt"
"github.com/reactivex/rxgo/v2"
"sync"
)

var waitSync sync.WaitGroup

func ProduceImagesWithScale(images chan rxgo.Item){
	for i:=0;i<10;i++ {
		images <- rxgo.Of("a"+string(i)+".jpg")
	}
	//Optionally these channel can be kept open if these process is never ending
	close(images)
	waitSync.Done()

}
func ClearImageWithScale(images chan rxgo.Item,outputChannel chan rxgo.Item) {
	observe := rxgo.FromChannel(images)
	for data := range observe.
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			// Can call an external api to clear the images
			newImage := fmt.Sprint(item) + " Colored_Image"
			return newImage, nil
		},
		    rxgo.WithCPUPool(),
			rxgo.WithBufferedChannel(1)).Observe() {
		    outputChannel <- data
	}
	waitSync.Done()
}

func UploadFilesWithScale(Output chan rxgo.Item){
	observe := rxgo.FromChannel(Output)
	for data := range observe.
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			//Real Uploading the file to a server and that also need some scale,using WithCPUPool() will scale these function
			return item, nil
		},
			rxgo.WithCPUPool(),
			rxgo.WithBufferedChannel(1)).Observe() {
			fmt.Println("Uploaded Completed ",data.V)
	}
	close(Output)
	waitSync.Done()
}
func main()  {
	images := make(chan rxgo.Item)
	OutputChannel := make(chan rxgo.Item)
	waitSync.Add(3)
	go ProduceImagesWithScale(images)
	go ClearImageWithScale(images,OutputChannel)
	go UploadFilesWithScale(OutputChannel)
	waitSync.Wait()
}
