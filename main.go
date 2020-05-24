package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/cranburyattic/GoQuadTree/quadtree"
)

func initQuadTree() *quadtree.Quadtree {
	rootBoundary := quadtree.NewBoundary(-2, 54.5, 8, 9)
	rootQT := quadtree.NewQuadtree(rootBoundary, 0)

	filePath := "./data.csv"

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		fmt.Println("Unable to locate data file")
		panic(1)
	}

	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if record[1] != "" {
			x, _ := strconv.ParseFloat(record[2], 64)
			y, _ := strconv.ParseFloat(record[1], 64)
			i, _ := strconv.ParseInt(record[0], 10, 64)
			rootQT.Insert(quadtree.Point{X: x, Y: y, I: i})
		}
	}

	fmt.Println("Data Loaded")
	return rootQT
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func rectsHandler(qt *quadtree.Quadtree) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		enableCors(&rw)

		rects := make([]quadtree.Boundary, 0)

		for _, qt := range qt.All() {
			boundary := qt.GetBoundary()
			rects = append(rects, boundary)
		}

		output, err := json.MarshalIndent(&rects, "", "\t")

		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}

		rw.Write(output)
	}
}

func queryHandler(qt *quadtree.Quadtree) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		enableCors(&rw)

		x, y, w, h, err := getQueryParaams(r.URL.Query())

		if err != nil {
			fmt.Println("Error set query params:", err)
		}

		points := qt.Query(quadtree.NewBoundary(x, y, w, h))
		output, err := json.MarshalIndent(&points, "", "\t")

		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}

		rw.Write(output)
	}
}

func getQueryParaams(params url.Values) (float64, float64, float64, float64, error) {
	paramX := params.Get("x")
	paramY := params.Get("y")
	paramW := params.Get("w")
	paramH := params.Get("h")

	fx, err := strconv.ParseFloat(paramX, 64)
	fy, err := strconv.ParseFloat(paramY, 64)
	fw, err := strconv.ParseFloat(paramW, 64)
	fh, err := strconv.ParseFloat(paramH, 64)

	return fx, fy, fw, fh, err
}

func main() {

	qt := initQuadTree()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/rects", rectsHandler(qt))
	http.HandleFunc("/query", queryHandler(qt))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
