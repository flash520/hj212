/**
 * @Author: koulei
 * @Description:
 * @File: HJ_524
 * @Version: 1.0.0
 * @Date: 2023/3/9 10:26
 */

package protocol

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// type Body map[string]DataPoint

// TimeBody CP基本字段
type TimeBody struct {
	DataTime time.Time `json:"dataTime,omitempty"`
	NoTime   bool      `json:"noTime"`
}

// DataPoint 数据点
type DataPoint struct {
	Cou  *float64 `json:"cou"`  // 累计值
	Min  *float64 `json:"min"`  // 最小值
	Avg  *float64 `json:"avg"`  // 平均值
	Max  *float64 `json:"max"`  // 最大值
	Flag byte     `json:"flag"` // 数据标记
}

// Body 检测因子
type Body struct {
	A00000 *DataPoint `json:"a00000,omitempty"` // 废气 B02 立方米/秒 立方米 N6.1
	A01001 *DataPoint `json:"a01001,omitempty"` // 温度 -- 摄氏度 N3.1
	A01002 *DataPoint `json:"a01002,omitempty"` // 湿度 -- % N3.1
	A01006 *DataPoint `json:"a01006,omitempty"` // 气压 -- 千帕 N5.3
	A01007 *DataPoint `json:"a01007,omitempty"` // 风速 -- 米/秒 N4.1
	A01008 *DataPoint `json:"a01008,omitempty"` // 风向 -- [角]度 N4
	A01010 *DataPoint `json:"a01010,omitempty"` // 林格曼黑度 37 无量纲 N1
	A01011 *DataPoint `json:"a01011,omitempty"` // 烟气流速 S02 米/秒 N5.2
	A01012 *DataPoint `json:"a01012,omitempty"` // 烟气温度 S03 摄氏度 N3.1
	A01013 *DataPoint `json:"a01013,omitempty"` // 烟气压力 S08 千帕 N5.3
	A01014 *DataPoint `json:"a01014,omitempty"` // 烟气湿度 S05 % N3.1
	A01015 *DataPoint `json:"a01015,omitempty"` // 制冷温度 S06 摄氏度 N3.1
	A01016 *DataPoint `json:"a01016,omitempty"` // 烟道截面积 S07 平方米 N4.2
	A01017 *DataPoint `json:"a01017,omitempty"` // 烟气动压 S04 千帕 N5.3
	A01901 *DataPoint `json:"a01901,omitempty"` // 垃圾焚烧炉膛内焚烧平均温度 -- 摄氏度 N4.1
	A01902 *DataPoint `json:"a01902,omitempty"` // 垃圾焚烧炉膛内	DCS 温度	-- 摄氏度 N4.1
	A05001 *DataPoint `json:"a05001,omitempty"` // 二氧化碳 30 毫克/立方米 千克 N3.3
	A05002 *DataPoint `json:"a05002,omitempty"` // 甲烷 -- 纳克/立方米 克 N4.1
	A05008 *DataPoint `json:"a05008,omitempty"` // 三氯一氟甲烷 -- 毫克/立方米 千克 N7.3
	A05009 *DataPoint `json:"a05009,omitempty"` // 二氯二氟甲烷 -- 毫克/立方米 千克 N7.3
	A05013 *DataPoint `json:"a05013,omitempty"` // 三氯三氟乙烷 -- 毫克/立方米 千克 N7.3
	A19001 *DataPoint `json:"a19001,omitempty"` // 氧气含量 S01 % N3.1
	A20007 *DataPoint `json:"a20007,omitempty"` // 砷 -- 纳克/立方米 克 N1.6
	A20016 *DataPoint `json:"a20016,omitempty"` // 铍及其化合物 36 毫克/立方米 千克 N4.4
	A20025 *DataPoint `json:"a20025,omitempty"` // 镉及其化合物 33 毫克/立方米 千克 N3.4
	A20026 *DataPoint `json:"a20026,omitempty"` // 镉 -- 纳克/立方米 克 N3.3
	A20043 *DataPoint `json:"a20043,omitempty"` // 铅及其化合物 32 毫克/立方米 千克 N2.4
	A20044 *DataPoint `json:"a20044,omitempty"` // 铅 -- 纳克/立方米 克 N3.3
	A20057 *DataPoint `json:"a20057,omitempty"` // 汞及其化合物 31 毫克/立方米 千克 N4.4
	A20058 *DataPoint `json:"a20058,omitempty"` // 汞 -- 纳克/立方米 克 N3.2
	A20063 *DataPoint `json:"a20063,omitempty"` // 镍及其化合物 35 毫克/立方米 千克 N3.3
	A20091 *DataPoint `json:"a20091,omitempty"` // 锡及其化合物 34 毫克/立方米 千克 N4.3
	A21001 *DataPoint `json:"a21001,omitempty"` // 氨（氨气） 10 纳克/立方米 克 N4.3
	A21002 *DataPoint `json:"a21002,omitempty"` // 氮氧化物 03 毫克/立方米 千克 N5.1
	A21003 *DataPoint `json:"a21003,omitempty"` // 一氧化氮 -- 毫克/立方米 千克 N4.1
	A21004 *DataPoint `json:"a21004,omitempty"` // 二氧化氮 -- 毫克/立方米 千克 N4.1
	A21005 *DataPoint `json:"a21005,omitempty"` // 一氧化碳 04 毫克/立方米 千克 N3.3
	A21017 *DataPoint `json:"a21017,omitempty"` // 氰化物 07 毫克/立方米 千克 N3.3
	A21018 *DataPoint `json:"a21018,omitempty"` // 氟化物 06 毫克/立方米 千克 N2.3
	A21022 *DataPoint `json:"a21022,omitempty"` // 氯气 11 毫克/立方米 千克 N4.3
	A21024 *DataPoint `json:"a21024,omitempty"` // 氯化氢 08 毫克/立方米 千克 N4.3
	A21026 *DataPoint `json:"a21026,omitempty"` // 二氧化硫 02 毫克/立方米 千克 N5.2
	A21028 *DataPoint `json:"a21028,omitempty"` // 硫化氢 05 毫克/立方米 千克 N3.2
	A23001 *DataPoint `json:"a23001,omitempty"` // 酚类 27 毫克/立方米 千克 N3.3
	A24003 *DataPoint `json:"a24003,omitempty"` // 二氯甲烷 -- 毫克/立方米 千克 N7.3
	A24004 *DataPoint `json:"a24004,omitempty"` // 三氯甲烷 -- 毫克/立方米 千克 N7.3
	A24005 *DataPoint `json:"a24005,omitempty"` // 四氯甲烷 -- 毫克/立方米 千克 N7.3
	A24006 *DataPoint `json:"a24006,omitempty"` // 二溴一氯甲烷 -- 毫克/立方米 千克 N7.3
	A24007 *DataPoint `json:"a24007,omitempty"` // 一溴二氯甲烷 -- 毫克/立方米 千克 N7.3
	A24008 *DataPoint `json:"a24008,omitempty"` // 溴甲烷 -- 毫克/立方米 千克 N7.3
	A24009 *DataPoint `json:"a24009,omitempty"` // 三溴甲烷 -- 毫克/立方米 千克 N7.3
	A24015 *DataPoint `json:"a24015,omitempty"` // 氯乙烷 -- 毫克/立方米 千克 N7.3
	A24016 *DataPoint `json:"a24016,omitempty"` // 1,1-二氯乙烷 -- 毫克/立方米 千克 N7.3
	A24017 *DataPoint `json:"a24017,omitempty"` // 1,2-二氯乙烷 -- 毫克/立方米 千克 N7.3
	A24018 *DataPoint `json:"a24018,omitempty"` // 1,1,1-三氯乙烷 -- 毫克/立方米 千克 N7.3
	A24019 *DataPoint `json:"a24019,omitempty"` // 1,1,2-三氯乙烷 -- 毫克/立方米 千克 N7.3
	A24020 *DataPoint `json:"a24020,omitempty"` // 1,1,2,2-四氯乙烷 -- 毫克/立方米 千克 N7.3
	A24027 *DataPoint `json:"a24027,omitempty"` // 1,2-二氯丙烷 -- 毫克/立方米 千克 N7.3
	A24034 *DataPoint `json:"a24034,omitempty"` // 1,2-二溴乙烷 -- 毫克/立方米 千克 N7.3
	A24036 *DataPoint `json:"a24036,omitempty"` // 环己烷 -- 毫克/立方米 千克 N7.3
	A24042 *DataPoint `json:"a24042,omitempty"` // 正己烷 -- 毫克/立方米 千克 N7.3
	A24043 *DataPoint `json:"a24043,omitempty"` // 正庚烷 -- 毫克/立方米 千克 N7.3
	A24046 *DataPoint `json:"a24046,omitempty"` // 氯乙烯 29 毫克/立方米 千克 N4.3
	A24047 *DataPoint `json:"a24047,omitempty"` // 1,1-二氯乙烯 -- 毫克/立方米 千克 N7.3
	A24049 *DataPoint `json:"a24049,omitempty"` // 三氯乙烯 -- 毫克/立方米 千克 N7.3
	A24050 *DataPoint `json:"a24050,omitempty"` // 四氯乙烯 -- 毫克/立方米 千克 N7.3
	A24053 *DataPoint `json:"a24053,omitempty"` // 丙烯 -- 毫克/立方米 千克 N7.3
	A24054 *DataPoint `json:"a24054,omitempty"` // 1,3-二氯丙烯 -- 毫克/立方米 千克 N7.3
	A24072 *DataPoint `json:"a24072,omitempty"` // 1,4-二恶烷 -- 毫克/立方米 千克 N7.3
	A24078 *DataPoint `json:"a24078,omitempty"` // 1,3-丁二烯 -- 毫克/立方米 千克 N7.3
	A24087 *DataPoint `json:"a24087,omitempty"` // 碳氢化合物 25 毫克/立方米 千克 N5.2
	A24088 *DataPoint `json:"a24088,omitempty"` // 非甲烷总烃 -- 毫克/立方米 千克 N7.3
	A24099 *DataPoint `json:"a24099,omitempty"` // 氯甲烷 -- 毫克/立方米 千克 N7.3
	A24110 *DataPoint `json:"a24110,omitempty"` // 反式-1,2-二氯乙烯 -- 毫克/立方米 千克 N7.3
	A24111 *DataPoint `json:"a24111,omitempty"` // 顺式-1,2-二氯乙烯 -- 毫克/立方米 千克 N7.3
	A24112 *DataPoint `json:"a24112,omitempty"` // 反式-1,3-二氯丙烯 -- 毫克/立方米 千克 N7.3
	A24113 *DataPoint `json:"a24113,omitempty"` // 六氯-1,3-丁二烯 -- 毫克/立方米 千克 N7.3
	A25002 *DataPoint `json:"a25002,omitempty"` // 苯 -- 毫克/立方米 千克 N7.3
	A25003 *DataPoint `json:"a25003,omitempty"` // 甲苯 17 毫克/立方米 千克 N4.2
	A25004 *DataPoint `json:"a25004,omitempty"` // 乙苯 -- 毫克/立方米 千克 N7.3
	A25005 *DataPoint `json:"a25005,omitempty"` // 二甲苯 18 毫克/立方米 千克 N4.2
	A25006 *DataPoint `json:"a25006,omitempty"` // 1,2-二甲基苯 -- 毫克/立方米 千克 N7.3
	A25007 *DataPoint `json:"a25007,omitempty"` // 1,3-二甲基苯 -- 毫克/立方米 千克 N7.3
	A25008 *DataPoint `json:"a25008,omitempty"` // 1,4-二甲基苯 -- 毫克/立方米 千克 N7.3
	A25010 *DataPoint `json:"a25010,omitempty"` // 氯苯 23 毫克/立方米 千克 N4.3
	A25011 *DataPoint `json:"a25011,omitempty"` // 1,2-二氯苯 -- 毫克/立方米 千克 N7.3
	A25012 *DataPoint `json:"a25012,omitempty"` // 1,3-二氯苯 -- 毫克/立方米 千克 N7.3
	A25013 *DataPoint `json:"a25013,omitempty"` // 1,4-二氯苯 -- 毫克/立方米 千克 N7.3
	A25014 *DataPoint `json:"a25014,omitempty"` // 1-乙基-4-甲基苯 -- 毫克/立方米 千克 N7.3
	A25015 *DataPoint `json:"a25015,omitempty"` // 1,2,4-三氯苯 -- 毫克/立方米 千克 N7.3
	A25019 *DataPoint `json:"a25019,omitempty"` // 1,2,4-三甲基苯 -- 毫克/立方米 千克 N7.3
	A25020 *DataPoint `json:"a25020,omitempty"` // 1,2,3-三甲基苯 -- 毫克/立方米 千克 N7.3
	A25021 *DataPoint `json:"a25021,omitempty"` // 1,3,5-三甲基苯 -- 毫克/立方米 千克 N7.3
	A25023 *DataPoint `json:"a25023,omitempty"` // 硝基苯 22 毫克/立方米 千克 N3.4
	A25038 *DataPoint `json:"a25038,omitempty"` // 乙烯基苯 -- 毫克/立方米 千克 N7.3
	A25044 *DataPoint `json:"a25044,omitempty"` // 苯并[a]芘 20 毫克/立方米 千克 N4.3
	A25072 *DataPoint `json:"a25072,omitempty"` // 四氢呋喃 -- 毫克/立方米 千克 N7.3
	A26001 *DataPoint `json:"a26001,omitempty"` // 苯胺类 21 毫克/立方米 千克 N4.3
	A29017 *DataPoint `json:"a29017,omitempty"` // 乙酸乙酯 -- 毫克/立方米 千克 N7.3
	A29026 *DataPoint `json:"a29026,omitempty"` // 乙酸乙烯酯 -- 毫克/立方米 千克 N7.3
	A30001 *DataPoint `json:"a30001,omitempty"` // 甲醇 28 毫克/立方米 千克 N4.3
	A30008 *DataPoint `json:"a30008,omitempty"` // 异丙醇 -- 毫克/立方米 千克 N7.3
	A30022 *DataPoint `json:"a30022,omitempty"` // 硫醇 13 毫克/立方米 千克 N4.3
	A31001 *DataPoint `json:"a31001,omitempty"` // 甲醛 19 毫克/立方米 千克 N3.3
	A31002 *DataPoint `json:"a31002,omitempty"` // 乙醛 26 毫克/立方米 千克 N3.4
	A31024 *DataPoint `json:"a31024,omitempty"` // 丙酮 -- 毫克/立方米 千克 N7.3
	A31025 *DataPoint `json:"a31025,omitempty"` // 2-丁酮 -- 毫克/立方米 千克 N7.3
	A31030 *DataPoint `json:"a31030,omitempty"` // 甲基异丁基甲酮 -- 毫克/立方米 千克 N7.3
	A34001 *DataPoint `json:"a34001,omitempty"` // 总悬浮颗粒物 TSP -- 纳克/立方米 克 N4.3
	A34002 *DataPoint `json:"a34002,omitempty"` // 可吸入颗粒物 PM10	-- 纳克/立方米 克 N3.3
	A34004 *DataPoint `json:"a34004,omitempty"` // 细微颗粒物 PM2.5 -- 纳克/立方米 克 N3.3
	A34005 *DataPoint `json:"a34005,omitempty"` // 亚微米颗粒物 PM1.0 -- 纳克/立方米 克 N3.3
	A34011 *DataPoint `json:"a34011,omitempty"` // 降尘 -- 吨/平方千米•月
	A34013 *DataPoint `json:"a34013,omitempty"` // 烟尘 01 毫克/立方米 千克 N4
	A34017 *DataPoint `json:"a34017,omitempty"` // 炭黑尘 -- 毫克/立方米 千克 N4
	A34038 *DataPoint `json:"a34038,omitempty"` // 沥青烟 09 毫克/立方米 千克 N4.3
	A34039 *DataPoint `json:"a34039,omitempty"` // 硫酸雾 14 毫克/立方米 千克 N4.3
	A34040 *DataPoint `json:"a34040,omitempty"` // 铬酸雾 15 毫克/立方米 千克 N2.3
	A99010 *DataPoint `json:"a99010,omitempty"` // 丙烯腈 -- 毫克/立方米 千克 N7.3
	A99049 *DataPoint `json:"a99049,omitempty"` // 光气 24 毫克/立方米 千克 N3.3
	A99051 *DataPoint `json:"a99051,omitempty"` // 二硫化碳 12 毫克/立方米 千克 N4.3
}

// MonitorAtmospheric 气监测因子编码表
type MonitorAtmospheric struct {
	TimeBody
	Body
}

func (entity *MonitorAtmospheric) Encode() ([]byte, error) {
	// 将数据反序列化到临时的map中
	temp := make(map[string]DataPoint)
	marshal, err := json.Marshal(entity.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(marshal, &temp); err != nil {
		return nil, err
	}

	var builder strings.Builder
	builder.WriteString("CP=&&")
	timeStr := fmt.Sprintf("DataTime=%s;", entity.DataTime.Format("20060102150405"))
	builder.WriteString(timeStr)

	// 检查map中各个字段，仅编码有值的字段
	for key, point := range temp {
		format := "%s-Avg=%v,%s-Flag=%s;"
		if point.Avg != nil {
			fields := fmt.Sprintf(format, key, *point.Avg, key, string([]byte{point.Flag}))
			builder.WriteString(fields)
		}
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
func (entity *MonitorAtmospheric) Decode(data []byte) error {
	dataStr := string(data)
	compile := regexp.MustCompile(`.*CP=&&(.*)&&.*`)
	bodyStr := compile.ReplaceAllString(dataStr, "${1}")

	// 解析数据到json字符串
	kvs := strings.Split(bodyStr, ";")

	var dataTimeMatched bool
	body := make(map[string]DataPoint)
	for _, kv := range kvs {
		fieldCompile := regexp.MustCompile(`(\w+)-Avg=(\d+),(\w+)-Flag=(\w+)`)
		subMatch := fieldCompile.FindAllStringSubmatch(kv, -1)
		for _, field := range subMatch {
			if len(field) != 5 {
				continue
			}
			num, err := strconv.ParseFloat(field[2], 64)
			if err != nil {
				return err
			}

			body[strings.ToUpper(field[1])] = DataPoint{
				Avg:  &num,
				Flag: field[4][0],
			}
		}

		// 如果
		if len(subMatch) > 0 || subMatch != nil {
			continue
		}

		// dataTimeMatched 解析dataTime字段
		if dataTimeMatched {
			continue
		}

		if strings.Contains(kv, "DataTime") {
			split := strings.Split(kv, "=")
			if len(split) == 2 {
				t, err := time.Parse("20060102150405", split[1])
				if err != nil {
					return err
				}
				entity.DataTime = t
				entity.NoTime = false
				// 标识dataTime已经解析了，后续不再执行此解析逻辑
				dataTimeMatched = true
			}
		} else {
			entity.NoTime = false
		}
	}
	if len(body) > 0 {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(b, &entity.Body); err != nil {
			return err
		}
	}

	return nil
}
