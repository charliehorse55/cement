package main

import (
    glfw "github.com/go-gl/glfw3"
	Ecem "github.com/charliehorse55/libcement"
)

var keys []glfw.Key

type keyboard struct {
	w *glfw.Window
	intensity []Ecem.RGB32f
} 

func (k *keyboard)Begin(w *glfw.Window, intensity []Ecem.RGB32f) error {
	k.w = w
    k.intensity = intensity
	keys = []glfw.Key{glfw.Key1, glfw.Key2, glfw.Key3, glfw.Key4, glfw.Key5, glfw.Key6, glfw.Key7, glfw.Key8, glfw.Key9, glfw.Key0}
	w.SetKeyCallback(k.keyPress)
	return nil
}

func (k *keyboard)keyPress(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	for i := range k.intensity {
		if k.w.GetKey(keys[i]) == glfw.Press {
			k.intensity[i].R = 1.0
			k.intensity[i].G = 1.0
			k.intensity[i].B = 1.0
		} else {
			k.intensity[i].R = 0.0
			k.intensity[i].G = 0.0
			k.intensity[i].B = 0.0
		}
	}
}

func (k *keyboard)ShouldSave() bool {
    return false
}


