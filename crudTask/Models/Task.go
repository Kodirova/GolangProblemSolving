package Models

import (
	"crudTask/Database"
	"fmt"
	"time"
)

const TableTaskCreation = "CREATE TABLE IF NOT EXISTS task(id SERIAL, name TEXT NOT NULL, status TEXT NOT NULL, priority TEXT NOT NULL, createdat TEXT, createdby TEXT, duedate TEXT)"

func CreateTask(task *Task) error {
	db := Database.ConnectDB()
	res, err1 := db.Exec(TableTaskCreation)
	if err1 != nil {
		return err1
	}
	fmt.Println(res)
	time := time.Now()
	task.CreatedAt = time.String()
	task.DueDate = time.String()
	sqlStatement := `INSERT INTO task(name, status, priority, createdat, createdby, duedate) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(sqlStatement, task.Name, task.Status, task.Priority, task.CreatedAt, task.CreatedBy, task.DueDate).Scan(&task.ID)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(t *Task, n string) error {
	db := Database.ConnectDB()
	res, err1 := db.Exec(TableTaskCreation)
	if err1 != nil {
		return err1
	}
	fmt.Println(res)
	sqlstatement := "update task set name=$1, status=$2, priority=$3, createdby=$4 where id=$5"
	res, err := db.Exec(sqlstatement, t.Name, t.Status, t.Priority, t.CreatedBy, n)
	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return nil
}

func DeleteTask(task *Task, n string) error {
	db := Database.ConnectDB()
	res, err1 := db.Exec(TableTaskCreation)
	if err1 != nil {
		return err1
	}
	fmt.Println(res)
	_, err := db.Exec("delete from task where id = $1", n)
	fmt.Println(err)
	fmt.Println("successfully deleted")
	return err
}

func GetTask(task *Task, n string) error {
	db := Database.ConnectDB()
	res, err1 := db.Exec(TableTaskCreation)
	if err1 != nil {
		return err1
	}
	fmt.Println(res)
	sqlstatement := "select name, status, priority, createdat, createdby, duedate from task where id =$1"
	row := db.QueryRow(sqlstatement, n)
	err := row.Scan(&task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)
	fmt.Println(task)
	return err
}

func ListTasks(task *[]Task) error {
	db := Database.ConnectDB()

	rows, err := db.Query("select * from task order by name")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Name, &t.Status, &t.Priority, &t.CreatedAt, &t.CreatedBy, &t.DueDate)
		if err != nil {
			return err
		}
		*task = append(*task, t)
	}

	return nil
}
