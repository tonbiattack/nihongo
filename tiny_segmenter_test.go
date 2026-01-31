package nihongo

import (
	"strings"
	"testing"
)

func TestSegmenter(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{"empty", "", []string{}},
		{"simple-japanese", "私は人間です", []string{"私", "は", "人間", "です"}},
		{"katakana", "私はロボットです", []string{"私", "は", "ロボット", "です"}},
		{"english-sentence", "This is an English sentence.", []string{"This", " ", "is", " ", "an", " ", "English", " ", "sentence", "."}},
		{"date-and-ai", "2024年にAIが進化した。", []string{"2", "0", "2", "4", "年", "に", "AI", "が", "進化", "し", "た", "。"}},
		{"tokyo", "東京都に行った。", []string{"東京都", "に", "行っ", "た", "。"}},
		{"weather", "今日は良い天気です。", []string{"今日", "は", "良い", "天気", "です", "。"}},
		{"mixed-kana-kanji", "カタカナとひらがなと漢字", []string{"カタカナ", "と", "ひら", "が", "なと", "漢字"}},
		{"halfwidth-kana", "半角ｶﾅと全角カナ", []string{"半角", "ｶﾅ", "と", "全角", "カナ"}},
		{"go-language", "Go言語は楽しい!", []string{"Go", "言語", "は", "楽しい", "!"}},
		{"hello-world", "Hello,世界!", []string{"Hello", ",", "世界", "!"}},
		{"tongue-twister", "すもももももももものうち", []string{"すも", "も", "も", "も", "も", "も", "も", "もの", "うち"}},
		{"middle-dot", "テスト・ケース", []string{"テスト", "・", "ケース"}},
		{"price", "価格は3,500円です", []string{"価格", "は", "3", ",", "5", "0", "0", "円", "です"}},
		{"slash", "A/Bテスト", []string{"A", "/", "B", "テスト"}},
		{"emoji", "😀絵文字テスト😀", []string{"😀", "絵文字", "テスト", "😀"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertTokens(t, tt.input, tt.want)
		})
	}
}

func TestTokenize_Properties(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"spaces", "  "},
		{"ascii", "abcXYZ"},
		{"numbers", "12345"},
		{"kana", "てすとテスト"},
		{"kanji", "国際連合"},
		{"mixed", "Go言語で2024年を学ぶ"},
		{"symbols", "。！？・,.-/"},
		{"emoji", "😀😺👍"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tokenize(tt.input)
			if tt.input == "" && len(got) != 0 {
				t.Fatalf("Tokenize(\"\") should return empty slice, got %v", got)
			}
			if tt.input != "" && len(got) == 0 {
				t.Fatalf("Tokenize(%q) returned empty slice", tt.input)
			}
			for i, token := range got {
				if token == "" {
					t.Fatalf("Tokenize(%q) contains empty token at index %d", tt.input, i)
				}
			}
			joined := strings.Join(got, "")
			if joined != tt.input {
				t.Fatalf("Tokenize(%q) tokens joined to %q", tt.input, joined)
			}
		})
	}
}

func TestTokenize_Deterministic(t *testing.T) {
	input := "私は人間です"
	first := Tokenize(input)
	second := Tokenize(input)
	testEqual(t, first, second)
}

func TestGetCType(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"hiragana", "あ", "I"},
		{"katakana", "ア", "K"},
		{"halfwidth-katakana", "ｱ", "K"},
		{"ascii", "A", "A"},
		{"number", "5", "N"},
		{"kanji", "龠", "H"},
		{"other", "😀", "O"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCType(tt.input)
			if got != tt.want {
				t.Fatalf("getCType(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func testEqual(t *testing.T, a1 []string, a2 []string) {
	if len(a1) != len(a2) {
		t.Errorf("Segments should be equals %q <-> %q", a1, a2)
		return
	}
	for i, str := range a1 {
		if str != a2[i] {
			t.Errorf("Index %d is not equals for %q <-> %q", i, a1, a2)
		}
	}
}

func assertTokens(t *testing.T, input string, want []string) {
	t.Helper()
	got := Tokenize(input)
	testEqual(t, got, want)
}
