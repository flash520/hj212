/**
 * @Author: koulei
 * @Description:
 * @File: checkout
 * @Version: 1.0.0
 * @Date: 2023/3/8 16:36
 */

package protocol

// CRCCheckout CRC16消息校验
func CRCCheckout(msg interface{}) uint16 {
	var bs []rune
	switch msg.(type) {
	case string:
		bs = []rune(msg.(string))
	case []rune:
		bs = msg.([]rune)
	case []byte:
		for bb := range msg.([]byte) {
			bs = append(bs, rune(bb))
		}
	}
	var crcReg uint16 = 0xFFFF
	for i := 0; i < len(bs); i++ {
		crcReg = (crcReg >> 8) ^ (uint16(bs[i]))
		for j := 0; j < 8; j++ {
			check := crcReg & 1
			crcReg >>= 0x1
			if check == 0x1 {
				crcReg ^= 0xA001
			}
		}
	}
	return crcReg
}
