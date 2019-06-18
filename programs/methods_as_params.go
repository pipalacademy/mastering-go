package main

type myFloat struct {
	f float64
}

func (myFloat mf) square() float64 {
	return mf.f * mf.f
}

func mapperFloat64(f func(float64) float64, alist []float64) []float64 {
	r := make([]float64, len(alist))
	for i, v := range alist {
		r[i] = f(v)
	}
	return r
}

func main() {

}
