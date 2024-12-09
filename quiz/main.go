package main

import (
	"encoding/csv" // For working with CSV files
	"flag"         // For parsing command-line flags
	"fmt"          // For formatted I/O
	"os"           // For file operations and exiting the program
	"strings"      // For string manipulations
	"time"         // For working with timers
)

func main() {
	// Define command-line flags for CSV file name and time limit
	csvFileName := flag.String("csv", "problems.csv", "a file with qna")
	timeLimit := flag.Int("limit", 30, "the time limit in secs")

	// Parse the command-line flags
	flag.Parse()

	// Open the CSV file specified by the user
	file, err := os.Open(*csvFileName)
	if err != nil {
		// Exit the program if the file cannot be opened
		exit(fmt.Sprintf("failed to open the CSV file %s", *csvFileName))
	}

	// Create a new CSV reader
	r := csv.NewReader(file)
	// Read all lines from the CSV file
	lines, err := r.ReadAll()
	if err != nil {
		// Exit the program if the CSV file cannot be parsed
		exit("failed to parse the provided csv file")
	}

	// Parse the lines from the CSV into a slice of problem structs
	problems := parseLines(lines)
	// Print the parsed problems for debugging purposes
	fmt.Println(problems)

	// Create a timer that runs for the specified time limit
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// Counter to track the number of correct answers
	var count int = 0

	// Iterate over each problem
	for i, p := range problems {
		// Print the current problem question
		fmt.Printf("problem #%d :  %s = \n", i+1, p.q)
		// Create a channel to handle user input asynchronously
		answerCh := make(chan string)

		// Launch a goroutine to read the user's answer
		go func() {
			var answer string
			fmt.Scanf("%s", &answer) // Read user input
			answerCh <- answer       // Send the answer to the channel
		}()

		// Use a select statement to handle either timer expiration or user input
		select {
		case <-timer.C:
			// If the timer expires, print the score and exit
			fmt.Printf("\n correct answers are : %d/%d", count, len(problems))
			return
		case answer := <-answerCh:
			// If an answer is received, check if it's correct
			if answer == p.a {
				count += 1
			}
		default:
			// Default case does nothing (can be omitted in this context)
		}
	}

	// Print the final score after all problems have been attempted
	fmt.Printf("\n correct answers are : %d/%d", count, len(problems))
}

// parseLines converts raw CSV data into a slice of problem structs
func parseLines(lines [][]string) []problem {
	// Create a slice of problem structs with the same length as the CSV lines
	ret := make([]problem, len(lines))

	// Iterate over the CSV lines
	for i, line := range lines {
		// Populate the problem struct with trimmed question and answer
		ret[i] = problem{
			q: strings.TrimSpace(line[0]), // Trim whitespace from the question
			a: strings.TrimSpace(line[1]), // Trim whitespace from the answer
		}
	}

	return ret // Return the slice of problems
}



// problem struct represents a single question-answer pair
type problem struct {
	q string // Question
	a string // Answer
}

// exit prints an error message and terminates the program
func exit(msg string) {
	fmt.Println(msg) // Print the error message
	os.Exit(1)       // Exit the program with a non-zero status
}
