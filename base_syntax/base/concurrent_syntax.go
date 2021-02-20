//example demo in book -- Go In Action
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func first_goroutine() {
	runtime.GOMAXPROCS(2)
	var wg1 sync.WaitGroup
	wg1.Add(2)
	go func() {
		defer wg1.Done()
		for i := 0; i < 3; i++ {
			for c := 'a'; c < 'a'+26; c++ {
				fmt.Printf(" %c", c)
			}
		}
	}()

	go func() {
		defer wg1.Done()
		for i := 0; i < 3; i++ {
			for c := 'A'; c < 'A'+26; c++ {
				fmt.Printf(" %c", c)
			}
		}
	}()
	fmt.Println("将会等待打印")
	wg1.Wait()
	fmt.Print("\n打印已完成")
}

func atomic_util() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(3)
	var counter = int64(0)
	count := func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
		}
	}
	go count()

	go count()

	go count()

	wg.Wait()
	fmt.Println(counter)
}

func atomic_read_write() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(3)
	var status = int32(1)
	count := func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			fmt.Printf("循环的status: %d \n", status)
			time.Sleep(200 * time.Millisecond)
			if atomic.LoadInt32(&status) == 0 {
				fmt.Printf("推出时的status: %d \n", status)
				break
			}
		}
	}
	go count()

	go count()

	go count()
	time.Sleep(100 * time.Millisecond)
	atomic.StoreInt32(&status, 0)
	wg.Wait()
}

func mutex_lock() {
	var mutex sync.Mutex
	counter := int64(1)
	var wg sync.WaitGroup
	wg.Add(3)
	runtime.GOMAXPROCS(3)
	count := func() {
		defer wg.Done()
		fmt.Printf("goroutine 中读取到的值 : %d \n", counter)
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
	go count()
	go count()
	go count()

	wg.Wait()
	fmt.Println(counter)
}

//无缓冲的channel 会让写、读channel的双方都进入阻塞状态，知道对方接收到数据后，才会继续执行后续代码
//var1 <- unBuffered，从channel中取值
//chan <- var1，向channel中写值
func unbuffered_chan() {
	unBuffered := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	runtime.GOMAXPROCS(2)
	go func() {
		defer wg.Done()
		fmt.Printf("1开始读取channel时间 %s \n", time.Now())
		local := <-unBuffered
		fmt.Println(local)
		fmt.Printf("1读取channel后时间 %s \n", time.Now())

		time.Sleep(2 * time.Second)
		unBuffered <- strconv.Itoa(rand.Int())
		fmt.Printf("1写入channel后时间 %s \n", time.Now())

	}()
	go func() {
		defer wg.Done()
		unBuffered <- strconv.Itoa(rand.Int())
		fmt.Printf("2写入channel后时间 %s \n", time.Now())

		fmt.Printf("2开始读取channel时间 %s \n", time.Now())
		local := <-unBuffered
		fmt.Println(local)
		fmt.Printf("2读取channel后时间 %s \n", time.Now())
	}()

	wg.Wait()
	close(unBuffered)
}

func buffered_chan() {
	buffered := make(chan string, 3)
	var mainWg sync.WaitGroup
	var goroutineWg sync.WaitGroup
	mainWg.Add(1)
	goroutineWg.Add(1)
	runtime.GOMAXPROCS(2)
	go func() {
		for i := 0; i < 4; i++ {
			if i == 3 {
				time.AfterFunc(10*time.Second, func() {
					goroutineWg.Done()
				})
			}
			fmt.Printf("第%d次写入channel\n", i+1)
			buffered <- strconv.Itoa(rand.Int())
			fmt.Printf("第%d次写入channel，完成\n", i+1)
		}

	}()

	go func() {
		defer mainWg.Done()
		goroutineWg.Wait()
		for i := 0; i < 4; i++ {
			read := <-buffered
			fmt.Println(read)
		}

	}()

	mainWg.Wait()
	close(buffered)
}

func main() {
	buffered_chan()
}
