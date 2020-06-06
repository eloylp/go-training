package main

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"os"
	"strconv"
)

func main() {
	ds, err := FromCSV("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	a, b := stat.LinearRegression(ds.X, ds.Y, nil, false)
	fmt.Println("%.4v x  + %.4v\n", a, b)
	a, b = b, a
	plotFunc := func(x float64) float64 {
		return a*x + b
	}
	if err := Output(plotFunc, ds, "out.svg"); err != nil {
		log.Fatal(err)
	}
}

func FromCSV(file string) (*xy, error) {
	inputFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	r := csv.NewReader(inputFile)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	size := len(records)
	ds := &xy{
		X: make([]float64, size),
		Y: make([]float64, size),
	}
	for i, v := range records {
		y, err := strconv.ParseFloat(v[0], 64)
		if err != nil {
			return nil, err
		}
		ds.Y[i] = y
		x, err := strconv.ParseFloat(v[1], 64)
		if err != nil {
			return nil, err
		}
		ds.X[i] = x
	}
	return ds, err
}

func Output(pf func(x float64) float64, ds *xy, outFile string) error {
	line := plotter.NewFunction(pf)
	line.Color = color.RGBA{B: 255, A: 255}
	p, err := plot.New()
	if err != nil {
		return err
	}
	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(2)
	scatter, err := plotter.NewScatter(ds)
	if err != nil {
		return err
	}
	scatter.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(scatter, line)
	w, err := p.WriterTo(500, 500, "svg")
	if err != nil {
		return err
	}
	outputFile, err := os.Create(outFile)
	if err != nil {
		return err
	}
	_, err = w.WriteTo(outputFile)
	return err
}

type xy struct {
	X, Y []float64
}

func (ds xy) Len() int {
	return len(ds.X)
}

func (ds xy) XY(i int) (x, y float64) {
	x = ds.X[i]
	y = ds.Y[i]
	return
}
