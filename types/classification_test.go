package types

import "testing"

func TestClassificationLevelComposite(t *testing.T) {
	tests := []struct {
		level ClassificationLevel
		want  string
	}{
		{ClassificationPublic, "0-public"},
		{ClassificationOperational, "1-operational"},
		{ClassificationCustomerContent, "2-customer-content"},
		{ClassificationRestricted, "3-restricted"},
		{ClassificationCritical, "4-critical"},
		{ClassificationLevel(""), ""},
		{ClassificationLevel("bogus"), ""},
	}
	for _, tt := range tests {
		if got := tt.level.Composite(); got != tt.want {
			t.Errorf("ClassificationLevel(%q).Composite() = %q, want %q", tt.level, got, tt.want)
		}
	}
}

func TestClassificationLevelMetadata(t *testing.T) {
	if got := ClassificationCustomerContent.Order(); got != 2 {
		t.Errorf("Order() = %d, want 2", got)
	}
	if got := ClassificationCustomerContent.Label(); got != "Customer Content" {
		t.Errorf("Label() = %q, want %q", got, "Customer Content")
	}
	if got := ClassificationCustomerContent.Color(); got != "yellow" {
		t.Errorf("Color() = %q, want %q", got, "yellow")
	}
	if !ClassificationRestricted.Valid() {
		t.Error("expected restricted to be valid")
	}
	if ClassificationLevel("bogus").Valid() {
		t.Error("expected bogus to be invalid")
	}
	if got := ClassificationLevel("bogus").Order(); got != -1 {
		t.Errorf("unknown Order() = %d, want -1", got)
	}
	if got := len(AllClassificationLevels()); got != 5 {
		t.Errorf("AllClassificationLevels() len = %d, want 5", got)
	}
}
