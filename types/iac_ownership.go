package types

import (
	"regexp"
	"sort"
	"strings"
)

// IacOwnershipConflict is what an IaC sync failure says when the repository it ran for tried
// to update blocks or env events that a different repository owns. The sync reports this as
// plain text in the intent workflow's status message; there is no error code that survives
// the trip through the engine. MatchIacOwnershipConflict recovers the structure from that
// text so the CLI and UI can point the user at the settings page where a stack owner or
// architect can change the owning repository, without changing the error message itself.
//
// The producer is nullfire (internal/blocks/errors.go for blocks, iac/update_events.go for
// events); its tests assert its messages match here, so the two cannot drift silently.
type IacOwnershipConflict struct {
	// Blocks maps each conflicting block name to the repository (owner/name) that owns it.
	Blocks map[string]string
	// Events maps each conflicting event name to the repository url that owns it.
	Events map[string]string
}

var (
	// iacBlockConflictHeader matches the first line of a block ownership conflict, for both
	// the normal ("IaC sync from X cannot update ...") and the blank-repo variant ("IaC sync
	// did not identify its source repository, so it cannot update ...").
	iacBlockConflictHeader = regexp.MustCompile(`cannot update \d+ blocks? owned by (?:another repository|other repositories):`)
	// iacBlockConflictLine matches one "\t<repo> owns a, b, c" line under that header.
	iacBlockConflictLine = regexp.MustCompile(`(?m)^\t(\S+) owns (.+)$`)
	// iacEventConflictLine matches one `event "x" is configured in another repository (<url>)`
	// entry, wherever multierror places it.
	iacEventConflictLine = regexp.MustCompile(`event "([^"]+)" is configured in another repository \(([^)]*)\)`)
)

// MatchIacOwnershipConflict returns the blocks and events an IaC sync failure message says
// are owned by another repository, or nil when the message is not an ownership conflict.
// The message is searched, not anchored: callers may have wrapped it with extra lines.
func MatchIacOwnershipConflict(msg string) *IacOwnershipConflict {
	var conflict *IacOwnershipConflict
	if iacBlockConflictHeader.MatchString(msg) {
		conflict = &IacOwnershipConflict{Blocks: map[string]string{}, Events: map[string]string{}}
		for _, m := range iacBlockConflictLine.FindAllStringSubmatch(msg, -1) {
			for _, name := range strings.Split(m[2], ",") {
				if name = strings.TrimSpace(name); name != "" {
					conflict.Blocks[name] = m[1]
				}
			}
		}
	}
	if matches := iacEventConflictLine.FindAllStringSubmatch(msg, -1); len(matches) > 0 {
		if conflict == nil {
			conflict = &IacOwnershipConflict{Blocks: map[string]string{}, Events: map[string]string{}}
		}
		for _, m := range matches {
			conflict.Events[m[1]] = m[2]
		}
	}
	return conflict
}

// BlockNames returns the conflicting block names, sorted.
func (c IacOwnershipConflict) BlockNames() []string {
	return sortedKeys(c.Blocks)
}

// EventNames returns the conflicting event names, sorted.
func (c IacOwnershipConflict) EventNames() []string {
	return sortedKeys(c.Events)
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
