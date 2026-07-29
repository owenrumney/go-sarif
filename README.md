# go-sarif

>[!IMPORTANT]
>go-sarif is true to the specifications with a single deviation to accomodate GitHub processing of SARIF reports. the `run.Results` attribute is not a required field per the spec but is required by GitHub to successfully process the report. To that end, the `omitEmpty` is removed on the attribute.


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

### Merging reports

`github.com/owenrumney/go-sarif/v3/pkg/report/merge` combines several reports of the same version into one, which is handy for collecting the output of several scanners, or of one scanner sharded across a build matrix.

```go
import "github.com/owenrumney/go-sarif/v3/pkg/report/merge"

merged, err := merge.V22(first, second, third)
// or merge.V210(first, second, third)
```

A SARIF log is a list of runs, so merging appends the runs of each report in order rather than folding them into a single run. That is what keeps it safe: SARIF refers to run level arrays by index from a dozen places (`result.ruleIndex`, `artifactLocation.index`, `threadFlowLocation.index` and so on), and leaving each run intact leaves every one of those references valid.

The merged report is given a fresh `guid`, since it is a new document. Report level properties are unioned with the earliest report winning any duplicated key, and tags are unioned and deduplicated. Note that the merged report shares run pointers with its inputs — it is not a deep copy.

### Converting between versions

`github.com/owenrumney/go-sarif/v3/pkg/report/convert` translates reports between `2.1.0` and `2.2`.

```go
import "github.com/owenrumney/go-sarif/v3/pkg/report/convert"

// 2.1.0 -> 2.2, lossless
upgraded, err := convert.ToV22(v210Report)

// 2.2 -> 2.1.0, drops anything 2.1.0 cannot represent
downgraded, err := convert.ToV210(v22Report)
```

Upgrading is lossless. Downgrading is not — `2.2` added a report level `guid` and `relatedLocations` on `notification`, neither of which `2.1.0` has a home for. By default those fields are dropped; pass `WithLossHandler` to see what went, or `WithStrictConversion` to fail instead.

```go
// report what was dropped
downgraded, err := convert.ToV210(v22Report, convert.WithLossHandler(func(l convert.Loss) {
    slog.Warn("dropped in conversion", "path", l.Path)
}))

// or refuse to drop anything
downgraded, err := convert.ToV210(v22Report, convert.WithStrictConversion())
// err is a *convert.LossyConversionError listing every field that could not be carried over
```

Note that `sarif.NewReport` for `2.2` always generates a report level `guid`, so a strict downgrade of such a report reports that `guid` as a loss. Clear `Report.Guid` first if that is not wanted.

### Example report

This example is taken directly from the [Microsoft SARIF pages](https://github.com/microsoft/sarif-tutorials/blob/master/docs/1-Introduction.md)

```json
{
    "version": "2.1.0",
    "$schema": "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
    "runs": [
      {
        "tool": {
          "driver": {
            "name": "ESLint",
            "informationUri": "https://eslint.org",
            "language": "en-US",
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
        "language": "en-US",
        "newlineSequences": ["\r\n", "\n"],
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



