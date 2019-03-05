package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var typeOfOscillation = []string{
	"Thiraktey",
	"Simat-tey",
	"Jhulastey",
	"Lachaktey",
	"Jhukey",
	"Machaltey",
}

var facialFeature = []string{
	"Nazron",
	"Hothon",
	"Nainon",
	"Gaalon",
	"Maathey",
	"Labon",
	"Zulfon",
}

var naturalPhenomena = []string{
	"Boond",
	"Oas",
	"Baarish",
	"Hawa",
	"Leheron",
}

var thingsThatHappenInBathroom = []string{
	"Phisalna",
	"Girna",
	"Tapakna",
	"Ulajhna",
	"Sulajhna",
	"Nikalna",
	"Sawarna",
	"Bikharna",
	"Behekna",
	"Sambhalna",
	"Sisakna",
	"Barasna",
}

var randomEmotion = []string{
	"Yaad",
	"Darata",
	"Hasata",
	"Rulata",
	"Lubhata",
	"Tadpata",
}

var movementOfUFO = []string{
	"Duur jaana",
	"Paas aana",
	"Roshni Phailana",
	"Gumm ho jaana",
	"Afsana",
}

func main() {
	count := flag.Int("c", 10, "number of lines of shaayari, default 10")
	flag.Parse()
	for i := 0; i < *count; i++ {
		s := randWord(typeOfOscillation) + " hue " + randWord(facialFeature) + " se " + randWord(naturalPhenomena) + " ka " + randWord(thingsThatHappenInBathroom) + ",\n" + randWord(randomEmotion) + " hai mujhe tera wahi " + randWord(movementOfUFO) + "\n"
		fmt.Println(s)
	}
}

func randWord(a []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Int63n(int64(len(a)))
	return a[r]
}
