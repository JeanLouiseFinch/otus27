package sql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JeanLouiseFinch/otus27/api/config"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Event struct {
	Title       string
	Description string
	Start       time.Time
	End         time.Time
}

func Connect() (*sqlx.DB, error) {
	cfg, err := config.GetConfigDB("")
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable", cfg.HostDB, cfg.UserDB, cfg.NameDB, cfg.PortDB, cfg.PasswordDB))
	if err != nil {
		log.Fatalln(err)
	}
	return db, nil
}

func NewEvent(ctx context.Context, db *sqlx.DB, calendar_id int, title, description string, start, end time.Time) (string, error) {
	var id uuid.UUID
	err := db.QueryRowContext(ctx, "INSERT INTO events (id,calendar_id,title,descr,start_time,end_time) VALUES($1,$2,$3,$4,$5,$6) RETURNING id", uuid.New(), calendar_id, title, description, start, end).Scan(&id)
	return id.String(), err
}

func GetAllEvents(ctx context.Context, db *sqlx.DB, calendar_id int) ([]Event, error) {
	rows, err := db.QueryContext(ctx, "SELECT title,descr,start_time,end_time FROM events WHERE calendar_id=$1", calendar_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]Event, 0)
	for rows.Next() {
		ev := Event{}
		err = rows.Scan(&ev.Title, &ev.Description, &ev.Start, &ev.End)
		if err != nil {
			continue
		}
		result = append(result, ev)
	}
	return result, nil
}

func GetEvent(ctx context.Context, db *sqlx.DB, id string) (*Event, error) {
	result := Event{}
	err := db.QueryRowContext(ctx, "SELECT title,descr,start_time,end_time FROM events WHERE id=$1", id).Scan(&result.Title, &result.Description, &result.Start, &result.End)
	return &result, err
}

func ModifyEvent(ctx context.Context, db *sqlx.DB, id string, title, description string, start, end time.Time) (*Event, error) {
	result := &Event{}
	err := db.QueryRowContext(ctx, "UPDATE events SET title=$1,descr=$2,start_time=$3,end_time=$4 WHERE id=$5 RETURNING title,descr,start_time,end_time", title, description, start, end, id).Scan(&result.Title, &result.Description, &result.Start, &result.End)
	return result, err
}
func RemoveEvent(ctx context.Context, db *sqlx.DB, id string) error {
	_, err := db.ExecContext(ctx, "DELETE FROM events WHERE id=$1", id)
	return err
}
func GetEventsByTime(ctx context.Context, db *sqlx.DB, duration time.Duration) ([]Event, error) {
	t1 := time.Now()
	t2 := t1.Add(duration)
	results, err := db.QueryContext(ctx, "SELECT title,descr,start_time,end_time FROM events WHERE start_time BETWEEN $1 AND $2", t1, t2)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	events := make([]Event, 0)
	for results.Next() {
		ev := Event{}
		err = results.Scan(&ev.Title, &ev.Description, &ev.Start, &ev.End)
		if err != nil {
			continue
		}
		events = append(events, ev)
	}
	return events, err
}
