package weather

type Location struct {
	Lat float64
	Lng float64
}

type Weather struct {
	Location Location
	Weather  string
}
