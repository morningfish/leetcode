package convert

import "testing"

func TestConvert(t *testing.T) {
	s := convert("LEETCODEISHIRING", 3)
	if s != "LCIRETOESIIGEDHN" {
		t.Fail()
	}
	t.Log(s)
}
