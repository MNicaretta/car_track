package main

import (
	"image"
	"math"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type movement struct {
	rotation float64
	vector   pixel.Vec
	end      pixel.Vec
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func getMovements() []movement {
	var movements []movement

	movements = append(movements, movement{
		rotation: 0.0,
		vector:   pixel.V(0, 1),
		end:      pixel.V(810, 570),
	})

	movements = append(movements, movement{
		rotation: math.Pi * 0.5,
		vector:   pixel.V(-1, 0),
		end:      pixel.V(570, 570),
	})

	movements = append(movements, movement{
		rotation: math.Pi,
		vector:   pixel.V(0, -1),
		end:      pixel.V(570, 370),
	})

	movements = append(movements, movement{
		rotation: math.Pi * 1.25,
		vector:   pixel.V(1.25, -1),
		end:      pixel.V(600, 350),
	})

	movements = append(movements, movement{
		rotation: math.Pi,
		vector:   pixel.V(0, -1),
		end:      pixel.V(600, 260),
	})

	movements = append(movements, movement{
		rotation: math.Pi * 0.75,
		vector:   pixel.V(-1.25, -1),
		end:      pixel.V(570, 240),
	})

	movements = append(movements, movement{
		rotation: math.Pi * 0.5,
		vector:   pixel.V(-1, -0),
		end:      pixel.V(470, 240),
	})

	movements = append(movements, movement{
		rotation: 0.0,
		vector:   pixel.V(0, 1),
		end:      pixel.V(470, 570),
	})

	movements = append(movements, movement{
		rotation: math.Pi * 0.5,
		vector:   pixel.V(-1, 0),
		end:      pixel.V(150, 570),
	})

	movements = append(movements, movement{
		rotation: math.Pi,
		vector:   pixel.V(0, -1),
		end:      pixel.V(150, 150),
	})

	return movements
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Animação Carro",
		Bounds: pixel.R(0, 0, 960, 720),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	pic, err := loadPicture("track.png")
	if err != nil {
		panic(err)
	}

	track := pixel.NewSprite(pic, pic.Bounds())

	pic, err = loadPicture("car.png")
	if err != nil {
		panic(err)
	}

	car := pixel.NewSprite(pic, pic.Bounds())

	move := 0.0
	last := time.Now()

	position := pixel.V(810, 150)

	movements := getMovements()
	currentMovement := 0

	for !win.Closed() {
		win.Clear(colornames.White)
		track.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		m := movements[currentMovement]

		mat := pixel.IM
		mat = mat.Moved(position)
		mat = mat.Rotated(position, m.rotation)

		car.Draw(win, mat)

		win.Update()

		dt := time.Since(last).Seconds()
		last = time.Now()

		move += 3 * dt

		if reachedTheEnd(position, m) {
			if currentMovement+1 < len(movements) {
				currentMovement++
				position = position.Add(pixel.V(move*m.vector.X, move*m.vector.Y))
			}
		} else {
			position = position.Add(pixel.V(move*m.vector.X, move*m.vector.Y))
		}
	}
}

func reachedTheEnd(position pixel.Vec, m movement) bool {
	var endX bool
	var endY bool

	if m.vector.X > 0 {
		endX = position.X >= m.end.X
	} else if m.vector.X < 0 {
		endX = position.X <= m.end.X
	} else {
		endX = true
	}

	if m.vector.Y > 0 {
		endY = position.Y >= m.end.Y
	} else if m.vector.Y < 0 {
		endY = position.Y <= m.end.Y
	} else {
		endY = true
	}

	return endX && endY
}

func main() {
	pixelgl.Run(run)
}
