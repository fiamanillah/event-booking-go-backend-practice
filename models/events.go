package models

import (
	"gin/db"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"` // Fixed typo
	UserID      int       `json:"user_id" binding:"required"`
}

func (e Event) Save() error {
	query := "INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close() // Important to close the statement

	// Execute the statement with parameters (removed the query string from Exec)
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = int(id)
	return nil
}

func GetEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, date_time, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}
