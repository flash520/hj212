/**
 * @Author: koulei
 * @Description:
 * @File: buffer
 * @Version: 1.0.0
 * @Date: 2023/3/12 17:26
 */

package hj212

import (
	"sync"
)

func init() {
	BuffPool = newBuffPool()
}

type buffPool struct {
	pool sync.Pool
}

var BuffPool *buffPool

func newBuffPool() *buffPool {
	return &buffPool{
		pool: sync.Pool{
			New: createBuff,
		},
	}
}

// createBuff 创建池buff实例
func createBuff() interface{} {
	return &[128]byte{}
}

func (buff *buffPool) Get() *[128]byte {
	b := buff.pool.Get().(*[128]byte)
	return b
}

func (buff *buffPool) Put(val interface{}) {
	buff.pool.Put(val)
}
