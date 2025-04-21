package proxy

import (
	"bytes"
	_ "embed"
	"testing"

	koutube_conv "koutube-tg-reply/internal/koutube-conv"
)

//go:embed test_data/example.html
var exampleHTML []byte

func TestExampleHTML(t *testing.T) {
	s := NewServer("8080", koutube_conv.NewConverter())
	info, err := s.parseHTML(exampleHTML)
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	if err := buildOG(buf, info); err != nil {
		t.Fatal(err)
	}
	t.Logf("buf: %s", buf.String())
}
