package db

import "FloatingBooks/model"

func GetPlaces () (places []model.PlaceInfo, err error) {
	err = Mysql.Table("locations").Select("id", "location_name").Where("is_suspended=0").Order("location_abbr").Scan(&places).Error
	return
}