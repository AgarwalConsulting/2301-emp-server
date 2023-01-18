package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"algogrit.com/emp_server/entities"
)

type sqlRepo struct {
	*sql.DB
}

func (repo *sqlRepo) ListAll() (allEmps []entities.Employee, err error) {
	rows, err := repo.DB.Query("SELECT * FROM employees;")

	if err != nil {
		return
	}

	allEmps = []entities.Employee{}

	for rows.Next() {
		var emp entities.Employee
		err = rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.ProjectID)

		if err != nil {
			return
		}

		allEmps = append(allEmps, emp)
	}

	return allEmps, nil
}

func (repo *sqlRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	_, err := repo.DB.Exec("INSERT INTO employees (name, department, project_id) VALUES ($1, $2, $3)", newEmp.Name, newEmp.Department, newEmp.ProjectID)

	if err != nil {
		return nil, err
	}

	// lastInsertedID, err := res.LastInsertId()

	// if err != nil {
	// 	return nil, err
	// }

	// newEmp.ID = int(lastInsertedID)

	return &newEmp, nil
}

func NewSQL(driverName string, connString string) EmployeeRepository {
	db, err := sql.Open(driverName, connString)

	if err != nil {
		log.Fatal("Unable to connect:", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS employees(id SERIAL primary key, name text, department text, project_id numeric);")

	if err != nil {
		log.Fatal("Unable to create table:", err)
	}

	return &sqlRepo{db}
}
