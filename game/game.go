package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	randNumber := rand.Intn(100-1) + 1
	for i := 0; i < 10; i++ {
		fmt.Print("Guess the number: ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		number, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			panic(err)
		}
		if number < randNumber {
			fmt.Println("LOW")
			continue
		} else if number > randNumber {
			fmt.Println("HIGH")
			continue
		} else {
			fmt.Println("CORRECT")
			return
		}
	}
	fmt.Println("Stupid bastard. You should've guessed", randNumber)

}
