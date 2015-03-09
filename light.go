package lux

type Light interface {
	SetColor(float32, float32, float32)
	Upload(UniformLocation)
	CastsShadow(bool)
}

type PointLight struct {
	X, Y, Z, R, G, B float32
}

func (this *PointLight) Move(x, y, z float32) {
	this.X = x
	this.Y = y
	this.Z = z
}

func (this *PointLight) SetColor(r, g, b float32) {
	this.R = r
	this.G = g
	this.B = b
}

func (this *PointLight) Upload(u UniformLocation) {
	panic("unimplemented")
}

func (this *PointLight) CastsShadow(cast bool) {
	panic("unimplemented")
}

type SpotLight struct {
	X, Y, Z, Tx, Ty, Tz, A, R, G, B float32
}

func (this *SpotLight) Move(x, y, z float32) {
	this.X = x
	this.Y = y
	this.Z = z
}

func (this *SpotLight) SetColor(r, g, b float32) {
	this.R = r
	this.G = g
	this.B = b
}

func (this *SpotLight) Upload(u UniformLocation) {
	panic("unimplemented")
}

func (this *SpotLight) CastsShadow(cast bool) {
	panic("unimplemented")
}

//Radian
func (this *SpotLight) SetAngle(a float32) {
	this.A = a
}

type DirectionalLight struct {
	Dx, Dy, Dz, R, G, B float32
}

func (this *DirectionalLight) SetColor(r, g, b float32) {
	this.R = r
	this.G = g
	this.B = b
}

func (this *DirectionalLight) Upload(u UniformLocation) {
	panic("unimplemented")
}

func (this *DirectionalLight) CastsShadow(cast bool) {
	panic("unimplemented")
}
