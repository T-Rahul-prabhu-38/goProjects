package main

import (
	"encoding/csv" // Package to work with CSV files
	"flag"         // Package to handle command-line flags
	"fmt"          // Package for formatted I/O
	"os"           // Package for file operations and program control
)

func main() {
	// Define a flag to accept the CSV file name from the command line.
	// Default is "problems.csv". Description is for the help message.
	csvFileName := flag.String("csv", "problems.csv", "a file with qna")
	// Parse the command-line flags so the program can access their values.
	flag.Parse()

	// Open the CSV file specified by the flag.
	// Dereference `csvFileName` because it is a pointer.
	file, err := os.Open(*csvFileName)
	if err != nil {
		// If the file cannot be opened, print an error message and exit the program.
		exit(fmt.Sprintf("failed to open the CSV file %s", *csvFileName))
	}

	// Create a new CSV reader to read the file line by line.
	r := csv.NewReader(file)
	// Read all lines from the CSV file into a slice of string slices.
	lines, err := r.ReadAll()
	if err != nil {
		// If the CSV file cannot be parsed, print an error message and exit the program.
		exit("failed to parse the provided csv file")
	}

	// Parse the CSV lines into a slice of `problem` structs.
	problems := parseLines(lines)
	// Print the parsed problems (for debugging or verification purposes).
	fmt.Println(problems)

	// Initialize a counter for correct answers.
	var count int = 0

	// Iterate through each problem in the slice.
	for i, p := range problems {
		// Display the problem to the user with a formatted string.
		fmt.Printf("problem #%d :  %s = \n", i+1, p.q)
		// Variable to store the user's answer.
		var answer string
		// Read the user's input (answer) from the console.
		fmt.Scanf("%s", &answer)
		// Compare the user's answer with the correct answer.
		if answer == p.a {
			// If correct, increment the counter and print a success message.
			fmt.Println("correct answer")
			count += 1
		} else {
			// If incorrect, print a failure message.
			fmt.Println("incorrect answer")
		}
	}

	// After all problems are answered, display the total number of correct answers.
	fmt.Printf("correct answers are : %d\n", count)
}

// Helper function to convert lines from the CSV into a slice of `problem` structs.
func parseLines(lines [][]string) []problem {
	// Create a slice of `problem` structs with the same length as the input.
	ret := make([]problem, len(lines))

	// Loop through each line and populate the slice of `problem` structs.
	for i, line := range lines {
		// Each line should have two entries: question (q) and answer (a).
		ret[i] = problem{
			q: line[0], // First column is the question.
			a: line[1], // Second column is the answer.
		}
	}

	// Return the slice of `problem` structs.
	return ret
}

// Struct to represent a single problem with a question and an answer.
type problem struct {
	q string // The question text.
	a string // The answer text.
}

// Helper function to handle errors by printing a message and exiting the program.
func exit(msg string) {
	// Print the error message to the console.
	fmt.Println(msg)
	// Exit the program with a non-zero status (indicating failure).
	os.Exit(1)
}
