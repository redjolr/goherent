package internal

import "testing"

func Test_DecodeGoherentTetstMessage(t *testing.T) {
	type TestCase struct {
		name                       string
		expectedDecodedMessage     string
		encodedGoherentTestMessage string
	}
	testCases := []TestCase{
		{
			name:                       "Empty_string",
			expectedDecodedMessage:     "",
			encodedGoherentTestMessage: "",
		},
		{
			name:                       "Single_alpha_character",
			expectedDecodedMessage:     "a",
			encodedGoherentTestMessage: "a",
		},
		{
			name:                       "Multiple_alpha_characters",
			expectedDecodedMessage:     "aasdsa",
			encodedGoherentTestMessage: "aasdsa",
		},
		{
			name:                       "Multiple_alpha_characters_with_trailing_whitespace",
			expectedDecodedMessage:     "aasdsa ",
			encodedGoherentTestMessage: "aasdsa" + ENCODED_WHITESPACE,
		},
		{
			name:                       "Multiple_alpha_characters_with_leading_whitespace",
			expectedDecodedMessage:     " aasdsa",
			encodedGoherentTestMessage: ENCODED_WHITESPACE + "aasdsa",
		},
		{
			name:                       "Multiple_alpha_characters_with_2trailing_whitespaces",
			expectedDecodedMessage:     "aasdsa  ",
			encodedGoherentTestMessage: "aasdsa" + ENCODED_WHITESPACE + ENCODED_WHITESPACE,
		},
		{
			name:                       "Multiple_alpha_characters_with_2leading_whitespaces",
			expectedDecodedMessage:     "  aasdsa",
			encodedGoherentTestMessage: ENCODED_WHITESPACE + ENCODED_WHITESPACE + "aasdsa",
		},
		{
			name:                       "Multiple_alpha_characters_with_whitespace_in_between",
			expectedDecodedMessage:     "aas dsa",
			encodedGoherentTestMessage: "aas" + ENCODED_WHITESPACE + "dsa",
		},
		{
			name:                       "Multiple_alpha_characters_with_3whitespaces_in_between",
			expectedDecodedMessage:     "aas   dsa",
			encodedGoherentTestMessage: "aas" + ENCODED_WHITESPACE + ENCODED_WHITESPACE + ENCODED_WHITESPACE + "dsa",
		},
		{
			name:                       "Multiple_alpha_characters_with_whitespaces_in_between_leading_and_traling",
			expectedDecodedMessage:     " aas dsa ",
			encodedGoherentTestMessage: ENCODED_WHITESPACE + "aas" + ENCODED_WHITESPACE + "dsa" + ENCODED_WHITESPACE,
		},
		{
			name:                       "Newline_only",
			expectedDecodedMessage:     "\n",
			encodedGoherentTestMessage: ENCODED_NEWLINE,
		},
		{
			name:                       "Newline_only3",
			expectedDecodedMessage:     "\n",
			encodedGoherentTestMessage: ENCODED_NEWLINE,
		},
		{
			name:                       "Multiple_alpha_chars_with_leading_newline",
			expectedDecodedMessage:     "asd\n",
			encodedGoherentTestMessage: "asd" + ENCODED_NEWLINE,
		},
		{
			name:                       "Multiple_alpha_chars_with_2leading_newlines_and_2whitespaces",
			expectedDecodedMessage:     "asd\n \n ",
			encodedGoherentTestMessage: "asd" + ENCODED_NEWLINE + ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE,
		},
		{
			name:                   "Multiple_alpha_chars_with_leading_trailing_in_between_newlines_and_whitespaces",
			expectedDecodedMessage: "\n asd \n \n dsa\n ",
			encodedGoherentTestMessage: ENCODED_NEWLINE + ENCODED_WHITESPACE +
				"asd" +
				ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE + ENCODED_NEWLINE + ENCODED_WHITESPACE +
				"dsa" + ENCODED_NEWLINE + ENCODED_WHITESPACE,
		},
		{
			name:                       "Single_tab1",
			expectedDecodedMessage:     "\t",
			encodedGoherentTestMessage: ENCODED_TAB,
		},
		{
			name:                       "Single_tab2",
			expectedDecodedMessage:     "	",
			encodedGoherentTestMessage: ENCODED_TAB,
		},
		{
			name:                       "Multiple_alpha_characters_with_trailing_tab",
			expectedDecodedMessage:     "aasdsa	",
			encodedGoherentTestMessage: "aasdsa" + ENCODED_TAB,
		},
		{
			name:                       "Multiple_alpha_characters_with_leading_tab",
			expectedDecodedMessage:     "	aasdsa",
			encodedGoherentTestMessage: ENCODED_TAB + "aasdsa",
		},
		{
			name:                       "Multiple_alpha_characters_with_tab_in_between",
			expectedDecodedMessage:     "aas\tdsa",
			encodedGoherentTestMessage: "aas" + ENCODED_TAB + "dsa",
		},
		{
			name:                   "Multiple_alpha_chars_with_newlines_tabs_and_whitespaces",
			expectedDecodedMessage: "\t \nasd \t\ndsa\n\t ",
			encodedGoherentTestMessage: ENCODED_TAB + ENCODED_WHITESPACE + ENCODED_NEWLINE +
				"asd" +
				ENCODED_WHITESPACE + ENCODED_TAB + ENCODED_NEWLINE +
				"dsa" + ENCODED_NEWLINE + ENCODED_TAB + ENCODED_WHITESPACE,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			decoded := DecodeGoherentTestMessage(testCase.encodedGoherentTestMessage)
			if decoded != testCase.expectedDecodedMessage {
				t.Errorf("Expected `%s`, got `%s`", testCase.expectedDecodedMessage, decoded)
			}
		})
	}
}
