package eago

import (
	"fmt"
	"testing"
)

func Compare(d1, d2 *DNA) int {
	// higher the fitness, the better
	switch {
	case d1.Fitness() > d2.Fitness():
		return 1
	case d1.Fitness() < d2.Fitness():
		return -1
	default:
		return 0
	}
}

func Eval(d *DNA) float64 {
	gene := d.Gene()
	size := d.Size()
	score := 0.0
	for i := 0; i < size; i++ {
		bit := string(gene[i])
		if bit == "1" {
			score++
		}
	}
	return score
}

func TestGA(t *testing.T) {
	// new config
	c := &Config{
		DNALen:       30,
		PopSize:      100,
		NumGen:       500,
		MutationRate: 0.5,
		Compare:      Compare,
		Evaluate:     Eval,
		Select:       TSelect,
		Crossover:    Crossover1P,
	}

	ga := NewGA(c)
	ga.Run()

	best := ga.Best()
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %d\n", int(best.Fitness()))
}

func TestES(t *testing.T) {

}
