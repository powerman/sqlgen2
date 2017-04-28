// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

package demo

import (
	"database/sql"
	"fmt"
	"encoding/json"
)

// IssueTableName is the default name for this table.
const IssueTableName = "issues"

// IssueTable holds a given table name with the database reference, providing access methods below.
type IssueTable struct {
	Name string
	Db   *sql.DB
}

// ScanIssue reads a database record into a single value.
func ScanIssue(row *sql.Row) (*Issue, error) {
	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,

	)
	if err != nil {
		return nil, err
	}

	v := &Issue{}
	v.Id = v0
	v.Number = v1
	v.Title = v2
	v.Body = v3
	v.Assignee = v4
	v.State = v5
	json.Unmarshal(v6, &v.Labels)


	return v, nil
}

// ScanIssues reads database records into a slice of values.
func ScanIssues(rows *sql.Rows) ([]*Issue, error) {
	var err error
	var vv []*Issue

	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,

		)
		if err != nil {
			return vv, err
		}

		v := &Issue{}
		v.Id = v0
		v.Number = v1
		v.Title = v2
		v.Body = v3
		v.Assignee = v4
		v.State = v5
		json.Unmarshal(v6, &v.Labels)

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceIssue(v *Issue) []interface{} {
	var v0 int64
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	v0 = v.Id
	v1 = v.Number
	v2 = v.Title
	v3 = v.Body
	v4 = v.Assignee
	v5 = v.State
	v6, _ = json.Marshal(&v.Labels)

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,

	}
}

func SliceIssueWithoutPk(v *Issue) []interface{} {
	var v1 int
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 []byte

	v1 = v.Number
	v2 = v.Title
	v3 = v.Body
	v4 = v.Assignee
	v5 = v.State
	v6, _ = json.Marshal(&v.Labels)

	return []interface{}{
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,

	}
}

func (tbl IssueTable) SelectOne(query string, args ...interface{}) (*Issue, error) {
	row := tbl.Db.QueryRow(query, args...)
	return ScanIssue(row)
}

func (tbl IssueTable) Select(query string, args ...interface{}) ([]*Issue, error) {
	rows, err := tbl.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ScanIssues(rows)
}

func (tbl IssueTable) Insert(v *Issue) error {
	query := fmt.Sprintf(sInsertIssueStmt, tbl.Name)
	res, err := tbl.Db.Exec(query, SliceIssueWithoutPk(v)...)
	if err != nil {
		return err
	}

	v.Id, err = res.LastInsertId()
	return err
}

// Update updates a record. It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl IssueTable) Update(v *Issue) (int64, error) {
	query := fmt.Sprintf(sUpdateIssueByPkStmt, tbl.Name)
	args := SliceIssueWithoutPk(v)
	args = append(args, v.Id)
	return tbl.Exec(query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// It returns the number of rows affected.
// Not every database or database driver may support this.
func (tbl IssueTable) Exec(query string, args ...interface{}) (int64, error) {
	res, err := tbl.Db.Exec(query, args...)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

//--------------------------------------------------------------------------------

const sCreateIssueStmt = `
CREATE TABLE IF NOT EXISTS %s (
 id       SERIAL PRIMARY KEY ,
 number   INTEGER,
 title    VARCHAR(512),
 bigbody  VARCHAR(2048),
 assignee VARCHAR(512),
 state    VARCHAR(50),
 labels   BYTEA
);
`

func CreateIssueStmt(tableName string) string {
	return fmt.Sprintf(sCreateIssueStmt, tableName)
}

const sInsertIssueStmt = `
INSERT INTO %s (
 number,
 title,
 bigbody,
 assignee,
 state,
 labels
) VALUES ($1,$2,$3,$4,$5,$6)
`

func InsertIssueStmt(tableName string) string {
	return fmt.Sprintf(sInsertIssueStmt, tableName)
}

const sSelectIssueStmt = `
SELECT 
 id,
 number,
 title,
 bigbody,
 assignee,
 state,
 labels
FROM %s
`

func SelectIssueStmt(tableName string) string {
	return fmt.Sprintf(sSelectIssueStmt, tableName)
}

const sSelectIssueRangeStmt = `
SELECT 
 id,
 number,
 title,
 bigbody,
 assignee,
 state,
 labels
FROM %s
LIMIT $1 OFFSET $2
`

func SelectIssueRangeStmt(tableName string) string {
	return fmt.Sprintf(sSelectIssueRangeStmt, tableName)
}

const sSelectIssueCountStmt = `
SELECT count(1)
FROM %s 
`

func SelectIssueCountStmt(tableName string) string {
	return fmt.Sprintf(sSelectIssueCountStmt, tableName)
}

const sSelectIssueByPkStmt = `
SELECT 
 id,
 number,
 title,
 bigbody,
 assignee,
 state,
 labels
FROM %s
 WHERE id=$1
`

func SelectIssueByPkStmt(tableName string) string {
	return fmt.Sprintf(sSelectIssueByPkStmt, tableName)
}

const sUpdateIssueByPkStmt = `
UPDATE %s SET 
 id=$1,
 number=$2,
 title=$3,
 bigbody=$4,
 assignee=$5,
 state=$6,
 labels=$7 
 WHERE id=$8
`

func UpdateIssueByPkStmt(tableName string) string {
	return fmt.Sprintf(sUpdateIssueByPkStmt, tableName)
}

const sDeleteIssueByPkeyStmt = `
DELETE FROM %s
 WHERE id=$1
`

func DeleteIssueByPkeyStmt(tableName string) string {
	return fmt.Sprintf(sDeleteIssueByPkeyStmt, tableName)
}

//--------------------------------------------------------------------------------

const sCreateIssueAssigneeStmt = `
CREATE INDEX IF NOT EXISTS issue_assignee ON %s (assignee)
`

func CreateIssueAssigneeStmt(tableName string) string {
	return fmt.Sprintf(sCreateIssueAssigneeStmt, tableName)
}

const sSelectIssueAssigneeStmt = `
SELECT 
 id,
 number,
 title,
 bigbody,
 assignee,
 state,
 labels
FROM %s
 WHERE assignee=$1
`

func SelectIssueAssigneeStmt(tableName string) string {
	return fmt.Sprintf(sSelectIssueAssigneeStmt, tableName)
}

const sSelectIssueAssigneeRangeStmt = `
SELECT 
 id,
 number,
 title,
 bigbody,
 assignee,
 state,
 labels
FROM %s
 WHERE assignee=$1
LIMIT $2 OFFSET $3
`

func SelectIssueAssigneeRangeStmt(tableName string) string {
	return fmt.Sprintf(sSelectIssueAssigneeRangeStmt, tableName)
}

const sSelectIssueAssigneeCountStmt = `
SELECT count(1)
FROM %s 
 WHERE assignee=$1
`

func SelectIssueAssigneeCountStmt(tableName string) string {
	return fmt.Sprintf(sSelectIssueAssigneeCountStmt, tableName)
}
