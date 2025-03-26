# go-sarif
[![Go Report Card](https://goreportcard.com/badge/github.com/owenrumney/go-sarif/v3)](https://goreportcard.com/report/github.com/owenrumney/go-sarif/v3)
[![Github Release](https://img.shields.io/github/release/owenrumney/go-sarif.svg)](https://github.com/owenrumney/go-sarif/releases)

## Overview

SARIF is the Static Analysis Results Interchange Format, this project seeks to provide a simple interface to generate reports in the SARIF format.

## Usage

Add an import to `go get github.com/owenrumney/go-sarif/v3`

### Parsing a SARIF report

There are a number of ways to load in the content of a SARIF report.

For a `v2.1.0` report use `import "github.com/owenrumney/go-sarif/v3/pkg/report/v210/sarif"`

For a `v2.2` report, use `import "github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"`


#### Open

`sarif.Open` takes a file path and loads the SARIF from that location. Returns a report and any corresponding error

#### FromBytes

`sarif.FromBytes` takes a slice of byte and returns a report and any corresponding error.

#### FromString

`sarif.FromString` takes a string of the SARIF content and returns a report and any corresponding error.

### Validating a Report

Once you have the report object, you can call `valid, err := report.Validate()` to get a list of any issues. This will evaluate the report against the schema.

### Creating a new report

Creating a new SARIF report can be done directly with the `sarif` package or using the `report` package at `github.com/owenrumney/go-sarif/v3/pkg/report`

for a detailed example check the example folder [example/main.go](example/main.go)

```go

import (
  "github.com/owenrumney/go-sarif/v3/pkg/report"
  "github.com/owenrumney/go-sarif/v3/pkg/report/v22/sarif"
)

...

// create the basic report shell
rep := report.NewV22Report()

// create a run 
run := sarif.NewRunWithInformationURI("my tool", "https://mytool.com")

// create a failed Rule
run.AddRule("rule#1").
  WithDescription("This rule is a really important one").
  WithHelpURI("https://mytool.com/rules/rule1").
  WithMarkdownHelp("# Try not to make this mistake")

// add the location an artifact
run.AddDistinctArtifact("file:///Users/me/code/myCode/terraform/main.tf")

// crete a result for the rule
run.CreateResultForRule("rule#1").
  WithLevel("high").
  WithMessage(sarif.NewTextMessage("This rule was breached in the file")).
  AddLocation(
    sarif.NewLocationWithPhysicalLocation(
      sarif.NewPhysicalLocation().
        WithArtifactLocation(
          sarif.NewSimpleArtifactLocation("file:///Users/me/code/myCode/terraform/main.tf")
        ).WithRegion(
          // set the line numbers of the issue
          sarif.NewSimpleRegion(1, 4)
        ),
    ),
  )
  
// add the run to the report
rep.AddRun(run)

// validate the report
if err := rep.Validate(); err != nil {
  println(err)
}





```

### Example report

This example is taken directly from the [Microsoft SARIF pages](https://github.com/microsoft/sarif-tutorials/blob/master/docs/1-Introduction.md)

```json
{
  "version": "2.1.0",
  "$schema": "(https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json)",
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


## More information about SARIF
For more information about SARIF, you can visit the [Oasis Open](https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=sarif) site.



