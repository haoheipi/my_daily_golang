package main

import (
	"day2_client"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	day2_client.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	client, _ := day2_client.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup
	channel := make(chan *day2_client.Call, 5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			//if err := client.Call("Foo.Sum", args, &reply); err != nil {
			//	log.Fatal("call Foo.Sum error:", err)
			//}

			client.Go("Foo.Sum", args, &reply, channel)

		}(i)
	}
	wg.Wait()
	for {
		<-channel
		log.Println("reply:")
	}

}
