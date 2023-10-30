package random

import (
	"company-service/pkg/domain"
	"company-service/pkg/file"
	"company-service/pkg/store"
	"company-service/pkg/utils"
	"sync"
)

const (
	namesSheetFileName = "./inputs/names.xls"
	namesSheetName     = "names"
)

type RandomGenerator interface {
	GenerateEmployee(minSalary, maxSalary float64) *domain.Employee
	GetRandomTeamNameByDepartmentName(department string) string
	GetRandomRoleByTeamName(team string) string
}

type randomGenerator struct {
	mu               sync.RWMutex
	names            []string
	departmentToRole map[string][]string
	teamToRole       map[string][]string
}

func NewRandomGenerator() (RandomGenerator, error) {

	// get the names for random generation
	names, err := file.GetAllNamesFromSheet(namesSheetFileName, namesSheetName)
	if err != nil {
		return nil, err
	}

	return &randomGenerator{
		names:            names,
		mu:               sync.RWMutex{},
		departmentToRole: store.GetDepartmentToTeamsMap(),
		teamToRole:       store.GetTeamToRolesMap(),
	}, nil
}

// To get a random name from random generator with thread safe
func (r *randomGenerator) getRandomName() string {

	// lock the mutex for read.
	r.mu.RLock()
	defer r.mu.RUnlock()

	// return a random name from the name
	return r.names[utils.GetRandomIndex(len(r.names))]
}
