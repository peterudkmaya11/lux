package lux

import (
	glm "github.com/go-gl/mathgl/mgl32"
	"luxengine.net/gl"
)

type ParticleSystem struct {
	root, direction    glm.Vec3
	transformfeedbacks []gl.TransformFeedback //size 2
	buffers            []gl.Buffer            //size 2
	program            gl.Program
	/*
		public:
		    ParticleSystem();

		    ~ParticleSystem();

		    bool InitParticleSystem(const Vector3f& Pos);

		    void Render(int DeltaTimeMillis, const Matrix4f& VP, const Vector3f& CameraPos);

		private:

		    bool m_isFirst;
		    unsigned int m_currVB;
		    unsigned int m_currTFB;
		    GLuint m_particleBuffer[2];
		    GLuint m_transformFeedback[2];
		    PSUpdateTechnique m_updateTechnique;
		    BillboardTechnique m_billboardTechnique;
		    RandomTexture m_randomTexture;
		    Texture* m_pTexture;
		    int m_time;
	*/
}

type Particle struct {
	Type               int32
	Position, Velocity glm.Vec3
	Lifetime           float32
}

func NewParticleSystem(position, direction glm.Vec3) (ps ParticleSystem) {
	ps.root = position
	ps.direction = direction
	ps.transformfeedbacks = gl.GenTransformFeedbacks(2)
	ps.buffers = gl.GenBuffers(2)
	return
}
