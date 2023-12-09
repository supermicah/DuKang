package waitgroup

import (
	"fmt"
	"time"

	"github.com/supermicah/dionysus/waitgroup"
)

func Do() {
	wg := &waitgroup.WaitGroup{}
	for i := 0; i < 10; i++ {
		second := time.Duration(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * second)
		}()
	}
	success := wg.WaitSuccess(time.Second * 5)
	fmt.Println(success)
}
