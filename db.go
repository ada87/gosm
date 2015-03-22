package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type fieldValue struct {
	Value string
	Desc  string
}
type field struct {
	Field_desc string
	Values     []fieldValue
}

func (f *field) appendVal(val, desc string) {
	f.Values = append(f.Values, fieldValue{val, desc})
}

var Fields = make(map[string]field)

func init() {
	db, err := sql.Open("sqlite3", "gosm-field")
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT f.field_code,f.field_desc,v.field_value,v.value_desc FROM Field AS f LEFT JOIN  Field_Value AS v  ON f.field_id = v.field_id")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		var field_code, field_desc, field_value, value_desc sql.NullString
		rows.Scan(&field_code, &field_desc, &field_value, &value_desc)
		//		checkErr(err)

		key := cString(field_code)
		f, has := Fields[key]
		if has {
			f.appendVal(cString(field_value), cString(value_desc))
			Fields[key] = f
		} else {
			if field_value.Valid {
				Fields[key] = field{cString(field_desc), []fieldValue{fieldValue{cString(field_value), cString(value_desc)}}}
			} else {
				Fields[key] = field{cString(field_desc), []fieldValue{}}
			}
		}
	}
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
