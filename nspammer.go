package nspammer

import (
	"math"
	"strings"
)

type SpamClassifier struct {
	Dataset map[string]bool

	// Training phase results
	vocab                    map[string]bool
	positiveCount            map[string]float64
	negativeCount            map[string]float64
	totalPositiveWords       float64
	totalNegativeWords       float64
	pTrue                    float64
	pFalse                   float64
	laplaceSmoothingConstant float64
	trained                  bool
}

// Train preprocesses the dataset and calculates all necessary probabilities
func (s *SpamClassifier) Train() {
	// Calculate class priors p(spam) and p(not spam)
	counterTrue := 0.0
	counterFalse := 0.0
	for _, v := range s.Dataset {
		if v == true {
			counterTrue += 1
		} else {
			counterFalse += 1
		}
	}

	s.pTrue = counterTrue / float64(len(s.Dataset))
	s.pFalse = counterFalse / float64(len(s.Dataset))

	// Build vocabulary and count word occurrences
	s.vocab = map[string]bool{}
	s.totalPositiveWords = 0.0
	s.totalNegativeWords = 0.0
	s.positiveCount = map[string]float64{}
	s.negativeCount = map[string]float64{}

	for observation, isPositive := range s.Dataset {
		observationWords := strings.Split(observation, " ")
		for _, w := range observationWords {
			s.vocab[w] = true
			if isPositive {
				s.positiveCount[w] += 1
				s.totalPositiveWords += 1
			} else {
				s.negativeCount[w] += 1
				s.totalNegativeWords += 1
			}
		}
	}

	s.laplaceSmoothingConstant = 1.0
	s.trained = true
}

// Classify uses the trained model to classify input text as spam or not spam
func (s *SpamClassifier) Classify(input string) bool {
	// Train automatically if not already trained (for backward compatibility)
	if !s.trained {
		s.Train()
	}

	// Calculate positive score: log(P(spam)) + sum(log(P(word|spam)))
	positiveScore := math.Log(s.pTrue)
	for _, w := range strings.Split(input, " ") {
		numerator := s.positiveCount[w] + s.laplaceSmoothingConstant
		denominator := s.laplaceSmoothingConstant*float64(len(s.vocab)) + s.totalPositiveWords
		positiveScore += math.Log(numerator / denominator)
	}

	// Calculate negative score: log(P(not spam)) + sum(log(P(word|not spam)))
	negativeScore := math.Log(s.pFalse)
	for _, w := range strings.Split(input, " ") {
		numerator := s.negativeCount[w] + s.laplaceSmoothingConstant
		denominator := s.laplaceSmoothingConstant*float64(len(s.vocab)) + s.totalNegativeWords
		negativeScore += math.Log(numerator / denominator)
	}

	return positiveScore > negativeScore
}
