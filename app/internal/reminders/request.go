package reminders

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Reminder struct {
	Id          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Priority    int    `json:"priority" db:"priority"`
	RemindOn    string `json:"remind_on" db:"remind_on"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

func NewReminderFromHttpRequest(r *http.Request) (reminder *Reminder, err error) {
	err = json.NewDecoder(r.Body).Decode(&reminder)
	return
}

func NewReminderFromDbRow(rows *sql.Rows) (reminder *Reminder, err error) {
	reminder = new(Reminder)
	err = rows.Scan(&reminder.Id, &reminder.Title, &reminder.Description, &reminder.Priority, &reminder.RemindOn, &reminder.CreatedAt, &reminder.UpdatedAt)
	return

}

const (
	ErrEmptyTitle    = "title cannot be empty"
	ErrEmptyDesc     = "description cannot be empty"
	ErrEmptyRemindOn = "Remind On cannot be empty and must be of format DD-MM-YYYY HH:MM "
)

func (r *Reminder) Validate() (err error) {
	if r.Title == "" {
		err = errors.New(ErrEmptyTitle)
		return
	}
	_, err = time.Parse("02-01-2006 15:04", r.RemindOn)
	if err != nil {
		err = errors.New(ErrEmptyRemindOn)
		return
	}
	if r.Description == "" {
		err = errors.New(ErrEmptyDesc)
		return
	}
	return
}

const (
	ErrEmptyId = "id cannot be empty"
)

type ID string

func NewIdFromHttpRequest(r *http.Request) (id ID) {
	id = ID(r.URL.Query().Get("id"))
	return
}

func (id ID) Validate() (err error) {
	if id == "" {
		err = errors.New(ErrEmptyId)
		return
	}
	return
}
