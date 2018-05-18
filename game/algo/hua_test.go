/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-05-18 16:53:22
 * Filename      : hua_test.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

import (
	"testing"
)

// 测试
func TestHua(t *testing.T) {
	cs := []uint32{0x33, 0x23, 0x17}
	t.Log(toHands(cs))
	cs = []uint32{0x1d, 0x4b, 0x41}
	t.Log(toHands(cs))
	cs = []uint32{0x4d, 0x4c, 0x41}
	t.Log(Hua(cs))
	cs = []uint32{0x1d, 0x2d, 0x3d}
	t.Log(Hua(cs))
	//
	cs1 := []uint32{0x19, 0x25, 0x35}
	cs2 := []uint32{0x12, 0x15, 0x45}
	t.Log(HuaCompare(cs1, cs2))
	cs1 = []uint32{0x19, 0x25, 0x35}
	cs2 = []uint32{0x12, 0x11, 0x41}
	t.Log(HuaCompare(cs1, cs2))
}
