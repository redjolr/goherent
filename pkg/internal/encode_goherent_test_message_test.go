package internal

import "testing"

func Test_MessageToTestName(t *testing.T) {
	type TestCase struct {
		name                      string
		goherentTestMessage       string
		expectedEncodedGoTestName string
	}
	testCases := []TestCase{
		{
			name:                      "Empty_string",
			goherentTestMessage:       "",
			expectedEncodedGoTestName: "",
		},
		{
			name:                      "Single_alpha_character",
			goherentTestMessage:       "a",
			expectedEncodedGoTestName: "a",
		},
		{
			name:                      "Multiple_alpha_characters",
			goherentTestMessage:       "aasdsa",
			expectedEncodedGoTestName: "aasdsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_trailing_whitespace",
			goherentTestMessage:       "aasdsa ",
			expectedEncodedGoTestName: "aasdsa" + ENCODED_WHITESPACE,
		},
		{
			name:                      "Multiple_alpha_characters_with_leading_whitespace",
			goherentTestMessage:       " aasdsa",
			expectedEncodedGoTestName: ENCODED_WHITESPACE + "aasdsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_2trailing_whitespaces",
			goherentTestMessage:       "aasdsa  ",
			expectedEncodedGoTestName: "aasdsa" + ENCODED_WHITESPACE + ENCODED_WHITESPACE,
		},
		{
			name:                      "Multiple_alpha_characters_with_2leading_whitespaces",
			goherentTestMessage:       "  aasdsa",
			expectedEncodedGoTestName: ENCODED_WHITESPACE + ENCODED_WHITESPACE + "aasdsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_whitespace_in_between",
			goherentTestMessage:       "aas dsa",
			expectedEncodedGoTestName: "aas" + ENCODED_WHITESPACE + "dsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_3whitespaces_in_between",
			goherentTestMessage:       "aas   dsa",
			expectedEncodedGoTestName: "aas" + ENCODED_WHITESPACE + ENCODED_WHITESPACE + ENCODED_WHITESPACE + "dsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_whitespaces_in_between_leading_and_traling",
			goherentTestMessage:       " aas dsa ",
			expectedEncodedGoTestName: ENCODED_WHITESPACE + "aas" + ENCODED_WHITESPACE + "dsa" + ENCODED_WHITESPACE,
		},
		{
			name:                      "Newline_only",
			goherentTestMessage:       "\n",
			expectedEncodedGoTestName: ENCODED_NEWLINE,
		},
		{
			name: "Newline_only2",
			goherentTestMessage: `
`,
			expectedEncodedGoTestName: ENCODED_NEWLINE,
		},
		{
			name:                      "Newline_only3",
			goherentTestMessage:       "\r\n",
			expectedEncodedGoTestName: ENCODED_NEWLINE,
		},
		{
			name:                      "Multiple_alpha_chars_with_leading_newline",
			goherentTestMessage:       "asd\n",
			expectedEncodedGoTestName: "asd" + ENCODED_NEWLINE,
		},
		{
			name:                      "Multiple_alpha_chars_with_2leading_newlines_and_2whitespaces",
			goherentTestMessage:       "asd\n \n ",
			expectedEncodedGoTestName: "asd" + ENCODED_NEWLINE + ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE,
		},
		{
			name:                "Multiple_alpha_chars_with_leading_trailing_in_between_newlines_and_whitespaces",
			goherentTestMessage: "\n asd \n \n dsa\n ",
			expectedEncodedGoTestName: ENCODED_NEWLINE + ENCODED_WHITESPACE +
				"asd" +
				ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE +
				"dsa" + ENCODED_NEWLINE + ENCODED_WHITESPACE,
		},
		{
			name:                      "Single_tab1",
			goherentTestMessage:       "\t",
			expectedEncodedGoTestName: ENCODED_TAB,
		},
		{
			name:                      "Single_tab2",
			goherentTestMessage:       "	",
			expectedEncodedGoTestName: ENCODED_TAB,
		},
		{
			name:                      "Multiple_alpha_characters_with_trailing_tab",
			goherentTestMessage:       "aasdsa	",
			expectedEncodedGoTestName: "aasdsa" + ENCODED_TAB,
		},
		{
			name:                      "Multiple_alpha_characters_with_leading_tab",
			goherentTestMessage:       "	aasdsa",
			expectedEncodedGoTestName: ENCODED_TAB + "aasdsa",
		},
		{
			name:                      "Multiple_alpha_characters_with_tab_in_between",
			goherentTestMessage:       "aas\tdsa",
			expectedEncodedGoTestName: "aas" + ENCODED_TAB + "dsa",
		},
		{
			name:                "Multiple_alpha_chars_with_newlines_tabs_and_whitespaces",
			goherentTestMessage: "\t \nasd \t\ndsa\n\t ",
			expectedEncodedGoTestName: ENCODED_TAB + ENCODED_WHITESPACE + ENCODED_NEWLINE +
				"asd" +
				ENCODED_WHITESPACE + ENCODED_TAB + ENCODED_NEWLINE +
				"dsa" + ENCODED_NEWLINE + ENCODED_TAB + ENCODED_WHITESPACE,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			encodedGoTestName := EncodeGoherentTestMessage(testCase.goherentTestMessage)
			if encodedGoTestName != testCase.expectedEncodedGoTestName {
				t.Errorf("Expected `%s`, got `%s`", testCase.expectedEncodedGoTestName, encodedGoTestName)
			}
		})
	}
}
