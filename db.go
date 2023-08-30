package main

import (
	"database/sql"
	"os"
	//        "fmt"
	clist "github.com/charmbracelet/bubbles/list"
	//        "time"
	_ "github.com/mattn/go-sqlite3"
)

func db() error{
fileName:=getWorkingDirectory()+"/xnotepro.db"
file,err:=os.OpenFile(fileName,os.O_WRONLY|os.O_APPEND|os.O_CREATE,0644)
if err!=nil{
	return err
}
defer file.Close()

//db, err := sql.Open("sqlite3", dir)
 //       checkErr(err)
	//	stmt,err:=db.Prepare("INSERT INTO tasks(task,description) values(?,?)")
	//	checkErr(err)

	//	res,err:=stmt.Exec("nothing","nothing")
	//	checkErr(err)
	//	id, err := res.LastInsertId()
    //    checkErr(err)
	//	fmt.Println(id)

//		db.Close()
}
func checkErr(err error) {
        if err != nil {
            panic(err)
        }
}

func getTasks() []clist.Item{

result:=[]clist.Item{}
dir:=getWorkingDirectory()+"/xnotepro.db"
db, err := sql.Open("sqlite3", dir)
        checkErr(err)
rows,err:=db.Query("Select task,description From tasks")
		checkErr(err)
var task string
var description string
for rows.Next() {
	err=rows.Scan(&task,&description)
	checkErr(err)
	result=append(result, TimeCell{title:task,description:description})
}

db.Close()	

return	result 
}

