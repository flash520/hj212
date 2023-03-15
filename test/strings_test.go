/**
 * @Author: koulei
 * @Description:
 * @File: strings_test
 * @Version: 1.0.0
 * @Date: 2023/3/14 18:17
 */

package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	s := "DataTime=20220513162600;a21004-Avg=4,a21004-Flag=N;a05024-Avg=63,a05024-Flag=N;a21026-Avg=6,a21026-Flag=N;a21005-Avg=0.427,a21005-Flag=N;a34004-Avg=14,a34004-Flag=N;a34002-Avg=31,a34002-Flag=N;a01007-Avg=1,a01007-Flag=N;a01008-Avg=11,a01008-Flag=N;a01001-Avg=17,a01001-Flag=N;a01002-Avg=85,a01002-Flag=N;a01006-Avg=95,a01006-Flag=N"

	items := strings.Split(s, ";")
	for _, item := range items {
		parts := strings.Split(item, ",")
		for _, part := range parts {
			kv := strings.Split(part, "=")
			fmt.Println(kv[0])
		}
	}
}
