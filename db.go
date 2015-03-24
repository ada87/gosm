package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type fieldValue struct {
	Id    int
	Value string
	Desc  string
}
type field struct {
	Field_id   int
	Field_desc string
	Values     []fieldValue
}

func (f *field) appendVal(id int, val, desc string) {
	f.Values = append(f.Values, fieldValue{id, val, desc})
}

var Fields = make(map[string]field)

func insertVal(fid, fval, fdes string) error {
	db, err := sql.Open("sqlite3", "gosm-field")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO Field_Value(field_id,field_value,value_desc) VALUES (?,?,?)")
	checkErr(err)
	defer stmt.Close()
	_, rtnerr := stmt.Exec(fid, fval, fdes)
	reflush()
	return rtnerr
}
func updateVal(vid, fval, fdes string) error {
	db, err := sql.Open("sqlite3", "gosm-field")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("UPDATE Field_Value SET field_value=? , value_desc = ? WHERE id = ?")
	checkErr(err)
	defer stmt.Close()
	_, rtnerr := stmt.Exec(fval, fdes, vid)
	reflush()
	return rtnerr
}

func reflush() {
	for k, _ := range Fields {
		delete(Fields, k)
	}
	db, err := sql.Open("sqlite3", "gosm-field")
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT f.field_id,f.field_code,f.field_desc,v.id,v.field_value,v.value_desc FROM Field AS f LEFT JOIN  Field_Value AS v  ON f.field_id = v.field_id")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		var field_code, field_desc, field_value, value_desc sql.NullString
		var field_id, value_id sql.NullInt64
		rows.Scan(&field_id, &field_code, &field_desc, &value_id, &field_value, &value_desc)
		key := cString(field_code)
		f, has := Fields[key]
		if has {
			f.appendVal(cInt(value_id), cString(field_value), cString(value_desc))
			Fields[key] = f
		} else {
			if value_id.Valid {
				Fields[key] = field{cInt(field_id), cString(field_desc), []fieldValue{fieldValue{cInt(value_id), cString(field_value), cString(value_desc)}}}
			} else {
				Fields[key] = field{cInt(field_id), cString(field_desc), []fieldValue{}}
			}
		}
	}

}

func init() {
	reflush()
}

func cString(str sql.NullString) string {
	if str.Valid {
		return str.String
	} else {
		return ""
	}
}

func cInt(it sql.NullInt64) int {
	if it.Valid {
		return int(it.Int64)
	} else {
		return 0
	}
}

func cFloat(fl sql.NullFloat64) float64 {
	if fl.Valid {
		return fl.Float64
	} else {
		return 0.0
	}
}
