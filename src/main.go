package main

import (
	"fmt"

	"github.com/dynzhdjl/AoC2021/binary_diagnostic"
	"github.com/dynzhdjl/AoC2021/dive"
	"github.com/dynzhdjl/AoC2021/sonar_sweep"
)

func main() {
	fmt.Println("AoC2021")
	fmt.Println("-------")
	fmt.Println("Day 1, Sonar Sweep: ", sonar_sweep.Sweep())
	fmt.Println("Day 1, Sonar Sweep (Sliding): ", sonar_sweep.SlidingSweep())
	fmt.Println("Day 2, Dive: ", dive.Navigate())
	fmt.Println("Day 2, Dive (With Aim): ", dive.NavigateWithAim())
	fmt.Println("Day 3, Diagnostic (Energy Consumption): ", binary_diagnostic.GetEnergyConsumption())
	fmt.Println("Day 3, Diagnostic (Life Support): ", binary_diagnostic.GetLifeSupportRating())
}
