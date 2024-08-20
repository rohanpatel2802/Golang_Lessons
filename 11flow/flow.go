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
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Tutorial for control flow")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a movie rating between 0-5")
	str, _ := reader.ReadString('\n')
	rat, _ := strconv.ParseFloat(strings.TrimSpace(str), 64)
	fmt.Println("Movie rating is:", rat)
	if rat >= 3.5 {
		fmt.Println("According to the rating, the movie is good.")
	} else if rat >= 2.5 {
		fmt.Println("According to the rating, the movie is average.")
	} else {
		fmt.Println("According to the rating, the movie is bad.")
	}

	pos := 1
	for pos < 100 {
		fmt.Println("Press enter to roll the dice, current position is:", pos)
		str1, _ := reader.ReadString('\n')
		fmt.Println(str1)
		roll := rand.Intn(6) + 1
		fmt.Println("You've roll a :", roll)
		if pos+int(roll) == 100 {
			fmt.Println("You won!!!!!")
			pos = 100
			break
		} else if pos+int(roll) > 100 {
			fmt.Println("Invalid move")
		} else {
			pos = pos + int(roll)
			fmt.Println("Position after roll is:", pos)
			switch pos {
			case 7:
				fmt.Println("Climbed a ladder to 28")
				pos = 28
			case 25:
				fmt.Println("Climbed a ladder to 51")
				pos = 51
			case 32:
				fmt.Println("Climbed a ladder to 85")
				pos = 85
			case 74:
				fmt.Println("Climbed a ladder to 91")
				pos = 91
			case 16:
				fmt.Println("Climbed a ladder to 45")
				pos = 45
			case 15:
				fmt.Println("Bit by a snake, fell to 3")
				pos = 3
			case 77:
				fmt.Println("Bit by a snake, fell to 56")
				pos = 56
			case 44:
				fmt.Println("Bit by a snake, fell to 22")
				pos = 22
			case 23:
				fmt.Println("Bit by a snake, fell to 10")
				pos = 10
			case 98:
				fmt.Println("Bit by a snake, fell to 12")
				pos = 12
				fallthrough
			case 12:
				fmt.Println("Climbed a ladder to 19")
			}
		}
	}
}
