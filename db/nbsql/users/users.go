// Code generated by gnorm, DO NOT EDIT!

package nbsql_users

import (
	"github.com/mwbrown/nagbot/db/nbsql"
	"github.com/pkg/errors"
)

// Row represents a row from 'users'.
type Row struct {
	ID         int    // id (PK)
	Username   string // username
	PwHash     string // pw_hash
	PwSalt     string // pw_salt
	IsEnabled  bool   // is_enabled
	IsAdmin    bool   // is_admin
	MinSessID  int    // min_sess_id
	NextSessID int    // next_sess_id
}

// Field values for every column in Users.
var (
	IDCol         nbsql.IntField    = "id"
	UsernameCol   nbsql.StringField = "username"
	PwHashCol     nbsql.StringField = "pw_hash"
	PwSaltCol     nbsql.StringField = "pw_salt"
	IsEnabledCol  nbsql.BoolField   = "is_enabled"
	IsAdminCol    nbsql.BoolField   = "is_admin"
	MinSessIDCol  nbsql.IntField    = "min_sess_id"
	NextSessIDCol nbsql.IntField    = "next_sess_id"
)

// Query retrieves rows from 'users' as a slice of Row.
func Query(db nbsql.DB, where nbsql.WhereClause) ([]*Row, error) {

	var sqlstr string
	var wherevals []interface{}

	const origsqlstr = `SELECT 
		id, username, pw_hash, pw_salt, is_enabled, is_admin, min_sess_id, next_sess_id
		FROM public.users`

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
		err := q.Scan(&r.ID, &r.Username, &r.PwHash, &r.PwSalt, &r.IsEnabled, &r.IsAdmin, &r.MinSessID, &r.NextSessID)
		if err != nil {
			return nil, err
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// One retrieve one row from 'users'.
func One(db nbsql.DB, where nbsql.WhereClause) ([]*Row, error) {
	const origsqlstr = `SELECT 
		id, username, pw_hash, pw_salt, is_enabled, is_admin, min_sess_id, next_sess_id
		FROM public.users WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	var vals []*Row
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		r := Row{}
		err := q.Scan(&r.ID, &r.Username, &r.PwHash, &r.PwSalt, &r.IsEnabled, &r.IsAdmin, &r.MinSessID, &r.NextSessID)
		if err != nil {
			return nil, err
		}
		vals = append(vals, &r)
	}
	return vals, nil
}

// Insert inserts the row into the database.
func Insert(db nbsql.DB, r *Row) error {
	const sqlstr = `INSERT INTO users (
			id, username, pw_hash, pw_salt, is_enabled, is_admin, min_sess_id, next_sess_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)`
	_, err := db.Exec(sqlstr, r.ID, r.Username, r.PwHash, r.PwSalt, r.IsEnabled, r.IsAdmin, r.MinSessID, r.NextSessID)
	return errors.Wrap(err, "insert Users")
}

// Update updates the Row in the database.
func Update(db nbsql.DB, r *Row) error {
	const sqlstr = `UPDATE users SET (
			username, pw_hash, pw_salt, is_enabled, is_admin, min_sess_id, next_sess_id		
		) = ( 
			$1, $2, $3, $4, $5, $6, $7
		) WHERE
			id = $8
		`
	_, err := db.Exec(sqlstr, r.Username, r.PwHash, r.PwSalt, r.IsEnabled, r.IsAdmin, r.MinSessID, r.NextSessID, r.ID)
	return errors.Wrap(err, "update Users:")
}

// Upsert performs an upsert for Users.
//
// NOTE: PostgreSQL 9.5+ only
func Upsert(db nbsql.DB, r *Row) error {
	const sqlstr = `INSERT INTO users (
		username, pw_hash, pw_salt, is_enabled, is_admin, min_sess_id, next_sess_id, id
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8
	) ON CONFLICT (id) DO UPDATE SET (
		username, pw_hash, pw_salt, is_enabled, is_admin, min_sess_id, next_sess_id
	) = ( 
		$1, $2, $3, $4, $5, $6, $7
	)`

	_, err := db.Exec(sqlstr, r.Username, r.PwHash, r.PwSalt, r.IsEnabled, r.IsAdmin, r.MinSessID, r.NextSessID, r.ID)
	return errors.Wrap(err, "upsert Users")
}

// Delete deletes the Row from the database.
func Delete(
	db nbsql.DB,
	id int,
) error {
	const sqlstr = `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(
		sqlstr,
		id,
	)
	return errors.Wrap(err, "delete Users")
}