package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	//firstOne()
	//secondOne()
	// thirdOne()
	// dupLinesTwo()
	// dupLinesThree()
	// lissaJous(os.Stdout)
	// createGIF()
	// request()
	// fetchAll()
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/handle2", handlerTwo)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		lissaJous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func firstOne() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		// fmt.Println(os.Args[0] + "\n")
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func secondOne() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func thirdOne() {
	// fmt.Println(strings.Join(os.Args[0:], " "))

	var s string
	for i := 1; i < len(os.Args); i++ {
		s += fmt.Sprintf("%s %d, \n", os.Args[i], i)
	}
	fmt.Println(s)
}

func dupLinesOne() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dupLinesTwo() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dupLinesThree() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		if len(counts) > 1 {
			fmt.Print(filename)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{R: 64, G: 224, B: 208, A: 255},
}

const (
	whiteIndex     = 0
	blackIndex     = 1
	turquoiseIndex = 2
)

func createGIF() {
	file, err := os.Create("lissajous.gif")
	if err != nil {
		panic(err)
		//08162498903
	}
	defer file.Close()
	// lissaJous(file)
}

// func lissaJous(out *os.File) {
func lissaJous(out http.ResponseWriter, r *http.Request) {
	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nFrames = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	cyclesParams := r.URL.Query().Get("cycles")
	cycles, err := strconv.Atoi(cyclesParams)
	if err != nil {
		cycles = 5 // default value if parsing fails
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0 //phase difference
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := blackIndex
		if i%2 == 0 {
			colorIndex = turquoiseIndex
		}
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
			// img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), turquoiseIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err2 := gif.EncodeAll(out, &anim)

	if err2 != nil {
		fmt.Print(err2)
	}
}

func request() {
	for i, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		// b, err := io.ReadAll(resp.Body)

		fileName := fmt.Sprintf("Response_%d.json", i)
		file, er := os.Create(fileName)
		if er != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", er)
			os.Exit(1)
		}
		defer file.Close()

		bytesWritten, err := io.Copy(file, resp.Body)
		// resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("The status code for the request is %d \n", resp.StatusCode)
		fmt.Printf("Wrote %d bytes to %s\n", bytesWritten, fileName)
	}
}

func fetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go requestTwo(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func requestTwo(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nBytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nBytes, url)
}

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handlerTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
