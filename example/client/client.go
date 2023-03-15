/**
 * @Author: koulei
 * @Description:
 * @File: client
 * @Version: 1.0.0
 * @Date: 2023/3/8 10:45
 */

package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", "tb.cdqidi.cn:8005")
			if err != nil {
				panic(err)
			}
			var sendCount int

			go func() {
				buff := make([]byte, 1036)
				for {
					count, _ := conn.Read(buff[:])
					fmt.Println(string(buff[:count]))
				}
			}()

			ticker := time.NewTicker(time.Microsecond * 1)
			for {
				select {
				case <-ticker.C:
					_, err := conn.Write([]byte("##0411QN=20220513162608370;ST=22;CN=2051;PW=123456;MN=JLYP01_AIR01_0001;Flag=4;CP=&&DataTime=20220513162600;a21004-Avg=4,a21004-Flag=N;a05024-Avg=63,a05024-Flag=N;a21026-Avg=6,a21026-Flag=N;a21005-Avg=0.427,a21005-Flag=N;a34004-Avg=14,a34004-Flag=N;a34002-Avg=31,a34002-Flag=N;a01007-Avg=1,a01007-Flag=N;a01008-Avg=11,a01008-Flag=N;a01001-Avg=17,a01001-Flag=N;a01002-Avg=85,a01002-Flag=N;a01006-Avg=95,a01006-Flag=N&&D641\r\n"))
					// fmt.Println(n)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					sendCount++
					fmt.Println(sendCount)
				case <-ctx.Done():
					cancel()
					return
				}
			}
		}()
		// time.Sleep(time.Millisecond * 5)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
