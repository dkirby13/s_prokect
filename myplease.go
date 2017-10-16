
package main

import (
  "html/template"
  "net/http"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Entry struct {
	Value         string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/update", UpdateTable)
	http.ListenAndServe(":7002", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("databasepage.html")
    checkError(err)
    err = t.Execute(w, getRows())
    checkError(err)
}

func UpdateTable(w http.ResponseWriter, r *http.Request){
     addRow(r.FormValue("value"))
     http.Redirect(w, r, "localhost:7002", 303)
}

func getRows() []Entry{
     con, err := sql.Open("mysql", "sql3199737:9M6ewCdtUn@tcp(sql3.freemysqlhosting.net:3306)/sql3199737")
     checkError(err)
     rows, err := con.Query("SELECT * FROM infotable")
     checkError(err)
     defer con.Close()
     r := []Entry{}
     row := Entry{}
     for rows.Next(){
     	 var value string
	 rows.Scan(&value)
	 row.Value = value
	 r = append(r, row)
      }
      return r
}

func addRow(value string){
      con, err := sql.Open("mysql", "sql3199737:9M6ewCdtUn@tcp(sql3.freemysqlhosting.net:3306)/sql3199737")
      _, err = con.Exec("INSERT INTO infotable(value) values(?)", value)
     checkError(err)
}	 

func checkError(err error){
     if err != nil {
     	panic(err)
     }
}