package noname

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
)

var _ = gl.ACTIVE_ATOMIC_COUNTER_BUFFERS

type GBuffer struct {
	framebuffer                                  Framebuffer
	program                                      Program
	PUni, VUni, MUni, NUni, MVPUni, DiffuseUni   UniformLocation
	DiffuseTex, NormalTex, PositionTex, DepthTex Texture
	AggregateFramebuffer                         AggregateFB
	vp, view                                     glm.Mat4
	width, height                                int32
	//shadow
	ShadowMapUni, ShadowMatUni UniformLocation
	//lights
	NumPointLightUni, PointLightPosUni, PointLightColUni UniformLocation
}

type AggregateFB struct {
	framebuffer                          Framebuffer
	program                              Program
	DiffUni, NormalUni, PosUni, DepthUni UniformLocation
	Out                                  Texture
}

func NewGBuffer(width, height int32) (gbuffer GBuffer, err error) {
	gbuffer.width, gbuffer.height = width, height
	fb := GenFramebuffer()
	fb.Bind()
	defer fb.Unbind()
	gbuffer.framebuffer = fb
	//depthbuffer := GenRenderBuffer()
	//depthbuffer.Bind()
	//depthbuffer.Storage(gl.DEPTH24_STENCIL8, width, height)
	//fb.RenderBuffer(DepthStencilAttachement, depthbuffer)

	depthtex := GenTexture()
	depthtex.Bind(gl.TEXTURE_2D)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH24_STENCIL8, width, height, 0, gl.DEPTH_STENCIL, gl.UNSIGNED_INT_24_8, nil)

	diffuseTex := GenTexture()
	diffuseTex.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, width, height, 0, gl.RGBA, gl.FLOAT, nil)

	normalTex := GenTexture()
	normalTex.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA16F, width, height, 0, gl.RGB, gl.FLOAT, nil)

	positionTex := GenTexture()
	positionTex.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA16F, width, height, 0, gl.RGB, gl.FLOAT, nil)

	fb.DrawBuffers(ColorAttachement0, ColorAttachement1, ColorAttachement2)

	fb.Texture(ReadDrawFramebuffer, ColorAttachement0, diffuseTex, 0 /*level*/)
	fb.Texture(ReadDrawFramebuffer, ColorAttachement1, normalTex, 0 /*level*/)
	fb.Texture(ReadDrawFramebuffer, ColorAttachement2, positionTex, 0 /*level*/)
	fb.Texture(ReadDrawFramebuffer, DepthStencilAttachement, depthtex, 0 /*level*/)

	gbuffer.DiffuseTex = diffuseTex
	gbuffer.NormalTex = normalTex
	gbuffer.PositionTex = positionTex
	gbuffer.DepthTex = depthtex

	vs, err := CompileShader(_gbuffer_vertex_shader_source, VertexShader)
	if err != nil {
		return
	}
	fs, err := CompileShader(_gbuffer_fragment_shader_source, FragmentShader)
	if err != nil {
		return
	}
	prog, err := NewProgram(vs, fs)
	if err != nil {
		return
	}
	gbuffer.program = prog

	prog.Use()
	defer prog.Unuse()
	gbuffer.PUni = prog.GetUniformLocation("P")
	gbuffer.VUni = prog.GetUniformLocation("V")
	gbuffer.MUni = prog.GetUniformLocation("M")
	gbuffer.NUni = prog.GetUniformLocation("N")
	gbuffer.MVPUni = prog.GetUniformLocation("MVP")
	gbuffer.DiffuseUni = prog.GetUniformLocation("diffuse")
	//shadow map

	prog.BindFragDataLocation(0, "")
	prog.BindFragDataLocation(1, "")
	prog.BindFragDataLocation(2, "")

	//Aggregated fb and textures, essentially a special post process effect
	aggfb := AggregateFB{}
	gbuffer.AggregateFramebuffer = aggfb

	avs, err := CompileShader(_fullscreen_vertex_shader, VertexShader)
	if err != nil {
		return
	}
	afs, err := CompileShader(_gbuffer_aggregate_fragment_shader, FragmentShader)
	if err != nil {
		return
	}
	aprog, err := NewProgram(avs, afs)
	if err != nil {
		return
	}
	aggfb.program = aprog

	aggfb.framebuffer = GenFramebuffer()
	aggfb.framebuffer.Bind()

	aggfb.Out = GenTexture()
	aggfb.Out.Bind(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, width, height, 0, gl.RGBA, gl.FLOAT, nil)

	aggfb.DiffUni = aprog.GetUniformLocation("diffusetex")
	aggfb.NormalUni = aprog.GetUniformLocation("normaltex")
	aggfb.PosUni = aprog.GetUniformLocation("postex")
	aggfb.DepthUni = aprog.GetUniformLocation("depthtex")
	gbuffer.ShadowMapUni = aprog.GetUniformLocation("shadowmap")
	gbuffer.ShadowMatUni = aprog.GetUniformLocation("shadowmat")

	gbuffer.NumPointLightUni = aprog.GetUniformLocation("NUM_POINT_LIGHT")
	gbuffer.PointLightPosUni = aprog.GetUniformLocation("point_light_pos")
	gbuffer.PointLightColUni = aprog.GetUniformLocation("point_light_color")

	aggfb.framebuffer.DrawBuffers(ColorAttachement0)
	aggfb.framebuffer.Texture(ReadDrawFramebuffer, ColorAttachement0, aggfb.Out, 0)

	gbuffer.AggregateFramebuffer = aggfb
	return
}

func (this *GBuffer) Bind(cam *Camera) {
	this.framebuffer.Bind()
	this.program.Use()

	this.vp = cam.Projection.Mul4(cam.View)
	this.view = cam.View

	ViewPortChange(this.width, this.height)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
}

//testing in progress
func (this *GBuffer) Render(cam *Camera, mesh Mesh, tex Texture, t *Transform) {

	model := t.Mat4()
	mvp := this.vp.Mul4(model)
	this.MVPUni.UniformMatrix4fv(1, false, &mvp[0])

	this.MUni.UniformMatrix4fv(1, false, &model[0])

	normal := cam.View.Mul4(model).Inv()
	this.NUni.UniformMatrix4fv(1, true, &normal[0])

	gl.ActiveTexture(TextureUnitDiffuse)
	tex.Bind(gl.TEXTURE_2D)
	this.DiffuseUni.Uniform1i(TextureUniformDiffuse)

	mesh.Bind()
	mesh.DrawCall()
}

func (this *GBuffer) Aggregate(plights []*PointLight, shadowmat glm.Mat4, tex Texture) {
	this.AggregateFramebuffer.framebuffer.Bind()

	this.AggregateFramebuffer.program.Use()

	gl.ActiveTexture(gl.TEXTURE0)
	this.DiffuseTex.Bind(gl.TEXTURE_2D)
	this.AggregateFramebuffer.DiffUni.Uniform1i(0)

	gl.ActiveTexture(gl.TEXTURE1)
	this.NormalTex.Bind(gl.TEXTURE_2D)
	this.AggregateFramebuffer.NormalUni.Uniform1i(1)

	gl.ActiveTexture(gl.TEXTURE2)
	this.PositionTex.Bind(gl.TEXTURE_2D)
	this.AggregateFramebuffer.PosUni.Uniform1i(2)

	gl.ActiveTexture(gl.TEXTURE3)
	this.DepthTex.Bind(gl.TEXTURE_2D)
	this.AggregateFramebuffer.DepthUni.Uniform1i(3)

	//point lights
	this.NumPointLightUni.Uniform1i(int32(len(plights)))
	plightpos := make([]float32, len(plights)*3)
	plightcol := make([]float32, len(plights)*3)
	for i, light := range plights {
		plightpos[i] = light.X
		plightpos[i+1] = light.Y
		plightpos[i+2] = light.Z

		plightcol[i] = light.R
		plightcol[i] = light.G
		plightcol[i] = light.B
	}
	this.PointLightPosUni.Uniform3fv(int32(len(plights)), &plightpos[0])
	this.PointLightColUni.Uniform3fv(int32(len(plights)), &plightcol[0])

	//=====shadow=====//
	gl.ActiveTexture(gl.TEXTURE4)
	tex.Bind(gl.TEXTURE_2D)
	this.ShadowMapUni.Uniform1i(4)

	this.ShadowMatUni.UniformMatrix4fv(1, false, &shadowmat[0])
	//================//

	Fstri()
}

var _gbuffer_vertex_shader_source = `
#version 330
uniform mat4 M;
uniform mat4 MVP;

layout (location=0) in vec3 vert;
layout (location=1) in vec2 vertTexCoord;
layout (location=2) in vec3 vertNormal;

out vec2 fragTexCoord;
out vec3 normal;
out vec3 world_pos;

void main() {
	normal = vertNormal;
	fragTexCoord = vertTexCoord;
	world_pos=(M*vec4(vert,1)).xyz;
	gl_Position = MVP * vec4(vert, 1);
}
` + "\x00"

var _gbuffer_fragment_shader_source = `
#version 330
uniform sampler2D diffuse;
uniform mat4 N;

in vec2 fragTexCoord;
in vec3 normal;
in vec3 world_pos;
layout (location=0) out vec3 outColor;
layout (location=1) out vec3 outNormal;
layout (location=2) out vec3 outPosition;
void main() {
	outColor = texture(diffuse, fragTexCoord).rgb;
	outNormal = (N*vec4(normalize(normal),1)).xyz;
	outPosition = world_pos;
}
` + "\x00"

//shouldnt be there
var _gbuffer_aggregate_fragment_shader = `
#version 330
#define MAX_POINT_LIGHT 8
uniform sampler2D diffusetex;
uniform sampler2D normaltex;
uniform sampler2D postex;
uniform sampler2D depthtex;

//lights
uniform int NUM_POINT_LIGHT;
uniform vec3 point_light_pos[MAX_POINT_LIGHT];
uniform vec3 point_light_color[MAX_POINT_LIGHT];

uniform sampler2DShadow shadowmap;
uniform mat4 shadowmat;

in vec2 uv;

layout (location=0) out vec4 outColor;


void main(){
	vec3 normal = normalize(texture(normaltex, uv).xyz);
	vec3 world_position = texture(postex, uv).xyz;

	vec4 shadowcoord = shadowmat*vec4(world_position, 1);
	shadowcoord.z+=0.005;
	float shadow = texture(shadowmap, shadowcoord.xyz,0.5);


	float light =0;
	for(int i = 0; i < NUM_POINT_LIGHT; i++){
		light+=max(0,dot(normal, point_light_pos[i]-world_position));
	}
	float luma = 0.3;
	//luma = max(0.3,shadow);//ambient light
	
	luma = max(luma,light);
	luma = clamp(luma,0,1);
	luma = min(shadow,luma);
	outColor = texture(diffusetex, uv)*luma;

}
` + "\x00"

/*
sobel operation
	//experimenting with sobel

	float i00   = texture2D(depthtex, uv).r;
	float im1m1 = texture2D(depthtex, uv+vec2(-pixwidth,-pixheight)).r;
	float ip1p1 = texture2D(depthtex, uv+vec2(pixwidth,pixheight)).r;
	float im1p1 = texture2D(depthtex, uv+vec2(-pixwidth,pixheight)).r;
	float ip1m1 = texture2D(depthtex, uv+vec2(pixwidth,-pixheight)).r;
	float im10 = texture2D(depthtex, uv+vec2(-pixwidth,0)).r;
	float ip10 = texture2D(depthtex, uv+vec2(pixwidth,0)).r;
	float i0m1 = texture2D(depthtex, uv+vec2(0,-pixheight)).r;
	float i0p1 = texture2D(depthtex, uv+vec2(0,pixheight)).r;
	float h = -im1p1 - 32.0 * i0p1 - ip1p1 + im1m1 + 32.0 * i0m1 + ip1m1;
	float v = -im1m1 - 32.0 * im10 - im1p1 + ip1m1 + 32.0 * ip10 + ip1p1;

	float mag = 1-length(vec2(h, v));

	//outColor = vec4(vec3(mag),1);


*/
