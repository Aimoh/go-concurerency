package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("I'm a miner:", n, "my working day is over")
			return
		default:
			fmt.Println("I'm a miner number:", n, "started to finish off the coal")
			time.Sleep(1 * time.Second)
			fmt.Println("I'm a miner number:", n, "finished off the coal:", power)

			transferPoint <- power
			fmt.Println("I'm a miner number:", n, "donated coal:", power)
		}
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
