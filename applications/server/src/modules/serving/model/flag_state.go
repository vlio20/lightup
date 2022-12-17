package model

type FlagStateParams struct {
	FlagName   string   `form:"flagName" binding:"required"`
	Tags       []string `form:"tags" binding:"required"`
	Identifier string   `form:"identifier" binding:"required"`
}
