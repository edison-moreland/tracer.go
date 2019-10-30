package main

import (
	"fmt"
	"os"
)

func main() {
	out, err := os.Create("traced.ppm")
	if err != nil { panic(err.Error()) }
	defer out.Close()

	nx := 200
	ny := 100
	_, _ = fmt.Fprintf(out, "P3\n%d %d\n255\n", nx, ny)
	for j := ny-1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2
			ir := int(255.99*r)
			ig := int(255.99*g)
			ib := int(255.99*b)
			_, _ = fmt.Fprintf(out, "%d %d %d\n", ir, ig, ib)
		}
	}
	
}