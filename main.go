package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var typeOfOscillation = []string{
	"Thiraktey",
	"Simat-tey",
	"Jhulastey",
}

var facialFeature = []string{
	"Nazron",
	"Hothon",
	"Nainon",
	"Gaalon",
	"Maathey",
}

var naturalPhenomena = []string{
	"Boond",
	"Oas",
	"Baarish",
	"Hawa",
}

var thingsThatHappenInBathroom = []string{
	"Phisalna",
	"Girna",
	"Tapakna",
	"Ulajhna",
	"Sulajhna",
	"Nikalna",
	"Sawarna",
}

var randomEmotion = []string{
	"Yaad",
	"Darata",
	"Hasata",
	"Rulata",
}

var movementOfUFO = []string{
	"Duur jaana",
	"Paas aana",
	"Roshni Phailana",
	"Gumm ho jaana",
}

func main() {
	count := flag.Int("c", 10, "number of lines of shaayari, default 10")
	flag.Parse()
	println(*count)
	for i := 0; i < *count; i++ {
		s := randWord(typeOfOscillation) + " hue " + randWord(facialFeature) + " se " + randWord(naturalPhenomena) + " ka " + randWord(thingsThatHappenInBathroom) + ",\n" + randWord(randomEmotion) + " hai mujhe tera wahi " + randWord(movementOfUFO) + "\n"
		fmt.Println(s)
	}
}

func randWord(a []string) string {
	r := rand.Int63n(int64(len(a)))
	return a[r]
}
