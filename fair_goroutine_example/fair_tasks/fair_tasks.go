package fairtasks

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/neutrin/fair_gouroutine_example/constants"
)

var curNo int64

func runTask(goRutineNo int, wg *sync.WaitGroup) {
	defer wg.Done()
	curTime := time.Now()
	for {
		x := atomic.AddInt64(&curNo, 1)
		if x >= constants.MaxRange {
			break
		}

		checkPrime(int(x))
	}

	fmt.Printf(" go routine = %d took = %d ms to complete \n", goRutineNo, time.Since(curTime).Milliseconds())

}

func checkPrime(no int) {
	if no <= 1 {
		return
	}
	if no == 2 {
		atomic.AddInt64(&constants.PrimeCount, 1)
		return
	}
	for curDiv := 3; curDiv <= int(math.Sqrt(float64(no))); curDiv++ {
		if no%curDiv == 0 {
			return
		}
	}
	atomic.AddInt64(&constants.PrimeCount, 1)
}

func Execute() {
	curNo = 1
	curTime := time.Now()
	var wg sync.WaitGroup
	for goRoutine := 0; goRoutine < int(constants.MaxRoutine); goRoutine++ {
		wg.Add(1)
		go runTask(goRoutine, &wg)

	}
	wg.Wait()
	fmt.Printf(" fair task execution completed in = %f time and found = %d prime number  \n", time.Since(curTime).Seconds(), constants.PrimeCount)
}
