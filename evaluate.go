package goneat

import (
	"math"
)

const mutateConnectOdds int = 25
const weightChangeOdds int = 80
const perturbWeightOdds int = 90
const mutateNodeOdds int = 50
const enableOdds int = 20
const disableOdds int = 40
const crossoverOdds int = 75

const stepSize float64 = 0.1

func mutate(genome *Genome) {
	connectCand := getEnableDisable(*genome)

	if len(connectCand[0]) > 0 {
		if globalRand.Intn(101) < enableOdds {
			tmp := genome.connects[connectCand[0][globalRand.Intn(len(connectCand[0]))]]
			tmp.enable = true
			genome.connects[connectCand[0][globalRand.Intn(len(connectCand[0]))]] = tmp
		}
	}
	if len(connectCand[1]) > 0 {
		if globalRand.Intn(101) < disableOdds {
			tmp := genome.connects[connectCand[1][globalRand.Intn(len(connectCand[1]))]]
			tmp.enable= false
			genome.connects[connectCand[1][globalRand.Intn(len(connectCand[1]))]] = tmp
		}
	}
	if globalRand.Intn(101) < mutateConnectOdds {
		genome.MutateGenomeConnect()
	}
	if globalRand.Intn(101) <  mutateNodeOdds {
		genome.MutateGenomeNode()
	}
	for idx := range genome.connects {
		if globalRand.Intn(101) < weightChangeOdds {
			tmp := genome.connects[idx]
			if globalRand.Intn(101) <  perturbWeightOdds {
				tmp.weight = tmp.weight+globalRand.Float64()*stepSize*2-stepSize
			} else {
				tmp.weight = globalRand.Float64()*2-1
			}
			genome.connects[idx] = tmp
		}
	}
}

func breedSpecies(sIdx int, nextGenomes []Genome, genomes []Genome) []Genome {
	var child Genome
	// TODO create a better function to select wich genome in a species reproduce
	g1 := genomes[speciesList[sIdx].genomes[globalRand.Intn(len(speciesList[sIdx].genomes))]]
	g2 := genomes[speciesList[sIdx].genomes[globalRand.Intn(len(speciesList[sIdx].genomes))]]
	if globalRand.Intn(101) < crossoverOdds {
		child = CrossOver(g1, g2)
	} else {
		child = g1
	}
	mutate(&child)
	return append(nextGenomes, child)
}

func NextGeneration(genomes []Genome) {
	var nextGenomes []Genome

	checkSpeciesProgess(genomes)
	removeExtinctSpecies()
	setAdjAvgFitness(genomes)
	fitTotal := getTotalFitness()
	if len(speciesList) > 0 {
		for sIdx := 0; sIdx < len(speciesList); sIdx++ {
			offspring := 0
			if fitTotal > 0 {
				offspring = int(math.Floor(speciesList[sIdx].averageFitness/fitTotal*float64(populationSize-len(speciesList))))
			}
			for i := 0; i < offspring; i++ {
				nextGenomes = breedSpecies(sIdx, nextGenomes, genomes)
			}
			nextGenomes = append(nextGenomes, genomes[getStrongest(sIdx, genomes)])
		}
		for len(nextGenomes) < populationSize {
			nextGenomes = breedSpecies(globalRand.Intn(len(speciesList)), nextGenomes, genomes)
		}
	} else {
		for gIdx := 0; len(nextGenomes) < populationSize; gIdx++ {
			nextGenomes = append(nextGenomes, genomes[gIdx])
		}
	}
	assignMascots(genomes)
	genomes = nextGenomes
	assignSpecies(genomes)
	genration++
}

// TODO: IMPROVE THIS!!!
// THIS IS HORRIBLE DON'T FORGET TO CHANGE THIS PLEASE
func GetOutput(genome *Genome) []float64 {
	var outputs []float64
	var nodes []int
	var nextNodes []int
	loop := make(map[int]bool)

	for i := 0; i < numberInputs; i++ {
		nextNodes = append(nextNodes, i)
	}
	for len(nextNodes) > 0 {
		nodes = nextNodes
		nextNodes = nil
		for i := 0; i < len(nodes); i++ {
			loop[nodes[i]] = true
			for n := range genome.nodes {
				if connect, ok := genome.connects[inOut{nodes[i], n}]; ok {
					if connect.enable {
						tmp := genome.nodes[n]
						tmp.val += genome.nodes[i].val * connect.weight
						genome.nodes[n] = tmp
						if _, ok := loop[n]; !ok {
							exist := false
							for e := 0; e < len(nextNodes); e++ {
								if nextNodes[e] == n {
									exist = true
									break
								}
							}
							if !exist {
								nextNodes = append(nextNodes, n)
							}
						}
					}
				}
			}
		}
	}
	for i := numberInputs; i < numberInputs+numberOuputs; i++ {
		tmp := genome.nodes[i]
		tmp.val = sigmoid(tmp.val)
		outputs = append(outputs, tmp.val)
		genome.nodes[i] = tmp
	}
	return outputs
}
