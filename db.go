package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
	"strconv"
)

func DbQuery(query string) {
	db, err := sql.Open("sqlite3", "gosm")
	checkErr(err)
	stmt, err := db.Prepare(query)
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		checkErr(err)
		fmt.Println(id, name)
	}

	stmt.Close()
	db.Close()
}

//type FieldValue struct {
//	Value string
//}
type Field struct {
	Formid     string
	Formtype   int
	Field_desc string
	Values     []string
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
		var field_code, field_formid, field_desc, field_value interface{}
		var field_formtype int
		rows.Scan(&field_code, &field_formid, &field_formtype, &field_desc, &field_value)
		//		checkErr(err)

		key := asString(field_code)
		_, has := Fields[key]

		if has {
			//			Fields[key].Values = append(Fields[key].Values, asString(field_value))
		} else {
			Fields[key] = Field{asString(field_formid), field_formtype, asString(field_desc), []string{asString(field_value)}}
		}
	}
	fmt.Println(Fields)
	stmt.Close()
	db.Close()
}
func asString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}
