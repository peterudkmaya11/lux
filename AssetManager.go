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

func (this *AssetManager) LoadModel(name, iname string) {
	if strings.Contains(name, ".obj") {
		this.Models[iname] = NewWavefrontModelFromFile(this.modelsDir + name)
	} else {
		log.Fatal("cannot find " + name)
	}
}

func (this *AssetManager) LoadRenderProgram(vertexShader, fragmentShader, iname string) {
	//program, err := LoadProgram(this.shadersDir+"standard.vert", this.shadersDir+"standard.frag")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//this.Programs[iname] = program
}

func (this *AssetManager) LoadTexture(name, iname string) {
	if strings.Contains(name, ".png") {
		pic, err := LoadPng(this.texturesDir + name)
		if err != nil {
			log.Fatal(err)
		}
		this.Textures[iname] = pic
	} else {
		log.Fatal("unable to find texture " + (this.modelsDir + name))
	}
}

func (this *AssetManager) Clean() {
	for name, model := range this.Models {
		model.Delete()
		delete(this.Models, name)
	}
	for name, tex := range this.Textures {
		tex.Delete()
		delete(this.Textures, name)
	}
	for name, prog := range this.Programs {
		prog.Delete()
		delete(this.Programs, name)
	}
}
