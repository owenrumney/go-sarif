package test

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	go_sarif "github.com/owenrumney/go-sarif"
	"testing"
)

type sarifHarness struct {
	t           *testing.T
	sarifReport *go_sarif.SarifReport
	content     string
}

func (h *sarifHarness) a_new_sarif_report() {
	h.sarifReport = go_sarif.New()
}

func (h *sarifHarness) the_report_is_written_to_string() {
	buf := new(bytes.Buffer)
	err := h.sarifReport.Write(buf)
	if err != nil {
		h.t.Error(err)
	}
	h.content = buf.String()
}

func (h *sarifHarness) content_should_be(expected string) {
	assert.Equal(h.t, h.content, expected)
}

func CreateNewSarifHarness(t *testing.T) (*sarifHarness, *sarifHarness, *sarifHarness) {
	s := &sarifHarness{
		t: t,
	}
	return s, s, s
}
