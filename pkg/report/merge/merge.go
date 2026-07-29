// Package merge combines several SARIF reports of the same version into one.
//
// A SARIF log is a list of runs, so merging appends the runs of every input
// report in order. Each run keeps its own artifacts, rules and other run level
// arrays, which is what makes this safe: SARIF refers to those arrays by index
// from a dozen different places (result.ruleIndex, artifactLocation.index,
// threadFlowLocation.index and so on), and leaving the runs intact leaves every
// one of those references pointing where it always did.
//
// Coalescing several runs into a single run is deliberately not offered. It
// would require remapping every index reference in the report, and a single
// missed reference silently repoints a result at the wrong rule.
//
// The merged report shares run pointers with its inputs; it is not a deep copy.
// Mutating a run of the merged report therefore also mutates the report it came
// from.
package merge

import (
	"errors"
	"fmt"

	v210 "github.com/owenrumney/go-sarif/v3/pkg/report/v210/sarif"
	v22 "github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"
)

// V22 merges 2.2 reports into a single 2.2 report, in the order given.
//
// The merged report is a new document, so it is given a freshly generated guid
// rather than inheriting one. Report level properties are unioned, with the
// earliest report winning any key that appears more than once; tags are unioned
// and deduplicated.
func V22(reports ...*v22.Report) (*v22.Report, error) {
	if len(reports) == 0 {
		return nil, errors.New("cannot merge zero reports")
	}

	merged := v22.NewReport()
	for i, report := range reports {
		if report == nil {
			return nil, fmt.Errorf("cannot merge a nil report at position %d", i)
		}

		merged.Runs = append(merged.Runs, report.Runs...)
		merged.InlineExternalProperties = append(merged.InlineExternalProperties, report.InlineExternalProperties...)
		mergePropertiesV22(merged.Properties, report.Properties)
	}

	return merged, nil
}

// V210 merges 2.1.0 reports into a single 2.1.0 report, in the order given.
//
// Report level properties are unioned, with the earliest report winning any key
// that appears more than once; tags are unioned and deduplicated.
func V210(reports ...*v210.Report) (*v210.Report, error) {
	if len(reports) == 0 {
		return nil, errors.New("cannot merge zero reports")
	}

	merged := v210.NewReport()
	for i, report := range reports {
		if report == nil {
			return nil, fmt.Errorf("cannot merge a nil report at position %d", i)
		}

		merged.Runs = append(merged.Runs, report.Runs...)
		merged.InlineExternalProperties = append(merged.InlineExternalProperties, report.InlineExternalProperties...)
		mergePropertiesV210(&merged.Properties, &report.Properties)
	}

	return merged, nil
}

func mergePropertiesV22(into, from *v22.PropertyBag) {
	if into == nil || from == nil {
		return
	}
	for key, value := range from.Properties {
		if _, exists := into.Properties[key]; !exists {
			into.Properties[key] = value
		}
	}
	into.Tags = appendNewTags(into.Tags, from.Tags)
}

func mergePropertiesV210(into, from *v210.PropertyBag) {
	if into == nil || from == nil {
		return
	}
	for key, value := range from.Properties {
		if _, exists := into.Properties[key]; !exists {
			into.Properties[key] = value
		}
	}
	into.Tags = appendNewTags(into.Tags, from.Tags)
}

// appendNewTags adds the tags not already present, preserving the order tags
// were first seen in.
func appendNewTags(into, from []string) []string {
	seen := make(map[string]struct{}, len(into))
	for _, tag := range into {
		seen[tag] = struct{}{}
	}
	for _, tag := range from {
		if _, ok := seen[tag]; ok {
			continue
		}
		seen[tag] = struct{}{}
		into = append(into, tag)
	}
	return into
}
