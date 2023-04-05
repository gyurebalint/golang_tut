package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tsp float64
type TBs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}

func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {
	excerciseToTestWhatYouLearned_1()
	definedTypes()

}
func definedTypes() {
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsp = %.2f ML\n", ml1)
	ml2 := ML(TBs(3)) * 14.79
	fmt.Printf("3 tbs = %.2f ML\n", ml2)

	fmt.Printf("3tsp = %.2f mL\n", tspToML(3))
	fmt.Printf("3tbs = %.2f mL\n", TBToML(3))

	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())
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
