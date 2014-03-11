package main

import (
    glfw "github.com/go-gl/glfw3"
	Ecem "github.com/charliehorse55/libcement"
)

type RGB64f struct {
	R, G, B float64
}

type keyScroll struct {
	p Ecem.Painting
	selected int
	r, g, b bool
	save bool
} 

func clamp(x float32) float32 {
	if x > 1.0 {
		return 1.0
	}
	if x < 0.0 {
		return 0.0
	}
	return x
}

func (k *keyScroll)didScroll(w *glfw.Window, xoff float64, yoff float64) {
	diff := float32(-yoff/100.0)
	i := k.selected
	if k.r {
		k.p[i].Intensity.R = clamp(k.p[i].Intensity.R + diff)
	}                                            
	if k.g {                                     
		k.p[i].Intensity.G = clamp(k.p[i].Intensity.G + diff)
	}                                            
	if k.b {                                     
		k.p[i].Intensity.B = clamp(k.p[i].Intensity.B + diff)
	}
}

func (k *keyScroll)keyPress(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if w.GetAttribute(glfw.Focused) == glfw.True && action == glfw.Release {
		index := k.selected; 
		switch key {
		case glfw.Key1:
			index = 0
		case glfw.Key2:
			index = 1
		case glfw.Key3:
			index = 2
		case glfw.Key4:
			index = 3
		case glfw.Key5:
			index = 4
		case glfw.Key6:
			index = 5
		case glfw.Key7:
			index = 6
		case glfw.Key8:
			index = 7
		case glfw.Key9:
			index = 8
		case glfw.Key0:
			index = 9
		case glfw.KeyQ:
			k.r, k.g, k.b = true, false, false
		case glfw.KeyW:
			k.r, k.g, k.b = false, true, false
		case glfw.KeyE:
			k.r, k.g, k.b = false, false, true
		case glfw.KeyR:
			k.r, k.g, k.b = true, true, true
		case glfw.KeyS:
			if (mods & glfw.ModControl) > 0 {
				k.save = true 
			}
		}
		if index < len(k.p) {
			k.selected = index
		}
	}
}

func (k *keyScroll)ShouldSave() bool {
	tmp := k.save
	k.save = false
	return tmp
}

func (k *keyScroll)Begin(w *glfw.Window, p Ecem.Painting) error {
	w.SetScrollCallback(k.didScroll)
	w.SetKeyCallback(k.keyPress)
	k.p = p
	k.r, k.g, k.b = true, true, true
	return nil
}


