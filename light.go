package lux

type Light interface {
	SetColor(float32, float32, float32)
	Upload(UniformLocation)
	CastsShadow(bool)
}

type PointLight struct {
	X, Y, Z, R, G, B float32
}

func (pl *PointLight) Move(x, y, z float32) {
	pl.X = x
	pl.Y = y
	pl.Z = z
}

func (pl *PointLight) SetColor(r, g, b float32) {
	pl.R = r
	pl.G = g
	pl.B = b
}

func (pl *PointLight) Upload(u UniformLocation) {
	panic("unimplemented")
}

func (pl *PointLight) CastsShadow(cast bool) {
	panic("unimplemented")
}

type SpotLight struct {
	X, Y, Z, Tx, Ty, Tz, A, R, G, B float32
}

func (sl *SpotLight) Move(x, y, z float32) {
	sl.X = x
	sl.Y = y
	sl.Z = z
}

func (sl *SpotLight) SetColor(r, g, b float32) {
	sl.R = r
	sl.G = g
	sl.B = b
}

func (sl *SpotLight) Upload(u UniformLocation) {
	panic("unimplemented")
}

func (sl *SpotLight) CastsShadow(cast bool) {
	panic("unimplemented")
}

//Radian
func (sl *SpotLight) SetAngle(a float32) {
	sl.A = a
}

type DirectionalLight struct {
	Dx, Dy, Dz, R, G, B float32
}

func (dl *DirectionalLight) SetColor(r, g, b float32) {
	dl.R = r
	dl.G = g
	dl.B = b
}

func (dl *DirectionalLight) Upload(u UniformLocation) {
	panic("unimplemented")
}

func (dl *DirectionalLight) CastsShadow(cast bool) {
	panic("unimplemented")
}
