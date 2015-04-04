package gotween

import (
	"testing"
)

type easingTestData struct {
	input         float64
	output        float64
	errorExpected bool
}

type easingTest struct {
	easingFunc EasingFunc
	tests      []*easingTestData
}

var easingTests = []*easingTest{
	&easingTest{
		easingFunc: Linear,
		tests: []*easingTestData{
			&easingTestData{0.0, 0.0, false},
			&easingTestData{0.1, 0.1, false},
			&easingTestData{0.2, 0.2, false},
			&easingTestData{0.3, 0.3, false},
			&easingTestData{0.4, 0.4, false},
			&easingTestData{0.5, 0.5, false},
			&easingTestData{0.6, 0.6, false},
			&easingTestData{0.7, 0.7, false},
			&easingTestData{0.8, 0.8, false},
			&easingTestData{0.9, 0.9, false},
			&easingTestData{1.0, 1.0, false},
			&easingTestData{1.1, 0.0, true},
			&easingTestData{-0.1, 0.0, true},
		},
	},
}

// Test all easing functions with a simple table driven test.
func TestEasingFuncs(t *testing.T) {

	for _, test := range easingTests {
		for _, testData := range test.tests {
			output, err := test.easingFunc(testData.input)
			if err != nil {
				if !testData.errorExpected {
					t.Errorf("Unexpected error occured: %v", err)
				}
				continue
			}
			if testData.errorExpected {
				t.Error("Expected error but none occured")
				continue
			}
			if output != testData.output {
				t.Errorf("Output did not match expected: actual: %d, expected: %d", output, testData.output)
			}
		}
	}

}
