/**
 * @Author: koulei
 * @Description:
 * @File: sync_test
 * @Version: 1.0.0
 * @Date: 2023/3/12 15:37
 */

package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var numCalcCreated uint32
var currentWorked uint32

type Buff struct {
	pool sync.Pool
}

func (b *Buff) Get() interface{} {
	object := b.pool.Get()
	if object != nil {
		// fmt.Println("from cache")
		return object
	}

	// fmt.Println("new object")
	atomic.AddUint32(&numCalcCreated, 1)
	newObject := make(map[string]interface{})
	return newObject
}

func (b *Buff) Put(object interface{}) {
	b.pool.Put(object)
}

func BenchmarkPool(b *testing.B) {
	buff := new(Buff)
	// runNum := 1024 * 1024
	// var wg sync.WaitGroup
	// wg.Add(runNum)

	start := time.Now()
	// for i := 0; i < runNum; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		atomic.AddUint32(&currentWorked, 1)
	// 		m1 := buff.Get().(map[string]interface{})
	// 		m1["a"] = 1
	// 		// fmt.Println("m1: ", m1)
	// 		buff.Put(m1)
	//
	// 		// m2 := buff.Get()
	// 		// fmt.Println("m2: ", m2)
	// 		// buff.Put(m2)
	// 	}()
	// }
	// wg.Wait()

	for i := 0; i < b.N; i++ {
		atomic.AddUint32(&currentWorked, 1)
		m1 := buff.Get().(map[string]interface{})
		m1["a"] = 1
		// fmt.Println("m1: ", m1)
		buff.Put(m1)
	}

	fmt.Printf("numCalcCreated: %d, currentWorked: %d, execution time: %s\n", numCalcCreated, currentWorked, time.Now().Sub(start))
}
