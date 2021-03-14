package sarif

type Suppression struct {
	Kind          string    `json:"kind"`
	Status        *string   `json:"status"`
	Location      *Location `json:"location"`
	Guid          *string   `json:"guid"`
	Justification *string   `json:"justification"`
}
