package lux

import (
	"log"
	"strings"
)

//keeps track of loaded textures, models and programs
type AssetManager struct {
	modelsDir, shadersDir, texturesDir string
	Models                             map[string]Mesh
	Textures                           map[string]Texture2D
	Programs                           map[string]Program
}

//Make a new asset manager
//	-root: the root of all the other folders. eg. "assets"
//	-models: location of models. eg. "models", located at "assets/models"
//	-shaders: location of shaders. Not really used right now becasue everything is hard coded :\. eg. "shaders", located at "assets/shaders"
//	-textures: location of texture. eg. "textures", located at "assets/textures"
func NewAssetManager(root, models, shaders, textures string) (out AssetManager) {
	out.modelsDir, out.shadersDir, out.texturesDir = root+models, root+shaders, root+textures
	out.Models = make(map[string]Mesh)
	out.Textures = make(map[string]Texture2D)
	out.Programs = make(map[string]Program)
	return out
}

//Load a single model. Only wavefront available for now. iname is the internal name to be set in the map.
func (am *AssetManager) LoadModel(name, iname string) {
	if strings.Contains(name, ".obj") {
		am.Models[iname] = NewWavefrontModelFromFile(am.modelsDir + name)
	} else {
		log.Fatal("cannot find " + name)
	}
}

//Do not use
func (am *AssetManager) LoadRenderProgram(vertexShader, fragmentShader, iname string) {
	//program, err := LoadProgram(am.shadersDir+"standard.vert", am.shadersDir+"standard.frag")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//am.Programs[iname] = program
}

//Load an image as a texture2D. iname is the internal name to be set in the map.
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

//delete/release everything loaded
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
