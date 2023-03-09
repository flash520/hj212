/**
 * @Author: koulei
 * @Description:
 * @File: entity
 * @Version: 1.0.0
 * @Date: 2023/3/7 12:25
 */

package protocol

// Entity Body实体
type Entity interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}
