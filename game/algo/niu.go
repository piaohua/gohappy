/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-03-23 21:54:36
 * Filename      : niu.go
 * Description   : 玩牌算法
 * *******************************************************/
package algo

import "sort"

const (
	HgihCard uint32 = iota + 0x00
	Niu1
	Niu2
	Niu3
	Niu4
	Niu5
	Niu6
	Niu7
	Niu8
	Niu9
	NiuNiu
	Straight
	FullHouse
	Flush
	FiveFlower
	Bomb
	StraightFlush
	FiveTiny
)

// rank
const (
	Ace uint32 = iota + 0x01
	Deuce
	Trey
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King

	RankMask uint32 = 0x0F
)

// suit
const (
	Spade   uint32 = 0x40 //黑桃
	Heart   uint32 = 0x30 //红桃
	Club    uint32 = 0x20 //梅花
	Diamond uint32 = 0x10 //方块

	SuitMask uint32 = 0xF0
)

const (
	NumCard = 52
)

func Rank(card uint32) uint32 {
	return card & RankMask
}

func Suit(card uint32) uint32 {
	return card & SuitMask
}

var NiuCARDS = []uint32{
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d,
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d,
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d,
	0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d,
}

var NIUS [][]int = [][]int{{0, 1, 2}, {0, 1, 3}, {0, 2, 3}, {1, 2, 3}, {0, 1, 4}, {0, 2, 4}, {1, 2, 4}, {0, 3, 4}, {1, 3, 4}, {2, 3, 4}}
var NIUL [][]int = [][]int{{3, 4}, {2, 4}, {1, 4}, {0, 4}, {2, 3}, {1, 3}, {0, 3}, {1, 2}, {0, 2}, {0, 1}}

//Algo []uint32{1, 5, 8, 9, K}
func Algo(mode uint32, cs []uint32) uint32 {
	if len(cs) != 5 {
		return 0
	}
	descSort(cs)
	var niu uint32 = HgihCard
	if mode != 0 {
		niu = Algo1(cs)
		if niu != 0 {
			return niu
		}
	}
	for k, v := range NIUS {
		if ((Trunc(cs[v[0]]) + Trunc(cs[v[1]]) + Trunc(cs[v[2]])) % 10) != 0 {
			continue
		}
		switch (Trunc(cs[NIUL[k][0]]) + Trunc(cs[NIUL[k][1]])) % 10 {
		case 0:
			return NiuNiu
		case 1:
			niu = max(niu, Niu1)
		case 2:
			niu = max(niu, Niu2)
		case 3:
			niu = max(niu, Niu3)
		case 4:
			niu = max(niu, Niu4)
		case 5:
			niu = max(niu, Niu5)
		case 6:
			niu = max(niu, Niu6)
		case 7:
			niu = max(niu, Niu7)
		case 8:
			niu = max(niu, Niu8)
		case 9:
			niu = max(niu, Niu9)
		}
	}
	return niu
}

//Algo1 原有特殊玩法
func Algo1(cs []uint32) uint32 {
	bomb_n := make(map[uint32]int)
	var tiny_n int
	var tiny_v uint32
	var flower int
	var ten int
	for _, v := range cs {
		bomb_n[Rank(v)] += 1
		switch Rank(v) {
		case Jack, Queen, King:
			flower++
		case Ten:
			ten++
		case Ace, Deuce, Trey, Four, Five, Six:
			tiny_n++
			tiny_v += Rank(v)
		}
	}
	if tiny_n == 5 && tiny_v <= Ten {
		return FiveTiny
	}
	niu := Algo2(cs)
	if niu == StraightFlush {
		return niu
	}
	for _, v := range bomb_n {
		if v == 4 {
			return Bomb
		}
	}
	if flower == 5 {
		return FiveFlower
	}
	return niu
}

//Algo2 新加特殊玩法
func Algo2(cs []uint32) uint32 {
	var straight bool
	var flush bool
	cards := make([]hands, len(cs))
	for k, v := range cs {
		cards[k].Suit = Suit(v)
		cards[k].Rank = Rank(v)
	}
	ascSortHands(cards)
	if cards[0].Suit == cards[1].Suit &&
		cards[1].Suit == cards[2].Suit &&
		cards[2].Suit == cards[3].Suit &&
		cards[3].Suit == cards[4].Suit {
		flush = true
	}
	if (cards[0].Rank+1) == cards[1].Rank &&
		(cards[1].Rank+1) == cards[2].Rank &&
		(cards[2].Rank+1) == cards[3].Rank &&
		(cards[3].Rank+1) == cards[4].Rank {
		straight = true
	}
	if straight && flush {
		return StraightFlush
	}
	if flush {
		return Flush
	}
	if (cards[0].Rank == cards[1].Rank &&
		cards[1].Rank == cards[2].Rank &&
		cards[3].Rank == cards[4].Rank) ||
		(cards[1].Rank == cards[2].Rank &&
			cards[2].Rank == cards[3].Rank &&
			cards[0].Rank == cards[4].Rank) ||
		(cards[2].Rank == cards[3].Rank &&
			cards[3].Rank == cards[4].Rank &&
			cards[0].Rank == cards[1].Rank) {
		return FullHouse
	}
	if straight {
		return Straight
	}
	return 0
}

//Trunc 取整
func Trunc(n uint32) uint32 {
	if Rank(n) > Ten {
		return Ten
	}
	return Rank(n)
}

//取大值
func max(n, m uint32) uint32 {
	if n > m {
		return n
	}
	return m
}

//降序排序
func descSort(cards []uint32) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] >= cards[j]
	})
}

//Equal 比较 a == b
func Equal(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	descSort(a)
	descSort(b)
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

//比较 a >= b
func Compare2(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	descSort(a)
	descSort(b)
	for i, v := range a {
		if v < b[i] {
			return false
		}
	}

	return true
}

type hands struct {
	Suit uint32
	Rank uint32
}

//降序排序
func descSortHands(cards []hands) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Rank == cards[j].Rank {
			return cards[i].Suit > cards[j].Suit
		}
		return cards[i].Rank > cards[j].Rank
	})
}

//升序排序
func ascSortHands(cards []hands) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Rank == cards[j].Rank {
			return cards[i].Suit < cards[j].Suit
		}
		return cards[i].Rank < cards[j].Rank
	})
}

//Compare 比较 a >= b (先比较牌值,牌值相同再比较花)
//同等牛的比其中牌值最大的一个,如果最大的一个牌值一样,则比花色(相同牌花色永远不同)
func Compare(a, b []uint32) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	a1 := make([]hands, len(a))
	b1 := make([]hands, len(b))
	for i, v := range a {
		a1[i].Suit = Suit(v)
		a1[i].Rank = Rank(v)
	}
	for i, v := range b {
		b1[i].Suit = Suit(v)
		b1[i].Rank = Rank(v)
	}

	descSortHands(a1)
	descSortHands(b1)
	//牌值比较
	if a1[0].Rank == b1[0].Rank {
		return a1[0].Suit > b1[0].Suit
	}
	return a1[0].Rank > b1[0].Rank
}

//Compare3 比较 a >= b (先比较牌值,牌值相同再比较花)
func Compare3(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	a1 := make([]hands, len(a))
	b1 := make([]hands, len(b))
	for i, v := range a {
		a1[i].Suit = Suit(v)
		a1[i].Rank = Rank(v)
	}
	for i, v := range b {
		b1[i].Suit = Suit(v)
		b1[i].Rank = Rank(v)
	}

	descSortHands(a1)
	descSortHands(b1)
	//牌值比较
	for i, v := range a1 {
		if v.Rank < b1[i].Rank {
			return false
		} else if v.Rank > b1[i].Rank {
			return true
		}
	}

	//牌值相同,比较花色
	for i, v := range a1 {
		if v.Suit < b1[i].Suit {
			return false
		} else if v.Suit > b1[i].Suit {
			return true
		}
	}

	return true
}

//Multiple 积分倍数
func Multiple(mode, n uint32) uint32 {
	if mode != 0 {
		return Multiple1(n)
	}
	switch n {
	case HgihCard, Niu1, Niu2, Niu3, Niu4, Niu5, Niu6, Niu7:
		return 1
	case Niu8:
		return 2
	case Niu9:
		return 3
	case NiuNiu:
		return 4
	}
	return 1
}

//Multiple1 积分倍数
func Multiple1(n uint32) uint32 {
	switch n {
	case HgihCard, Niu1:
		return 1
	case Niu2:
		return 2
	case Niu3:
		return 3
	case Niu4:
		return 4
	case Niu5:
		return 5
	case Niu6:
		return 6
	case Niu7:
		return 7
	case Niu8:
		return 8
	case Niu9:
		return 9
	case NiuNiu:
		return 10
	case Straight:
		return 11
	case FullHouse:
		return 12
	case Flush:
		return 13
	case FiveFlower:
		return 14
	case Bomb:
		return 15
	case StraightFlush:
		return 16
	case FiveTiny:
		return 17
	}
	return 1
}

//AlgoVerify []uint32{1, 5, 8, 9, K}
func AlgoVerify(cs []uint32, val uint32) bool {
	if len(cs) != 5 {
		return false
	}
	descSort(cs)
	bomb_n := make(map[uint32]int)
	var tiny_n int
	var tiny_v uint32
	var flower int
	var ten int
	for _, v := range cs {
		bomb_n[Rank(v)] += 1
		switch Rank(v) {
		case Jack, Queen, King:
			flower++
		case Ten:
			ten++
		case Ace, Deuce, Trey, Four, Five, Six:
			tiny_n++
			tiny_v += Rank(v)
		}
	}
	if tiny_n == 5 && tiny_v <= Ten && val == FiveTiny {
		return true
	}
	for _, v := range bomb_n {
		if v == 4 && val == Bomb {
			return true
		}
	}
	if flower == 5 && val == FiveFlower {
		return true
	}
	for k, v := range NIUS {
		if ((Trunc(cs[v[0]]) + Trunc(cs[v[1]]) + Trunc(cs[v[2]])) % 10) != 0 {
			continue
		}
		switch (Trunc(cs[NIUL[k][0]]) + Trunc(cs[NIUL[k][1]])) % 10 {
		case 0:
			if val == NiuNiu {
				return true
			}
		case 1:
			if val == Niu1 {
				return true
			}
		case 2:
			if val == Niu2 {
				return true
			}
		case 3:
			if val == Niu3 {
				return true
			}
		case 4:
			if val == Niu4 {
				return true
			}
		case 5:
			if val == Niu5 {
				return true
			}
		case 6:
			if val == Niu6 {
				return true
			}
		case 7:
			if val == Niu7 {
				return true
			}
		case 8:
			if val == Niu8 {
				return true
			}
		case 9:
			if val == Niu9 {
				return true
			}
		}
	}
	return false
}

//Remove 移除一个牌
func Remove(c uint32, cs []uint32) []uint32 {
	for i, v := range cs {
		if c == v {
			cs = append(cs[:i], cs[i+1:]...)
			break
		}
	}
	return cs
}

//SameCard 移除一个牌
func SameCard(cs, hs []uint32) bool {
	for _, c := range cs {
		for _, h := range hs {
			if c == h {
				return true
			}
		}
	}
	return false
}

//VerifyCard 验证手牌(设置时存在0)
func VerifyCard(cs []uint32) bool {
	for _, c := range cs {
		if c == 0 {
			continue
		}
		switch Suit(c) {
		case Club, Heart, Spade, Diamond:
		default:
			return false
		}
		switch Rank(c) {
		case Ace, Deuce, Trey, Four, Five, Six:
		case Seven, Eight, Nine, Ten, Jack, Queen, King:
		default:
			return false
		}
	}
	return true
}
