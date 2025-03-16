package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/owenrumney/go-sarif/v3/pkg/report"
	"github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"
)

// TfsecResults is a simple structure for the output of tfsec
type TfsecResults struct {
	Results []struct {
		RuleID          string `json:"rule_id"`
		RuleDescription string `json:"rule_description"`
		RuleProvider    string `json:"rule_provider"`
		Link            string `json:"link"`
		Location        struct {
			Filename  string `json:"filename"`
			StartLine int    `json:"start_line"`
			EndLine   int    `json:"end_line"`
		} `json:"location"`
		Description string `json:"description"`
		Impact      string `json:"impact"`
		Resolution  string `json:"resolution"`
		Severity    string `json:"severity"`
		Passed      bool   `json:"passed"`
	} `json:"results"`
}

func main() {

	// Get the results from the results file
	tfsecResults, err := loadTfsecResults()
	if err != nil {
		panic(err)
	}

	// create a new report object
	report := report.NewV22Report()

	// create a run for tfsec
	run := sarif.NewRunWithInformationURI("tfsec", "https://tfsec.dev")

	// for each result, add the
	for _, r := range tfsecResults.Results {

		// create a property bag for the non standard stuff
		pb := sarif.NewPropertyBag()
		pb.Add("impact", r.Impact)
		pb.Add("resolution", r.Resolution)

		// create a new rule for each rule id
		run.AddRule(r.RuleID).
			WithDescription(r.Description).
			WithHelpURI(r.Link).
			WithProperties(pb).
			WithMarkdownHelp("# markdown")

		// add the location as a unique artifact
		run.AddDistinctArtifact(r.Location.Filename)

		// add each of the results with the details of where the issue occurred
		run.CreateResultForRule(r.RuleID).
			WithLevel(strings.ToLower(r.Severity)).
			WithMessage(sarif.NewTextMessage(r.Description)).
			AddLocation(
				sarif.NewLocationWithPhysicalLocation(
					sarif.NewPhysicalLocation().
						WithArtifactLocation(
							sarif.NewSimpleArtifactLocation(r.Location.Filename),
						).WithRegion(
						sarif.NewSimpleRegion(r.Location.StartLine, r.Location.EndLine),
					),
				),
			)
	}

	// add the run to the report
	report.AddRun(run)

	isValid, err := report.Validate()
	if err != nil {
		panic(err)
	}

	if !isValid {
		panic("report is not valid")
	}

	println("Report is valid")

	// print the report to stdout
	// _ = report.PrettyWrite(os.Stdout)

	// save the report
	if err := report.WriteFile("example-report.sarif"); err != nil {
		panic(err)
	}

}

// load the example results file
func loadTfsecResults() (TfsecResults, error) {

	jsonResult, err := os.ReadFile("./example/results.json")
	if err != nil {
		panic(err)
	}

	var results TfsecResults

	err = json.Unmarshal(jsonResult, &results)
	return results, err
}
