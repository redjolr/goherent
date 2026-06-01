package utils

// DisplayWidth returns the number of terminal columns a string occupies when
// printed, summing the cell width of each rune: 2 for East Asian Wide/Fullwidth
// characters and emoji (⏳ ✅ ❌ ⏩ 🚀 📦 📋, CJK), 1 otherwise.
//
// Terminal cursor movement is measured in cells, not runes, so any code that
// computes how far to move the cursor over printed text must size with this
// rather than len() / rune counts. This is a pragmatic subset of Unicode's East
// Asian Width / emoji data — enough for the glyphs this project renders. For
// production-grade coverage use github.com/mattn/go-runewidth.
func DisplayWidth(s string) int {
	width := 0
	for _, r := range s {
		if isWideRune(r) {
			width += 2
		} else {
			width += 1
		}
	}
	return width
}

func isWideRune(r rune) bool {
	switch {
	// Emoji with default emoji presentation that live in the Misc Technical,
	// Dingbats and Misc Symbols blocks (covers ⏩ U+23E9, ⏳ U+23F3, ✅ U+2705,
	// ❌ U+274C). These blocks also contain narrow symbols, so they are matched
	// point-by-point rather than as whole ranges.
	case r == 0x231A || r == 0x231B,
		r >= 0x23E9 && r <= 0x23F3,
		r >= 0x23F8 && r <= 0x23FA,
		r == 0x2705,
		r >= 0x270A && r <= 0x270B,
		r == 0x2728,
		r == 0x274C,
		r == 0x274E,
		r >= 0x2753 && r <= 0x2755,
		r == 0x2757,
		r >= 0x2795 && r <= 0x2797,
		r == 0x27B0,
		r == 0x27BF:
		return true
	// East Asian Wide / Fullwidth blocks (CJK and friends).
	case r >= 0x1100 && r <= 0x115F, // Hangul Jamo
		r >= 0x2E80 && r <= 0x303E, // CJK radicals .. CJK symbols
		r >= 0x3041 && r <= 0x33FF, // Hiragana .. CJK compatibility
		r >= 0x3400 && r <= 0x4DBF, // CJK Unified Ideographs Ext A
		r >= 0x4E00 && r <= 0x9FFF, // CJK Unified Ideographs
		r >= 0xA000 && r <= 0xA4CF, // Yi
		r >= 0xAC00 && r <= 0xD7A3, // Hangul Syllables
		r >= 0xF900 && r <= 0xFAFF, // CJK Compatibility Ideographs
		r >= 0xFE30 && r <= 0xFE4F, // CJK Compatibility Forms
		r >= 0xFF00 && r <= 0xFF60, // Fullwidth Forms
		r >= 0xFFE0 && r <= 0xFFE6: // Fullwidth signs
		return true
	// Emoji & Pictographs and supplementary ideographs.
	case r >= 0x1F300 && r <= 0x1FAFF, // covers 🚀 📦 📋 and most emoji
		r >= 0x20000 && r <= 0x3FFFD: // CJK Unified Ideographs Ext B+
		return true
	}
	return false
}
