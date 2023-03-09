/**
 * @Author: koulei
 * @Description:
 * @File: register
 * @Version: 1.0.0
 * @Date: 2023/3/9 00:31
 */

package protocol

import "sync"

var entityMapper = map[uint16]Entity{
	22: new(MonitorAtmospheric),
}

func RegisterEntity(typ uint16, entity Entity) {
	entityMapper[typ] = entity
}

func RemoveEntity(typ uint16) {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := entityMapper[typ]; ok {
		delete(entityMapper, typ)
	}
}
