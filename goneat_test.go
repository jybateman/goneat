package goneat

import (
	"testing"
)

func TestGenomeCreation(t *testing.T) {
	g := InitNEAT(1, 10, 5)
	ninput := 0
	nhidden := 0
	noutput := 0

	for _, node := range g[0].nodes {
		switch node.types {
		case INPUT:
			ninput++
		case HIDDEN:
			nhidden++
		case OUTPUT:
			noutput++
		}
	}
	if ninput != 11 || nhidden!= 0 || noutput != 5 {
		t.Error("Expected 11 0 5, got ", ninput, nhidden, noutput)
	}
	// for i := range g.input {
	// 	if i > 9 {
	// 		t.Errorf("Expected id to be less than 10, got %d", i)
	// 	}
	// }
	// for i := range g.output {
	// 	if i < 10 {
	// 		t.Errorf("Expected id to be greater than 9, got %d", i)
	// 	}
	// }
}

func TestGenomeMutateConnectInputOutput(t *testing.T) {
	InitNEAT(1, 10, 5)
	g := CreateGenome(10, 5)
	for i := 0; i < 10; i++ {
		g.MutateGenomeConnect()
		if len(g.connects) != i+1 {
			t.Errorf("Expected %d, got %d", i+1, len(g.connects))
		}
	}
	for io := range g.connects {
		if io.input != g.connects[io].input || io.output != g.connects[io].output {
			t.Error("Connection Key/Map error:", io.input, g.connects[io].input, io.output, g.connects[io].output)
		}
		if g.nodes[io.input].types == OUTPUT || g.nodes[io.output].types == INPUT {
			t.Error("Connection Input/Output error:", io.input, io.output)
		}
	}
}

func TestGenomeMutateNodeInputOutput(t *testing.T) {
	InitNEAT(1, 10, 5)
	g := CreateGenome(10, 5)
	for i := 1; i < 10; i++ {
		g.MutateGenomeConnect()
		g.MutateGenomeNode()
		if len(g.nodes) != 16+i {
			t.Errorf("Expected %d, got %d", i+16, len(g.nodes))
		}
	}
}

func TestInnov(t *testing.T) {
	g := InitNEAT(2, 2, 1)
	g[0].MutateGenomeConnect()
	g[0].MutateGenomeConnect()
	g[0].MutateGenomeConnect()
	g[1].MutateGenomeConnect()
	g[1].MutateGenomeConnect()
	g[1].MutateGenomeConnect()
	if len(innovDB) != 3 || innovID != 3 || len(innovMap) != 3 {
		t.Error("Innovation global variable error, expected 3 3 3, got:", len(innovDB), innovID, len(innovMap))
	}
}

func TestDisjointExcessMatching(t *testing.T) {
	g := InitNEAT(2, 2, 1)
	addConnect(&g[0], inOut{input:0, output:2}, -2)
	addConnect(&g[1], inOut{input:1, output:2}, -2)
	dem := getDisjointsExcessMatching(g[0], g[1])
	if len(dem[0]) != 1 || len(dem[1]) != 1 || len(dem[2]) != 0 {
		t.Error("Expected 1 1 0, got ", len(dem[0]), len(dem[1]), len(dem[2]))
	}

	addConnect(&g[0], inOut{input:1, output:2}, -2)
	addNode(&g[1], HIDDEN)
	addConnect(&g[1], inOut{input:1, output:3}, -2)
	addConnect(&g[1], inOut{input:3, output:2}, -2)

	dem = getDisjointsExcessMatching(g[0], g[1])
	if len(dem[0]) != 1 || len(dem[1]) != 2 || len(dem[2]) != 1 {
		t.Error("Expected 1 2 1, got ", len(dem[0]), len(dem[1]), len(dem[2]))
	}

	addNode(&g[0], HIDDEN)
	addConnect(&g[0], inOut{input:0, output:5}, -2)
	addConnect(&g[0], inOut{input:5, output:2}, -2)
	dem = getDisjointsExcessMatching(g[1], g[0])
	if len(dem[0]) != 3 || len(dem[1]) != 2  || len(dem[2]) != 1 {
		t.Error("Expected 3 2 1, got ", len(dem[0]), len(dem[1]), len(dem[2]))
	}

	dem = getDisjointsExcessMatching(g[1], g[1])
	if len(dem[0]) != 0 || len(dem[1]) != 0 || len(dem[2]) != 3 {
		t.Error("Expected 0 0 3, got ", len(dem[0]), len(dem[1]), len(dem[2]))
	}

	dem = getDisjointsExcessMatching(g[0], g[0])
	if len(dem[0]) != 0 || len(dem[1]) != 0 || len(dem[2]) != 4 {
		t.Error("Expected 0 0 4, got ", len(dem[0]), len(dem[1]), len(dem[2]))
	}

	g = InitNEAT(2, 2, 1)
	addConnect(&g[0], inOut{input:0, output:2}, -2)
	addConnect(&g[1], inOut{input:1, output:2}, -2)
	addConnect(&g[0], inOut{input:1, output:2}, -2)
	addNode(&g[0], HIDDEN)
	addConnect(&g[0], inOut{input:0, output:5}, -2)
	addConnect(&g[0], inOut{input:5, output:2}, -2)
	dem = getDisjointsExcessMatching(g[0], g[1])
	if len(dem[0]) != 1 || len(dem[1]) != 2 || len(dem[2]) != 1 {
		t.Error("Expected 1 2 1, got ", len(dem[0]), len(dem[1]), len(dem[2]))
	}
}

func TestCrossover(t *testing.T) {

}


func TestCrossoverRand(t *testing.T) {

}
