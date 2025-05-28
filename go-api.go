package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "modernc.org/sqlite"
)

type Tree struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Height int    `json:"height"`
}

type TempReading struct {
	ID          int `json:"id"`
	TempReading int `json:"temp_reading"`
}

func main() {

	// Connect to the Database
	db, err := sql.Open("sqlite", "./temp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query data
	rows, err := db.Query("SELECT id, temp_reading FROM temp")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/temps", func(c echo.Context) error {

		var readings []TempReading
		for rows.Next() {

			var tr TempReading
			if err := rows.Scan(&tr.ID, &tr.TempReading); err != nil {
				return err
			}
			readings = append(readings, tr)

		}
		return c.JSON(http.StatusOK, readings)
	})

	e.GET("/tree/tallest", func(c echo.Context) error {
		tree := Tree{
			ID:     45,
			Type:   "Maple",
			Height: 60,
		}
		return c.JSON(http.StatusOK, tree)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
