// Code generated by gnorm, DO NOT EDIT!

package nbsql_config

import (
	"github.com/mwbrown/nagbot/ndb/nbsql"
	"github.com/pkg/errors"
)

// Row represents a row from 'config'.
type Row struct {
	SchemaVer int // schema_ver
}

// Field values for every column in Config.
var (
	SchemaVerCol nbsql.IntField = "schema_ver"
)

// Query retrieves rows from 'config' as a slice of Row.
func Query(db nbsql.DB, where nbsql.WhereClause) ([]*Row, error) {

	var sqlstr string
	var wherevals []interface{}

	const origsqlstr = `SELECT 
		schema_ver
		FROM public.config`

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
		err := q.Scan(&r.SchemaVer)
		if err != nil {
			return nil, err
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// One retrieve one row from 'config'.
func One(db nbsql.DB, where nbsql.WhereClause) ([]*Row, error) {
	const origsqlstr = `SELECT 
		schema_ver
		FROM public.config WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	var vals []*Row
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		r := Row{}
		err := q.Scan(&r.SchemaVer)
		if err != nil {
			return nil, err
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// Insert inserts the row into the database.
func Insert(db nbsql.DB, r *Row) error {
	const sqlstr = `INSERT INTO config (
			schema_ver
		) VALUES (
			$1
		)`
	_, err := db.Exec(sqlstr, r.SchemaVer)
	return errors.Wrap(err, "insert Config")
}
