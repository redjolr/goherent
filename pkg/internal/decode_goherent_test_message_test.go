package internal

import "testing"

func Test_DecodeGoherentTetstMessage(t *testing.T) {
	type TestCase struct {
		name                    string
		expectedDecodedTestName string
		encodedGoherentTestName string
	}
	testCases := []TestCase{
		{
			name:                    "Empty_string",
			expectedDecodedTestName: "",
			encodedGoherentTestName: "",
		},
		{
			name:                    "Single_alpha_character",
			expectedDecodedTestName: "a",
			encodedGoherentTestName: "a",
		},
		{
			name:                    "Multiple_alpha_characters",
			expectedDecodedTestName: "aasdsa",
			encodedGoherentTestName: "aasdsa",
		},
		{
			name:                    "Multiple_alpha_characters_with_trailing_whitespace",
			expectedDecodedTestName: "aasdsa ",
			encodedGoherentTestName: "aasdsa" + ENCODED_WHITESPACE,
		},
		{
			name:                    "Multiple_alpha_characters_with_leading_whitespace",
			expectedDecodedTestName: " aasdsa",
			encodedGoherentTestName: ENCODED_WHITESPACE + "aasdsa",
		},
		{
			name:                    "Multiple_alpha_characters_with_2trailing_whitespaces",
			expectedDecodedTestName: "aasdsa  ",
			encodedGoherentTestName: "aasdsa" + ENCODED_WHITESPACE + ENCODED_WHITESPACE,
		},
		{
			name:                    "Multiple_alpha_characters_with_2leading_whitespaces",
			expectedDecodedTestName: "  aasdsa",
			encodedGoherentTestName: ENCODED_WHITESPACE + ENCODED_WHITESPACE + "aasdsa",
		},
		{
			name:                    "Multiple_alpha_characters_with_whitespace_in_between",
			expectedDecodedTestName: "aas dsa",
			encodedGoherentTestName: "aas" + ENCODED_WHITESPACE + "dsa",
		},
		{
			name:                    "Multiple_alpha_characters_with_3whitespaces_in_between",
			expectedDecodedTestName: "aas   dsa",
			encodedGoherentTestName: "aas" + ENCODED_WHITESPACE + ENCODED_WHITESPACE + ENCODED_WHITESPACE + "dsa",
		},
		{
			name:                    "Multiple_alpha_characters_with_whitespaces_in_between_leading_and_traling",
			expectedDecodedTestName: " aas dsa ",
			encodedGoherentTestName: ENCODED_WHITESPACE + "aas" + ENCODED_WHITESPACE + "dsa" + ENCODED_WHITESPACE,
		},
		{
			name:                    "Newline_only",
			expectedDecodedTestName: "\n",
			encodedGoherentTestName: ENCODED_NEWLINE,
		},
		{
			name:                    "Newline_only3",
			expectedDecodedTestName: "\n",
			encodedGoherentTestName: ENCODED_NEWLINE,
		},
		{
			name:                    "Multiple_alpha_chars_with_leading_newline",
			expectedDecodedTestName: "asd\n",
			encodedGoherentTestName: "asd" + ENCODED_NEWLINE,
		},
		{
			name:                    "Multiple_alpha_chars_with_2leading_newlines_and_2whitespaces",
			expectedDecodedTestName: "asd\n \n ",
			encodedGoherentTestName: "asd" + ENCODED_NEWLINE + ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE,
		},
		{
			name:                    "Multiple_alpha_chars_with_leading_trailing_in_between_newlines_and_whitespaces",
			expectedDecodedTestName: "\n asd \n \n dsa\n ",
			encodedGoherentTestName: ENCODED_NEWLINE + ENCODED_WHITESPACE +
				"asd" +
				ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE +
				"dsa" + ENCODED_NEWLINE + ENCODED_WHITESPACE,
		},
		{
			name:                    "Single_tab1",
			expectedDecodedTestName: "\t",
			encodedGoherentTestName: ENCODED_TAB,
		},
		{
			name:                    "Single_tab2",
			expectedDecodedTestName: "	",
			encodedGoherentTestName: ENCODED_TAB,
		},
		{
			name:                    "Multiple_alpha_characters_with_trailing_tab",
			expectedDecodedTestName: "aasdsa	",
			encodedGoherentTestName: "aasdsa" + ENCODED_TAB,
		},
		{
			name:                    "Multiple_alpha_characters_with_leading_tab",
			expectedDecodedTestName: "	aasdsa",
			encodedGoherentTestName: ENCODED_TAB + "aasdsa",
		},
		{
			name:                    "Multiple_alpha_characters_with_tab_in_between",
			expectedDecodedTestName: "aas\tdsa",
			encodedGoherentTestName: "aas" + ENCODED_TAB + "dsa",
		},
		{
			name:                    "Multiple_alpha_chars_with_newlines_tabs_and_whitespaces",
			expectedDecodedTestName: "\t \nasd \t\ndsa\n\t ",
			encodedGoherentTestName: ENCODED_TAB + ENCODED_WHITESPACE + ENCODED_NEWLINE +
				"asd" +
				ENCODED_WHITESPACE + ENCODED_TAB + ENCODED_NEWLINE +
				"dsa" + ENCODED_NEWLINE + ENCODED_TAB + ENCODED_WHITESPACE,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			decoded := DecodeGoherentTestMessage(testCase.encodedGoherentTestName)
			if decoded != testCase.expectedDecodedTestName {
				t.Errorf("Expected `%s`, got `%s`", testCase.expectedDecodedTestName, decoded)
			}
		})
	}
}
