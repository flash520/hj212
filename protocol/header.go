/**
 * @Author: koulei
 * @Description:
 * @File: header
 * @Version: 1.0.0
 * @Date: 2023/3/7 12:24
 */

package protocol

import "time"

type Header struct {
	QN     time.Time
	ST     uint16
	CN     uint16
	MN     string
	PW     string
	Flag   byte
	Packet *Packet
}
