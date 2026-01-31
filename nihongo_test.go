package nihongo

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"no-change", "テストてすと", "テストてすと"},
		{"halfwidth-kana-and-symbols", "テストﾃｽﾄ／＋", "テストテスト/+"},
		{"halfwidth-kana", "ﾊﾝｶｸｶﾀｶﾅ", "ハンカクカタカナ"},
		{"fullwidth-ascii", "ＡＢＣ１２３", "ABC123"},
		{"halfwidth-kana-dakuten", "ｶﾞｷﾞｸﾞｹﾞｺﾞ", "ガギグゲゴ"},
		{"combining-dakuten", "ばびぶべぼ", "ばびぶべぼ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Normalize(tt.input)
			if got != tt.want {
				t.Errorf("Normalize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToHiragana(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"katakana-only", "テスト", "てすと"},
		{"mixed-kana", "テスト混合てすと", "てすと混合てすと"},
		{"mixed-english", "Englishテスト混合", "Englishてすと混合"},
		{"katakana-chart", "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ", "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽ"},
		{"prolonged-and-numbers", "テストー123", "てすとー123"},
		{"halfwidth-with-normalize", Normalize("ﾃｽﾄｶﾅ"), "てすとかな"},
		{"non-kana-unchanged", "漢字ABC", "漢字ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToHiragana(tt.input)
			if got != tt.want {
				t.Errorf("ToHiragana(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToKatakana(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"hiragana-only", "てすと", "テスト"},
		{"mixed-kana", "てすと混合テスト", "テスト混合テスト"},
		{"mixed-english", "てすと混合English", "テスト混合English"},
		{"hiragana-chart", "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぷぺぽ", "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ"},
		{"prolonged-and-numbers", "てすとー123", "テストー123"},
		{"halfwidth-with-normalize", Normalize("てすとｶﾅ"), "テストカナ"},
		{"non-kana-unchanged", "漢字ABC", "漢字ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToKatakana(tt.input)
			if got != tt.want {
				t.Errorf("ToKatakana(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestContainsHiragana(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty", "", false},
		{"katakana", "テスト", false},
		{"english", "English", false},
		{"hiragana", "てすと", true},
		{"mixed", "テストてすと", true},
		{"numbers", "123", false},
		{"symbols", "。、・", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkContainsHiragana(t, tt.input, tt.expected)
		})
	}
}

func TestContainsKatakana(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty", "", false},
		{"katakana", "テスト", true},
		{"english", "English", false},
		{"hiragana", "てすと", false},
		{"mixed", "テストてすと", true},
		{"numbers", "123", false},
		{"symbols", "。、・", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkContainsKatakana(t, tt.input, tt.expected)
		})
	}
}

func checkContainsHiragana(t *testing.T, text string, expected bool) {
	contains := ContainsHiragana(text)
	if contains == expected {
		return
	}
	if contains {
		t.Errorf("ContainsHiragana detected hiragana on %v", text)
	} else {
		t.Errorf("ContainsHiragana did not detect hiragana on %v", text)
	}
}

func checkContainsKatakana(t *testing.T, text string, expected bool) {
	contains := ContainsKatakana(text)
	if contains == expected {
		return
	}
	if contains {
		t.Errorf("ContainsKatakana detected katakana on %v", text)
	} else {
		t.Errorf("ContainsKatakana did not detect katakana on %v", text)
	}
}
