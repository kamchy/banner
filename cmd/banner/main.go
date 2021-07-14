package main

import (
	"math/rand"
	"time"

	ba "github.com/kamchy/banner"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := ba.GetInput()
	ba.GenerateBanner(i)
}
