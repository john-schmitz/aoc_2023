package day14

import "testing"

func TestPartOne(t *testing.T) {
	testCases := []struct {
		file_path string
		expected  int
	}{
		{
			file_path: "sample.txt",
			expected:  136,
		},
		{
			file_path: "input.txt",
			expected:  136,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.file_path, func(t *testing.T) {
			actual := PartOne(tC.file_path)
			if actual != tC.expected {
				t.Errorf("Expected PartOne(%s) = %d. Got %d", tC.file_path, tC.expected, actual)
			}
		})
	}
}
