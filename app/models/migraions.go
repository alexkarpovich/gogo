package models

type PermissionsMap map[string]map[string][]string
type PermissionItem map[string][]string

func Migrations() ([]string, PermissionsMap, [][]string) {
	Roles := []string{"Admin", "Employee", "Guest"}

	Users := [][]string{
		{"admin@admin.com","Admin","Admin","admin","Admin"}}

	Permissions := PermissionsMap{
		"App": PermissionItem{
			"Index": []string{"Admin", "Employee", "Guest"}},
		"AccountUser": PermissionItem{
			"Signup":   []string{"Guest"},
			"Login":    []string{"Guest"},
			"Logout":   []string{"Admin", "Employee"},
			"Profile":  []string{"Admin", "Employee"},
			"Retrieve": []string{"Admin", "Employee"}},
		"AdminUser": PermissionItem{
			"List":   []string{"Admin", "Employee"},
			"Create": []string{"Admin"},
			"Update": []string{"Admin"},
			"Delete": []string{"Admin"}},
		"AdminArticle": PermissionItem{
			"List":   []string{"Admin", "Employee"},
			"Create": []string{"Admin", "Employee"},
			"Update": []string{"Admin"},
			"Delete": []string{"Admin"}},
		"AdminRole": PermissionItem{
			"List":   []string{"Admin", "Employee"},
			"Create": []string{"Admin"},
			"Update": []string{"Admin"},
			"Delete": []string{"Admin"}},
		"AdminPermission": PermissionItem{
			"List":   []string{"Admin"},
			"Create": []string{"Admin"},
			"Update": []string{"Admin"},
			"Delete": []string{"Admin"}}}
	return Roles, Permissions, Users
}
