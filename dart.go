package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
)

//each dart thrown is in the range (0.0, 0.0) -> (1.0, 1.0)
// (0.5, 0.5) is the center of the dart board.
// any dart where D(center, dart) < 0.5 is considered a hit
// any dart where D >0.5 is a miss
type Coord struct {
	x float64
	y float64
}
type Result struct {
	hits   int
	misses int
}

func init() {
	fmt.Println("num cpus = ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	var threadcount int
	var totalcount int
	var gameseach int
	var n_darts int
	var nres Result
	var res Result
	res.hits = 0
	res.misses = 0
	var gamecount int = 0
	var i int = 0
	var pi float64
	fmt.Println("Lets try throwing some darts")
	fmt.Println("How many threads at a given time?")
	_, err := fmt.Scanf("%d\n", &threadcount)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan Result, threadcount)
	fmt.Println("How many games each")
	_, err = fmt.Scanf("%d\n", &gameseach)
	if err != nil {
		log.Fatal(err)
	}
	totalcount = threadcount * gameseach
	fmt.Println("How many darts shall each throw?")
	_, err = fmt.Scanf("%d\n", &n_darts)
	if err != nil {
		log.Fatal(err)
	}

	for i < threadcount {
		go playdart(ch, n_darts, gameseach)
		i++

	}
	//receive results, launch another game if not finished
	for gamecount < totalcount {
		nres = <-ch
		res.hits += nres.hits
		res.misses += nres.misses
		//print current pi approximation
		pi = 4.0 * (float64(res.hits) / float64(res.hits+res.misses))
		fmt.Println("current value of pi: ", pi)
		gamecount++
	}
	fmt.Println("finished task")
	return
}

//play n_games of dart, throwing n_darts each time, and reporting results through channel each time
func playdart(ch chan Result, n_darts int, n_games int) {
	var i int
	var dart Coord
	var res Result
	var dist float64
	game_num := 0
	for game_num < n_games {

		i = 0
		res.hits = 0
		res.misses = 0
		for i < n_darts {
			//throw dart and translate hit relative to center of dart board
			dart.x = rand.Float64()
			dart.y = rand.Float64()
			dist = distance(dart)
			if dist < 1.0 {
				res.hits++
			} else {
				res.misses++
			}
			i++
		}
		ch <- res
		game_num++
	}
}

func distance(dart Coord) float64 {
	return math.Sqrt(math.Pow(dart.x, 2) + math.Pow(dart.y, 2))
}
