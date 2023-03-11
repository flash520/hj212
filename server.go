/**
 * @Author: koulei
 * @Description:
 * @File: server
 * @Version: 1.0.0
 * @Date: 2023/3/7 11:07
 */

package hj212

import (
	"net"
	"runtime/debug"
	"sync"

	"github.com/flash520/link"
	log "github.com/sirupsen/logrus"

	"github.com/flash520/hj212/consts"
	"github.com/flash520/hj212/errors"
	"github.com/flash520/hj212/protocol"
)

type Server struct {
	server          *link.Server
	mutex           sync.Mutex
	handler         sessionHandler
	sessions        map[uint64]*Session
	messageHandlers sync.Map
	closeHandler    func(session *Session)
	listenAddress   string
	sendChanSize    int
}

type Option struct {
	ListenAddress string
	SendChanSize  int
	Keepalive     int
	CloseHandler  func(session *Session)
}

type MessageHandler func(session *Session, message *protocol.Message)

func NewServer(option Option) *Server {
	if option.Keepalive == 0 {
		option.Keepalive = 60
	}

	server := &Server{
		closeHandler:  option.CloseHandler,
		listenAddress: option.ListenAddress,
		sendChanSize:  option.SendChanSize,
		sessions:      make(map[uint64]*Session),
	}
	server.handler.server = server
	return server
}

func (server *Server) Run() error {
	if server.server != nil {
		return errors.ErrAlreadyRunning
	}
	if server.listenAddress == "" {
		server.listenAddress = "0.0.0.0:8192"
	}

	listener, err := net.Listen("tcp", server.listenAddress)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(consts.ServerName)
	}

	p := &Protocol{}
	server.server = link.NewServer(listener, p, server.sendChanSize, server.handler)
	log.Infof("%s %s %s", consts.ServerName, "protocol server started on ", server.listenAddress)
	return server.server.Serve()
}

// GetSession 获取session
func (server *Server) GetSession(sessionID uint64) (*Session, bool) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	session, ok := server.sessions[sessionID]
	return session, ok
}

// AddHandler 添加消息处理器
func (server *Server) AddHandler(st uint16, handler MessageHandler) {
	if handler == nil {
		return
	}
	server.messageHandlers.Store(st, handler)
}

// RemoveHandler 移除消息处理器
func (server *Server) RemoveHandler(st, cn uint16) {
	server.messageHandlers.Delete(st + cn)
}

// Stop 停止服务
func (server *Server) Stop() {
	if server.server != nil {
		server.server.Stop()
		server.server = nil
	}
}

// dispatchMessage 分派消息
func (server *Server) dispatchMessage(session *Session, message *protocol.Message) {
	log.WithFields(log.Fields{
		"deviceID": session.deviceID,
		"st":       message.Header.ST,
		"cn":       message.Header.CN,
	}).Debug(consts.ServerName, "dispatch message")

	handler, ok := server.messageHandlers.Load(message.Header.ST)
	if !ok {
		log.WithFields(log.Fields{
			"deviceID": session.deviceID,
			"st":       message.Header.ST,
			"cn":       message.Header.CN,
			"reason":   "not found message handler",
		}).Warn(consts.ServerName, "dispatch message failed")
		return
	}

	handler.(MessageHandler)(session, message)
}

// handleClose 会话结束回调
func (server *Server) handleClose(session *Session) {
	server.mutex.Lock()
	delete(server.sessions, session.session.ID())
	server.mutex.Unlock()

	if server.closeHandler != nil {
		func() {
			if err := recover(); err != nil {
				debug.PrintStack()
			}
			server.closeHandler(session)
		}()
	}

	log.WithFields(log.Fields{
		"deviceID":  session.deviceID,
		"sessionID": session.session.ID(),
	}).Warn(consts.ServerName, "session closed")
}

func (server *Server) SessionCount() uint64 {
	return uint64(len(server.sessions))
}
