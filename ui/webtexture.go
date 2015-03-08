package ui

import (
	"github.com/go-gl/glow/gl-core/3.3/gl"
	"github.com/hydroflame/gosomium"
	engine "github.com/hydroflame/noname"
	"strings"
)

var standard_js_app_name = "app"
var standard_base_directory = "./public/"

func InitializeGosomiumDefault() {
	gosomium.WebcoreInitialize(true, true, false, "", "", "", "", "", gosomium.LL_NORMAL, false, "", true, "", "", "", "", "", "", true, 0, false, false, "")
	gosomium.SetBaseDirectory(standard_base_directory)
}

//automatically set to "./public/"
func SetBaseDirectory(path string) {
	gosomium.SetBaseDirectory(path)
}

type HTMLTexture struct {
	Texture engine.Texture
	Webview gosomium.Webview
	funcmap map[string]func(gosomium.JSArray)
}

func NewHTMLview(width, height int, url string) *HTMLTexture {
	this := HTMLTexture{}
	this.Texture = engine.GenTexture()
	this.Webview = gosomium.CreateWebview(width, height, false)
	if strings.Contains(url, "http://") || strings.Contains(url, "https://") {
		this.Webview.LoadURL(url, "", "", "")
	} else {
		this.Webview.LoadFile(url, "")
	}
	this.Webview.SetTransparent(true)
	this.Webview.CreateObject(standard_js_app_name)
	this.funcmap = make(map[string]func(gosomium.JSArray))
	this.Texture.Bind(gl.TEXTURE_2D)
	defer this.Texture.Unbind(gl.TEXTURE_2D)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(width), int32(height), 0, gl.BGRA, gl.UNSIGNED_BYTE, nil)
	return &this
}

func (this *HTMLTexture) jscallback(objname, cbname string, vals gosomium.JSArray) {
	this.funcmap[cbname](vals)
}

func (this *HTMLTexture) Upload() {
	this.Texture.Bind(gl.TEXTURE_2D)
	defer this.Texture.Unbind(gl.TEXTURE_2D)
	ren := this.Webview.Render()
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, int32(ren.Width()), int32(ren.Height()), gl.BGRA, gl.UNSIGNED_BYTE, engine.Ptr(ren.GetBuffer()))
}

func (this *HTMLTexture) AddFunc(cbname string, cbfunc func(gosomium.JSArray)) {
	this.Webview.SetObjectCallback(standard_js_app_name, cbname)
	this.funcmap[cbname] = cbfunc
}
