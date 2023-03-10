/**
 * @Author: koulei
 * @Description:
 * @File: message
 * @Version: 1.0.0
 * @Date: 2023/3/7 12:23
 */

package protocol

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/flash520/hj212/consts"
	"github.com/flash520/hj212/errors"
)

type Message struct {
	Header Header
	Body   Entity
}

// Encode 消息编码
func (message *Message) Encode() ([]byte, error) {

	return nil, nil
}

// Decode 消息解码
func (message *Message) Decode(data []byte) error {
	dataStr := string(data)

	// 包头解码
	headerIndex := strings.Index(dataStr, "CP")
	if err := message.Header.Decode([]byte(dataStr[6:headerIndex])); err != nil {
		log.WithFields(log.Fields{
			"data":   dataStr[6:headerIndex],
			"reason": err.Error(),
		}).Error(consts.ServerName, "parse header failed")
		return err
	}

	// 正文解码
	bodyStr := dataStr[headerIndex:]
	entity, err := message.decode(message.Header.ST, bodyStr)
	if err != nil {
		return err
	}
	message.Body = entity

	return nil
}

func (message *Message) decode(typ uint16, data string) (Entity, error) {
	entity, ok := entityMapper[typ]
	if !ok {
		return nil, errors.ErrEntityNotFound
	}

	if err := entity.Decode([]byte(data)); err != nil {
		return nil, err
	}

	return entity, nil
}
