/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-03-23 21:54:36
 * Filename      : ebg.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

/*
普通：全部平倍
疯狂：八点九点 2倍，二八杠 3倍，对子 4倍，对白板5倍
*/

const (
	EBG0 uint32 = iota + 0x00
	EBG1
	EBG2
	EBG3
	EBG4
	EBG5
	EBG6
	EBG7
	EBG8
	EBG9
	EBG10  //28杠
	EBGDui //对子
	BAIBAN //白板对
)

const (
	EbgNumCard = 40
)

var EbgCARDS = []uint32{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
	0x0a, 0x0a, 0x0a, 0x0a,
}

//Ebg 牌型
func Ebg(mode uint32, cs []uint32) uint32 {
	if len(cs) != 2 {
		return 0
	}
	if mode != 0 {
		if cs[0] == cs[1] {
			if cs[0] == 0x0a {
				return BAIBAN
			}
			return EBGDui
		}
		if (cs[0] == 0x02 && cs[1] == 0x08) ||
			(cs[0] == 0x08 && cs[1] == 0x02) {
			return EBG10
		}
	}
	return (cs[0] + cs[1]) % 10
}

//EbgMultiple 积分倍数
func EbgMultiple(mode, n uint32) uint32 {
	if mode != 0 {
		switch n {
		case EBG10:
			return 3
		case EBGDui:
			return 4
		case BAIBAN:
			return 5
		}
		return 1
	}
	switch n {
	case EBG8, EBG9:
		return 2
	}
	return 1
}
