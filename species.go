package goneat

const speciesThreshold float64 = 3.0

const extinctSpecies int = 15

var speciesList []Species

type Species struct {
	stall int
	mascot Genome
	maxFitness float64
	genomes []int
	averageFitness float64
	name string
	protect bool
}

func GetSpecies() []Species {
	return speciesList
}

// returns the fittest genome in a species (sIdx)
func getStrongest(sIdx int, genomes []Genome) int {
	strong := 0
	fitness := 0
	for i := range speciesList[sIdx].genomes {
		if genomes[speciesList[sIdx].genomes[i]].fitness > fitness {
			fitness = genomes[speciesList[sIdx].genomes[i]].fitness
			strong = speciesList[sIdx].genomes[i]
		}
	}
	return strong
}

func getTotalFitness() float64 {
	var total float64

	for sIdx := 0; sIdx < len(speciesList); sIdx++ {
		total += speciesList[sIdx].averageFitness
	}
	return total
}

// sets the adjusted for each genome and average fitness for the species
func setAdjAvgFitness(genomes []Genome) {
	for sIdx := 0; sIdx < len(speciesList); sIdx++ {
		total := 0.0
		for gIdx := 0; gIdx < len(speciesList[sIdx].genomes); gIdx++ {
			genomes[speciesList[sIdx].genomes[gIdx]].adjFitness = float64(genomes[speciesList[sIdx].genomes[gIdx]].fitness) / float64(len(speciesList[sIdx].genomes))
			total += genomes[speciesList[sIdx].genomes[gIdx]].adjFitness
		}
		speciesList[sIdx].averageFitness = total / float64(len(speciesList[sIdx].genomes))
	}
}

// Remove species that have not made any progress for extinctSpecies
func removeExtinctSpecies() {
	for sIdx := 0; sIdx < len(speciesList); sIdx++ {
		if speciesList[sIdx].stall > extinctSpecies || speciesList[sIdx].genomes == nil || len(speciesList[sIdx].genomes) < 1 {
			speciesList = append(speciesList[:sIdx], speciesList[sIdx+1:]...)
			sIdx--
		}
	}
}

// Updates the stall counter for each species
func checkSpeciesProgess(genomes []Genome) {
	for sIdx := 0; sIdx < len(speciesList); sIdx++ {
		progress := false
		for gIdx := 0; gIdx < len(speciesList[sIdx].genomes); gIdx++ {
			if genomes[speciesList[sIdx].genomes[gIdx]].adjFitness > speciesList[sIdx].maxFitness {
				speciesList[sIdx].maxFitness = genomes[speciesList[sIdx].genomes[gIdx]].adjFitness
				progress = true
			}
		}
		if progress == false {
			speciesList[sIdx].stall++
		} else {
			speciesList[sIdx].stall = 0
		}
	}
}

func assignMascots(genomes []Genome) {
	for sIdx := 0; sIdx < len(speciesList); sIdx++ {
		speciesList[sIdx].mascot = genomes[speciesList[sIdx].genomes[globalRand.Intn(len(speciesList[sIdx].genomes))]]
	}
}

func assignSpecies(genomes []Genome) {
	var sIdx int

	for sIdx = 0; sIdx < len(speciesList); sIdx++ {
		speciesList[sIdx].genomes = nil
	}
	for gIdx := 0; gIdx < len(genomes); gIdx++ {
		found := false
		for sIdx = 0; sIdx < len(speciesList); sIdx++ {
			if getDistance(speciesList[sIdx].mascot, genomes[gIdx]) < speciesThreshold {
				found = true
				break
			}
		}
		if !found {
			speciesList = append(speciesList, Species{mascot: genomes[gIdx], maxFitness: 0.0, stall: 0})
			speciesList[sIdx].genomes = append(speciesList[sIdx].genomes, gIdx)
		} else {
			speciesList[sIdx].genomes = append(speciesList[sIdx].genomes, gIdx)
		}
	}
}
