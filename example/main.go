package main

import (
	"fmt"

	"github.com/rhidayat1980/goflow"
	flow "github.com/trustmaster/goflow"
)

func main() {

	net := goflow.NewEncodingApp()

	in := make(chan string)
	net.SetInPort("In", in)

	wait := flow.Run(net)

	for i := 0; i < 20; i++ {
		in <- fmt.Sprint("message", i)
	}

	close(in)
	<-wait
}
