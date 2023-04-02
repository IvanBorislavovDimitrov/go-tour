package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	grade, err := strconv.ParseFloat(strings.TrimSpace(input), 123)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(grade)
}
