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
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id" binding:"required"`
}

// Save is a method on Event that saves it to the database
func (e *Event) Save() error {
	query := "INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

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

// GetEvents is a function that returns all events from the database
func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events"
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

// GetEvent is a function that returns a single event by ID
func GetEvent(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Update is a method on Event that updates it in the database
func (e Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, date_time = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err

	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	return err

}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err

	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}
