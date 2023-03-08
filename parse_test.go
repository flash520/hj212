/**
 * @Author: koulei
 * @Description:
 * @File: parse_test
 * @Version: 1.0.0
 * @Date: 2023/3/8 21:45
 */

package hj212

import (
	"fmt"
	"strings"
	"testing"
)

type Entity struct {
	QN   string // 采样时间
	ST   string // 系统编码
	CN   string // 命令编码
	PW   string // 密码
	MN   string // 设备唯一标识符
	Flag string // 数据段加密方式
	CP   string // 数据内容
}

type DataPoint struct {
	Avg  string `json:"avg"`  // 平均值
	Flag string `json:"flag"` // 状态标志位
}

type Body struct {
	DataTime string               // 数据时间
	Points   map[string]DataPoint // 测点数据
}

func TestParse(t *testing.T) {
	msg := "##0411QN=20220513162608370;ST=22;CN=2051;PW=123456;MN=JLYP01_AIR01_0001;Flag=4;CP=&&DataTime=20220513162600;a21004-Avg=4,a21004-Flag=N;a05024-Avg=63,a05024-Flag=N;a21026-Avg=6,a21026-Flag=N;a21005-Avg=0.427,a21005-Flag=N;a34004-Avg=14,a34004-Flag=N;a34002-Avg=31,a34002-Flag=N;a01007-Avg=1,a01007-Flag=N;a01008-Avg=11,a01008-Flag=N;a01001-Avg=17,a01001-Flag=N;a01002-Avg=85,a01002-Flag=N;a01006-Avg=95,a01006-Flag=N&&D641"

	entity, body, err := ParseMsg(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Entity: %#v\n", entity)
	fmt.Printf("Body: %#v\n", body)
}

func ParseMsg(msg string) (*Entity, *Body, error) {
	index := strings.Index(msg, "CP")

	_ = msg[6:index]

	return nil, nil, nil
}
