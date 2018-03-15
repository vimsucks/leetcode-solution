package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
	fmt.Println(countSegments("Hello,    my name is John"))
	fmt.Println(countSegments("Hello,    my name is John!"))
	fmt.Println(countSegments("love live! mu'sic forever"))
	fmt.Println(countSegments("The one-hour drama series Westworld is a dark odyssey about the dawn of artificial consciousness and the evolution of sin. Set at the intersection of the near future and the reimagined past, it explores a world in which every human appetite, no matter how noble or depraved, can be indulged."))
	fmt.Println(countSegments(", , , ,        a, eaefa"))
}

func countSegments(s string) int {
	re := regexp.MustCompile("\\S+")
	return len(re.FindAllString(s, -1))
}
