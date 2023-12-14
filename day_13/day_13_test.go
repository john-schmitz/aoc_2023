package day_13

import "testing"

func TestPartOne(t *testing.T) {
	testCases := []struct {
		input_file string
		expected   int
	}{
		{
			input_file: "sample.txt",
			expected:   405,
		},
		{
			input_file: "input.txt",
			expected:   30535,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input_file, func(t *testing.T) {
			actual := PartOne(tC.input_file)

			if actual != tC.expected {
				t.Errorf("Expected PartOne(%s) = %d. Got %d", tC.input_file, tC.expected, actual)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		input_file string
		expected   int
	}{
		{
			input_file: "sample.txt",
			expected:   405,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.input_file, func(t *testing.T) {
			actual := PartTwo(tC.input_file)

			if actual != tC.expected {
				t.Errorf("Expected PartTwo(%s) = %d. Got %d", tC.input_file, tC.expected, actual)
			}
		})
	}
}

func TestFindReflection(t *testing.T) {
	testCases := []struct {
		grid     []string
		expected int
	}{
		{
			grid: []string{
				"#........",
				"....##...",
				"##......#",
				"##......#",
				"....##...",
				".........",
				"#...##...",
			},
			expected: 5,
		},
		{
			grid: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			expected: 5,
		},
		{
			grid: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			expected: 400,
		},
	}

	for _, tC := range testCases {
		t.Run("TestFindReflection", func(t *testing.T) {
			actual := FindReflection(tC.grid)

			if actual != tC.expected {
				t.Errorf("Expected PartTwo(%s) = %d. Got %d", tC.grid, tC.expected, actual)
			}
		})
	}
}
