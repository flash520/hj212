/**
 * @Author: koulei
 * @Description:
 * @File: consts
 * @Version: 1.0.0
 * @Date: 2023/3/7 22:09
 */

package protocol

var (
	// 报文前缀标识与结束标识符
	PrefixID = []byte{'#', '#'}
	SuffixID = []byte{'\r', '\n'}
)

// CNTyp 命令编码类型
type CNTyp uint16

// CN (Host command and slave req command)
const (
	// 初始化命令
	HSetTimeoutRepeats CNTyp = 1000 // 设置超时时间及重发次数
	HGetSlaveTime      CNTyp = 1011 // 提取现场机时间
	SupDate            CNTyp = 1011 // 上传现场机时间
	HSetSlaveTime      CNTyp = 1012 // 设置现场机时间
	SSetTimeReq        CNTyp = 1013 // 现场机时间校准请求
	HGetRealInterval   CNTyp = 1061 // 提取实时数据间隔
	SupRealInterval    CNTyp = 1061 // 上传实时数据间隔
	HSetRealInterval   CNTyp = 1062 // 设置实时数据间隔
	HGetMinuteInterval CNTyp = 1063 // 提取分钟数据间隔
	SupMinuteInterval  CNTyp = 1063 // 上传分钟数据间隔
	HSetMinuteInterval CNTyp = 1064 // 设置分钟数据间隔
	HSetSlaveRestart   CNTyp = 1070
	HResetSlavePasswd  CNTyp = 1072 // 设置现场机密码

	// 实时数据指令
	HGetRealTimeData  CNTyp = 2011 // 取污染物实时数据
	SupRealTimeData   CNTyp = 2011 // 上传污染物实时数据
	HStopRealTimeData CNTyp = 2012 // 停止察看污染物实时数据
	// 设备状态
	HGetDevStat  CNTyp = 2021 // 取设备运行状态数据
	SupDevStat   CNTyp = 2021 // 上传设备运行状态数据
	HStopDevStat CNTyp = 2022 // 停止察看设备运行状态
	// 日数据
	HGetDayHistory           CNTyp = 2031 // 取污染物日历史数据
	SupDayHistory            CNTyp = 2031 // 上传污染物日历史数据
	HGetDevRunTimeDayHistory CNTyp = 2041 // 取设备运行时间日历史数据
	SupDevRunTimeDayHistory  CNTyp = 2041 // 上传设备运行时间日历史数据
	// 分钟数据
	HGetMinHistory CNTyp = 2051 // 取污染物分钟数据
	SupMinHistory  CNTyp = 2051 // 上传污染物分钟数据
	// 小时数据
	HGetHourHistory CNTyp = 2061 // 取污染物小时数据
	SupHourHistory  CNTyp = 2061 // 上传污染物小时数据
	// 其它数据
	SupSCYupTime CNTyp = 2081 // 上传数采仪开机时间
	// 控制指令
	HZeroCal      CNTyp = 3011 // 零点校准量程校准
	HRealSampling CNTyp = 3012 // 即时采样
	HCleanStart   CNTyp = 3013 // 启动清洗/反吹
	HCompSampling CNTyp = 3014 // 比对采样
	HKeepSample   CNTyp = 3015 // 超标留样
	// 设备信息指令
	HGetDevID          CNTyp = 3019 // 提取设备唯一标识
	SupDevID           CNTyp = 3019 // 上传设备唯一标识
	HGetSlaveInfo      CNTyp = 3020 // 提取现场机信息
	SupSaveInfo        CNTyp = 3020 // 上传现场机信息
	HSetSlaveParameter CNTyp = 3021 // 设置现场机参数
	// 上位机与现场机交互指令
	SResponse   CNTyp = 9011 // 请求应答
	SExecResult CNTyp = 9012 // 执行结果
	HNoticeACK  CNTyp = 9013 // 通知应答
	HDataACK    CNTyp = 9014 // 数据应答
	HGetConfig  CNTyp = 9017
	HSetConfig  CNTyp = 9018
	SPutConfig  CNTyp = 9018
)

// ST 系统编码对照表
var systemCodeStr = map[uint16]string{
	21: "地表水质量监测",
	22: "空气质量监测",
	23: "声环境质量监测",
	31: "大气环境污染源",
	32: "地表水质污染源",
	99: "餐饮油烟污染源",
	91: "系统交互",
}

// execResultCodeStr 执行结果状态对照表
var execResultCodeStr = map[uint16]string{
	1:   "执行成功",
	2:   "执行失败,但不知道原因",
	3:   "命令请求条件错误",
	100: "没有数据",
}

// reqCmdReturn 请求指令响应状态对照表
var reqCmdReturn = map[uint16]string{
	1:  "准备执行请求",
	2:  "请求被拒绝",
	3:  "PW密码错误",
	4:  "MN设备编号错误",
	5:  "ST系统编码错误",
	6:  "Flag标记错误",
	7:  "QN请求唯一标识错误",
	8:  "CN指令错误",
	9:  "CRC校验错误",
	10: "未知错误",
}

// dataTag 数据标签对照表
var DataTag = map[byte]string{
	'N': "在线监控仪器仪表工作正常",
	'P': "电源故障",
	'p': "电源故障",
	'B': "监测仪器发生故障",
	'b': "监测仪器发生故障",
	'D': "数据采集通道关闭",
	'd': "数据采集通道关闭",
	'C': "监测仪处于校准状态",
	'c': "监测仪处于校准状态",
	'H': "数据超出仪器量程上限",
	'h': "数据超出仪器量程上限",
	'L': "数据低于仪器量程上限",
	'l': "数据低于仪器量程上限",
	']': "数据高于用户自定义设定的上限",
	')': "数据高于用户自定义设定的上限",
	'[': "数据低于用户自定义设定的上限",
	'(': "数据低于用户自定义设定的上限",
	'F': "在线监控仪器仪表停运",
}

// CP 指令参数对照表
var commandParams = map[string]string{
	"SystemTime":  "系统时间",
	"QN":          "请求编号",
	"QnRtn":       "请求回应码",
	"ExeRtn":      "执行结果回应代码",
	"RtdInterval": "实时数据上报间隔",
}
