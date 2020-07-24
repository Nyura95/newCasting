# newBroadcasting

```shell
go get github.com/Nyura95/nyubroadcasting
```

## Basic usage

```go
package main

import "log"
import "github.com/Nyura95/nyubroadcasting"

func main() {
  broadcasting := nyubroadcasting.NewBroadcasting()
	broadcasting.Start()

	go listener(broadcasting.CreateNewListener("hash2")) // start any listeners
	go listener(broadcasting.CreateNewListener("hash1"))

	broadcasting.Broadcaster <- "Hello world" // send to all listeners "hello world"
	time.Sleep(2 * time.Second)

	broadcasting.StopListener("hash1") // stop listener hash1
	time.Sleep(2 * time.Second)

	broadcasting.StopAllListener() // stop all listeners
	time.Sleep(2 * time.Second)

	broadcasting.Stop() // stop broadcaster
}

func listener(listener <-chan nyubroadcasting.ExternalCommunication) {
	for {
		com, alive := <-listener
		if !alive {
			log.Println("channel closed")
			break
		}
		log.Println(com)
	}
}
```
