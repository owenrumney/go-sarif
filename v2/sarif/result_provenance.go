package sarif

import "time"

type ResultProvenance struct {
	PropertyBag
	ConversionSources     []*PhysicalLocation `json:"conversionSources,omitempty"`
	FirstDetectionRunGUID *string             `json:"firstDetectionRunGuid,omitempty"`
	FirstDetectionTimeUTC *time.Time          `json:"firstDetectionTimeUtc,omitempty"`
	InvocationIndex       *int                `json:"invocationIndex,omitempty"`
	LastDetectionRunGUID  *string             `json:"lastDetectionRunGuid,omitempty"`
	LastDetectionTimeUTC  *time.Time          `json:"lastDetectionTimeUtc,omitempty"`
}

func NewResultProvenance() *ResultProvenance {
	return &ResultProvenance{}
}

func (resultProvenance *ResultProvenance) WithConversionSources(conversionSources []*PhysicalLocation) *ResultProvenance {
	resultProvenance.ConversionSources = conversionSources
	return resultProvenance
}

func (resultProvenance *ResultProvenance) AddConversionSource(conversionSource *PhysicalLocation) {
	resultProvenance.ConversionSources = append(resultProvenance.ConversionSources, conversionSource)
}

func (resultProvenance *ResultProvenance) WithFirstDetectionRunGUID(firstDetectionRunGUID string) *ResultProvenance {
	resultProvenance.FirstDetectionRunGUID = &firstDetectionRunGUID
	return resultProvenance
}

func (resultProvenance *ResultProvenance) WithFirstDetectionTimeUTC(firstDetectionTimeUTC *time.Time) *ResultProvenance {
	resultProvenance.FirstDetectionTimeUTC = firstDetectionTimeUTC
	return resultProvenance
}

func (resultProvenance *ResultProvenance) WithInvocationIndex(invocationIndex int) *ResultProvenance {
	resultProvenance.InvocationIndex = &invocationIndex
	return resultProvenance
}

func (resultProvenance *ResultProvenance) WithLastDetectionRunGUID(lastDetectionRunGUID string) *ResultProvenance {
	resultProvenance.LastDetectionRunGUID = &lastDetectionRunGUID
	return resultProvenance
}

func (resultProvenance *ResultProvenance) WithLastDetectionTimeUTC(lastDetectionTimeUTC *time.Time) *ResultProvenance {
	resultProvenance.LastDetectionTimeUTC = lastDetectionTimeUTC
	return resultProvenance
}
