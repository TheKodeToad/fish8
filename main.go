package main

import (
	"fmt"
	"math"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/thekodetoad/fish8/system"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: <path>")
		os.Exit(1)
	}

	program, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	sys, err := system.New(program)
	if err != nil {
		panic(err)
	}

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(
		system.DisplayWidth*10,
		system.DisplayHeight*10,
		"fish8",
	)

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		shouldDraw := sys.Tick()

		sys.UpdateKeypad(getHeldKeys())

		if !shouldDraw {
			continue
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		scale := math.Min(
			float64(rl.GetScreenWidth())/system.DisplayWidth,
			float64(rl.GetScreenHeight())/system.DisplayHeight,
		)

		rl.Translatef(
			float32(rl.GetScreenWidth())/2-system.DisplayWidth*float32(scale)/2,
			float32(rl.GetScreenHeight())/2-system.DisplayHeight*float32(scale)/2,
			0,
		)
		rl.Scalef(float32(scale), float32(scale), 0)

		sys.ReadDisplay(func(x, y int, on bool) {
			if on {
				rl.DrawRectangle(int32(x), int32(y), 1, 1, rl.White)
			}
		})

		rl.EndDrawing()
	}
}

var keyMap = map[int32]system.Key{
	rl.KeyOne:   0x1,
	rl.KeyTwo:   0x2,
	rl.KeyThree: 0x3,
	rl.KeyFour:  0xC,

	rl.KeyQ: 0x4,
	rl.KeyW: 0x5,
	rl.KeyE: 0x6,
	rl.KeyR: 0xD,

	rl.KeyA: 0x7,
	rl.KeyS: 0x8,
	rl.KeyD: 0x9,
	rl.KeyF: 0xE,

	rl.KeyZ: 0xA,
	rl.KeyX: 0x0,
	rl.KeyC: 0xB,
	rl.KeyV: 0xF,
}

func getHeldKeys() system.KeySet {
	var result system.KeySet

	for rlKey, key := range keyMap {
		if rl.IsKeyDown(rlKey) {
			result |= key.ToKeySet()
		}
	}

	return result
}
