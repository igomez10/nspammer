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
