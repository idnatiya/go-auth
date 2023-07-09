package requests

type PermissionRequest struct {
	PermissionName string `form:"permission_name" json:"permission_name" binding:"required,min=3,max=255"`
}
