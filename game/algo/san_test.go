/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-03-23 21:54:36
 * Filename      : niu.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

import (
	"testing"
)

// 测试
func TestSan(t *testing.T) {
	cs := []uint32{0x33, 0x23, 0x17}
	t.Log(San(cs))
	cs = []uint32{0x1d, 0x4b, 0x4d}
	t.Log(San(cs))
	cs = []uint32{0x28, 0x39, 0x31}
	t.Log(San(cs))
	cs = []uint32{0x39, 0x32, 0x4c}
	t.Log(San(cs))
	cs = []uint32{0x1d, 0x2d, 0x3d}
	t.Log(San(cs))
	cs = []uint32{0x11, 0x21, 0x31}
	t.Log(San(cs))
}
