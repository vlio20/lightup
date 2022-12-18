package model

type FlagStateParams struct {
	FlagName string `form:"flagName" binding:"required"`
}
