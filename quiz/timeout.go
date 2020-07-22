package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1)

    // Run your long running function in it's own goroutine and pass back it's
    // response into our channel.
    go func() {
        text := LongRunningProcess()
        c1 <- text
    }()

    // Listen on our channel AND a timeout channel - which ever happens first.
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("out of time :(")
    }

}

func LongRunningProcess() string {
    time.Sleep(5 * time.Second)
    return "My golangcode.com example is finished :)"
}
