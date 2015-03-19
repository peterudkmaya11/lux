package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glm "github.com/go-gl/mathgl/mgl32"
	"luxengine.net/lux/utils"
)

type Mesh interface {
	Bind()
	Delete()
	Size() int
	DrawCall()
}

type VUNMesh struct { //Vertex, Uv, Normal Model
	VAO                              VertexArray
	Indices, Positions, Uvs, Normals Buffer
	Msize                            int
}

func NewWavefrontModelFromFile(file string) Mesh {
	//load object
	meshObj := utils.LoadObject(file, false)

	//prepare indices //to optimise later
	indices, indexedVertices, indexedUvs, indexedNormals := utils.IndexVBOSlow(meshObj.Vertices, meshObj.UVs, meshObj.Normals)
	return NewVUNModel(indices, indexedVertices, indexedUvs, indexedNormals)
}

func NewVUNModel(indices []uint16, indexedVertices []glm.Vec3, indexedUvs []glm.Vec2, indexedNormals []glm.Vec3) Mesh {

	m := VUNMesh{}
	m.VAO = GenVertexArray()
	m.VAO.Bind()
	defer m.VAO.Unbind()

	m.Msize = len(indices)
	//create a bunch of buffers and fill them
	//Positions
	m.Positions = GenBuffer()
	m.Positions.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(indexedVertices)*3*4, ptr(indexedVertices), gl.STATIC_DRAW)

	//Uvs
	m.Uvs = GenBuffer()
	m.Uvs.Bind(gl.ARRAY_BUFFER)
	// And yet, the weird length stuff doesn't seem to matter for UV or normal //<- wtf is this guy talking about
	gl.BufferData(gl.ARRAY_BUFFER, len(indexedUvs)*2*4, ptr(indexedUvs), gl.STATIC_DRAW)

	//Normals
	m.Normals = GenBuffer()
	m.Normals.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(indexedNormals)*3*4, ptr(indexedNormals), gl.STATIC_DRAW)

	//indices
	m.Indices = GenBuffer()
	m.Indices.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*2, ptr(indices), gl.STATIC_DRAW)

	return &m
}

func (m *VUNMesh) Bind() {
	m.VAO.Bind()

	gl.EnableVertexAttribArray(0)
	m.Positions.Bind(gl.ARRAY_BUFFER)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(1)
	m.Uvs.Bind(gl.ARRAY_BUFFER)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	gl.EnableVertexAttribArray(2)
	m.Normals.Bind(gl.ARRAY_BUFFER)
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)

	m.Indices.Bind(gl.ELEMENT_ARRAY_BUFFER)
}

func (m *VUNMesh) Unbind() {
	m.VAO.Unbind()
}

func (m VUNMesh) Delete() {
	defer m.Positions.Delete()
	defer m.Uvs.Delete()
	defer m.Normals.Delete()
	defer m.Indices.Delete()
	defer m.VAO.Delete()
}

func (m *VUNMesh) Size() int {
	return m.Msize
}

func (m *VUNMesh) DrawCall() {
	gl.DrawElements(gl.TRIANGLES, int32(m.Size()), gl.UNSIGNED_SHORT, nil)
}
