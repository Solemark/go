package console_apps

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Radius interface {
	float32 | float64
}
