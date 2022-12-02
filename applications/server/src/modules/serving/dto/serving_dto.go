package dto

type GetFeatureFlagStateParams struct {
	FeatureFlagName string `json:"featureFlagName" binding:"required"`
	ServiceName     string `json:"serviceName" binding:"required"`
	Identifier      string `json:"identifier" binding:"required"`
}

type FeatureFlagStateDto struct {
	ID              string `json:"id"`
	FeatureFlagName string `json:"featureFlagName"`
	Enabled         bool   `json:"enabled"`
}
