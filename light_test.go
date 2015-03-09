package lux

import (
	"testing"
)

func TestLightInterface(t *testing.T) {
	var l Light
	l = &PointLight{}
	l = &SpotLight{}
	l = &DirectionalLight{}
	l.SetColor(1, 1, 1)
}
