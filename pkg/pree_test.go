package pkg

import (
	"fmt"
	"sync"
	"testing"
	"time"
)
var w sync.WaitGroup

func aa(i interface{})  {
	time.Sleep(1*time.Second)
	fmt.Println(i)

	w.Done()
}

func TestA(t *testing.T)  {


	//

	t1 := time.Now().UnixNano()

	//p,_ := ants.NewPoolWithFunc(5, aa)
	for i:=0; i<20; i++ {
		w.Add(1)
		go aa(i)
		//go p.Invoke(i)
	}
	w.Wait()
	fmt.Println(time.Now().UnixNano()-t1)
}

func TestSchedule(t *testing.T)  {
	schedule(aa)
}

func schedule(f func(i interface{}))  {
	var durattion int = 10
	interval := 2

	totalbatch, batchcount := int(durattion/interval), 0
	tick := time.NewTicker(2*time.Second)
	defer func() {
		tick.Stop()
	}()
	for ; batchcount < totalbatch; batchcount++ {
		<-tick.C
		w.Add(1)
		fmt.Println("============")
		go f(batchcount)
	}
	w.Wait()

}