# go-sarif

[![Build Status](https://travis-ci.com/owenrumney/go-sarif.svg?branch=main)](https://travis-ci.com/owenrumney/go-sarif) 
[![Go Report Card](https://goreportcard.com/badge/github.com/owenrumney/go-sarif)](https://goreportcard.com/report/github.com/owenrumney/go-sarif)

### Important - the implementation has enough at this point to satisfy the requirements of [tfsec](https://tfsec.dev) - the example below is achievable

## What?

SARIF is the Static Analysis Results Interchange Format, this project seeks to provide a simple interface to generate reports in the SARIF format.

### Example report

This example is taken directly from the [Microsoft sarif pages](https://github.com/microsoft/sarif-tutorials/blob/master/docs/1-Introduction.md)

```json
{
  "version": "2.1.0",
  "$schema": "http://json.schemastore.org/sarif-2.1.0-rtm.4",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "ESLint",
          "informationUri": "https://eslint.org",
          "rules": [
            {
              "id": "no-unused-vars",
              "shortDescription": {
                "text": "disallow unused variables"
              },
              "helpUri": "https://eslint.org/docs/rules/no-unused-vars",
              "properties": {
                "category": "Variables"
              }
            }
          ]
        }
      },
      "artifacts": [
        {
          "location": {
            "uri": "file:///C:/dev/sarif/sarif-tutorials/samples/Introduction/simple-example.js"
          }
        }
      ],
      "results": [
        {
          "level": "error",
          "message": {
            "text": "'x' is assigned a value but never used."
          },
          "locations": [
            {
              "physicalLocation": {
                "artifactLocation": {
                  "uri": "file:///C:/dev/sarif/sarif-tutorials/samples/Introduction/simple-example.js",
                  "index": 0
                },
                "region": {
                  "startLine": 1,
                  "startColumn": 5
                }
              }
            }
          ],
          "ruleId": "no-unused-vars",
          "ruleIndex": 0
        }
      ]
    }
  ]
}
```

## What prompted this?
One of the projects I like to contribute to is [tfsec](https://tfsec.dev) - this is a static analysis tool for Terraform which produces output in many formats. Generating SARIF reports is missing functionality and felt like it warranted being moved out to a project of its own.

## More information about SARIF
For more information about SARIF, you can visit the [Oasis Open](https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=sarif) site.

## Usage

Add an import to `github.com/owenrumney/go-sarif/sarif`

Creating a new Sarif report is done by passing the version, the only supported at the moment is `2.1.0`

The example below is taken from the `tfsec` usage of `go-sarif`. For context, at the end of the process a slice of `Result` objects is returned with the relevant information about the check failures.

```go
// create the report object
report, err := sarif.New(sarif.Version210)
if err != nil {
	return err
}

// add a run to the report
run := report.AddRun("tfsec", "https://tfsec.dev")

// for each result add the rule, location and result to the report
for _, result := range results {
	rule := run.AddRule(string(result.RuleID)).
		WithDescription(result.Description).
		WithHelpUri(fmt.Sprintf("https://tfsec.dev/%s/%s", strings.ToLower(string(result.RuleProvider)), result.RuleID))

	ruleResult := run.AddResult(rule.Id).
		WithMessage(string(result.RuleDescription)).
		WithLevel(string(result.Severity)).
		WithLocationDetails(result.Range.Filename, result.Range.StartLine, 1)

	run.AddResultDetails(rule, ruleResult, result.Range.Filename)
}

// print the report to anything that implements `io.Writer`
return report.Write(w)
```

