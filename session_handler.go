/**
 * @Author: koulei
 * @Description:
 * @File: session_handler
 * @Version: 1.0.0
 * @Date: 2023/3/7 12:04
 */

package hj212

import (
	"github.com/flash520/link"
	log "github.com/sirupsen/logrus"

	"github.com/flash520/hj212/consts"
	"github.com/flash520/hj212/protocol"
)

type sessionHandler struct {
	server *Server
}

// HandleSession 消息处理器
func (handler sessionHandler) HandleSession(sess *link.Session) {
	log.WithFields(log.Fields{
		"sessionID": sess.ID(),
	}).Debug(consts.ServerName, "新会话创建成功")

	session := newSession(handler.server, sess)
	handler.server.mutex.Lock()
	handler.server.sessions[sess.ID()] = session
	handler.server.mutex.Unlock()

	sess.AddCloseCallback(nil, nil, func() {
		handler.server.handleClose(session)
	})

	for {
		msg, err := sess.Receive()
		if err != nil {
			sess.Close()
			break
		}

		message := msg.(protocol.Message)
		session.message(&message)
		handler.server.dispatchMessage(session, &message)
		continue
	}
}
