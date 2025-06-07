package stringer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type inner struct {
	Secret string `stringer:"masked,length=true"`
	ReallySecret string `stringer:"masked"`
	PointerToSecret *string `stringer:"masked"`
}

type testStruct struct {
	ID      string            `stringer:"include"`
	Secret  inner             `stringer:"nested"`
	Details map[string]string `stringer:"type"`
	SkipMe  string
	private string
}

func TestToStringWithTags(t *testing.T) {
	id := "abc123"
	secret := "sensitive"
	noshow := "no show"
	private := "shhhhh"
	obj := testStruct{
		ID:     id,
		Secret: inner{
			Secret: secret,
			PointerToSecret: nil,
		},
		Details: map[string]string{
			"foo": "bar",
		},
		SkipMe: noshow,
		private: private,
	}

	got := ToStringWithTags(obj)
	if got == "" {
		t.Errorf("expected string output, got empty string")
	}
	t.Logf("Output: %s", got)
	assert.Contains(t, got, fmt.Sprintf("%s,", id), "ID should be included")
	assert.Contains(t, got, fmt.Sprintf("***** (len=%d),", len(secret)), "Masked secret should be shown with length")
	assert.Contains(t, got, "\"\",", "Masked really secret should be shown as empty string without length")
	assert.Contains(t, got, "<nil>},", "Masked null secret should be shown as")
	assert.Contains(t, got, "<map[string]string>}", "Map type should be shown")
	assert.NotContains(t, got, noshow, "Unannotated fields should not be included")
	assert.NotContains(t, got, private, "Private fields should not be included")
}
