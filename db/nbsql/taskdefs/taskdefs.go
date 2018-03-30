// Code generated by gnorm, DO NOT EDIT!

package nbsql_taskdefs

import (
	"github.com/mwbrown/nagbot/db/nbsql"
	"github.com/pkg/errors"
)

// Row represents a row from 'task_defs'.
type Row struct {
	ID          int    // id (PK)
	Description string // description
	OwnerID     int    // owner_id
}

// Field values for every column in TaskDefs.
var (
	IDCol          nbsql.IntField    = "id"
	DescriptionCol nbsql.StringField = "description"
	OwnerIDCol     nbsql.IntField    = "owner_id"
)

// Query retrieves rows from 'task_defs' as a slice of Row.
func Query(db nbsql.DB, where nbsql.WhereClause) ([]*Row, error) {

	var sqlstr string
	var wherevals []interface{}

	const origsqlstr = `SELECT 
		id, description, owner_id
		FROM public.task_defs`

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
		err := q.Scan(&r.ID, &r.Description, &r.OwnerID)
		if err != nil {
			return nil, err
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// One retrieve one row from 'task_defs'.
func One(db nbsql.DB, where nbsql.WhereClause) ([]*Row, error) {
	const origsqlstr = `SELECT 
		id, description, owner_id
		FROM public.task_defs WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	var vals []*Row
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		r := Row{}
		err := q.Scan(&r.ID, &r.Description, &r.OwnerID)
		if err != nil {
			return nil, err
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// Insert inserts the row into the database.
func Insert(db nbsql.DB, r *Row) error {
	const sqlstr = `INSERT INTO task_defs (
			id, description, owner_id
		) VALUES (
			$1, $2, $3
		)`
	_, err := db.Exec(sqlstr, r.ID, r.Description, r.OwnerID)
	return errors.Wrap(err, "insert TaskDefs")
}

// Update updates the Row in the database.
func Update(db nbsql.DB, r *Row) error {
	const sqlstr = `UPDATE task_defs SET (
			description, owner_id		
		) = ( 
			$1, $2
		) WHERE
			id = $3
		`
	_, err := db.Exec(sqlstr, r.Description, r.OwnerID, r.ID)
	return errors.Wrap(err, "update TaskDefs:")
}

// Upsert performs an upsert for TaskDefs.
//
// NOTE: PostgreSQL 9.5+ only
func Upsert(db nbsql.DB, r *Row) error {
	const sqlstr = `INSERT INTO task_defs (
		description, owner_id, id
	) VALUES (
		$1, $2, $3
	) ON CONFLICT (id) DO UPDATE SET (
		description, owner_id
	) = ( 
		$1, $2
	)`

	_, err := db.Exec(sqlstr, r.Description, r.OwnerID, r.ID)
	return errors.Wrap(err, "upsert TaskDefs")
}

// Delete deletes the Row from the database.
func Delete(
	db nbsql.DB,
	id int,
) error {
	const sqlstr = `DELETE FROM task_defs WHERE id = $1`

	_, err := db.Exec(
		sqlstr,
		id,
	)
	return errors.Wrap(err, "delete TaskDefs")
}
