package versions

import (
	"modding_server/database"
	"time"
)

type Version struct {
	Id          int
	ModId       int
	UpdatedAt   time.Time
	ContentId   int
	IconId      int
	ReadmeId    int
	ChangelogId *int
}

const table string = "versions"

func getArgs(ver *Version) []any {
	return []any{&ver.Id, &ver.ModId, &ver.UpdatedAt, &ver.ContentId, &ver.IconId, &ver.ReadmeId, &ver.ChangelogId}
}

func Create(modId int, updatedAt time.Time, contentId int, iconId int, readmeId int, changelogId *int) (Version, error) {
	var ver Version = Version{ModId: modId, UpdatedAt: updatedAt, ContentId: contentId, IconId: iconId, ReadmeId: readmeId, ChangelogId: changelogId}
	args := getArgs(&ver)
	err := database.Create(table, "mod_id, updated_at, content_id, icon_id, readme_id, changelog_id", "$1, $2, $3, $4, $5, $6", args)
	return ver, err
}

func Read(id int) (Version, error) {
	var ver Version = Version{Id: id}
	args := getArgs(&ver)
	err := database.Read(table, args)
	return ver, err
}

func Update(version Version) (Version, error) {
	var ver Version = version
	args := getArgs(&ver)
	err := database.Update(table, "mod_id = $2, updated_at = $3, content_id = $4, icon_id = $5, readme_id = $6, changelog_id = $7", args)
	return ver, err
}

func Delete(id int) (Version, error) {
	var ver Version = Version{Id: id}
	args := getArgs(&ver)
	err := database.Delete(table, args)
	return ver, err
}
