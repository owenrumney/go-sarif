# go-sarif
[![Go Report Card](https://goreportcard.com/badge/github.com/owenrum/go-sarif)](https://goreportcard.com/report/github.com/owenrum/go-sarif) 
[![Github Release](https://img.shields.io/github/release/owenrum/go-sarif.svg)](https://github.com/owenrum/go-sarif/releases)

## What?

SARIF is the Static Analysis Results Interchange Format, this project seeks to provide a simple interface to generate reports in the SARIF format.

### Example report

This example is taken directly from the [Microsoft sarif pages](https://github.com/microsoft/sarif-tutorials/blob/master/docs/1-Introduction.md)

```json
{
  "version": "2.1.0",
  "$schema": "https://json.schemastore.org/sarif-2.1.0-rtm.5.json",
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

Add an import to `go get github.com/owenrum/go-sarif/sarif`
### Parsing a Sarif report

There are a number of ways to load in the content of a sarif report.

#### Open

`sarif.Open` takes a file path and loads the sarif from that location. Returns a report and any corresponding error

#### FromBytes

`sarif.FromBytes` takes a slice of byte and returns a report and any corresponding error.

#### FromString

`sarif.FromString` takes a string of the sarif content and returns a report and any corresponding error.

### Creating a new report

Creating a new Sarif report is done by passing the version, the only supported at the moment is `2.1.0`

for a detailed example check the example folder [example/main.go](example/main.go)

