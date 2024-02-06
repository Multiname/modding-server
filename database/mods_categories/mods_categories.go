package mods_categories

import (
	"modding_server/database"
)

type ModCategory struct {
	Id         int
	ModId      int
	CategoryId int
}

const table string = "mods_categories"

func getArgs(mc *ModCategory) []any {
	return []any{&mc.Id, &mc.ModId, &mc.CategoryId}
}

func Create(modId int, categoryId int) (ModCategory, error) {
	var mc ModCategory = ModCategory{ModId: modId, CategoryId: categoryId}
	args := getArgs(&mc)
	err := database.Create(table, "mod_id, category_id", "$1, $2", args)
	return mc, err
}

func Read(id int) (ModCategory, error) {
	var mc ModCategory = ModCategory{Id: id}
	args := getArgs(&mc)
	err := database.Read(table, args)
	return mc, err
}

func Update(modCategory ModCategory) (ModCategory, error) {
	var mc ModCategory = modCategory
	args := getArgs(&mc)
	err := database.Update(table, "mod_id = $2, category_id = $3", args)
	return mc, err
}

func Delete(id int) (ModCategory, error) {
	var mc ModCategory = ModCategory{Id: id}
	args := getArgs(&mc)
	err := database.Delete(table, args)
	return mc, err
}
