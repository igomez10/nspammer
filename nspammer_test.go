package nspammer

import (
	"fmt"
	"testing"
)

func TestNoCode(t *testing.T) {
	n := 1000.0
	nspam := 200.0
	nnonspam := 800.0
	lipitorSpam := 20.0
	lipitorNonSpam := 1.0
	HelloSpam := 190.0
	HelloNonSpam := 700.0
	SirSpam := 10.0
	SirNonSpam := 400.0

	pSpam := nspam / n
	pNonSpam := nnonspam / n

	pLipitorSpam := lipitorSpam / nspam
	pHelloSpam := HelloSpam / nspam
	pSirSpam := SirSpam / nspam

	pLipitorNonSpam := lipitorNonSpam / nnonspam
	pHelloNonSpam := HelloNonSpam / nnonspam
	pSirNonSPam := SirNonSpam / nnonspam

	// posterior
	pSpamGivenLipitorHelloSir := pSpam * pLipitorSpam * pHelloSpam * pSirSpam
	fmt.Printf("%.10f\n", pSpamGivenLipitorHelloSir)

	pNonSpamGivenLipitorHelloSir := pNonSpam * pLipitorNonSpam * pHelloNonSpam * pSirNonSPam
	fmt.Printf("%.10f\n", pNonSpamGivenLipitorHelloSir)
}

func TestSpamClassifier_Classify(t *testing.T) {
	tests := []struct {
		name         string // description of this test case
		trainingData map[string]bool
		input        string
		want         bool
	}{
		{
			name: "simple case",
			trainingData: map[string]bool{
				"spam":    true,
				"spam2":   true,
				"spam3":   true,
				"notpsam": false,
			},
			input: "spam spam2 spam3",
			want:  true,
		},
		{
			name: "simple negative case",
			trainingData: map[string]bool{
				"spam":     true,
				"spam2":    true,
				"spam3":    true,
				"notpsam":  false,
				"notpsam2": false,
				"notpsam3": false,
				"notpsam4": false,
				"notpsam5": false,
			},
			input: "notpsam",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSpamClassifier(tt.trainingData)
			s.Dataset = tt.trainingData
			got := s.Classify(tt.input)
			if got != tt.want {
				t.Errorf("Classify() = %v, want %v", got, tt.want)
			}
		})
	}
}
