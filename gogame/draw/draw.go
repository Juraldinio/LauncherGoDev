package draw

import (
//	"fmt"
//	"math"
	"github.com/veandco/go-sdl2/sdl"
	"../color"
	
)

func AARoundRect() {
	
}

func Point(surf *sdl.Surface, c color.Color, x,y int) {
	pixels := surf.Pixels()
	bytes_per_pixel := surf.BytesPerPixel()
	
	addr := y * int(surf.Pitch) + x*bytes_per_pixel // 1 2 3 4

	color_bytes := c.ToBytes()

	surf.Lock()
	
	if bytes_per_pixel == 1 {
		pixels[addr] = color_bytes[0]
	}

	if bytes_per_pixel == 2 {
		for i :=0; i < bytes_per_pixel; i++ {
			pixels[addr+i] = color_bytes[i]
		}	
	}

	if bytes_per_pixel == 3 {
		for i :=0; i < bytes_per_pixel; i++ {
			pixels[addr+i] = color_bytes[i]
		}	
	}

	if bytes_per_pixel == 4 {
		for i :=0; i < bytes_per_pixel; i++ {
			pixels[addr+i] = color_bytes[i]
		}
	}
	
	surf.Unlock()
}
