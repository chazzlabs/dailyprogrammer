package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
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

func round(f float64, places int) (float64) {
    shift := math.Pow(10, float64(places))
    return math.Floor(f * shift) / shift;
}

func convertToFloatArray(stringValues []string) []float64 {
    var floatValues []float64

    for _, stringValue := range stringValues {
        floatValue, err := strconv.ParseFloat(stringValue, 64)

        if err != nil {
            break
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
