package sarif

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/owenrumney/go-sarif/models"
	"io"
)

var versions = map[string]string{
	"2.1.0": "http://json.schemastore.org/sarif-2.1.0-rtm.4",
}

type SarifReport struct {
	Version string        `json:"version"`
	Schema  string        `json:"$schema"`
	Runs    []*models.Run `json:"runs"`
}

func New(version string) (*SarifReport, error) {
	schema, err := getVersionSchema(version)
	if err != nil {
		return nil, err
	}
	return &SarifReport{
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

func (sarif *SarifReport) AddRun(run *models.Run) {
	sarif.Runs = append(sarif.Runs, run)
}

func (sarif *SarifReport) Write(w io.Writer) error {
	marshal, err := json.Marshal(sarif)
	if err != nil {
		return err
	}
	_, err = w.Write(marshal)
	return err
}
