package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
)

type GBuffer struct {
	framebuffer                                  Framebuffer
	program                                      Program
	PUni, VUni, MUni, NUni, MVPUni, DiffuseUni   UniformLocation
	DiffuseTex, NormalTex, PositionTex, DepthTex Texture2D
	AggregateFramebuffer                         AggregateFB
	vp, view                                     glm.Mat4
	width, height                                int32
	//shadow
	ShadowMapUni, ShadowMatUni UniformLocation
	//lights
	NumPointLightUni, PointLightPosUni, PointLightColUni UniformLocation
	//View uniforms
	CamPosUni UniformLocation
	//cook torrance
	Cook_roughnessValue, Cook_F0, Cook_k UniformLocation
}

type AggregateFB struct {
	framebuffer                          Framebuffer
	program                              Program
	DiffUni, NormalUni, PosUni, DepthUni UniformLocation
	Out                                  Texture2D
}

func NewGBuffer(width, height int32) (gbuffer GBuffer, err error) {
	gbuffer.width, gbuffer.height = width, height
	fb := GenFramebuffer()
	fb.Bind()
	defer fb.Unbind()
	gbuffer.framebuffer = fb

	depthtex := GenTexture2D()
	depthtex.Bind()
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH24_STENCIL8, width, height, 0, gl.DEPTH_STENCIL, gl.UNSIGNED_INT_24_8, nil)

	diffuseTex := GenTexture2D()
	diffuseTex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)

	normalTex := GenTexture2D()
	normalTex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA16F, width, height, 0, gl.RGB, gl.FLOAT, nil)

	positionTex := GenTexture2D()
	positionTex.Bind()
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

	aggfb.Out = GenTexture2D()
	aggfb.Out.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)

	aggfb.DiffUni = aprog.GetUniformLocation("diffusetex")
	aggfb.NormalUni = aprog.GetUniformLocation("normaltex")
	aggfb.PosUni = aprog.GetUniformLocation("postex")
	aggfb.DepthUni = aprog.GetUniformLocation("depthtex")
	gbuffer.ShadowMapUni = aprog.GetUniformLocation("shadowmap")
	gbuffer.ShadowMatUni = aprog.GetUniformLocation("shadowmat")

	gbuffer.NumPointLightUni = aprog.GetUniformLocation("NUM_POINT_LIGHT")
	gbuffer.PointLightPosUni = aprog.GetUniformLocation("point_light_pos")
	gbuffer.PointLightColUni = aprog.GetUniformLocation("point_light_color")
	gbuffer.CamPosUni = aprog.GetUniformLocation("cam_pos")

	//test data for cook torrance shader
	gbuffer.Cook_roughnessValue = aprog.GetUniformLocation("roughnessValue")
	gbuffer.Cook_F0 = aprog.GetUniformLocation("F0")
	gbuffer.Cook_k = aprog.GetUniformLocation("k")

	aggfb.framebuffer.DrawBuffers(ColorAttachement0)
	aggfb.framebuffer.Texture(ReadDrawFramebuffer, ColorAttachement0, aggfb.Out, 0)

	gbuffer.AggregateFramebuffer = aggfb
	return
}

func (gb *GBuffer) Bind(cam *Camera) {
	gb.framebuffer.Bind()
	gb.program.Use()

	gb.vp = cam.Projection.Mul4(cam.View)
	gb.view = cam.View

	ViewPortChange(gb.width, gb.height)

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
}

//testing in progress
func (gb *GBuffer) Render(cam *Camera, mesh Mesh, tex Texture2D, t *Transform) {

	model := t.Mat4()
	mvp := gb.vp.Mul4(model)
	gb.MVPUni.UniformMatrix4fv(1, false, &mvp[0])

	gb.MUni.UniformMatrix4fv(1, false, &model[0])

	normal := model.Inv()
	gb.NUni.UniformMatrix4fv(1, true, &normal[0])

	gl.ActiveTexture(TextureUnitDiffuse)
	tex.Bind()
	gb.DiffuseUni.Uniform1i(TextureUniformDiffuse)

	mesh.Bind()
	mesh.DrawCall()
}

func (gb *GBuffer) Aggregate(cam *Camera, plights []*PointLight, shadowmat glm.Mat4, tex Texture2D, f1, f2, f3 float32) {
	gb.AggregateFramebuffer.framebuffer.Bind()

	gb.AggregateFramebuffer.program.Use()

	gb.Cook_roughnessValue.Uniform1f(f1)
	gb.Cook_F0.Uniform1f(f2)
	gb.Cook_k.Uniform1f(f3)

	gl.ActiveTexture(gl.TEXTURE0)
	gb.DiffuseTex.Bind()
	gb.AggregateFramebuffer.DiffUni.Uniform1i(0)

	gl.ActiveTexture(gl.TEXTURE1)
	gb.NormalTex.Bind()
	gb.AggregateFramebuffer.NormalUni.Uniform1i(1)

	gl.ActiveTexture(gl.TEXTURE2)
	gb.PositionTex.Bind()
	gb.AggregateFramebuffer.PosUni.Uniform1i(2)

	gl.ActiveTexture(gl.TEXTURE3)
	gb.DepthTex.Bind()
	gb.AggregateFramebuffer.DepthUni.Uniform1i(3)

	//point lights
	gb.NumPointLightUni.Uniform1i(int32(len(plights)))
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
	gb.PointLightPosUni.Uniform3fv(int32(len(plights)), &plightpos[0])
	gb.PointLightColUni.Uniform3fv(int32(len(plights)), &plightcol[0])

	gb.CamPosUni.Uniform3fv(1, &cam.Pos[0])

	//=====shadow=====//
	gl.ActiveTexture(gl.TEXTURE4)
	tex.Bind()
	gb.ShadowMapUni.Uniform1i(4)

	gb.ShadowMatUni.UniformMatrix4fv(1, false, &shadowmat[0])
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
	outNormal = (N*vec4(normal,1)).xyz;
	outPosition = world_pos;
}
` + "\x00"

//shouldnt be there
var _gbuffer_aggregate_fragment_shader = `
#version 330
#define MAX_POINT_LIGHT 8

//GBuffer textures
uniform sampler2D diffusetex;
uniform sampler2D normaltex;
uniform sampler2D postex;
uniform sampler2D depthtex;

//cook
uniform float roughnessValue;
uniform float F0;
uniform float k;

//Lights
uniform int NUM_POINT_LIGHT;
uniform vec3 point_light_pos[MAX_POINT_LIGHT];
uniform vec3 point_light_color[MAX_POINT_LIGHT];

//Shadows
uniform sampler2DShadow shadowmap;
uniform mat4 shadowmat;

//View
uniform vec3 cam_pos;

in vec2 uv;

layout (location=0) out vec4 outColor;

void main(){
	vec3 normal = normalize(texture(normaltex, uv).xyz);
	vec3 world_position = texture(postex, uv).xyz;

	vec4 shadowcoord = shadowmat*vec4(world_position, 1);
	shadowcoord.z+=0.005;
	float shadow = texture(shadowmap, shadowcoord.xyz,0);


	//float light =0;
	//for(int i = 0; i < NUM_POINT_LIGHT; i++){
	//	light+=max(0,dot(normal, point_light_pos[i]-world_position));
	//}
	//float luma = 0.3;
	
	//luma = max(luma,light);
	//luma = min(shadow,luma);
	//luma = clamp(luma,0.3,1);
	//outColor = texture(diffusetex, uv)*luma;

	//////cook torrance

	//material values
	//float roughnessValue = 0.1;
	//float F0 = 0.8; //fresnel reflectance at normal incidence
	//float k = 0.2; //fraction of diffuse reflection (specular reflection = 1 - k)
	vec3 lightColor = vec3(0.9,0.1,0.1);

	vec3 world_pos = texture(postex, uv).xyz;
	vec3 lightDir = point_light_pos[0]-world_pos;

	float NdL = max(dot(normal, lightDir), 0);

	float lux = shadow;
	if(shadow > 0){
		float specular = 0.0;
		if(NdL > 0.0){
			vec3 eyeDir = normalize(cam_pos-world_pos);

			vec3 halfVec = normalize(lightDir+eyeDir);
			float NdH = max(0,dot(normal,halfVec));
			float NdV = max(0,dot(normal, eyeDir));
			float VdH = max(0,dot(eyeDir, halfVec));
			float mSqu = roughnessValue*roughnessValue;

			float NH2 = 2.0*NdH;
			float geoAtt = min(1.0,min((NH2*NdV)/VdH,(NH2*NdL)/VdH));
			float roughness = (1.0 / ( 4.0 * mSqu * pow(NdH, 4.0)))*exp((NdH * NdH - 1.0) / (mSqu * NdH * NdH));
			float fresnel = pow(1.0 - VdH, 5.0)*(1.0 - F0)+F0;
			specular = (fresnel*geoAtt*roughness)/(NdV*NdL*3.14);
		}
		lux=NdL * (k + specular * (1.0 - k));
	}
	outColor = texture(diffusetex, uv)*lux;
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

/*
precision highp float; //set default precision in glsl es 2.0

uniform vec3 lightDirection;

varying vec3 varNormal;
varying vec3 varEyeDir;

void main()
{
    // set important material values
    float roughnessValue = 0.3; // 0 : smooth, 1: rough
    float F0 = 0.8; // fresnel reflectance at normal incidence
    float k = 0.2; // fraction of diffuse reflection (specular reflection = 1 - k)
    vec3 lightColor = vec3(0.9, 0.1, 0.1);

    // interpolating normals will change the length of the normal, so renormalize the normal.
    vec3 normal = normalize(varNormal);

    // do the lighting calculation for each fragment.
    float NdotL = max(dot(normal, lightDirection), 0.0);

    float specular = 0.0;
    if(NdotL > 0.0)
    {
        vec3 eyeDir = normalize(varEyeDir);

        // calculate intermediary values
        vec3 halfVector = normalize(lightDirection + eyeDir);
        float NdotH = max(dot(normal, halfVector), 0.0);
        float NdotV = max(dot(normal, eyeDir), 0.0); // note: this could also be NdotL, which is the same value
        float VdotH = max(dot(eyeDir, halfVector), 0.0);
        float mSquared = roughnessValue * roughnessValue;

        // geometric attenuation
        float NH2 = 2.0 * NdotH;
        float g1 = (NH2 * NdotV) / VdotH;
        float g2 = (NH2 * NdotL) / VdotH;
        float geoAtt = min(1.0, min(g1, g2));

        // roughness (or: microfacet distribution function)
        // beckmann distribution function
        float r1 = 1.0 / ( 4.0 * mSquared * pow(NdotH, 4.0));
        float r2 = (NdotH * NdotH - 1.0) / (mSquared * NdotH * NdotH);
        float roughness = r1 * exp(r2);

        // fresnel
        // Schlick approximation
        float fresnel = pow(1.0 - VdotH, 5.0);
        fresnel *= (1.0 - F0);
        fresnel += F0;

        specular = (fresnel * geoAtt * roughness) / (NdotV * NdotL * 3.14);
    }

    vec3 finalValue = lightColor * NdotL * (k + specular * (1.0 - k);
    gl_FragColor = vec4(finalValue, 1.0);
}
*/
