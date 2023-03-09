/**
 * @Author: koulei
 * @Description:
 * @File: protocol
 * @Version: 1.0.0
 * @Date: 2023/3/7 11:09
 */

package hj212

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"

	"github.com/flash520/link"
	log "github.com/sirupsen/logrus"

	"github.com/flash520/hj212/consts"
	"github.com/flash520/hj212/errors"
	"github.com/flash520/hj212/protocol"
)

type Protocol struct {
}

// NewCodec 创建编解器
func (protocol *Protocol) NewCodec(rw io.ReadWriter) (link.Codec, error) {
	codec := &ProtocolCodec{
		w: rw,
		r: rw,

		buffReceiving: bytes.NewBuffer(nil),
	}
	codec.closer = rw.(io.Closer)

	return codec, nil
}

type ProtocolCodec struct {
	w io.Writer
	r io.Reader

	closer        io.Closer
	buffReceiving *bytes.Buffer
}

// Send 发送消息
func (codec *ProtocolCodec) Send(msg interface{}) error {
	message, ok := msg.(protocol.Message)
	if !ok {
		log.WithFields(log.Fields{"reason": errors.ErrInvalidMessage}).Error(consts.ServerName, "write message failed")
		return errors.ErrInvalidMessage
	}

	data, err := message.Encode()
	if err != nil {
		log.WithFields(log.Fields{
			"deviceID": message.Header.MN,
			"reason":   err.Error(),
		}).Error(consts.ServerName, "write message failed")
		return err
	}

	count, err := codec.w.Write(data)
	if err != nil {
		log.WithFields(log.Fields{
			"deviceID": message.Header.MN,
			"reason":   err.Error(),
		}).Error(consts.ServerName, "write message failed")
		return err
	}

	log.WithFields(log.Fields{
		"deviceID": message.Header.MN,
		"bytes":    count,
	}).Debug(consts.ServerName, "write message successful")
	return nil
}

// Receive 接收消息
func (codec *ProtocolCodec) Receive() (interface{}, error) {
	message, ok, err := codec.readFromBuff()
	if ok {
		return message, nil
	}
	if err != nil {
		return nil, err
	}

	var buff [128]byte
	for {
		count, err := io.ReadAtLeast(codec.r, buff[:], 1)
		if err != nil {
			return nil, err
		}

		codec.buffReceiving.Write(buff[:count])
		if codec.buffReceiving.Len() == 0 {
			continue
		}

		message, ok, err = codec.readFromBuff()
		if ok {
			return message, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

// readFromBuff 从buff解析数据
func (codec *ProtocolCodec) readFromBuff() (protocol.Message, bool, error) {
	if codec.buffReceiving.Len() == 0 {
		return protocol.Message{}, false, nil
	}

	data := codec.buffReceiving.Bytes()
	if len(data) < 16 {
		log.WithFields(log.Fields{
			"data":   string(data),
			"reason": errors.ErrInvalidMessage,
		}).Error(consts.ServerName, consts.ReadMessageFailed)

		return protocol.Message{}, false, errors.ErrInvalidBody
	}

	// 检查报文头部标识符
	if !bytes.HasPrefix(data, protocol.PrefixID) {
		// codec.buffReceiving.Next(2)
		log.WithFields(log.Fields{
			"data":   hex.EncodeToString(data),
			"reason": errors.ErrInvalidPrefixID,
		}).Error(consts.ServerName, consts.ReadMessageFailed)

		return protocol.Message{}, false, errors.ErrInvalidPrefixID
	}

	// 获取报文数据上度
	dataLenStr := string(data[2:6])
	dataLen, err := strconv.Atoi(dataLenStr)
	if err != nil {
		return protocol.Message{}, false, errors.ErrInvalidMessage
	}
	if len(data) < (dataLen + 12) {
		return protocol.Message{}, false, nil
	}

	// 检查报文尾部标识符
	if !bytes.HasSuffix(data[dataLen+10:dataLen+12], protocol.SuffixID) {
		codec.buffReceiving.Next(dataLen + 12)
		log.WithFields(log.Fields{
			"data":   data[dataLen+10 : dataLen+12],
			"reason": errors.ErrInvalidSuffixID,
		}).Error(consts.ServerName, consts.ReadMessageFailed)
		return protocol.Message{}, false, errors.ErrInvalidSuffixID
	}

	// 消息CRC校验
	hexStr := string(data[dataLen+6 : dataLen+10])
	decimal, err := strconv.ParseUint(hexStr, 16, 16)
	if err != nil {
		return protocol.Message{}, false, err
	}
	checkout := protocol.CRCCheckout(string(data[6 : dataLen+6]))
	if uint16(decimal) != checkout {
		return protocol.Message{}, false, errors.ErrInvalidCheckSum
	}

	// 消息解码
	message := protocol.Message{}
	if err = message.Decode(data[:dataLen+12]); err != nil {
		log.WithFields(log.Fields{
			"data":   fmt.Sprintf("0x%x", hex.EncodeToString(data[:dataLen+12])),
			"reason": err,
		}).Error(consts.ServerName, "receive message failed")
		return protocol.Message{}, false, err
	}

	codec.buffReceiving.Next(dataLen + 12)

	log.WithFields(log.Fields{
		"data": data[dataLen+10 : dataLen+12],
		"len":  len(data),
		"buff": codec.buffReceiving.Len(),
	}).Debug(consts.ServerName, "receive message successful")
	return message, true, nil
}

// Close 关闭编解器
func (codec *ProtocolCodec) Close() error {
	return codec.closer.Close()
}
