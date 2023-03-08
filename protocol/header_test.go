/**
 * @Author: koulei
 * @Description:
 * @File: header_test.go
 * @Version: 1.0.0
 * @Date: 2023/3/8 23:43
 */

package protocol

import (
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	header := Header{
		Flag: 7,
	}

	mask := byte(1)
	for i := 0; i < 8; i++ {
		if header.Flag&mask == mask {
			fmt.Printf("bit: %d, mask: %04x, result: true\n", i, mask)
		} else {
			fmt.Printf("bit: %d, mask: %04x, result: false\n", i, mask)
		}
		mask = mask << 1
	}
	fmt.Println(header.Flag&0x1 == 0x1)
	fmt.Println(header.Flag&0x8 == 0x8)
}
