package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	duration time.Time
	slice    []time.Duration
}

func (t *Stopwatch) Start() {
	t.slice = nil
	t.duration = time.Now()
}
func (t *Stopwatch) GetResults() []time.Duration {
	return t.slice
}
func (t *Stopwatch) SaveSplit() {
	t.slice = append(t.slice, time.Since(t.duration))
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
