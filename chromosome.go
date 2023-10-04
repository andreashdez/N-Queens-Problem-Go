package main

import (
	"math/rand"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Chromosome struct {
	positions    []int
	conflicts    []int
	conflictsSum int
	fitness      float64
}

func NewChromosome(positions []int) *Chromosome {
	conflicts := countConflicts(positions)
	conflictsSum := sumConflicts(conflicts)
	fitness := 0.0
	log.
		Debug().
		Int("conflictsSum", conflictsSum).
		Msg("chromosome conflicts sum")
	return &Chromosome{positions, conflicts, conflictsSum, fitness}
}

func (c *Chromosome) SetFitness(fitness float64) {
	c.fitness = fitness
}

func GenerateDistinctRandomValues(size int) []int {
	return rand.Perm(size)
}

func countConflicts(positions []int) []int {
	size := len(positions)
	conflicts := make([]int, size)
	for xTwo := 0; xTwo < size-1; xTwo++ {
		for xOne := xTwo + 1; xOne < size; xOne++ {
			distance := xOne - xTwo
			yOne := positions[xOne]
			yTwo := positions[xTwo]
			if diff(yOne, yTwo) == distance {
				log.
					Trace().
					Str("chromosomeOne", "("+strconv.Itoa(xOne)+","+strconv.Itoa(yOne)+")").
					Str("chromosomeTwo", "("+strconv.Itoa(xTwo)+","+strconv.Itoa(yTwo)+")").
					Msg("found conflicts")
				conflicts[xOne] += 1
				conflicts[xTwo] += 1
			}
		}
	}
	return conflicts
}

func sumConflicts(conflicts []int) int {
	conflictsSum := 0
	for _, c := range conflicts {
		conflictsSum += c
	}
	return conflictsSum / 2
}

func diff(one int, two int) int {
	if one > two {
		return one - two
	}
	return two - one
}
