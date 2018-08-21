/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-03-23 21:54:36
 * Filename      : lhd.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

import (
	"testing"
)

// 测试
func TestLhd(t *testing.T) {
	cs1 := []uint32{0x0a}
	cs2 := []uint32{0x0a}
	t.Log(Lhd(cs1, cs2))
}