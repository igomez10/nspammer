package nspammer

import (
	"math"
	"strings"
)

type SpamClassifier struct {
	Dataset map[string]bool
}

func (s *SpamClassifier) Classify(input string) bool {
	// calculate p(true)
	counterTrue := 0.0
	counterFalse := 0.0
	for _, v := range s.Dataset {
		if v == true {
			counterTrue += 1
		} else {
			counterFalse += 1
		}
	}

	// initialize dict
	totalPositiveWords := 0.0
	totalNegativeWords := 0.0
	vocab := map[string]bool{}
	positiveCount := map[string]float64{}
	negativeCount := map[string]float64{}
	for observation, isPositive := range s.Dataset {
		observationWords := strings.Split(observation, " ")
		for _, w := range observationWords {
			vocab[w] = true
			if isPositive {
				positiveCount[w] += 1
				totalPositiveWords += 1
			} else {
				negativeCount[w] += 1
				totalNegativeWords += 1
			}
		}
	}

	laplaceSmoothingConstant := 1.0

	// start with ptrue
	// this is the prior
	positiveScore := 0.0
	var pTrue float64 = counterTrue / float64(len(s.Dataset))
	positiveScore += math.Log(pTrue)
	// multiply each pWord|true to find the posterior
	// posterior is p(spam|words)
	for _, w := range strings.Split(input, " ") {
		numerator := positiveCount[w] + laplaceSmoothingConstant
		denominator := laplaceSmoothingConstant * float64(len(vocab))
		denominator += (totalPositiveWords)
		positiveScore += math.Log(numerator / denominator)
	}

	// calculate the other posterior
	// posterior p(nonspam|words)
	negativeScore := 0.0
	var pFalse float64 = counterFalse / float64(len(s.Dataset))
	negativeScore += math.Log(pFalse)
	for _, w := range strings.Split(input, " ") {
		numerator := negativeCount[w] + laplaceSmoothingConstant
		denominator := laplaceSmoothingConstant * float64(len(vocab))
		denominator += (totalNegativeWords)
		negativeScore += math.Log(numerator / denominator)
	}

	return positiveScore > negativeScore
}
