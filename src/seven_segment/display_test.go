package seven_segment

import (
	"fmt"
	"testing"
)

func TestBitPatternCreation(t *testing.T) {
	a := bitPattern("abcdefg")
	b := bitPattern("cabdefg")
	if a != b {
		t.Errorf("a:%b, b:%b", a, b)
	}
	if a != 127 {
		t.Error()
	}
}

func TestSignalToDigit(t *testing.T) {
	signals := []string{
		"be",      //1
		"cfbegad", //8
		"cbdgef",  //9
		"fgaecd",  //6
		"cgeb",    //4
		"fdcge",   //5
		"agebfd",  //0
		"fecdb",   //3
		"fabcd",   //2
		"edb",     //7
	}
	m := signalToDigit(signals)
	fmt.Println(m[bitPattern("be")])
	fmt.Println(m[bitPattern("cfbegad")])
	fmt.Println(m[bitPattern("cbdgef")])
	fmt.Println(m[bitPattern("fgaecd")])
	fmt.Println(m[bitPattern("cgeb")])
	fmt.Println(m[bitPattern("fdcge")])
	fmt.Println(m[bitPattern("agebfd")])
	fmt.Println(m[bitPattern("fecdb")])
	fmt.Println(m[bitPattern("fabcd")])
	fmt.Println(m[bitPattern("edb")])
}
