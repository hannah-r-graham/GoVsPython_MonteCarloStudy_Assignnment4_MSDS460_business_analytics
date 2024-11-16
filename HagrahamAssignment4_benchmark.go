package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Mathematical functions
func add(a, b float64) float64 {
	return a + b
}

func multiply(a, b float64) float64 {
	return a * b
}

func exponent(a, b float64) float64 {
	return math.Pow(a, b)
}

func sqrt(a, b float64) float64 {
	return b * math.Sqrt(a)
}

func matrixMultiply(a, b [][]float64) ([][]float64, error) {
	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("number of columns in matrix a must be equal to the number of rows in matrix b")
	}

	result := make([][]float64, len(a))
	for i := range result {
		result[i] = make([]float64, len(b[0]))
	}

	for i := range a {
		for j := range b[0] {
			for k := range b {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result, nil
}

// String manipulation functions
func generateRandomString(size int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, size)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func concatenateStrings(a, b string) string {
	return a + b
}

func searchSubstring(s, sub string) bool {
	return strings.Contains(s, sub)
}

func regexMatch(s, pattern string) bool {
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}

// Monte Carlo simulation for mathematical functions
func MathmonteCarloSimulation(numSamples, matrixSize int, rng *rand.Rand) ([]float64, []float64, []float64, []float64, []float64) {
	addTimes := make([]float64, numSamples)
	multiplyTimes := make([]float64, numSamples)
	exponentTimes := make([]float64, numSamples)
	sqrtTimes := make([]float64, numSamples)
	matrixMultiplyTimes := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		a := rng.Float64()
		b := rng.Float64()
		A := generateRandomMatrix(matrixSize, matrixSize, rng)
		B := generateRandomMatrix(matrixSize, matrixSize, rng)

		startTime := time.Now()
		_ = add(a, b)
		addTimes[i] = time.Since(startTime).Seconds()

		startTime = time.Now()
		_ = multiply(a, b)
		multiplyTimes[i] = time.Since(startTime).Seconds()

		startTime = time.Now()
		_ = exponent(a, b)
		exponentTimes[i] = time.Since(startTime).Seconds()

		startTime = time.Now()
		_ = sqrt(a, b)
		sqrtTimes[i] = time.Since(startTime).Seconds()

		startTime = time.Now()
		_, err := matrixMultiply(A, B)
		if err != nil {
			fmt.Println("Matrix multiplication error:", err)
			continue
		}
		matrixMultiplyTimes[i] = time.Since(startTime).Seconds()
	}

	return addTimes, multiplyTimes, exponentTimes, sqrtTimes, matrixMultiplyTimes
}

// Monte Carlo simulation for string functions
func StringmonteCarloSimulation(numSamples, stringSize int) ([]float64, []float64, []float64) {
	concatTimes := make([]float64, numSamples)
	searchTimes := make([]float64, numSamples)
	regexTimes := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		a := generateRandomString(stringSize)
		b := generateRandomString(stringSize)
		pattern := `[A-Za-z0-9]{5}`

		startTime := time.Now()
		_ = concatenateStrings(a, b)
		concatTimes[i] = time.Since(startTime).Seconds()

		startTime = time.Now()
		_ = searchSubstring(a, "test")
		searchTimes[i] = time.Since(startTime).Seconds()

		startTime = time.Now()
		_ = regexMatch(a, pattern)
		regexTimes[i] = time.Since(startTime).Seconds()
	}

	return concatTimes, searchTimes, regexTimes
}

// Helper function to generate a random matrix
func generateRandomMatrix(rows, cols int, rng *rand.Rand) [][]float64 {
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
		for j := range matrix[i] {
			matrix[i][j] = rng.Float64()
		}
	}
	return matrix
}

// Helper function to calculate the sum of a slice of float64
func sum(times []float64) float64 {
	total := 0.0
	for _, t := range times {
		total += t
	}
	return total
}

// Execute math functions and string functions for Monte Carlo
func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	numSamples := 100000
	matrixSize := 10
	stringSize := 1000

	// Run Monte Carlo simulation for mathematical functions
	addTimes, multiplyTimes, exponentTimes, sqrtTimes, matrixMultiplyTimes := MathmonteCarloSimulation(numSamples, matrixSize, rng)

	// Run Monte Carlo simulation for string functions
	concatTimes, searchTimes, regexTimes := StringmonteCarloSimulation(numSamples, stringSize)

	// Print mathematical results
	fmt.Printf("Average Addition Time: %e seconds\n", sum(addTimes)/float64(numSamples))
	fmt.Printf("Average Multiplication Time: %e seconds\n", sum(multiplyTimes)/float64(numSamples))
	fmt.Printf("Average Exponentiation Time: %e seconds\n", sum(exponentTimes)/float64(numSamples))
	fmt.Printf("Average Square Root Time: %e seconds\n", sum(sqrtTimes)/float64(numSamples))
	fmt.Printf("Average Matrix Multiplication Time: %e seconds\n", sum(matrixMultiplyTimes)/float64(numSamples))

	// Print string manipulation results
	fmt.Printf("Average Concatenation Time: %e seconds\n", sum(concatTimes)/float64(numSamples))
	fmt.Printf("Average Search Time: %e seconds\n", sum(searchTimes)/float64(numSamples))
	fmt.Printf("Average Regex Match Time: %e seconds\n", sum(regexTimes)/float64(numSamples))

	// Create a struct to store results
	type Result struct {
		Operation   string
		AverageTime float64
	}

	// Combine results from both simulations
	results := []Result{
		{"Addition", sum(addTimes) / float64(numSamples)},
		{"Multiplication", sum(multiplyTimes) / float64(numSamples)},
		{"Exponentiation", sum(exponentTimes) / float64(numSamples)},
		{"Square Root", sum(sqrtTimes) / float64(numSamples)},
		{"Matrix Multiplication", sum(matrixMultiplyTimes) / float64(numSamples)},
		{"Concatenation", sum(concatTimes) / float64(numSamples)},
		{"Search", sum(searchTimes) / float64(numSamples)},
		{"Regex Match", sum(regexTimes) / float64(numSamples)},
	}

	fmt.Println("Results:")
	for _, result := range results {
		fmt.Printf("%s: %e seconds\n", result.Operation, result.AverageTime)
	}

	// Save results to CSV
	file, err := os.Create("GoCombinedResults.csv")
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Operation", "Average Time (seconds)"})
	for _, result := range results {
		writer.Write([]string{result.Operation, fmt.Sprintf("%e", result.AverageTime)})
	}

	// Plot results
	p := plot.New()
	p.Title.Text = "Average Execution Time for Go Mathematical and String Operations"
	p.X.Label.Text = "Operation"
	p.Y.Label.Text = "Average Time (seconds)"

	points := make(plotter.XYs, len(results))
	for i, result := range results {
		points[i].X = float64(i)
		points[i].Y = result.AverageTime
	}

	line, pointsPlot, err := plotter.NewLinePoints(points)
	if err != nil {
		log.Panic(err)
	}
	p.Add(line, pointsPlot)

	p.NominalX("Addition", "Multiplication", "Exponentiation", "Square Root", "Matrix Multiplication", "Concatenation", "Search", "Regex Match")

	if err := p.Save(10*vg.Inch, 6*vg.Inch, "GoCombinedResults.png"); err != nil {
		log.Panic(err)
	}
}
