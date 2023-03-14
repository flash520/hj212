/**
 * @Author: koulei
 * @Description:
 * @File: header
 * @Version: 1.0.0
 * @Date: 2023/3/7 12:24
 */

package protocol

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/flash520/hj212/errors"
)

type Header struct {
	QN     time.Time
	ST     uint16
	CN     uint16
	MN     string
	PW     string
	Flag   byte
	Packet *Packet
}

// Encode 包头编码
func (header *Header) Encode() ([]byte, error) {
	builder := strings.Builder{}

	// 获取QN值
	qn := header.QN.Format("20060102150405.000")
	qn = strings.ReplaceAll(qn, ".", "")

	// 获取header字符串，并写入buff
	h := fmt.Sprintf("QN=%s;ST=%d;CN=%d;PW=%s;MN=%s;Flag=%v;", qn, header.ST, header.CN, header.PW, header.MN, header.Flag)
	builder.WriteString(h)

	return []byte(builder.String()), nil
}

// Decode 包头解码
func (header *Header) Decode(data string) error {
	fields := strings.Split(data, ";")
	if len(fields) == 0 {
		return errors.ErrInvalidHeader
	}

	for _, field := range fields {
		items := strings.Split(field, "=")
		switch items[0] {
		case "QN":
			timeStr := items[1][:14] + "." + items[1][14:]
			parse, err := time.Parse("20060102150405.000", timeStr)
			if err != nil {
				return err
			}
			header.QN = parse
		case "ST":
			parseUint, err := strconv.ParseUint(items[1], 10, 16)
			if err != nil {
				return err
			}
			header.ST = uint16(parseUint)
		case "CN":
			parseUint, err := strconv.ParseUint(items[1], 10, 16)
			if err != nil {
				return err
			}
			header.CN = uint16(parseUint)
		case "PW":
			header.PW = items[1]
		case "MN":
			header.MN = items[1]
		case "Flag":
			header.Flag = items[1][0]
		case "PNUM":
			parseUint, err := strconv.ParseUint(items[1], 10, 32)
			if err != nil {
				return err
			}
			header.Packet.PNUM = uint32(parseUint)
		case "PNO":
			parseUint, err := strconv.ParseUint(items[1], 10, 32)
			if err != nil {
				return err
			}
			header.Packet.PNUM = uint32(parseUint)
		}
	}

	return nil
}

func (header *Header) GetID() string {
	return header.QN.Format("20060102150405")
}

func (header *Header) GetQN() time.Time {
	return header.QN
}

func (header *Header) GetST() uint16 {
	return header.ST
}

func (header *Header) GetCN() uint16 {
	return header.CN
}

func (header *Header) GetMN() string {
	return header.MN
}

func (header *Header) GetPW() string {
	return header.PW
}

func (header *Header) GetFlag() byte {
	return header.Flag
}

// 是否需要确认
func (header *Header) NeedAck() bool {
	return header.Flag&0x1 == 0x1
}

// HasPacket 是否有分包
func (header *Header) HasPacket() bool {
	return header.Flag&(1<<1) == 0x2
}
