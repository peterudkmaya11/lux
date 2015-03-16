package lux

import (
	"log"
	"strings"
)

type AssetManager struct {
	modelsDir, shadersDir, texturesDir string
	Models                             map[string]Mesh
	Textures                           map[string]Texture2D
	Programs                           map[string]Program
}

func NewAssetManager(root, models, shaders, textures string) (out AssetManager) {
	out.modelsDir, out.shadersDir, out.texturesDir = root+models, root+shaders, root+textures
	out.Models = make(map[string]Mesh)
	out.Textures = make(map[string]Texture2D)
	out.Programs = make(map[string]Program)
	return out
}

func (am *AssetManager) LoadModel(name, iname string) {
	if strings.Contains(name, ".obj") {
		am.Models[iname] = NewWavefrontModelFromFile(am.modelsDir + name)
	} else {
		log.Fatal("cannot find " + name)
	}
}

func (am *AssetManager) LoadRenderProgram(vertexShader, fragmentShader, iname string) {
	//program, err := LoadProgram(am.shadersDir+"standard.vert", am.shadersDir+"standard.frag")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//am.Programs[iname] = program
}

func (am *AssetManager) LoadTexture(name, iname string) {
	if strings.Contains(name, ".png") {
		pic, err := LoadPng(am.texturesDir + name)
		if err != nil {
			log.Fatal(err)
		}
		am.Textures[iname] = pic
	} else {
		log.Fatal("unable to find texture " + (am.modelsDir + name))
	}
}

func (am *AssetManager) Clean() {
	for name, model := range am.Models {
		model.Delete()
		delete(am.Models, name)
	}
	for name, tex := range am.Textures {
		tex.Delete()
		delete(am.Textures, name)
	}
	for name, prog := range am.Programs {
		prog.Delete()
		delete(am.Programs, name)
	}
}
