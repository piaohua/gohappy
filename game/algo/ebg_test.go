/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-03-23 21:54:36
 * Filename      : ebg.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

import (
	"testing"
)

// 测试
func TestEbg(t *testing.T) {
	cs := []uint32{0x0a, 0x0a}
	if BAIBAN != Ebg(cs) {
		t.Log("BAIBAN failed")
	} else {
		t.Log("BAIBAN successful")
	}
}