package main

import (
	"bufio"
	"calculator-console/calculator"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("Welcome to CalculatorConsole!")

	calc := calculator.NewCalculator()
	continueProcessing := true
	reader := bufio.NewReader(os.Stdin)
	var calculationResult string
	var line string
	var err error

	for continueProcessing {

		fmt.Println("Enter a new command? (Y/N)")

		line, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err) // should not be reached in standard usage - user input from console *should* end in \n
		}

		line = strings.TrimSpace(line)

		if line == "y" || line == "Y" {

			fmt.Println("Please enter a new calculation request")

			line, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err) // should not be reached in standard usage - user input from console *should* end in \n
			}

			if calc == nil { // if the calculator reference is nil for some reason, get a new one (safety check)
				calc = calculator.NewCalculator()
			}

			line = strings.TrimSpace(line)
			calc.UpdateCommandString(line)

			calculationResult, err = calc.GetResult()

			if err != nil {
				fmt.Printf("Error calculating result result: %s\n", err)
			} else {
				fmt.Printf("Result: %s\n", calculationResult)
			}

		} else if line == "n" || line == "N" {

			continueProcessing = false

		} else {

			fmt.Println("Invalid input")

		}

	}

}
