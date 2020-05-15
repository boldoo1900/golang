package action

import (
	"api/types"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var tasks = types.Tasks{
	{
		ID:          "1",
		Title:       "dfasdfasdf",
		Description: "descrskaldjflasd",
	},
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("select title, description from tasks")
	if err != nil {
		panic(err.Error())
	}

	var taskss []types.Task
	for results.Next() {
		var task types.Task

		err = results.Scan(&task.Title, &task.Description)
		if err != nil {
			panic(err.Error())
		}

		taskss = append(taskss, task)
	}

	// json.NewEncoder(w).Encode(events)
	json.NewEncoder(w).Encode(taskss)
}

func GetTaskOne(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)["id"]
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	var task types.Task
	err = db.QueryRow("select title, description from tasks where id = "+taskId).Scan(&task.Title, &task.Description)
	if err != nil {
		panic(err.Error())
	}

	// for _, singleEvent := range tasks {
	// 	if singleEvent.ID == taskId {
	// 		json.NewEncoder(w).Encode(singleEvent)
	// 	}
	// }

	json.NewEncoder(w).Encode(tasks)
}
