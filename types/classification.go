package types

import "fmt"

// ClassificationLevel is the sensitivity axis of a workspace's data
// classification. It is an ordered ladder stored as a slug. The fixed Nullstone
// taxonomy below is the single source of truth other services import.
type ClassificationLevel string

const (
	ClassificationPublic          ClassificationLevel = "public"
	ClassificationOperational     ClassificationLevel = "operational"
	ClassificationCustomerContent ClassificationLevel = "customer-content"
	ClassificationRestricted      ClassificationLevel = "restricted"
	ClassificationCritical        ClassificationLevel = "critical"
)

// classificationLevelInfo carries the fixed metadata for a sensitivity level.
type classificationLevelInfo struct {
	order int
	label string
	color string
}

// classificationLevels is the ordered taxonomy. Order, label, and color all live
// here so there is exactly one place that defines the standard.
var classificationLevels = []struct {
	slug ClassificationLevel
	classificationLevelInfo
}{
	{ClassificationPublic, classificationLevelInfo{0, "Public", "green"}},
	{ClassificationOperational, classificationLevelInfo{1, "Operational", "blue"}},
	{ClassificationCustomerContent, classificationLevelInfo{2, "Customer Content", "yellow"}},
	{ClassificationRestricted, classificationLevelInfo{3, "Restricted", "red"}},
	{ClassificationCritical, classificationLevelInfo{4, "Critical", "purple"}},
}

func (l ClassificationLevel) info() (classificationLevelInfo, bool) {
	for _, entry := range classificationLevels {
		if entry.slug == l {
			return entry.classificationLevelInfo, true
		}
	}
	return classificationLevelInfo{}, false
}

// Valid reports whether l is a known sensitivity level slug.
func (l ClassificationLevel) Valid() bool {
	_, ok := l.info()
	return ok
}

// Order returns the numeric position of the level in the ladder (0-4), or -1 if
// the level is unknown.
func (l ClassificationLevel) Order() int {
	if info, ok := l.info(); ok {
		return info.order
	}
	return -1
}

// Label returns the human-readable name of the level (e.g. "Customer Content").
func (l ClassificationLevel) Label() string {
	if info, ok := l.info(); ok {
		return info.label
	}
	return ""
}

// Color returns the fixed display color of the level (green/blue/yellow/red/purple).
func (l ClassificationLevel) Color() string {
	if info, ok := l.info(); ok {
		return info.color
	}
	return ""
}

// Composite returns the cloud-tag value for the level as "<#>-<slug>" (e.g.
// "2-customer-content"). This is the one place the level -> cloud-tag value
// mapping lives. An empty/unknown level returns "" so no tag is emitted.
func (l ClassificationLevel) Composite() string {
	info, ok := l.info()
	if !ok {
		return ""
	}
	return fmt.Sprintf("%d-%s", info.order, l)
}

// AllClassificationLevels returns the taxonomy in ladder order.
func AllClassificationLevels() []ClassificationLevel {
	out := make([]ClassificationLevel, len(classificationLevels))
	for i, entry := range classificationLevels {
		out[i] = entry.slug
	}
	return out
}
