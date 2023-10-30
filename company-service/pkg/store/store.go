package store

// To get all department names and employee percentage in department
func GetAllDepartmentNamesAndDetails() [len(departmentsDetails)]departmentDetail {
	mu.RLock()
	defer mu.RUnlock()

	return departmentsDetails
}

// To get depart names to it's roles in a map
func GetDepartmentToTeamsMap() map[string][]string {
	mu.RLock()
	defer mu.RUnlock()

	return departmentTeamNames
}

// To get team to it's roles in map
func GetTeamToRolesMap() map[string][]string {
	mu.RLock()
	defer mu.RUnlock()
	return teamRoles
}
