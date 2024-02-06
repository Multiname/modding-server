package dependencies

import (
	"modding_server/database"
)

type Dependency struct {
	Id             int
	DependentModId int
	RequiredModId  int
}

const table string = "dependencies"

func getArgs(dep *Dependency) []any {
	return []any{&dep.Id, &dep.DependentModId, &dep.RequiredModId}
}

func Create(dependentModId int, requiredModId int) (Dependency, error) {
	var dep Dependency = Dependency{DependentModId: dependentModId, RequiredModId: requiredModId}
	args := getArgs(&dep)
	err := database.Create(table, "dependent_mod_id, required_mod_id", "$1, $2", args)
	return dep, err
}

func Read(id int) (Dependency, error) {
	var dep Dependency = Dependency{Id: id}
	args := getArgs(&dep)
	err := database.Read(table, args)
	return dep, err
}

func Update(dependency Dependency) (Dependency, error) {
	var dep Dependency = dependency
	args := getArgs(&dep)
	err := database.Update(table, "dependent_mod_id = $2, required_mod_id = $3", args)
	return dep, err
}

func Delete(id int) (Dependency, error) {
	var dep Dependency = Dependency{Id: id}
	args := getArgs(&dep)
	err := database.Delete(table, args)
	return dep, err
}
