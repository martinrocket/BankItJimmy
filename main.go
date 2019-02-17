package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const version = "1.01"
const nameOfTheGame = "Bank it Jimmy"
const header = "You are playing " + nameOfTheGame

var myOS = getOS()
var ThisGame = players{}

type players struct {
	list []player
}

type player struct {
	name  string
	score int64
}

func main() {

	var exit bool = false
	var text string = "nil"

	BIJLog, err := os.Create("log.txt")
	for err != nil {
		fmt.Println("error creating log file")
	}

	BIJLog.WriteString("Starting Log " + time.Now().String() + "\n")
	BIJLog.WriteString("---\n")

	for exit != true {
		clearScreen()
		getMenu1()
		if text != "nil" {
			log.Printf("You chose %v.\n", text)
		}
		text = getChoice("Enter text")
		switch strings.ToLower(text) {
		case "x":
			clearScreen()
			getMenu1()
			fmt.Println("Exiting")
			exit = true
			break
		case "n":
			clearScreen()
			fmt.Println("New Game Menu")
			getMenuNewGame()
		case "rd":
			d1, d2 := getMenuRoll()
			fmt.Printf("Dice 1 = %v \nDice 2 = %v \n", d1, d2)
			getChoice("Press Enter")
		default:
			clearScreen()
			fmt.Println("No Menu Yet")
			getMenu1()

		}

	}
	x := new(player)
	x.name = "martin"
	fmt.Println(x.name, x.score)
}

func getOS() string {
	if runtime.GOOS == "darwin" {
		return "OSX"
	} else {
		return runtime.GOOS
	}
}

func clearScreen() {
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func getMenu1() {
	fmt.Printf(header + " - " + myOS + "\n")
	fmt.Println()
	fmt.Print(`


		 ______              _        _
		(____  \            | |      (_)_
		 ____)  ) ____ ____ | |  _    _| |_
		|  __  ( / _  |  _ \| | / )  | |  _)
		| |__)  | ( | | | | | |< (   | | |__
		|______/ \_||_|_| |_|_| \_)  |_|\___)

		   _____ _                   _
		  (_____|_)                 | |
		     _   _ ____  ____  _   _| |
		    | | | |    \|    \| | | |_|
		 ___| | | | | | | | | | |_| |_
		(____/  |_|_|_|_|_|_|_|\__  |_|
		                      (____/


  `)
	fmt.Println()
	fmt.Print(`
n) New Game
r) Resume Game
p) Player Rankings
x) exit
    `)
	fmt.Println()
}

func getMenuNewGame() {
	done := false
	for done == false {
		clearScreen()
		fmt.Printf("Number of Players: %v.\n", len(ThisGame.list))
		listPlayers()
		fmt.Printf("Choose your player options: \n")
		fmt.Print(`
a) Add Player
d) Delete Player
l) List Players
p) Play
d) done
    `)
		fmt.Println()
		playerText := getChoice("Enter text")
		switch strings.ToLower(playerText) {
		case "a":
			ThisGame.list = append(ThisGame.list, addPlayer())
			listPlayers()
		case "d":
			fmt.Println("d")
			done = true
		case "p":
			playGame()
		default:
			fmt.Println("other")
		}
	}

}

func getChoice(s string) string {
	fmt.Printf("%s: ", s)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	//text = strings.ToLower(strings.Trim(text, " \r\n"))
	text = strings.Trim(text, " \r\n")
	return text
}

func getMenuRoll() (int64, int64) {
	dice1 := rand.Int63n(6) + 1
	dice2 := rand.Int63n(6) + 1
	return dice1, dice2

}

func addPlayer() player {
	fmt.Println("Adding a player")
	firstName := getChoice("First Name")
	lastName := getChoice("Last Name")
	fmt.Println("Your new Player name is", firstName, lastName)
	getChoice("Press Enter")
	p := player{name: firstName + " " + lastName}
	return p
}

func playGame() {
	done := false //flag to stop playing the game when done equals true.
	for done != true {
		clearScreen()
		fmt.Println("Play Game!") //Header

		listPlayers()                             //List of the current players
		d := getChoice("Press Enter to Continue") //Wait for player to hit enter
		if strings.ToLower(d) == "d" {
			done = true
		}
		for i := 0; i < len(ThisGame.list); i++ {
			stop := false
			//var playInCounter int64
			//var playInSelect string
			var score, thisScore1, thisScore2, playInScoreTotal, playInScore1, playInScore2 int64
			for stop == false {
				if playInScore1 == 0 {
					fmt.Printf("%v, ", ThisGame.list[i].name) //Print Name
					getChoice("Let's roll the dice for you play in. Press Enter")
					playInScore1, playInScore2 = getMenuRoll()
					playInScoreTotal = playInScore1 + playInScore2
					fmt.Printf("%v, ", ThisGame.list[i].name) //Print Name
					fmt.Printf("You rolled a %v and %v for a total of %v.\n", playInScore1, playInScore2, playInScoreTotal)
					fmt.Printf("%v, ", ThisGame.list[i].name) //Print Name
					higherOrLower := getChoice("Will your roll be higher or lower? (h/l)")
					for strings.ToLower(higherOrLower) == "h" || strings.ToLower(higherOrLower) == "l" {
						break
					}

				}
				fmt.Printf("Player %v please roll:", ThisGame.list[i].name)
				getChoice("Enter to roll")
				thisScore1, thisScore2 = getMenuRoll()
				fmt.Printf("You rolled %v and %v.\n", thisScore1, thisScore2)
				fmt.Print(`
b) bank it
r) roll again
`)
				q := getChoice("What now?")
				switch strings.ToLower(q) {
				case "b":
					score = score + thisScore1 + thisScore2
					ThisGame.list[i].score = score
					stop = true
				case "r":
					score = score + thisScore1 + thisScore2

				}
			}
		}
	}
	return
}

func listPlayers() {
	for i := 0; i < len(ThisGame.list); i++ {
		fmt.Printf("%v %v \n", ThisGame.list[i].name, ThisGame.list[i].score)
	}
}

func playGame2() {
	playGameBool := false
	for playGameBool == false {
		fmt.Println("Play Bank it Jimmy!")
		fmt.Println("Players:")
		listPlayers()
		//playMenu()
		playTurnBool := false
		for playTurnBool == false {
			for i := 0; i < len(ThisGame.list); i++ {
				x := getPlayIn(ThisGame.list[i].name)
				if x != 0 {
					//choice := getChoice("%v, do you want to bank %v or roll again? (b/r)", ThisGame.list[i].name, x)
					choice := getChoice("Do you want to bank %v or roll again? (b/r)?")
					fmt.Println(choice)
					// for strings.ToLower(choice) == "b" {
					// 	append(ThisGame.list[i], ThisGame.list[i].score+x)
					//
					// }
				} else {
					break
				}
				//playTurn()
				//nextPlayer
				//nextRound
				//endThisGame
			}
		}
	}
}

func getPlayIn(name string) int64 {
	fmt.Printf("%v, you are rolling for play in. Roll the dice!", name)
	getChoice("Press Enter")
	d1, d2 := getMenuRoll()
	t1 := d1 + d2
	fmt.Printf("%v, you rolled for a %v and %v for a total of %v.\n", d1, d2, t1)
	guess := getChoice("Will you roll higher or lower? (h/l)")
	for strings.ToLower(guess) != "h" || strings.ToLower(guess) != "l" {
		guess = getChoice("Will you roll higher or lower? (h/l)")
	}
	fmt.Printf("%v, you say that your next roll will be %v than %v.\n", guess, t1)
	getChoice("Press Enter")
	d3, d4 := getMenuRoll()
	t2 := d3 + d4
	fmt.Printf("%v, you rolled for a %v and %v for a total of %v.\n", d3, d4, t2)
	if (strings.ToLower(guess) == "h" && t2 > t1) || (strings.ToLower(guess) == "l" && t2 < t1) {
		fmt.Printf("You win the play in!")
		return t2
	} else {
		fmt.Printf("You lose the play in")
		return 0

	}

	return 0

}
