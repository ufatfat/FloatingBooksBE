package db

import "FloatingBooks/model"

func GetPlaces () (places []model.PlaceInfo, err error) {
	err = Mysql.Table("locations").Select("id", "location_name").Order("localtion_abbr").Scan(&places).Error
	return
}