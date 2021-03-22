package examples_test

import (
	"fmt"

	"github.com/avito-tech/normalize"
)

func ExampleNormalize() {
	fuzzy := "VAG-1101"
	clean := normalize.Normalize(fuzzy)
	fmt.Print(clean) // Output: vag1101
}

func ExampleNormalizeMany() {
	manyFuzzy := []string{"VAG-1101", "VAG-1102"}
	manyClean := normalize.Many(manyFuzzy)
	fmt.Print(manyClean) // Output: [vag1101 vag1102]
}

func ExampleNormalize_withOptions() {
	fuzzy := "АВ-123"
	clean := normalize.Normalize(fuzzy, normalize.WithLowerCase(), normalize.WithCyrillicToLatinLookAlike())
	fmt.Print(clean)
	// Output: ab-123
}

func ExampleAreStringsSimilar() {
	// nolint:goconst
	fuzzy := "Hyundai-Kia"
	// nolint:goconst
	otherFuzzy := "HYUNDAI"
	similarityThreshold := 0.3
	result := normalize.AreStringsSimilar(fuzzy, otherFuzzy, similarityThreshold)
	fmt.Print(result)
	// Output: true
}

func ExampleAreStringsSimilar_withOptions() {
	fuzzy := "Hyundai-Kia"
	otherFuzzy := "HYUNDAI"
	similarityThreshold := 0.3
	result := normalize.AreStringsSimilar(fuzzy, otherFuzzy, similarityThreshold, normalize.WithLowerCase(), normalize.WithCyrillicToLatinLookAlike())
	fmt.Print(result)
	// Output: false
}
