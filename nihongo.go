package nihongo

import (
	"bytes"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

// Normalize は日本語テキストを NFKC 正規化する。
// 例: 半角カナ -> 全角カナ、全角記号の一部 -> 半角記号に統一される。
// 見た目の揺れを減らし、後続の判定や変換の前処理として使う。
func Normalize(text string) string {
	return norm.NFKC.String(text)
}

// ContainsHiragana は文字列内にひらがなが含まれていれば true を返す。
// 1 文字ずつ rune で走査し、unicode.Hiragana に属するか判定する。
func ContainsHiragana(text string) bool {
	for _, r := range text {
		// unicode.In は対象 rune が指定した Unicode 範囲に含まれるかを判定する。
		if unicode.In(r, unicode.Hiragana) {
			return true
		}
	}
	return false
}

// ContainsKatakana は文字列内にカタカナが含まれていれば true を返す。
// 1 文字ずつ rune で走査し、unicode.Katakana に属するか判定する。
func ContainsKatakana(text string) bool {
	for _, r := range text {
		// unicode.In は対象 rune が指定した Unicode 範囲に含まれるかを判定する。
		if unicode.In(r, unicode.Katakana) {
			return true
		}
	}
	return false
}

// ToHiragana はカタカナをひらがなに変換する。
// 事前に Normalize で正規化しておくと、半角カナなども確実に変換できる。
func ToHiragana(text string) string {
	var buf bytes.Buffer
	for _, r := range text {
		if unicode.In(r, unicode.Katakana) {
			// カタカナとひらがなは Unicode 上で 0x60 だけ離れて配置されている。
			// そのためコードポイントを 0x60 減算すると対応するひらがなになる。
			r -= 0x60
		}
		// 変換対象外の文字（漢字・英数字・記号など）はそのまま保持する。
		buf.WriteRune(r)
	}
	return buf.String()
}

// ToKatakana はひらがなをカタカナに変換する。
// 事前に Normalize で正規化しておくと、濁点の分解などの揺れが抑えられる。
func ToKatakana(text string) string {
	var buf bytes.Buffer
	//buf := bytes.NewBuffer(make([]byte, len(text)))
	for _, r := range text {
		if unicode.In(r, unicode.Hiragana) {
			// ひらがなとカタカナは Unicode 上で 0x60 だけ離れて配置されている。
			// そのためコードポイントを 0x60 加算すると対応するカタカナになる。
			r += 0x60
		}
		// 変換対象外の文字はそのまま書き込む。
		buf.WriteRune(r)
	}
	return buf.String()
}
