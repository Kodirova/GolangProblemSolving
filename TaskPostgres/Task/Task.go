package Task

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

func CreateTask(db *sql.DB, t Task) error {
	time := time.Now()
	t.CreatedAt = time.String()
	t.DueDate = time.String()
	fmt.Println("we are here")
	sqlStatement := `INSERT INTO tasks(name, status, priority, createdat, createdby, duedate) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(sqlStatement, t.Name, t.Status, t.Priority, t.CreatedAt, t.CreatedBy, t.DueDate).Scan(&t.ID)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(db *sql.DB, t Task, n int) int64 {
	sqlstatement := "update tasks set name=$1, status=$2, priority=$3, createdby=$4 where id=$5"
	res, err := db.Exec(sqlstatement, t.Name, t.Status, t.Priority, t.CreatedBy, n)
	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteTask(db *sql.DB, n int) error {
	_, err := db.Exec("delete from tasks where id = $1", n)
	fmt.Println(err)
	fmt.Println("successfully deleted")
	return err
}

func GetTask(db *sql.DB, n int) (Task, error) {
	var task Task
	sqlstatement := "select name, status, priority, createdat, createdby, duedate from tasks where id =$1"
	row := db.QueryRow(sqlstatement, n)
	err := row.Scan(&task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)
	fmt.Println(task)
	return task, err
}

func ListTasks(db *sql.DB) ([]Task, error) {
	var tasks []Task
	rows, err := db.Query("select * from tasks order by name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Name, &t.Status, &t.Priority, &t.CreatedAt, &t.CreatedBy, &t.DueDate)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, t)
	}
	fmt.Println(tasks)
	return tasks, nil
}
