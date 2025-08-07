package models

import (
	"eventsManagement/db"
	"time"
)

type Event struct {
	Id          int64
	UserName    string    `binding : "required`
	Description string    `binding : "required`
	Location    string    `binding : "required`
	DateTime    time.Time `binding : "required`
	UserId      int64
}

func (e *Event) Save() error {
	var insertEventQuery = `
	INSERT INTO events (UserName, Description, Location, DateTime, UserId)
	VALUES (?, ?, ?, ?, ?);
	`
	statement, err := db.DB.Prepare(insertEventQuery)

	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.UserName, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	e.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {
	var getAllEventsQuery = "SELECT * FROM events;"
	rows, err := db.DB.Query(getAllEventsQuery)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.UserName, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	getEventQuery := `
	SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(getEventQuery, id)

	var event Event

	err := row.Scan(&event.Id, &event.UserName, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) UpdateEventById(id int64) (int, error) {
	updateEventQuery := `
		UPDATE events
		SET
		username = ?,
		description = ?,
		location = ?,
		datetime = ?
		WHERE id = ?;
	`

	statement, err := db.DB.Prepare(updateEventQuery)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(e.UserName, e.Description, e.Location, e.DateTime, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()

	return int(rowsAffected), nil

}

func DeleteEventById(id int64) error {
	deleteEventQuery := `DELETE FROM events WHERE id=?`
	statement, err := db.DB.Prepare(deleteEventQuery)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Register(userId int64) error {
	registerEventQuery := `INSERT INTO registrations(event_id, user_id) VALUES (?, ?)`

	statement, err := db.DB.Prepare(registerEventQuery)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.Id, userId)
	if err != nil {
		return err
	}
	return nil

}

func (e Event) CancelRegistration(userId int64) error {
	cancelEventQuery := `DELETE INTO registrations WHERE event_id = ? AND user_id = ?`

	statement, err := db.DB.Prepare(cancelEventQuery)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.Id, userId)
	if err != nil {
		return err
	}
	return nil

}
