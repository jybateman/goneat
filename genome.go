package goneat

import (
	"log"
)

const (
	INPUT = 0
	HIDDEN = 1
	OUTPUT = 2
)

type inOut struct {
	input int
	output int
}

// A DB of all connections as well as if said connection as been mutated with a node by giving its id
// -1 if no node mutation
var innovDB map[inOut]int

// Mapping for innovDB between the key and id
var innovMap map[int]inOut

var innovID int
var nodeID int

type Genome struct {
	nodes map[int]Node
	connects map[inOut]Connect
	fitness int
	adjFitness float64
}

type Node struct {
	id int
	val float64
	types int
}

type Connect struct {
	input int
	output int
	weight float64
	enable bool
}

// Mutate Genone Node
func (g *Genome) MutateGenomeNode() {
	if len(g.connects) < 1 {
		// log.Println("Unable to add new node")
		return
	}
	connects := getPossibleNodes(g)
	chosen := connects[globalRand.Intn(len(connects))]

	id := innovDB[inOut{chosen.input, chosen.output}]
	if id == -1 {
		innovDB[inOut{chosen.input, chosen.output}] = nodeID
		id = nodeID
		getMiddlepoint(chosen)
		nodeID++
	}
	tmp := g.connects[chosen]
	tmp.enable = false
	g.connects[chosen] = tmp
	g.nodes[id] = Node{types:HIDDEN}
	addConnect(g, inOut{chosen.input, id}, 1)
	addConnect(g, inOut{id, chosen.output}, g.connects[chosen].weight)
}

// Mutate Genome Connection
func (g *Genome) MutateGenomeConnect() {
	connects := getPossibleConnects(g)
	if len(connects) == 0 {
		// log.Println("No connections possible")
		return
	}
	chosen := connects[globalRand.Intn(len(connects))]

	addConnect(g, chosen, -2)
}

// Sets the fitness of a given genome
func (g *Genome) SetFitness(fit int) {
	g.fitness = fit
}

// Sets the inputs of a given genome
func (g *Genome) SetInputs(val []float64) {
	if len(val) != numberInputs-1 {
		log.Println("Invalid number of inputs")
		return
	}
	for i := 0; i < len(val); i++ {
		tmp := g.nodes[i]
		tmp.val = val[i]
		g.nodes[i] = tmp
	}
	tmp := g.nodes[numberInputs-1]
	tmp.val = 1.0
	g.nodes[numberInputs-1] = tmp
}

// Create a Genome with the number of inputs and outputs
func CreateGenome(input, output int) Genome {
	var genome Genome

	genome.nodes = make(map[int]Node)
	genome.connects = make(map[inOut]Connect)
	for i := 0; i < input; i++ {
		globalNodesPos[nodeID] = graphPos{x:sqrSize, y:int32(i+1)*sqrSize+sqrSize*int32(i+1)}
		addNode(&genome, INPUT)
	}
	tmp := genome.nodes[input-1]
	tmp.val = 1.0
	genome.nodes[input-1] = tmp
	for i := 0; i < output; i++ {
		globalNodesPos[nodeID] = graphPos{x:winWidth-sqrSize*2, y:int32(i+1)*sqrSize+sqrSize*int32(i+1)}
		addNode(&genome, OUTPUT)
	}
	return genome
}
