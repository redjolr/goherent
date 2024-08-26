package internal

import (
	"strings"
	"testing"
)

func Test_EncodeGoherentTestName(t *testing.T) {
	type TestCase struct {
		name                      string
		goherentTestName          string
		expectedEncodedGoTestName string
	}
	testCases := []TestCase{
		{
			name:                      "Empty_string",
			goherentTestName:          "",
			expectedEncodedGoTestName: "",
		},
		{
			name:                      "Single_alpha_character",
			goherentTestName:          "a",
			expectedEncodedGoTestName: "a",
		},
		{
			name:                      "Multiple_alpha_characters",
			goherentTestName:          "aasdsa",
			expectedEncodedGoTestName: "aasdsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_trailing_whitespace",
			goherentTestName:          "aasdsa ",
			expectedEncodedGoTestName: "aasdsa" + ENCODED_WHITESPACE,
		},
		{
			name:                      "Multiple_alpha_characters_with_leading_whitespace",
			goherentTestName:          " aasdsa",
			expectedEncodedGoTestName: ENCODED_WHITESPACE + "aasdsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_2trailing_whitespaces",
			goherentTestName:          "aasdsa  ",
			expectedEncodedGoTestName: "aasdsa" + ENCODED_WHITESPACE + ENCODED_WHITESPACE,
		},
		{
			name:                      "Multiple_alpha_characters_with_2leading_whitespaces",
			goherentTestName:          "  aasdsa",
			expectedEncodedGoTestName: ENCODED_WHITESPACE + ENCODED_WHITESPACE + "aasdsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_whitespace_in_between",
			goherentTestName:          "aas dsa",
			expectedEncodedGoTestName: "aas" + ENCODED_WHITESPACE + "dsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_3whitespaces_in_between",
			goherentTestName:          "aas   dsa",
			expectedEncodedGoTestName: "aas" + ENCODED_WHITESPACE + ENCODED_WHITESPACE + ENCODED_WHITESPACE + "dsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_whitespaces_in_between_leading_and_traling",
			goherentTestName:          " aas dsa ",
			expectedEncodedGoTestName: ENCODED_WHITESPACE + "aas" + ENCODED_WHITESPACE + "dsa" + ENCODED_WHITESPACE,
		},
		{
			name:                      "Newline_only",
			goherentTestName:          "\n",
			expectedEncodedGoTestName: ENCODED_NEWLINE,
		},
		{
			name: "Newline_only2",
			goherentTestName: `
`,
			expectedEncodedGoTestName: ENCODED_NEWLINE,
		},
		{
			name:                      "Newline_only3",
			goherentTestName:          "\r\n",
			expectedEncodedGoTestName: ENCODED_NEWLINE,
		},
		{
			name:                      "Multiple_alpha_chars_with_leading_newline",
			goherentTestName:          "asd\n",
			expectedEncodedGoTestName: "asd" + ENCODED_NEWLINE,
		},
		{
			name:                      "Multiple_alpha_chars_with_2leading_newlines_and_2whitespaces",
			goherentTestName:          "asd\n \n ",
			expectedEncodedGoTestName: "asd" + ENCODED_NEWLINE + ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE,
		},
		{
			name:             "Multiple_alpha_chars_with_leading_trailing_in_between_newlines_and_whitespaces",
			goherentTestName: "\n asd \n \n dsa\n ",
			expectedEncodedGoTestName: ENCODED_NEWLINE + ENCODED_WHITESPACE +
				"asd" +
				ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE +
				"dsa" + ENCODED_NEWLINE + ENCODED_WHITESPACE,
		},
		{
			name:                      "Single_tab1",
			goherentTestName:          "\t",
			expectedEncodedGoTestName: strings.Repeat(ENCODED_WHITESPACE, 4),
		},
		{
			name:                      "Single_tab2",
			goherentTestName:          "	",
			expectedEncodedGoTestName: strings.Repeat(ENCODED_WHITESPACE, 4),
		},
		{
			name:                      "Multiple_alpha_characters_with_trailing_tab",
			goherentTestName:          "aasdsa	",
			expectedEncodedGoTestName: "aasdsa" + strings.Repeat(ENCODED_WHITESPACE, 4),
		},
		{
			name:                      "Multiple_alpha_characters_with_leading_tab",
			goherentTestName:          "	aasdsa",
			expectedEncodedGoTestName: strings.Repeat(ENCODED_WHITESPACE, 4) + "aasdsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_tab_in_between",
			goherentTestName:          "aas\tdsa",
			expectedEncodedGoTestName: "aas" + strings.Repeat(ENCODED_WHITESPACE, 4) + "dsa",
		},
		{
			name:             "Multiple_alpha_chars_with_newlines_tabs_and_whitespaces",
			goherentTestName: "\t \nasd \t\ndsa\n\t ",
			expectedEncodedGoTestName: strings.Repeat(ENCODED_WHITESPACE, 4) + ENCODED_WHITESPACE + ENCODED_NEWLINE +
				"asd" +
				ENCODED_WHITESPACE + strings.Repeat(ENCODED_WHITESPACE, 4) + ENCODED_NEWLINE +
				"dsa" + ENCODED_NEWLINE + strings.Repeat(ENCODED_WHITESPACE, 4) + ENCODED_WHITESPACE,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			encodedGoTestName := EncodeGoherentTestName(testCase.goherentTestName)
			if encodedGoTestName != testCase.expectedEncodedGoTestName {
				t.Errorf("Expected `%s`, got `%s`", testCase.expectedEncodedGoTestName, encodedGoTestName)
			}
		})
	}
}
