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
	excerciseToTestWhatYouLearned_1()

}

func excerciseToTestWhatYouLearned_1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is you age: ")
	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}

	if iAge < 5 {
		fmt.Println("Too young for school")
	} else if iAge == 5 {
		fmt.Println("Go to kindergarten")
	} else if (iAge > 5) && (iAge <= 17) {
		grade := iAge - 5
		fmt.Printf("Go to grade %d\n", grade)
	} else {
		fmt.Println("Go to collage")
	}
}
