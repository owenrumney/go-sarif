package sarif

import "time"

type VersionControlDetails struct {
	PropertyBag
	AsOfTimeUTC   *time.Time        `json:"asOfTimeUtc,omitempty"`
	Branch        *string           `json:"branch,omitempty"`
	MappedTo      *ArtifactLocation `json:"mappedTo,omitempty"`
	RepositoryURI *string           `json:"repositoryUri"`
	RevisionID    *string           `json:"revisionId,omitempty"`
	RevisionTag   *string           `json:"revisionTag,omitempty"`
}

func NewVersionControlDetails() *VersionControlDetails {
	return &VersionControlDetails{}
}

func (versionControlDetails *VersionControlDetails) WithAsOfTimeUTC(asOfTimeUTC *time.Time) *VersionControlDetails {
	versionControlDetails.AsOfTimeUTC = asOfTimeUTC
	return versionControlDetails
}

func (versionControlDetails *VersionControlDetails) WithBranch(branch string) *VersionControlDetails {
	versionControlDetails.Branch = &branch
	return versionControlDetails
}

func (versionControlDetails *VersionControlDetails) WithMappedTo(mappedTo *ArtifactLocation) *VersionControlDetails {
	versionControlDetails.MappedTo = mappedTo
	return versionControlDetails
}

func (versionControlDetails *VersionControlDetails) WithRepositoryURI(repositoryURI string) *VersionControlDetails {
	versionControlDetails.RepositoryURI = &repositoryURI
	return versionControlDetails
}

func (versionControlDetails *VersionControlDetails) WithRevisionID(revisionID string) *VersionControlDetails {
	versionControlDetails.RevisionID = &revisionID
	return versionControlDetails
}

func (versionControlDetails *VersionControlDetails) WithRevisionTag(revisionTag string) *VersionControlDetails {
	versionControlDetails.RevisionTag = &revisionTag
	return versionControlDetails
}
