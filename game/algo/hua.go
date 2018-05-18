/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2018-05-18 16:53:03
 * Filename      : hua.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

import (
	"sort"
)

/*
4. 牌型
豹子：三张同样大小的牌。
顺金：花色相同的三张连牌。
金花：三张花色相同的牌。
顺子：三张花色不全相同的连牌。
对子：三张牌中有两张同样大小的牌。
特殊：花色不同的 235 。
单张：除以上牌型的牌。
5. 牌型的比较
1. 豹子>顺金>金花>顺子>对子>散牌。特殊>豹子。特殊<散牌。
2．牌点中，2为最小，A为最大。从大到小依次为：A、K、Q、J、10、9、8、7、6、5、4、3、2
*/

const (
	Hua0    uint32 = iota + 0x00
	DuiZi          //对子
	ShunZi         //顺子
	JinHua         //金花
	ShunJin        //顺金
	BaoZi          //豹子
)

//Hua 牌型
func Hua(cs []uint32) (i uint32) {
	hs := toHands(cs)
	i = toPoint(hs)
	return
}

func is235(hs []hands) bool {
	if len(hs) != 3 {
		return false
	}
	if hs[0].Suit == hs[1].Suit && hs[1].Suit == hs[2].Suit {
		return false
	}
	if hs[0].Rank == Five &&
		hs[1].Rank == Trey &&
		hs[2].Rank == Deuce {
		return true
	}
	return false
}

func isShun(hs []hands) bool {
	if len(hs) != 3 {
		return false
	}
	if hs[0].Rank == (hs[1].Rank+1) &&
		hs[1].Rank == (hs[2].Rank+1) {
		return true
	}
	if hs[0].Rank == Ace &&
		hs[1].Rank == 0x0d &&
		hs[2].Rank == 0x0c {
		return true
	}
	return false
}

func isPair(hs []hands) bool {
	if len(hs) != 3 {
		return false
	}
	if hs[0].Rank == hs[1].Rank ||
		hs[1].Rank == hs[2].Rank ||
		hs[0].Rank == hs[2].Rank {
		return true
	}
	return false
}

func pairVal(hs []hands) (uint32, uint32) {
	if len(hs) != 3 {
		return 0, 0
	}
	if hs[0].Rank == hs[1].Rank {
		return hs[0].Rank, hs[2].Rank
	}
	if hs[1].Rank == hs[2].Rank {
		return hs[1].Rank, hs[0].Rank
	}
	if hs[0].Rank == hs[2].Rank {
		return hs[0].Rank, hs[1].Rank
	}
	return 0, 0
}

func toPoint(hs []hands) uint32 {
	if len(hs) != 3 {
		return 0
	}
	if hs[0].Rank == hs[1].Rank && hs[1].Rank == hs[2].Rank {
		return BaoZi
	}
	if hs[0].Suit == hs[1].Suit && hs[1].Suit == hs[2].Suit {
		if isShun(hs) {
			return ShunJin
		}
		return JinHua
	}
	if isShun(hs) {
		return ShunZi
	}
	if isPair(hs) {
		return DuiZi
	}
	return 0
}

func toHands(cs []uint32) (hs []hands) {
	hs = make([]hands, len(cs))
	for i, v := range cs {
		hs[i].Suit = Suit(v)
		hs[i].Rank = Rank(v)
	}
	descSort2Hands(hs)
	return
}

//降序排序
func descSort2Hands(cards []hands) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Rank == cards[j].Rank {
			return cards[i].Suit > cards[j].Suit
		}
		//A最大
		if cards[i].Rank == Ace {
			return true
		}
		if cards[j].Rank == Ace {
			return false
		}
		return cards[i].Rank > cards[j].Rank
	})
}

//HuaCompare 比较 a >= b (比较牌值,牌值相同先开牌者输,b=先开牌者)
func HuaCompare(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	a1 := toHands(a)
	b1 := toHands(b)
	//fmt.Printf("a1 - %#v\n", a1)
	//fmt.Printf("b1 - %#v\n", b1)

	a2 := toPoint(a1)
	b2 := toPoint(b1)
	//fmt.Printf("a2 - %d\n", a2)
	//fmt.Printf("b2 - %d\n", b2)

	//235>豹子
	if a2 == BaoZi && b2 == 0 {
		if is235(b1) {
			return false
		}
	}
	if b2 == BaoZi && a2 == 0 {
		if is235(a1) {
			return false
		}
	}

	//不同牌型
	if a2 != b2 {
		return a2 > b2
	}

	//对子
	if a2 == DuiZi && b2 == DuiZi {
		a3, av3 := pairVal(a1)
		b3, bv3 := pairVal(b1)
		if a3 != b3 {
			//A最大
			if a3 == Ace {
				return true
			}
			if b3 == Ace {
				return false
			}
			return a3 > b3
		}
		//A最大
		if av3 == Ace {
			return true
		}
		if bv3 == Ace {
			return false
		}
		return av3 > bv3
	}

	//牌值比较
	for i, v := range a1 {
		if v.Rank == b1[i].Rank {
			continue
		}
		//A最大
		if v.Rank == Ace {
			return true
		}
		if b1[i].Rank == Ace {
			return false
		}
		if v.Rank < b1[i].Rank {
			return false
		} else if v.Rank > b1[i].Rank {
			return true
		}
	}
	//先开牌者输
	return true
}

//HuaMultiple 积分倍数
func HuaMultiple(n uint32) uint32 {
	switch n {
	case DuiZi:
	case ShunZi:
		return 2
	case JinHua:
		return 3
	case ShunJin:
		return 5
	case BaoZi:
		return 10
	}
	return 1
}
