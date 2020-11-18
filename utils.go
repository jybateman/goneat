package goneat

import (
	"math"
)

func sigmoid(x float64) float64 {
	return 2/(1+math.Exp(-4.9*x))-1
}

// return a list of all the enable and disable connections
// connect[0][x]: enable
// connect[1][x]: disable
func getEnableDisable(genome Genome) [2][]inOut {
	var connects [2][]inOut
	for idx := range genome.connects {
		if genome.connects[idx].enable == true {
			connects[0] = append(connects[0], idx)
		} else {
			connects[1] = append(connects[1], idx)
		}
	}
	return connects
}

// Get the distance between two genomes
func getDistance(g1, g2 Genome) float64 {
	const c1 float64 = 1.0
	const c2 float64 = 1.0
	const c3 float64 = 0.4
	const nMin int = 20

	n := 1.0
	w := 0.0
	dem := getDisjointsExcessMatching(g1, g2)

	if len(g1.nodes) > nMin || len(g2.nodes) > nMin {
		n = float64(len(g1.nodes))
		if float64(len(g2.nodes)) > n {
			n = float64(len(g2.nodes))
		}
	}
	for i := 0; i < len(dem[2]); i++ {
		w += math.Abs(g1.connects[dem[2][i]].weight - g2.connects[dem[2][i]].weight)
	}
	if len(dem[2]) > 0 {
		w = w / float64(len(dem[2]))
	}
	return c1*float64(len(dem[1]))/n+c2*float64(len(dem[0]))/n+c3*w
}

// retrieve all disjoint, excess and matching connections
// disExcMatc[0]: Disjoints
// disExcMatc[1]: Excess
// disExcMatc[2]: Matching
func getDisjointsExcessMatching(g1, g2 Genome) [3][]inOut {
	var disExcMatc [3][]inOut
	len1 := len(g1.connects)
	len2 := len(g2.connects)

	for i := 0; i < innovID && (len1 > 0 || len2 > 0); i++ {
		_, ok1 := g1.connects[innovMap[i]]
		_, ok2 := g2.connects[innovMap[i]]
		if ok1 && !ok2 {
			if len2 < 1 {
				disExcMatc[1] = append(disExcMatc[1], innovMap[i])
			} else {
				disExcMatc[0] = append(disExcMatc[0], innovMap[i])
				len1--
			}
		} else if !ok1 && ok2 {
			if len1 < 1 {
				disExcMatc[1] = append(disExcMatc[1], innovMap[i])
			} else {
				disExcMatc[0] = append(disExcMatc[0], innovMap[i])
				len2--
			}
		} else if ok1 && ok2 {
			disExcMatc[2] = append(disExcMatc[2], innovMap[i])
			len1--
			len2--
		}
	}
	return disExcMatc
}

// Returns the point between two nodes for the genome graph
func getMiddlepoint(points inOut) {
	p1 := globalNodesPos[points.input]
	p2 := globalNodesPos[points.output]
	globalNodesPos[nodeID] = graphPos{x:int32((p1.x+p2.x)/2), y:int32((p1.y+p2.y)/2)+sqrSize}
}

// Add connection to innovation DB and to the given genome
func addConnect(g *Genome, io inOut, weight float64) {
	if weight == -2 {
		weight = globalRand.Float64()*2-1
	}
	if _, ok := innovDB[io]; !ok {
		innovMap[innovID] = io
		innovDB[io] = -1
		innovID++
	}
	g.connects[io] = Connect{input:io.input, output:io.output, weight:weight, enable:true}
}

// Add a node to the Genome
func addNode(g *Genome, nType int) {
	g.nodes[nodeID] = Node{types:nType}
	nodeID++
}

// Returns nodes that have not yet been mutated with a node
func getPossibleNodes(g *Genome) []inOut {
	var connects []inOut

	for key := range g.connects {
		if innovDB[key] == -1 {
			connects = append(connects, key)
		} else {
			if _, ok := g.nodes[innovDB[key]]; !ok {
				connects = append(connects, key)
			}
		}
	}
	return connects
}

// Returns a list of all possible connect
func getPossibleConnects(g *Genome) []inOut {
	var connects []inOut
	var tmp inOut

	for input := range g.nodes {
		if g.nodes[input].types != OUTPUT {
			for output := range g.nodes {
				if input != output && g.nodes[output].types != INPUT {
					tmp.input = input
					tmp.output = output
					if _, ok := g.connects[tmp]; !ok {
						connects = append(connects, tmp)
					}

				}
			}
		}
	}
	return connects
}
