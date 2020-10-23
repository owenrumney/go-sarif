# WIP - go-sarif

** This project is still early stages of development and not ready for use **

## What?

SARIF is the Static Analysis Results Interchange Format, this project seeks to provide a simple interface to generate reports in the SARIF format.

## What prompted this?
One of the projects I like to contribute to is [tfsec](https://tfsec.dev) - this is a static analysis tool for Terraform which produces output in many formats. Generating SARIF reports is missing functionality and felt like it warranted being moved out to a project of its own.

## More information about SARIF
For more information about SARIF, you can visit the [Oasis Open](https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=sarif) site.

## Usage

Add an import to `github.com/owenrumney/go-sarif/sarif`

Creating a new Sarif report is done by passing the version, the only supported at the moment is `2.1.0`

