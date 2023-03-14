/**
 * @Author: koulei
 * @Description:
 * @File: HJ_524
 * @Version: 1.0.0
 * @Date: 2023/3/9 10:26
 */

package protocol

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/flash520/hj212/consts"
)

// type Body map[string]DataPoint

// TimeBody CP基本字段
type TimeBody struct {
	DataTime time.Time `json:"dataTime,omitempty"`
	NoTime   bool      `json:"noTime"`
}

// DataPoint 数据点
type DataPoint struct {
	Cou  float64 `json:"cou"`  // 累计值
	Min  float64 `json:"min"`  // 最小值
	Avg  float64 `json:"avg"`  // 平均值
	Max  float64 `json:"max"`  // 最大值
	Flag byte    `json:"flag"` // 数据标记
}

// Body 检测因子
type Body map[string]DataPoint

// MonitorAtmospheric 气监测因子编码表
type MonitorAtmospheric struct {
	TimeBody
	Body
}

func (entity *MonitorAtmospheric) Encode() ([]byte, error) {
	var builder strings.Builder
	builder.WriteString("CP=&&")
	timeStr := fmt.Sprintf("DataTime=%s;", entity.DataTime.Format("20060102150405"))
	builder.WriteString(timeStr)

	// 检查map中各个字段，仅编码有值的字段
	format := "%s-Avg=%v,%s-Flag=%s;"
	for key, point := range entity.Body {
		fields := fmt.Sprintf(format, key, point.Avg, key, string([]byte{point.Flag}))
		builder.WriteString(fields)
	}

	// 去除CP最后一个字段后面的";"字符
	s := builder.String()
	l := builder.Len() - 1
	builder.Reset()
	builder.WriteString(s[:l])

	// 添加CP结束标识符
	builder.WriteString("&&")

	return []byte(builder.String()), nil
}

// Decode 数据解码
func (entity *MonitorAtmospheric) Decode(data string) error {
	// 初始化entity.body
	entity.Body = make(map[string]DataPoint)

	// 提取CP报文体并重新赋值给data
	start := strings.Index(data, "CP=&&") + len("CP=&&")
	end := strings.LastIndex(data, "&&")
	if start < 0 || end < 0 || end <= start {
		log.WithFields(log.Fields{}).Warn(consts.ServerName, "No data found in CP")
		return nil
	}
	data = data[start:end]

	var items []string
	if len(data) > 0 {
		items = strings.Split(data, ";")
	}

	var findTime bool
	for _, item := range items {
		parts := strings.Split(item, ",")

		// 解析DataTime
		if len(parts) == 1 && strings.Contains(parts[0], "DataTime") {
			kv := strings.Split(parts[0], "=")
			t, err := time.ParseInLocation("20060102150405", kv[1], time.Local)
			if err != nil {
				return err
			}
			entity.DataTime = t
			entity.NoTime = false
			findTime = true
			continue
		}

		// 解析污染物数据
		var key string
		var points DataPoint
		for _, part := range parts {
			kv := strings.Split(part, "=")
			if len(kv) == 2 {
				// 提取污染物检测因子名称和度量标记
				keys := strings.Split(kv[0], "-")
				var num float64
				if keys[1] != "Flag" {
					// 转换污染物值为浮点数
					n, err := strconv.ParseFloat(kv[1], 64)
					if err != nil {
						return err
					}
					num = n
				}

				if len(keys) == 2 {
					key = keys[0]
					switch keys[1] {
					case "Avg":
						points.Avg = num
					case "Cou":
						points.Cou = num
					case "Min":
						points.Min = num
					case "Max":
						points.Max = num
					case "Flag":
						points.Flag = kv[1][0]
					}
				}
			}
		}
		entity.Body[key] = points
	}

	// 更新NoTime字段
	if !findTime {
		entity.NoTime = true
	}

	return nil
}
