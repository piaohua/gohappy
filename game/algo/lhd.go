/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-03-23 21:54:36
 * Filename      : lhd.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

//Lhd 比较, 0:相同,1:大,2:小
func Lhd(cs1, cs2 []uint32) uint32 {
	if len(cs1) != len(cs2) {
		return 0
	}
	if len(cs1) != 1 || len(cs2) != 1 {
		return 0
	}
	if Rank(cs1[0]) > Rank(cs2[0]) {
		return 1
	}
	if Rank(cs1[0]) < Rank(cs2[0]) {
		return 2
	}
	return 0
}

//LhdRank 大小
func LhdRank(cs []uint32) uint32 {
	if len(cs) != 1 {
		return 0
	}
	return Rank(cs[0])
}