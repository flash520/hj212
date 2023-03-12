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

	log "github.com/sirupsen/logrus"

	"github.com/flash520/hj212/consts"
)

func init() {
	BuffPool = newBuffPool()
}

type buffPool struct {
	pool *sync.Pool
}

var BuffPool *buffPool

func newBuffPool() *buffPool {
	return &buffPool{
		pool: &sync.Pool{
			New: createBuff,
		},
	}
}

// createBuff 创建池buff实例
func createBuff() interface{} {
	log.WithFields(log.Fields{}).Warn(consts.ServerName, "new buff created")
	return [1036]byte{}
}

func (buff *buffPool) Get() [1036]byte {
	b := buff.pool.Get().([1036]byte)
	return b
}

func (buff *buffPool) Put(val interface{}) {
	buff.pool.Put(val)
}
