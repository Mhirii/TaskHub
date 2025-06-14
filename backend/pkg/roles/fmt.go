package roles

import "strings"

func RolesToString(roles []string) string {
	return strings.Join(roles, "--")
}

func StringToRoles(str string) []string {
	return strings.Split(str, "--")
}
