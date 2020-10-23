package sarif

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/owenrumney/go-sarif/models"
)

var versions = map[string]string{
	"2.1.0": "http://json.schemastore.org/sarif-2.1.0-rtm.4",
}

type Report struct {
	Version string        `json:"version"`
	Schema  string        `json:"$schema"`
	Runs    []*models.Run `json:"runs"`
}

func New(version string) (*Report, error) {
	schema, err := getVersionSchema(version)
	if err != nil {
		return nil, err
	}
	return &Report{
		Version: version,
		Schema:  schema,
		Runs:    []*models.Run{},
	}, nil
}

func getVersionSchema(version string) (string, error) {
	for ver, schema := range versions {
		if ver == version {
			return schema, nil
		}
	}
	return "", errors.New(fmt.Sprintf("version [%s] is not supported", version))
}

func (sarif *Report) AddRun(run *models.Run) {
	sarif.Runs = append(sarif.Runs, run)
}

func (sarif *Report) Write(w io.Writer) error {
	marshal, err := json.Marshal(sarif)
	if err != nil {
		return err
	}
	_, err = w.Write(marshal)
	return err
}
