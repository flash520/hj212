/**
 * @Author: koulei
 * @Description:
 * @File: header
 * @Version: 1.0.0
 * @Date: 2023/3/7 12:24
 */

package protocol

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/tidwall/gjson"

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
func (header *Header) Decode(data []byte) error {
	headerStr := string(data)

	// 匹配header正则表达式
	reg, err := regexp.Compile(`(?P<k>\w+)=(?P<v>\w+);`)
	if err != nil {
		return err
	}

	// 组装成json字符串
	headerFields := reg.ReplaceAllString(headerStr, "\"${k}\":\"${v}\",")
	headerFields = strings.TrimRight(headerFields, ",")
	jsonStr := fmt.Sprintf("{%s}\n", headerFields)

	// 验证json字符串是否正确
	if !json.Valid([]byte(jsonStr)) {
		return errors.ErrValidHeaderJsonFailed
	}

	// 解析数据到header对象
	result := gjson.Parse(jsonStr)
	temp := result.Map()
	header.ST = uint16(temp["ST"].Uint())
	header.CN = uint16(temp["CN"].Uint())
	header.MN = temp["MN"].String()
	header.PW = temp["PW"].String()
	header.Flag = byte(temp["Flag"].Int())

	timeStr := temp["QN"].String()
	timeStr = timeStr[:14] + "." + timeStr[14:]
	parseTime, err := time.Parse("20060102150405.000", timeStr)
	if err != nil {
		return err
	}
	header.QN = parseTime

	// 检查是否有分包信息
	if header.HasPacket() {
		pnum, ok := temp["PNUM"]
		if !ok {
			return errors.ErrNotFoundPNUM
		}
		header.Packet.PNUM = uint32(pnum.Uint())

		pno, ok := temp["PNO"]
		if !ok {
			return errors.ErrNotFoundPNO
		}
		header.Packet.PNO = uint32(pno.Uint())
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
