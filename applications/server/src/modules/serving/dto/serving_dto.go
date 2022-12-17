package dto

type FeatureFlagStateDto struct {
	ID       string `json:"id"`
	FlagName string `json:"featureFlagName"`
	Enabled  bool   `json:"enabled"`
}
