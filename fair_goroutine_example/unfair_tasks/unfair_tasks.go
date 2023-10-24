package unfair_tasks

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/neutrin/fair_gouroutine_example/constants"
)

func runTask(goRoutineNo, startRange, endRange int64, wg *sync.WaitGroup) {
	defer wg.Done()
	curTime := time.Now()
	for curNo := startRange; curNo <= endRange; curNo++ {
		checkPrime(int(curNo))
	}
	fmt.Printf(" go routine = %d took = %d ms to complete \n", goRoutineNo, time.Since(curTime).Milliseconds())

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
	batchSize := int(float64(constants.MaxRange) / float64(constants.MaxRoutine))
	curBatchRange := 1
	curTime := time.Now()
	var wg sync.WaitGroup
	for routineNo := 0; routineNo < int(constants.MaxRoutine); routineNo++ {
		wg.Add(1)
		go runTask(int64(routineNo), int64(curBatchRange),
			int64(math.Min(float64(int64(curBatchRange+int(batchSize)-1)), float64(constants.MaxRange))), &wg)
		curBatchRange += int(batchSize)

	}
	wg.Wait()
	fmt.Printf(" unfair total time Taken to complete job = %f seconds  and prime = %d \n", time.Since(curTime).Seconds(), constants.PrimeCount)
}
