package UI

import (

	"fmt"
	
	"github.com/veandco/go-sdl2/ttf"
	
	"github.com/cuu/gogame/font"	
	"github.com/cuu/LauncherGo/sysgo"
)

type Widget struct {
	PosX int
	PosY int
	Width int
	Height int
}

func (self *Widget) Size() (int,int) {
	return self.Width,self.Height
}

func (self *Widget) NewSize(w,h int) {
	self.Width = w
	self.Height = h
}

func (self *Widget) Coord() (int,int) {
	return self.PosX,self.PosY
}

func (self *Widget) NewCoord(x,y int) {
	self.PosX = x
	self.PosY = y
}


func Init() {
	font.Init()

	skinpath := "skin/"+sysgo.SKIN+"/truetype"

	Fonts = make(map[string]*ttf.Font)

	fonts_path := make(map[string]string)


	fonts_path["varela"] = fmt.Sprintf("%s/VarelaRound-Regular.ttf",skinpath)
	fonts_path["veramono"] = fmt.Sprintf("%s/VeraMono.ttf",skinpath)
	fonts_path["noto"]     = fmt.Sprintf("%s/NotoSansMono-Regular.ttf", skinpath)
	fonts_path["notocjk"]  = fmt.Sprintf("%s/NotoSansCJK-Regular.ttf" ,skinpath)

	for i:=12;i<41;i++ {
		keyname := fmt.Sprintf("varela%d",i)
		Fonts[ keyname ] = font.Font(fonts_path["varela"],i)
	}

	for i:=10;i<26;i++ {
		keyname := fmt.Sprintf("veramono%d", i)
		Fonts[keyname] = font.Font(fonts_path["veramono"],i)
	}

	for i:= 10;i<18;i++ {
		keyname := fmt.Sprintf("notosansmono%d", i)
		Fonts[keyname] = font.Font(fonts_path["noto"], i)
	}

	for i:=10;i<18;i++ {
		keyname := fmt.Sprintf("notosanscjk%d",i)
		Fonts[keyname] = font.Font(fonts_path["notocjk"],i)
	}
}
