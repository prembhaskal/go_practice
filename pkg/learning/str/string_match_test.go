package str

import(
	"testing"
)

func TestKMPMatch(t *testing.T) {
	text := "ABABDABACDABABCABAB"
	pat := "ABABCABAB"
	ans := kmpStringMatch(text, pat)
	t.Errorf("answer is %t", ans)

	text = "ABABDABACDABABCABAX"
	// pat := "ABABCABAB"
	ans = kmpStringMatch(text, pat)
	t.Errorf("answer is %t", ans)
}