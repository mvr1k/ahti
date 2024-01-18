package reminders

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (s *service) runMigration() (err error) {
	createStmt := `CREATE TABLE IF NOT EXISTS reminders (
		id TEXT PRIMARY KEY,
		title TEXT default 'NA',
		description TEXT default 'NA',
		priority INTEGER default 0,
		remind_on timestamp default CURRENT_TIMESTAMP,
		created_at timestamp default CURRENT_TIMESTAMP,
		updated_at timestamp default CURRENT_TIMESTAMP
	)`

	_, err = s.db.Exec(createStmt)

	if err != nil {
		return
	}

	return

}

func (s *service) createReminder(reminder *Reminder) (err error) {
	reminder.Id = uuid.New().String()

	dbTime, _ := time.Parse("02-01-2006 15:04", reminder.RemindOn)

	insertQuery := "INSERT INTO reminders (id, title, description, priority, remind_on) VALUES ($1, $2, $3, $4, $5)"

	_, err = s.db.Exec(insertQuery, reminder.Id, reminder.Title, reminder.Description, reminder.Priority, dbTime)

	return

}

func (s *service) updateReminder(id ID, reminder *Reminder) error {
	updateQuery := "UPDATE reminders SET title = $1, description = $2, priority = $3, remind_on = $4 , updated_at = CURRENT_TIMESTAMP WHERE id = $5"

	dbTime, _ := time.Parse("02-01-2006 15:04", reminder.RemindOn)

	_, err := s.db.Exec(updateQuery, reminder.Title, reminder.Description, reminder.Priority, dbTime, id)

	return err

}

func (s *service) deleteReminder(id ID) error {
	deleteQuery := "DELETE FROM reminders WHERE id = $1"

	_, err := s.db.Exec(deleteQuery, id)

	return err
}

func (s *service) getReminder(id ID) (reminders []*Reminder, err error) {

	selectQuery := "SELECT id, title, description, priority, remind_on , created_at, updated_at FROM reminders"

	if id != "" {
		selectQuery += fmt.Sprintf(" WHERE id = '%v'", id)
	}

	rows, err := s.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		reminder, err := NewReminderFromDbRow(rows)
		if err != nil {
			return nil, err
		}
		reminders = append(reminders, reminder)
	}

	return reminders, err

}

func (s *service) getJobsToBeTriggeredNow() (list []*Reminder, err error) {
	currentTime := time.Now()

	selectQuery := "SELECT id, title, description, priority, remind_on , created_at, updated_at FROM reminders WHERE TO_CHAR(remind_on, 'DD-MM-YYYY HH24:MI') = $1"

	rows, err := s.db.Query(selectQuery, currentTime.Format("02-01-2006 15:04"))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		reminder, err := NewReminderFromDbRow(rows)
		if err != nil {
			s.logger.Error(err)
			continue
		}
		list = append(list, reminder)
	}

	return
}
