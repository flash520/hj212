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
	ErrAlreadyRunning = errors.New("服务器已经运行中")

	// ErrInvalidMessage 无效消息格式
	ErrInvalidMessage = errors.New("无效的消息格式")

	// ErrInvalidBody 无效消息体
	ErrInvalidBody = errors.New("无效消息体")

	// ErrBodyTooLong 消息体过长
	ErrBodyTooLong = errors.New("消息体过长")

	// ErrInvalidHeader 无效消息头
	ErrInvalidHeader = errors.New("无效消息头")

	// ErrInvalidSuffixID 无效的消息尾部标识符
	ErrInvalidSuffixID = errors.New("无效的消息尾部标识符")

	// ErrInvalidPrefixID 未找到标识符
	ErrInvalidPrefixID = errors.New("无效的消息头部标识符")

	// ErrInvalidBCDTime 无效BCD时间
	ErrInvalidBCDTime = errors.New("无效BCD时间")

	// ErrInvalidCheckSum 无效消息校验和
	ErrInvalidCheckSum = errors.New("无效消息校验和")

	// ErrTypeNotRegistered 消息类型未注册
	ErrTypeNotRegistered = errors.New("消息类型未注册")

	// ErrInvalidExtraLength 附加信息长度错误
	ErrInvalidExtraLength = errors.New("附加信息长度错误")

	// ErrDecryptMessageFailed 消息解密失败
	ErrDecryptMessageFailed = errors.New("消息解密失败")

	// ErrValidHeaderJsonFailed 验证消息头json失败
	ErrValidHeaderJsonFailed = errors.New("验证消息头json失败")

	// ErrNotFoundPNUM 未找到分包信息
	ErrNotFoundPNUM = errors.New("未找到分包信息")

	// ErrNotFoundPNO 未找到分包序号
	ErrNotFoundPNO = errors.New("未找到分包序号")
)
