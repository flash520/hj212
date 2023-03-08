/**
 * @Author: koulei
 * @Description:
 * @File: message
 * @Version: 1.0.0
 * @Date: 2023/3/7 12:23
 */

package protocol

type Message struct {
	Header Header
	Body   Entity
}

func (message *Message) Encode() ([]byte, error) {

	return nil, nil
}

func (message *Message) Decode([]byte) error {

	return nil
}
