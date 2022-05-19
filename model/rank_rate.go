package model

import (
	"errors"
	"log"
	"os"
	"strconv"
)

const MaxRankRate = float64(100.0)

var (
	SRate RankRate
	ARate RankRate
	BRate RankRate
	CRate RankRate
)

func init() {
	s := os.Getenv("S_RATE")

	sRate, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalln("S rate is not a number")
	}

	SRate, err = NewRankRate(sRate)
	if err != nil {
		log.Fatalln(err.Error())
	}

	a := os.Getenv("A_RATE")

	aRate, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatalln("A rate is not a number")
	}

	ARate, err = NewRankRate(aRate)
	if err != nil {
		log.Fatalln(err.Error())
	}

	b := os.Getenv("B_RATE")

	bRate, err := strconv.ParseFloat(b, 64)
	if err != nil {
		log.Fatalln("B rate is not a number")
	}

	BRate, err = NewRankRate(bRate)
	if err != nil {
		log.Fatalln(err.Error())
	}

	CRate, err = NewRankRate(MaxRankRate - sRate - aRate - bRate)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type RankRate float64

func NewRankRate(
	value float64,
) (RankRate, error) {
	if 0 > value || value > MaxRankRate {
		return RankRate(-1), errors.New("max rate over")
	}
	return RankRate(value), nil
}
