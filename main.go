package main

import (
	"fmt"
	"go-chip-8/cpu"
)

func main() {
	new_cpu := cpu.NewCPU("roms/pong.ch8")
	fmt.Println(new_cpu.Mem)
	cpu.Reset(new_cpu)
	fmt.Println(new_cpu.Mem)
	// if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
	// 	panic(err)
	// }
	// defer sdl.Quit()

	// window, err := sdl.CreateWindow("SDL Tutorial", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	// if err != nil {
	// 	panic(err)
	// }
	// defer window.Destroy()

	// surface, err := window.GetSurface()
	// if err != nil {
	// 	panic(err)
	// }
	// surface.FillRect(nil, 0)

	// rect := sdl.Rect{0, 0, 200, 200}
	// color := sdl.Color{255, 0, 0, 80}
	// pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.G, color.A)
	// surface.FillRect(&rect, pixel)
	// window.UpdateSurface()

	// running := true
	// for running {
	// 	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch event.(type) {
	// 		case *sdl.QuitEvent:
	// 			running = false
	// 		}
	// 	}
	// }

}
