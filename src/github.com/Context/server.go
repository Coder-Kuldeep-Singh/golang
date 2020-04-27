package main

import (
	"context"
	"fmt"
)

type message struct {
	responseChan chan<- int
	parameter    string
	ctx          context.Context
}

func ProcessMessage(work <-chan message) {
	for job := range work {
		select {
		// If the context is finished, don't bother processing the message.
		case <-job.ctx.Done():
			continue
		default:
		}
		// Assume this takes a long time to calculate
		hardToCalculate := len(job.parameter)
		select {
		case <-job.ctx.Done():
		case job.responseChan <- hardToCalculate:
		}
	}
}

func newRequest(ctx context.Context, input string, q chan<- message) {
	r := make(chan int)
	select {
	case <-ctx.Done():
		fmt.Println("Context ended before q could see message")
		return
	case q <- message{
		responseChan: r,
		parameter:    input,
		ctx:          ctx,
	}:
	}
	select {
	case out := <-r:
		fmt.Printf("The len of %s is %d\n", input, out)
	case <-ctx.Done():
		fmt.Println("Context ended before q coulf process message")
	}
}

func main() {
	q := make(chan message)
	go ProcessMessage(q)
	ctx := context.Background()
	newRequest(ctx, "Response Golang", q)
	newRequest(ctx, "Response Again", q)
	close(q)
}
