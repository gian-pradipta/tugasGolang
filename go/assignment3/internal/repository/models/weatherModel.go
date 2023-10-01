package models

import (
	"assignment_3/internal/repository/database"
	"database/sql"
)

var db *sql.DB = database.New()

type Weather struct {
	Water float64 `json:"water"`
	Wind  float64 `json:"wind"`
}

func GetData() (*Weather, error) {
	var result Weather
	var err error
	rows, err := db.Query("SELECT water, wind FROM weather")
	var water float64
	var wind float64

	for rows.Next() {
		err = rows.Scan(&water, &wind)
		result.Water = water
		result.Wind = wind
	}
	defer rows.Close()

	return &result, err
}

func UpdateData(wind float64, water float64) (sql.Result, error) {
	result, err := db.Exec("UPDATE weather SET wind = ?, water = ?", water, wind)
	return result, err
}
