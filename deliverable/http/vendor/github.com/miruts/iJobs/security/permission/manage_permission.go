package permission

import (
	"strings"
)

type permission struct {
	roles   []string
	methods []string
}

type authority map[string]permission

var authorities = authority{
	"/jobseeker": permission{
		roles:   []string{"JOBSEEKER"},
		methods: []string{"GET", "POST"},
	},
	"/company": permission{
		roles:   []string{"COMPANY"},
		methods: []string{"GET", "POST"},
	},
}

// HasPermission checks if a given role has permission to access a given route for a given method
func HasPermission(path string, role string, method string) bool {
	if strings.HasPrefix(path, "/jobseeker") {
		path = "/jobseeker"
	}
	if strings.HasPrefix(path, "/company") {
		path = "/company"
	}
	perm := authorities[path]
	checkedRole := checkRole(role, perm.roles)
	checkedMethod := checkMethod(method, perm.methods)
	if !checkedRole || !checkedMethod {
		return false
	}
	return true
}

func checkRole(role string, roles []string) bool {
	for _, r := range roles {
		if strings.ToUpper(r) == strings.ToUpper(role) {
			return true
		}
	}
	return false
}

func checkMethod(method string, methods []string) bool {
	for _, m := range methods {
		if strings.ToUpper(m) == strings.ToUpper(method) {
			return true
		}
	}
	return false
}
