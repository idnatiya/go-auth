package requests

type CreatePermissionRequest struct {
	PermissionName string `form:"permission_name" json:"permission_name" binding:"required,min=3,max=255"`
}
