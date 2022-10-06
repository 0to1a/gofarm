package structure

type ModularStruct struct {
	Name       string           `json:"name"`
	Version    int              `json:"version"`
	MinVersion int              `json:"min_version"`
	MaxVersion int              `json:"max_version"`
	Depending  []*ModularStruct `json:"depending"`
}
