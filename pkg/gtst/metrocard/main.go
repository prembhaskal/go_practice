package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/prembhaskal/go_practice/pkg/gtst/metrocard/metro"
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	inventory := metro.NewInventory()

	for scanner.Scan() {
		inputType := scanner.Text()
		switch inputType {
		case "BALANCE":
			readAndAddCard(scanner, inventory)
		case "CHECK_IN":
			checkIn(scanner, inventory)
		case "PRINT_SUMMARY":
			printSummary(inventory)
		default:
			panic(fmt.Sprintf("unknown input: %s", inputType))
		}

	}
}

func GetNextWord(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return scanner.Text()
	}
	panic("no more data to read")
}

func readAndAddCard(scanner *bufio.Scanner, inventory *metro.Inventory) {
	cardID := GetNextWord(scanner)
	balance, err := strconv.ParseFloat(GetNextWord(scanner), 64)
	if err != nil {
		panic(fmt.Sprintf("float parse error: %v", err))
	}
	addCard(inventory, cardID, balance)
}

func checkIn(scanner *bufio.Scanner, inventory *metro.Inventory) {
	cardID := GetNextWord(scanner)
	passTypeName := GetNextWord(scanner)
	stationName := GetNextWord(scanner)

	card := inventory.GetCard(cardID)
	passType := inventory.GetPassType(passTypeName)
	fromStation := inventory.GetStation(stationName)
	toStation := inventory.GetToStation(stationName)

	fromStation.CheckIn(card, passType, toStation.Name)
}

func printSummary(inventory *metro.Inventory) {
	// fmt.Print("printing summary")
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()

	central := inventory.GetStation("CENTRAL")
	metro.PrintSummary(central, bw)
	airport := inventory.GetStation("AIRPORT")
	metro.PrintSummary(airport, bw)
}

func addCard(inventory *metro.Inventory, id string, balance float64) {
	inventory.AddCard(metro.NewTravelCard(id, balance))
}
