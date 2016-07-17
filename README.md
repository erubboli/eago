#EAGO
![alt text](https://github.com/jinseokYeom/eago/blob/master/eago.png "EAGO")
## Overview
EAGO is a simple framework for evolutionary computation (EC) in Go. You can use
this framework to run any kind of EC experiment, simply by two steps: configure
and run. 

## Algorithms
Following are algorithms that are available in this framework. More will be added
with future updates.
* (mu, lambda) Evolution Strategy
* (mu + lambda) Evolution Strategy
* Genetic Algorithm

## Example

```go
package main

import (
	"fmt"
    "github.com/jinseokYeom/eago"
)

// define an evaluation function
func Eval(d *eago.DNA) float64 {
	gene := d.Gene()
	score := 0.0
	for i := 0; i < 20; i++ {
		bit := string(gene[i])
		if bit == "1" {
			score++
		}
	}
	return score
}

func GeneticAlgorithm() {
	// Configure first
	c := &eago.Config{
		GeneLen:      20,                   // size of each DNA
		PopSize:      50,                   // population size (number of DNAs)
		NumGen:       100,                  // number of generations
		MutationRate: 0.3,                  // mutation rate
		Compare:      eago.DirectCompare,   // comparison method
		Evaluate:     Eval,                 // evaluation method
		Select:       eago.TSelect,         // selection method
		Crossover:    eago.UCrossover(0.3), // crossover method
	}

    // Create GA,
	ga := eago.NewGA(c)

    // ...and run!
    ga.Run()

    // Check the best DNA and its fitness.
	best := ga.Best()
	fmt.Printf("BEST: %s\n", best.Gene())
	fmt.Printf("SCORE: %d\n", int(best.Fitness()))
}
```
