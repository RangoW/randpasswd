package randpasswd

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	passwd, err := Generate(11)
	t.Log(err)
	t.Log(len(passwd))
	t.Log(passwd)
}
