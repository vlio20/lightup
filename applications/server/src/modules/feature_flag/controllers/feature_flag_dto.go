package feature_flag_ctrl

type CreateFeatureFlagDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
