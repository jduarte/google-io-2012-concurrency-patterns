package main

import (
        "fmt"
        "time"
        "math/rand"
)

// Generator Pattern
// # Function that return a channel

// Returns receive-only channel
func boring(msg string) <-chan string {
  c := make(chan string)

  // Launch the goroutine from inside 'boring' function
  go func() {
    for i := 0; ; i++ {
      c <- fmt.Sprintf("%s %d", msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()

  // Return the channel to the caller
  return c
}

func main() {
  // c := boring("Oh so boring ...")
  //
  // for i := 0; i < 5; i++ {
  //   fmt.Printf("You said: %q\n", <-c)
  // }
  //
  // fmt.Println("You're boring. Leaving ... ")



  // // Chanells as a handle on a service
  // joe := boring("Joe")
  // ann := boring("Ann")
  //
  // for i := 0; i < 5; i++ {
  //   fmt.Println(<-joe)
  //   fmt.Println(<-ann)
  // }
  //
  // fmt.Println("You're both boring. Leaving ... ")



  // // Multiplexing, limiting nr of inputs
  // // * Here Joe and Ann run in two completely independent connections

  // fanIn := func(input1, input2 <-chan string) <-chan string {
  //   cc := make(chan string)
  //   go func() {
  //     for {
  //       cc <- <-input1
  //     }
  //   }()
  //
  //   go func() {
  //     for {
  //       cc <- <-input2
  //     }
  //   }()
  //
  //   return cc
  // }
  //
  // c := fanIn(boring("Joe"), boring("Ann"))
  // for i := 0; i < 10; i++ {
  //   fmt.Println(<-c)
  // }
  //
  // fmt.Println("You're both boring. Leaving ... ")

  // // Restoring sequence

  fanIn := func(input1, input2 <-chan string) <-chan string {
    cc := make(chan string)
    go func() {
      for {
        cc <- <-input1
      }
    }()

    go func() {
      for {
        cc <- <-input2
      }
    }()

    return cc
  }

  c := fanIn(boring("Joe"), boring("Ann"))
  // We include a message structure with all required fields for the message
  // that we want a pass and a 'wait' channel to function as a signaler
  type Message struct {
    str string
    wait chan bool
  }

  for  i := 0; i < 5; i++ {
    msg1 := <- c
    fmt.Println(msg1.str)
    msg2 := <- c
    fmt.Println(msg2.str)
    msg1.wait <- true
    msg2.wait <- true
  }

  waitForIt := make(chan bool)

  c <- Message{ fmt.Sprintf("%s: %d", msg, i), waitForIt }
  time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
  <-waitForIt
}
