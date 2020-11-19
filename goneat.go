package goneat

import (
	// "log"
	"fmt"
	"time"
	"math/rand"
)

var numberInputs int
var numberOuputs int
var genration int
var populationSize int

var globalRand *rand.Rand

func GetGeneration() int {
	return genration
}

func DisplayInfo() {
	fmt.Printf("Number of Species: %d\n", len(speciesList))
}

// Standard corssover function with g1 being the most fit parent
func CrossOver(g1, g2 Genome) Genome {
	var child Genome

	if g1.fitness == g2.fitness {
		// log.Println("Both genomes have the same fitness.\nRedirecting to CrossOverRand function")
		return CrossOverRand(g1, g2)
	}
	if g1.fitness < g2.fitness {
		g1, g2 = g2, g1
	}
	child.nodes = make(map[int]Node)
	for innov := 0; innov < numberInputs+numberOuputs; innov++ {
		child.nodes[innov] = g1.nodes[innov]
	}
	child.connects = make(map[inOut]Connect)
	for key := range g1.connects {
		if _, ok := g2.connects[key]; ok && globalRand.Intn(2) == 0 {
			child.connects[key] = g2.connects[key]
			child.nodes[key.input] = g2.nodes[key.input]
			child.nodes[key.output] = g2.nodes[key.output]
		} else {
			child.connects[key] = g1.connects[key]
			child.nodes[key.input] = g1.nodes[key.input]
			child.nodes[key.output] = g1.nodes[key.output]
		}
	}
	return child
}

// This function is called when both parents have the same fitness
func CrossOverRand(g1, g2 Genome) Genome {
	var child Genome

	if g1.fitness != g2.fitness {
		// log.Println("Both genomes do not have the same fitness.\nRedirecting to CrossOver function")
		return CrossOver(g1, g2)
	}

	child.nodes = make(map[int]Node)
	for innov := 0; innov < numberInputs+numberOuputs; innov++ {
		child.nodes[innov] = g1.nodes[innov]
	}
	dem := getDisjointsExcessMatching(g1, g2)
	child.connects = make(map[inOut]Connect)
	for idx := 0; idx < 3; idx++ {
		for _, key := range dem[idx] {
			if _, ok := g2.connects[key]; ok && globalRand.Intn(2) == 0 {
				child.connects[key] = g2.connects[key]
				child.nodes[key.input] = g2.nodes[key.input]
				child.nodes[key.output] = g2.nodes[key.output]
			} else if _, ok := g1.connects[key]; ok {
				child.connects[key] = g1.connects[key]
				child.nodes[key.input] = g1.nodes[key.input]
				child.nodes[key.output] = g1.nodes[key.output]
			}
		}
	}
	return child
}

// Initialise the NEAT
// with the population size and number of input and output
func InitNEAT(pop, input, output int) []Genome {
	var genomes []Genome
	s := rand.NewSource(time.Now().UnixNano())
	globalRand = rand.New(s)

	//InitGraph()

	globalNodesPos = make(map[int]graphPos)
	genration = 0
	innovDB = make(map[inOut]int)
	innovMap = make(map[int]inOut)
	
	// Add an extra input for the bias node
	input++

	numberInputs = input
	numberOuputs = output
	populationSize = pop
	for i := 0; i < pop; i++ {
		nodeID = 0
		genomes = append(genomes, CreateGenome(input, output))
	}
	innovID = 0

	assignSpecies(genomes)
	assignMascots(genomes)

	return genomes
}
