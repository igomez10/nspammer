package nspammer

import (
	"math"
	"strings"
)

const laplaceSmoothingConstant = 1.0

type WordRecord struct {
	PositiveCount float64
	NegativeCount float64
}

func NewSpamClassifier(dataset map[string]bool) *SpamClassifier {
	s := &SpamClassifier{
		Dataset: dataset,
	}
	s.train()
	return s
}

type SpamClassifier struct {
	Dataset map[string]bool

	// Training phase results
	TotalWordsInPositive     float64
	TotalWordsInNegative     float64
	laplaceSmoothingConstant float64
	counterObservationsTrue  float64
	trained                  bool
	WordCounts               map[string]WordRecord
}

// Train preprocesses the dataset and calculates all necessary probabilities
func (s *SpamClassifier) train() {
	// Calculate class priors p(spam) and p(not spam)
	for _, v := range s.Dataset {
		if v == true {
			s.counterObservationsTrue += 1
		}
	}

	// Build vocabulary and count word occurrences
	s.TotalWordsInPositive = 0.0
	s.TotalWordsInNegative = 0.0
	s.WordCounts = map[string]WordRecord{}

	for observation, isPositiveObservation := range s.Dataset {
		observationWords := strings.Split(observation, " ")
		for _, w := range observationWords {
			if _, exists := s.WordCounts[w]; !exists {
				s.WordCounts[w] = WordRecord{
					PositiveCount: 0,
					NegativeCount: 0,
				}
			}

			if isPositiveObservation {
				s.WordCounts[w] = WordRecord{
					PositiveCount: s.WordCounts[w].PositiveCount + 1,
					NegativeCount: s.WordCounts[w].NegativeCount,
				}
				s.TotalWordsInPositive += 1
			} else {
				s.WordCounts[w] = WordRecord{
					PositiveCount: s.WordCounts[w].PositiveCount,
					NegativeCount: s.WordCounts[w].NegativeCount + 1,
				}
				s.TotalWordsInNegative += 1
			}
		}
	}

}

// Classify uses the trained model to classify input text as spam or not spam
func (s *SpamClassifier) Classify(input string) bool {

	// Calculate positive score: log(P(spam)) + sum(log(P(word|spam)))
	pTrue := s.counterObservationsTrue / float64(len(s.Dataset))
	positiveScore := math.Log(pTrue)
	for _, w := range strings.Split(input, " ") {
		numerator := s.WordCounts[w].PositiveCount + s.laplaceSmoothingConstant
		denominator := s.laplaceSmoothingConstant*float64(len(s.WordCounts)) + s.TotalWordsInPositive
		positiveScore += math.Log(numerator / denominator)
	}

	// Calculate negative score: log(P(not spam)) + sum(log(P(word|not spam)))
	pFalse := (float64(len(s.Dataset)) - s.counterObservationsTrue) / float64(len(s.Dataset))
	negativeScore := math.Log(pFalse)
	for _, w := range strings.Split(input, " ") {
		numerator := s.WordCounts[w].NegativeCount + s.laplaceSmoothingConstant
		denominator := s.laplaceSmoothingConstant*float64(len(s.WordCounts)) + s.TotalWordsInNegative
		negativeScore += math.Log(numerator / denominator)
	}

	return positiveScore > negativeScore
}
