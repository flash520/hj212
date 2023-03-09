/**
 * @Author: koulei
 * @Description:
 * @File: errors
 * @Version: 1.0.0
 * @Date: 2023/3/7 13:17
 */

package errors

import "errors"

var (
	// ErrAlreadyRunning 服务器已在运行中
	ErrAlreadyRunning = errors.New("server already running")

	// ErrInvalidMessage 无效消息格式
	ErrInvalidMessage = errors.New("invalid message format")

	// ErrInvalidBody 无效消息体
	ErrInvalidBody = errors.New("invalid message body")

	// ErrBodyTooLong 消息体过长
	ErrBodyTooLong = errors.New("message body too long")

	// ErrInvalidHeader 无效消息头
	ErrInvalidHeader = errors.New("invalid message header")

	// ErrInvalidSuffixID 无效的消息尾部标识符
	ErrInvalidSuffixID = errors.New("invalid message suffix")

	// ErrInvalidPrefixID 无效的消息头
	ErrInvalidPrefixID = errors.New("invalid message prefix")

	// ErrInvalidBCDTime 无效BCD时间
	ErrInvalidBCDTime = errors.New("invalid BCD time")

	// ErrInvalidCheckSum 无效消息校验和
	ErrInvalidCheckSum = errors.New("invalid message checksum")

	// ErrTypeNotRegistered 消息类型未注册
	ErrTypeNotRegistered = errors.New("message type not registered")

	// ErrInvalidExtraLength 附加信息长度错误
	ErrInvalidExtraLength = errors.New("extra message length error")

	// ErrDecryptMessageFailed 消息解密失败
	ErrDecryptMessageFailed = errors.New("message decryption failed")

	// ErrValidHeaderJsonFailed 验证消息头json失败
	ErrValidHeaderJsonFailed = errors.New("message header json validate failed")

	// ErrNotFoundPNUM 未找到分包信息
	ErrNotFoundPNUM = errors.New("not found multipart packet ")

	// ErrNotFoundPNO 未找到分包序号
	ErrNotFoundPNO = errors.New("not found multipart packet serial number")

	// ErrEntityNotFound 实体未找到
	ErrEntityNotFound = errors.New("not found entity")

	// ErrParseEntityFailed 解析实体错误
	ErrParseEntityFailed = errors.New("parse entity failed")
)
