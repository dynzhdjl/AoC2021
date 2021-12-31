package main

import (
	"fmt"

	"github.com/dynzhdjl/AoC2021/binary_diagnostic"
	"github.com/dynzhdjl/AoC2021/bingo"
	"github.com/dynzhdjl/AoC2021/dive"
	"github.com/dynzhdjl/AoC2021/dumbo_octopus"
	"github.com/dynzhdjl/AoC2021/extended_polymerization"
	"github.com/dynzhdjl/AoC2021/hydrothermal_venture"
	"github.com/dynzhdjl/AoC2021/lanternfish"
	"github.com/dynzhdjl/AoC2021/passage_pathing"
	"github.com/dynzhdjl/AoC2021/seven_segment"
	"github.com/dynzhdjl/AoC2021/smoke_basin"
	"github.com/dynzhdjl/AoC2021/sonar_sweep"
	"github.com/dynzhdjl/AoC2021/syntax_scoring"
	"github.com/dynzhdjl/AoC2021/transparent_origami"
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
	fmt.Println("Day 4, Bingo: ", bingo.WinnerScore())
	fmt.Println("Day 5, Hydrothermal Vents:", hydrothermal_venture.NumberOfVentLineIntersection(2))
	fmt.Println("Day 6, Lanternfish:", lanternfish.Population(256, 6, 2))
	fmt.Println("Day 7, Crazy Carbs Horizontal Alignment:", "*")
	fmt.Println("Day 8, Seven Segment Search: ", seven_segment.Decode())
	fmt.Println("Day 9, Smoke Basin (Risk): ", smoke_basin.CalculateRisk())
	fmt.Println("Day 9, Smoke Basin (Biggest Basin): ", smoke_basin.FindBiggestBasin())
	fmt.Println("Day 10, Syntax Scoring: ", syntax_scoring.CalculateSyntaxCheckerScore())
	fmt.Println("Day 10, Autocomplete Scoring: ", syntax_scoring.CalculateAutoCompleteScore())
	fmt.Println("Day 11, Dumbo Octopus (Flash Count): ", dumbo_octopus.CountFlashes(100))
	fmt.Println("Day 11, Dumbo Octopus (Sync Step): ", dumbo_octopus.FindSyncronizationStep())
	fmt.Println("Day 12, Passage Pathing (Unique Path Count): ", passage_pathing.PathCount())
	fmt.Println("Day 12, Passage Pathing (Unique Path Allowed Twice Count): ", passage_pathing.PathAllowedTwiceCount())
	transparent_origami.Fold()
	fmt.Println("Day 14, Extended Polymerization (10 Steps): ", extended_polymerization.Diff(10))
	fmt.Println("Day 14, Extended Polymerization (40 Steps): ", extended_polymerization.Diff(40))
}
