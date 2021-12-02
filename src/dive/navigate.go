package dive

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dynzhdjl/AoC2021/util"
)

type command struct {
	T string
	V int
}

type position struct {
	depth      int
	horizontal int
	aim        int
}

type navigator struct {
	position        *position
	updateProcedure func(p *position, cmd command)
}

func (n navigator) navigate(cmd command) {
	n.updateProcedure(n.position, cmd)
}

func updateWithAim(p *position, cmd command) {
	switch cmd.T {
	case "forward":
		p.horizontal += cmd.V
		p.depth += (p.aim * cmd.V)
	case "up":
		p.aim -= cmd.V
	case "down":
		p.aim += cmd.V
	}
}

func updateWithOutAim(p *position, cmd command) {
	switch cmd.T {
	case "forward":
		p.horizontal += cmd.V
	case "up":
		p.depth -= cmd.V
	case "down":
		p.depth += cmd.V
	}
}

func navigate(commands []command, updateStrategy func(p *position, cmd command)) int {
	p := position{0, 0, 0}
	n := navigator{&p, updateStrategy}
	for _, c := range commands {
		n.navigate(c)
	}
	return p.depth * p.horizontal
}

func getCommands() []command {
	lines := util.Read("dive/input.txt")
	input := make([]command, len(lines))
	for _, l := range lines {
		v := strings.Split(l, " ")
		d, err := strconv.Atoi(v[1])
		if err != nil {
			panic(fmt.Sprintf("could parse the input: %s", l))
		}
		input = append(input, command{v[0], d})
	}
	return input
}

func Navigate() int {
	return navigate(getCommands(), updateWithOutAim)
}

func NavigateWithAim() int {
	return navigate(getCommands(), updateWithAim)
}
