package model

import (
	"math/rand"
	"time"
)

type Rank string

const (
	S = Rank("s")
	A = Rank("a")
	B = Rank("b")
	C = Rank("c")
)

func LotteryRankRate() Rank {
	rand.Seed(time.Now().UnixNano())
	value := rand.Float64()
	switch {
	case value < SRate.Value()/MaxRankRate:
		return S
	case SRate.Value()/MaxRankRate < value || value > (SRate.Value()+ARate.Value())/MaxRankRate:
		return A
	case (SRate.Value()+ARate.Value())/MaxRankRate < value || value > (SRate.Value()+ARate.Value()+BRate.Value())/MaxRankRate:
		return B
	default:
		return C
	}
}

func (r Rank) String() string {
	return string(r)
}

func (r Rank) Rate() float64 {
	switch r {
	case S:
		return SRate.Value()
	case A:
		return ARate.Value()
	case B:
		return BRate.Value()
	default:
		return CRate.Value()
	}
}
