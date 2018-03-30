// Code generated by gnorm, DO NOT EDIT!

package nbsql_taskschedules

import (
	"database/sql"

	"github.com/mwbrown/nagbot/db/nbsql"
	"github.com/pkg/errors"
)

// Row represents a row from 'task_schedules'.
type Row struct {
	ID           int           // id (PK)
	TaskID       int           // task_id
	OwnerID      int           // owner_id
	Type         int           // type
	ExactOnly    bool          // exact_only
	SchedTime    int           // sched_time
	SchedWeekday sql.NullInt64 // sched_weekday
	NextDue      int64         // next_due
	IsActive     bool          // is_active
}

// Field values for every column in TaskSchedules.
var (
	IDCol           nbsql.IntField          = "id"
	TaskIDCol       nbsql.IntField          = "task_id"
	OwnerIDCol      nbsql.IntField          = "owner_id"
	TypeCol         nbsql.IntField          = "type"
	ExactOnlyCol    nbsql.BoolField         = "exact_only"
	SchedTimeCol    nbsql.IntField          = "sched_time"
	SchedWeekdayCol nbsql.SqlNullInt64Field = "sched_weekday"
	NextDueCol      nbsql.Int64Field        = "next_due"
	IsActiveCol     nbsql.BoolField         = "is_active"
)

// Query retrieves rows from 'task_schedules' as a slice of Row.
func Query(db nbsql.DB, where nbsql.WhereClause) ([]*Row, error) {

	var sqlstr string
	var wherevals []interface{}

	const origsqlstr = `SELECT 
		id, task_id, owner_id, type, exact_only, sched_time, sched_weekday, next_due, is_active
		FROM public.task_schedules`

	// Allow a nil WhereClause to select all rows.
	if where != nil {
		idx := 1
		sqlstr = origsqlstr + " WHERE (" + where.String(&idx) + ") "
		wherevals = where.Values()
	} else {
		sqlstr = origsqlstr
		wherevals = []interface{}{}
	}

	var vals []*Row

	q, err := db.Query(sqlstr, wherevals...)

	if err != nil {
		return nil, err
	}
	for q.Next() {
		r := Row{}
		err := q.Scan(&r.ID, &r.TaskID, &r.OwnerID, &r.Type, &r.ExactOnly, &r.SchedTime, &r.SchedWeekday, &r.NextDue, &r.IsActive)
		if err != nil {
			return nil, err
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// One retrieve one row from 'task_schedules'.
func One(db nbsql.DB, where nbsql.WhereClause) ([]*Row, error) {
	const origsqlstr = `SELECT 
		id, task_id, owner_id, type, exact_only, sched_time, sched_weekday, next_due, is_active
		FROM public.task_schedules WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	var vals []*Row
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		r := Row{}
		err := q.Scan(&r.ID, &r.TaskID, &r.OwnerID, &r.Type, &r.ExactOnly, &r.SchedTime, &r.SchedWeekday, &r.NextDue, &r.IsActive)
		if err != nil {
			return nil, err
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// Insert inserts the row into the database.
func Insert(db nbsql.DB, r *Row) error {
	const sqlstr = `INSERT INTO task_schedules (
			id, task_id, owner_id, type, exact_only, sched_time, sched_weekday, next_due, is_active
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)`
	_, err := db.Exec(sqlstr, r.ID, r.TaskID, r.OwnerID, r.Type, r.ExactOnly, r.SchedTime, r.SchedWeekday, r.NextDue, r.IsActive)
	return errors.Wrap(err, "insert TaskSchedules")
}

// Update updates the Row in the database.
func Update(db nbsql.DB, r *Row) error {
	const sqlstr = `UPDATE task_schedules SET (
			task_id, owner_id, type, exact_only, sched_time, sched_weekday, next_due, is_active		
		) = ( 
			$1, $2, $3, $4, $5, $6, $7, $8
		) WHERE
			id = $9
		`
	_, err := db.Exec(sqlstr, r.TaskID, r.OwnerID, r.Type, r.ExactOnly, r.SchedTime, r.SchedWeekday, r.NextDue, r.IsActive, r.ID)
	return errors.Wrap(err, "update TaskSchedules:")
}

// Upsert performs an upsert for TaskSchedules.
//
// NOTE: PostgreSQL 9.5+ only
func Upsert(db nbsql.DB, r *Row) error {
	const sqlstr = `INSERT INTO task_schedules (
		task_id, owner_id, type, exact_only, sched_time, sched_weekday, next_due, is_active, id
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	) ON CONFLICT (id) DO UPDATE SET (
		task_id, owner_id, type, exact_only, sched_time, sched_weekday, next_due, is_active
	) = ( 
		$1, $2, $3, $4, $5, $6, $7, $8
	)`

	_, err := db.Exec(sqlstr, r.TaskID, r.OwnerID, r.Type, r.ExactOnly, r.SchedTime, r.SchedWeekday, r.NextDue, r.IsActive, r.ID)
	return errors.Wrap(err, "upsert TaskSchedules")
}

// Delete deletes the Row from the database.
func Delete(
	db nbsql.DB,
	id int,
) error {
	const sqlstr = `DELETE FROM task_schedules WHERE id = $1`

	_, err := db.Exec(
		sqlstr,
		id,
	)
	return errors.Wrap(err, "delete TaskSchedules")
}
