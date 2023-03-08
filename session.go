/**
 * @Author: koulei
 * @Description:
 * @File: session
 * @Version: 1.0.0
 * @Date: 2023/3/7 11:39
 */

package hj212

import (
	"github.com/flash520/link"
	log "github.com/sirupsen/logrus"

	"github.com/flash520/hj212/consts"
	"github.com/flash520/hj212/protocol"
)

type Session struct {
	deviceID string
	session  *link.Session
	server   *Server
}

// newSession 创建session
func newSession(server *Server, sess *link.Session) *Session {
	return &Session{
		session: sess,
		server:  server,
	}
}

// ID 获取sessionID
func (session *Session) ID() uint64 {
	return session.session.ID()
}

// Send 发送消息
func (session *Session) Send(entity protocol.Entity) error {
	message := protocol.Message{
		Header: protocol.Header{},
		Body:   entity,
	}

	return session.session.Send(message)
}

// message 消息事件
func (session *Session) message(message *protocol.Message) {
	if message.Header.MN != "" {
		old := session.deviceID
		if old != "" && message.Header.MN != old {
			log.WithFields(log.Fields{
				"sessionID": session.session.ID(),
				"old":       old,
				"new":       message.Header.MN,
			}).Warn(consts.ServerName, "设备ID不一致")
		}

		session.deviceID = message.Header.MN
	}
}
