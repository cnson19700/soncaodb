package main

import (
	"context"
	"log"
	"time"

	"github.com/soncaodb/client/mysql"
	"github.com/soncaodb/config"

	//serviceHttp "github.com/soncaodb/delivery/http"
	"github.com/soncaodb/migration"
	"github.com/soncaodb/repository"
	//"github.com/soncaodb/usecase"
)

func main() {
	cfg := config.GetConfig()

	// setup locale
	{
		loc, _ := time.LoadLocation(cfg.TimeZone)
		time.Local = loc
	}
	mysql.Init()
	migration.Up()

	repo := repository.New(mysql.GetClient)
	//student, err := repo.Student.GetById(context.Background(), 1)
	students, err := repo.Student.GetAll(context.Background())

	//fmt.Println(students)
	//log.Fatal(student)
	repo.Student.Delete(context.Background(), 4)
	//log.Fatal(students)
	// var student model.Student
	// student.ID = 4
	// student.Name = "Link"
	// student.MathSco = 8.6
	// student.PhysicSco = 7.4
	// student.ChemiSco = 8.4
	// repo.Student.Create(context.Background(), &student)
	//fmt.Println(students)
	repo.Student.ShowStudent(context.Background(), students)

	//fmt.Println(students)

	if err != nil {
		log.Fatal(err)
	}
	//ucase := usecase.New(repo)

	//executeServer(repo, ucase)

}

// func executeServer(repo *repository.Repository, ucase *usecase.UseCase) {
// 	cfg := config.GetConfig()

// 	l, err := net.Listen("tcp", ":"+cfg.Port)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	errs := make(chan error)

// 	// http
// 	h := serviceHttp.NewHTTPHandler(repo, ucase)

// 	go func() {
// 		h.Listener = l

// 		log.Printf("Server is running on http://localhost:%s", cfg.Port)
// 		errs <- h.Start("")
// 	}()

// 	log.Println("exit", <-errs)
// }
