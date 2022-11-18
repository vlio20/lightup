package model

type FeatureFlagConfig struct {
	Identifier string  `bson:"indentifier" json:"identifier" binding:"required"`
	Percentage float32 `bson:"percentage" json:"percentage" binding:"required"`
}

type CreateFeatureFlagDto struct {
	Name        string
	Description string
	Archived    bool              `bson:"archived"`
	Config      FeatureFlagConfig `bson:"config"`
}
