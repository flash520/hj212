/**
 * @Author: koulei
 * @Description:
 * @File: pool_test
 * @Version: 1.0.0
 * @Date: 2023/3/13 23:31
 */

package test

import (
	"fmt"
	"sync"
	"testing"
)

type Worker func()

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func NewPool(numWorkers int) *Pool {
	p := &Pool{
		work: make(chan Worker, 1000),
	}
	p.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go p.worker()
	}
	return p
}

func (p *Pool) worker() {
	defer p.wg.Done()
	for w := range p.work {
		w()
	}
}

func (p *Pool) Submit(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

func TestPool(t *testing.T) {
	pool := NewPool(5)
	defer pool.Shutdown()

	for i := 0; i < 10; i++ {
		id := i
		pool.Submit(func() {
			fmt.Printf("Worker %d is processing\n", id)
		})
	}
}
