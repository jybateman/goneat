package goneat

import (
	"log"
	// "strconv"

    "github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type graphPos struct {
	x int32
	y int32
}

var globalWindow *sdl.Window
var globalRenderer *sdl.Renderer
var globalNodesPos map[int]graphPos

var winWidth int32 = 1000
var winHeight int32 = 600
var sqrSize int32 = 20
var winTitle string = "Test"


// TEST
var GlobalFont *ttf.Font
// END TEST

func GetRenderer() *sdl.Renderer {
	return globalRenderer
}

func SetCustomInputCoord() {
	x := sqrSize
	y := sqrSize
	for i := 0; i < numberInputs; i++ {
		if i > 0 && i%7 == 0 {
			x = sqrSize
			y += sqrSize+2
		}
		globalNodesPos[i] = graphPos{x:x, y:y}
		x += sqrSize+2
	}
}

func DrawGraph(g *Genome) {
	globalRenderer.SetDrawColor(0, 0, 0, 255)
	globalRenderer.Clear()
	globalRenderer.SetDrawColor(255, 0, 0, 255)
	for key := range g.nodes {
		pos := globalNodesPos[key]
		if g.nodes[key].val > 0.0 {
			globalRenderer.SetDrawColor(0, 255, 0, 255)
		} else if g.nodes[key].val < 0.0 {
			globalRenderer.SetDrawColor(255, 0, 0, 255)
		} else {
			globalRenderer.SetDrawColor(255, 255, 0, 255)
		}

		globalRenderer.FillRect(&sdl.Rect{X:pos.x, Y:pos.y, H:sqrSize, W:sqrSize})

		// TEST
		// idSurf, err := GlobalFont.RenderUTF8Solid(strconv.FormatFloat(g.nodes[key].val, 'f', 2, 64), sdl.Color{255, 255, 255, 255})
		// if err != nil {
		// 	log.Fatalf("Failed to create surface: %s\n", err)
		// }
		// idText, err := globalRenderer.CreateTextureFromSurface(idSurf)
		// if err != nil {
		// 	log.Fatalf("Failed to create texture: %s\n", err)
		// }
		// rect := sdl.Rect{X:pos.x, Y:pos.y, H:sqrSize, W:sqrSize}
		// globalRenderer.Copy(idText, nil, &rect)
		// END TEST

	}
	for key := range g.connects {
		if g.connects[key].enable {
			if g.connects[key].weight >= 0 {
				globalRenderer.SetDrawColor(0, 255, 0, 255)
			} else {
				globalRenderer.SetDrawColor(255, 0, 0, 255)
			}
			p1 := globalNodesPos[key.input]
			p2 := globalNodesPos[key.output]
			globalRenderer.DrawLine(p1.x+sqrSize/2, p1.y+sqrSize/2, p2.x+sqrSize/2, p2.y+sqrSize/2)

			// TEST
			// p1 = globalNodesPos[g.connects[key].input]
			// p2 = globalNodesPos[g.connects[key].output]
			// x := int32((p1.x+p2.x)/2)
			// y := int32((p1.y+p2.y)/2)

			// idSurf, err := GlobalFont.RenderUTF8Solid(strconv.FormatFloat(g.connects[key].weight, 'f', 2, 64), sdl.Color{255, 255, 255, 255})
			// if err != nil {
			// 	log.Fatalf("Failed to create surface: %s\n", err)
			// }
			// idText, err := globalRenderer.CreateTextureFromSurface(idSurf)
			// if err != nil {
			// 	log.Fatalf("Failed to create texture: %s\n", err)
			// }
			// rect := sdl.Rect{X:x, Y:y, H:sqrSize, W:sqrSize}
			// globalRenderer.Copy(idText, nil, &rect)
			// END TEST

		}
	}
	globalRenderer.Present()
}

// Initilise the window and rederer
func InitGraph() {
    globalWindow, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
        winWidth, winHeight, sdl.WINDOW_SHOWN)
    if err != nil {
        log.Fatalf("Failed to create window: %s\n", err)
    }
    // GlobalWindow.SetWindowOpacity(0.5)
    // defer GlobalWindow.Destroy()

    globalRenderer, err = sdl.CreateRenderer(globalWindow, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        log.Fatalf("Failed to create renderer: %s\n", err)

    }
    err = globalRenderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
    if err != nil {
        log.Fatalf("Failed to set blend mode: %s\n", err)

    }

	// globalNodesPos = make(map[int]graphPos)

	// TEST
	ttf.Init()
    if err != nil {
        log.Fatalf("Failed to create font: %s\n", err)
    }

    GlobalFont, err = ttf.OpenFont("/usr/share/fonts/truetype/open-sans-elementary/OpenSans-Regular.ttf", 25)
    if err != nil {
        log.Fatalf("Failed to create font: %s\n", err)
    }
	// END TEST
}
