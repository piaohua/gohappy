/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-03-23 21:54:36
 * Filename      : san.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

//大三公 相同的三张公仔牌 大三公牌型 9倍分 KKK、QQQ、JJJ
//小三公 相同的三张数字牌 小三公牌型 7倍分 101010、999、222、AAA
//混三公 公仔牌的差异配合 混三公牌型 5倍分 KQJ、QQJ、JJK
//特性数 肆意8点、9点的牌型 3倍分 K108、A25、K27、432
//单牌 肆意0－7点的牌型 1倍分 QJ10、A38、678、K107

const (
	San0 uint32 = iota + 0x00
	San1
	San2
	San3
	San4
	San5
	San6
	San7
	San8
	San9
	Gong1 //混三公
	Gong2 //小三公
	Gong3 //大三公
)

//San 三公点数
func San(cs []uint32) (i uint32) {
	if len(cs) != 3 {
		return
	}
	r0 := Rank(cs[0])
	r1 := Rank(cs[1])
	r2 := Rank(cs[2])
	if r0 == r1 && r1 == r2 {
		switch r0 {
		case Jack, Queen, King:
			i = Gong3
		default:
			i = Gong2
		}
		return
	}
	if r0 >= Jack && r1 >= Jack && r2 >= Jack {
		i = Gong1
		return
	}
	i = (Trunc(r0) + Trunc(r1) + Trunc(r2)) % 10
	return
}

//SanMultiple 积分倍数
func SanMultiple(n uint32) uint32 {
	switch n {
	case San0, San1, San2, San3, San4, San5, San6:
		return 1
	case San7:
		return 2
	case San8, San9:
		return 3
	case Gong1:
		return 5
	case Gong2:
		return 7
	case Gong3:
		return 9
	}
	return 1
}
