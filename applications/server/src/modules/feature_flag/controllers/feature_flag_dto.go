package feature_flag_ctrl

type CreateFeatureFlagDto struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"descriptio" binding:"required"`
}
