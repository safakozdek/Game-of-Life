package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
	"math/rand"
)

const (
	SCALE  = 8
	WIDTH  = 50
	HEIGHT = 50
)

var (
	black  = color.RGBA{R: 0, G: 0, B: 0, A: 255}       //0,0,0
	white  = color.RGBA{R: 200, G: 200, B: 255, A: 255} //200,200,255
	grid   = [WIDTH][HEIGHT]uint8{}
	buffer = [WIDTH][HEIGHT]uint8{}
	mode   = 0 //default value for random grid
)
var (
	empty [][]int

	gliderX = []int{1, 2, 3, 3, 3}
	gliderY = []int{2, 3, 1, 2, 3}
	glider  = [][]int{gliderX, gliderY}

	gliderGunX = []int{1, 1, 2, 2, 11, 11, 11, 12, 12, 13, 13, 14, 14, 15, 16, 16, 17, 17, 17, 18,
		21, 21, 21, 22, 22, 22, 23, 23, 25, 25, 25, 25, 35, 35, 36, 36}
	gliderGunY = []int{5, 6, 5, 6, 5, 6, 7, 4, 8, 3, 9, 3, 9, 6, 4, 8, 5, 6, 7, 6,
		3, 4, 5, 3, 4, 5, 2, 6, 1, 2, 6, 7, 3, 4, 3, 4}
	gliderGun = [][]int{gliderGunX, gliderGunY}


	defaultScreens = [][][]int{empty, glider, gliderGun}
)

type Game struct {
}

func (g *Game) Update(screen *ebiten.Image) error {

	for x := 1; x < WIDTH-1; x++ {
		for y := 1; y < HEIGHT-1; y++ {
			buffer[x][y] = 0
			n := g.sumAdjacents(x, y)
			if grid[x][y] == 0 && n == 3 {
				buffer[x][y] = 1
			} else if n > 3 || n < 2 {
				buffer[x][y] = 0
			} else {
				buffer[x][y] = grid[x][y]
			}
		}
	}
	temp := buffer
	buffer = grid
	grid = temp
	var err = g.Draw(screen)
	return err
}

func (g *Game) sumAdjacents(x, y int) uint8 {
	var sum uint8
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				sum += grid[x+i][y+j]
			}
		}
	}
	return sum
}

func (g *Game) Render(screen *ebiten.Image) error {
	var err error = nil
	err = screen.Fill(white)
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if grid[x][y] > 0 {
				for x1 := 0; x1 < SCALE; x1++ {
					for y1 := 0; y1 < SCALE; y1++ {
						screen.Set((x*SCALE)+x1, (y*SCALE)+y1, black)
					}
				}
			}
		}
	}
	return err
}

func (g *Game) Draw(screen *ebiten.Image) error {
	var err error = nil
	if !ebiten.IsDrawingSkipped() {
		err = g.Render(screen)
	}
	return err
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func initializeGrid() {
	if mode == 0 {
		for x := 1; x < WIDTH-1; x++ {
			for y := 1; y < HEIGHT-1; y++ {
				if rand.Float32() <= 0.5 {
					grid[x][y] = 1
				}
			}
		}
	} else {
		for i := 0; i < len(defaultScreens[mode][0]); i++ {
			grid[defaultScreens[mode][0][i]][defaultScreens[mode][1][i]] = 1
		}
	}
}

func main() {
	initializeGrid()
	ebiten.SetMaxTPS(20)
	ebiten.SetWindowSize(WIDTH*SCALE, HEIGHT*SCALE)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

