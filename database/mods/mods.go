package mods

import (
	"modding_server/database"
)

type Mod struct {
	Id        int
	Name      string
	OwnerId   int
	Downloads int
	Rating    int
	Active    bool
}

const table string = "mods"

func getArgs(md *Mod) []any {
	return []any{&md.Id, &md.Name, &md.OwnerId, &md.Downloads, &md.Rating, &md.Active}
}

func Create(name string, ownerId int, downloads int, rating int, active bool) (Mod, error) {
	var md Mod = Mod{Name: name, OwnerId: ownerId, Downloads: downloads, Rating: rating, Active: active}
	args := getArgs(&md)
	err := database.Create(table, "name, owner_id, downloads, rating, active", "$1, $2, $3, $4, $5", args)
	return md, err
}

func Read(id int) (Mod, error) {
	var md Mod = Mod{Id: id}
	args := getArgs(&md)
	err := database.Read(table, args)
	return md, err
}

func Update(mod Mod) (Mod, error) {
	var md Mod = mod
	args := getArgs(&md)
	err := database.Update(table, "name = $2, owner_id = $3, downloads = $4, rating = $5, active = $6", args)
	return md, err
}

func Delete(id int) (Mod, error) {
	var md Mod = Mod{Id: id}
	args := getArgs(&md)
	err := database.Delete(table, args)
	return md, err
}

func GetCategories(id int) ([]int, error) {
	return database.GetList("mods_categories", "category_id", "mod_id", id)
}

func GetDependencies(id int) ([]int, error) {
	return database.GetList("dependencies", "required_mod_id", "dependent_mod_id", id)
}

func GetVersions(id int) ([]int, error) {
	return database.GetList("versions", "id", "mod_id", id)
}
