package service

import (
	"errors"
	"time"

	"github.com/eduardonakaidev/go-tasks/config"
	"github.com/eduardonakaidev/go-tasks/models"
	"github.com/eduardonakaidev/go-tasks/util"
)

func CreateTask(idUser string,taskDTO models.TaskInputPublic) error {

	var err error
	// connect to database
	conn, _ := config.ConnectPostgresDB()

	defer conn.Close()

	//verify idUser exists

	_,err = conn.Query(`SELECT id from users id=$1`,idUser)
	if err != nil {
		return errors.New("User already exists")
	}
	var task models.Task

	task.ID = util.NewUUID()
	task.IsCompleted = false
	task.IdUser = idUser
	task.CreatedAt = time.Now()
	
	sqlQueryRow := ``
	_, err = conn.Query(sqlQueryRow,task.ID,task.Title,task.Description,task.IsCompleted,task.IdUser,task.CreatedAt)

	return err
}
func UpdatedTask(idUser string, taskupdated models.TaskInputUpdate) (task models.Task, err error) {

	// connect to database
	conn, _ := config.ConnectPostgresDB()

	defer conn.Close()

	//verify idUser exists

	_,err = conn.Query(`SELECT id from users id=$1`,idUser)
	if err != nil {
		return task,errors.New("User already exists")
	}

	var taskQuery models.Task
	 err = conn.QueryRow(`SELECT id,title,description,iscompleted,iduser,createdat FROM task WHERE id=$1`,taskupdated.ID).Scan(
		&taskQuery.ID,
		&taskQuery.Title,
		&taskQuery.Description,
		&taskQuery.IsCompleted,
		&taskQuery.IdUser,
		&taskQuery.CreatedAt,
	 )
	 if err != nil {
		return task,errors.New("Task não encontrada")
	 }
	 taskQuery.Title = taskupdated.Title
	 taskQuery.Description = taskupdated.Description
	 taskQuery.IsCompleted = taskupdated.IsCompleted
	 _, err = conn.Exec(`
	UPDATE task
	SET title=$1,description=$2,iscompleted=$3 WHERE id=$4
	 `,taskQuery.Title,taskQuery.Description,taskQuery.IsCompleted,taskQuery.ID)
	 if err != nil {
		return task,errors.New("Erro ao atualiza a task")
	 }
	 return taskQuery, err


}