package main

import (
	"fmt"

	"github.com/dynzhdjl/AoC2021/dive"
	"github.com/dynzhdjl/AoC2021/sonar_sweep"
)

func main() {
	fmt.Println("Hello AoC2021")
	fmt.Println("Sonar Sweep: ", sonar_sweep.Sweep())
	fmt.Println("Sonar Sweep (Sliding): ", sonar_sweep.SlidingSweep())
	fmt.Println("Dive: ", dive.Navigate())
	fmt.Println("Dive (With Aim): ", dive.NavigateWithAim())
}
