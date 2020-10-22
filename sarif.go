package go_sarif

import (
	"encoding/json"
	"github.com/owenrumney/go-sarif/models"
	"io"
)

type SarifReport struct {
	Runs []*models.Run `json:"runs"`
}

func New() *SarifReport {
	return &SarifReport{
		Runs: []*models.Run{},
	}
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
