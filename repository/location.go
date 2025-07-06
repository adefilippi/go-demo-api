package repository

import (
	"fmt"

	"github.com/adefilippi/go-demo-api/entity"
	"github.com/syneido/go-api-core/database"
)

func GetAllLocations(parameters map[string]interface{}) ([]entity.Location, error) {
	codeCE := parameters["path"].(map[string]interface{})["codeCE"]
	_db := database.GetDB("crm")

	var locations []entity.Location
	var location entity.Location
	rows, err := _db.Raw("EXEC ps_API_L_RESEAU @code_sitedevente = ?", codeCE).Rows()
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return locations, err
	}

	for rows.Next() {
		_db.ScanRows(rows, &location)
		locations = append(locations, location)
	}
	return locations, nil
}
