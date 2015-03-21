package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Field struct {
	Formid     string
	Formtype   int
	Field_desc string
	Values     []string
}

func (f *Field) appendVal(val string) {
	f.Values = append(f.Values, val)
}

var Fields = make(map[string]Field)

func init() {
	db, err := sql.Open("sqlite3", "gosm-field")
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT f.field_code,f.field_formid,f.field_formtype,f.field_desc,v.field_value FROM Field AS f LEFT JOIN  Field_Value AS v ")
	checkErr(err)
	defer stmt.Close()
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		var field_code, field_formid, field_desc, field_value sql.NullString
		var field_formtype int
		rows.Scan(&field_code, &field_formid, &field_formtype, &field_desc, &field_value)
		//		checkErr(err)

		key := cString(field_code)
		_, has := Fields[key]
		if has {
			f := Fields[key]
			f.appendVal(cString(field_value))
			Fields[key] = f
		} else {
			Fields[key] = Field{cString(field_formid), field_formtype, cString(field_desc), []string{cString(field_value)}}
		}
	}
	stmt.Close()
	db.Close()
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
