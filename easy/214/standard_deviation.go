package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
    "log"
)

func calculateMean(values []float64) float64 {
    sum := 0.0

    for _, value := range values {
        sum += value
    }

    return sum / float64(len(values))
}

func calculateVariance(values []float64, mean float64) float64 {
    var differences []float64

    for _, value := range values {
        difference := value - mean
        differences = append(differences, difference * difference)
    }

    return calculateMean(differences)
}

func calculateStdDeviation(values []float64) float64 {
    mean := calculateMean(values)
    variance := calculateVariance(values, mean)
    return math.Sqrt(variance)
}

// From: https://gist.github.com/indytechcook/5706434
func round(x float64, prec int) float64 {
    var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	intermed += .5
	x = .5
	if frac < 0.0 {
		x = -.5
		intermed -= 1
	}
	if frac >= x {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func convertToFloatArray(stringValues []string) []float64 {
    var floatValues []float64

    for _, stringValue := range stringValues {
        floatValue, err := strconv.ParseFloat(stringValue, 64)

        if err != nil {
            log.Fatal(err)
        }

        floatValues = append(floatValues, floatValue)
    }

    return floatValues
}

func main() {
    values := convertToFloatArray(os.Args[1:])

    standardDeviation := calculateStdDeviation(values)

    fmt.Println(round(standardDeviation, 4))
}
