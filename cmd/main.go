package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/soncaodb/client/mysql"
	"github.com/soncaodb/config"
	"github.com/soncaodb/model"
	"github.com/soncaodb/package/keyboard"
	"github.com/soncaodb/package/namestand"

	//serviceHttp "github.com/soncaodb/delivery/http"
	"github.com/soncaodb/migration"
	"github.com/soncaodb/repository"
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

	fmt.Println()
	fmt.Print("-----------------------MENU-----------------------\n")
	fmt.Print("--  1. Show All Student.\n")
	fmt.Print("--  2. Add student.\n")
	fmt.Print("--  3. Find student.\n")
	fmt.Print("--  4. Delete Student.\n")
	fmt.Print("--  5. Update Student.\n")
	fmt.Print("--  0. Exit\n")
	fmt.Print("--------------------------------------------------\n")
	fmt.Println()

Menu:
	for {
		listStu, err := repo.Student.GetAll(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Enter your number: ")
		key := keyboard.GetIntNumber()
		switch key {
		case 1:
			{
				fmt.Println("Show all Students")
				repo.Student.ShowStudent(context.Background(), listStu)
				continue
			}
		case 2:
			{
				fmt.Println("Add Student")
				var st model.Student
				fmt.Print("Type ID: ")
				st.ID = int64(keyboard.GetIntNumber())
				fmt.Print("Type your name: ")
				st.Name = namestand.Check(keyboard.GetText())
				fmt.Print("Type Math Score: ")
				st.MathSco = keyboard.GetFloatNumber()
				fmt.Print("Type Physic Score: ")
				st.PhysicSco = keyboard.GetFloatNumber()
				fmt.Print("Type Literature Score: ")
				st.ChemiSco = keyboard.GetFloatNumber()
				repo.Student.Create(context.Background(), &st)
				continue
			}
		case 3:
			{
				fmt.Print("Enter Student Name to find: ")
				txt := keyboard.GetText()
				studentFind, err := repo.Student.FindStudent(context.Background(), txt)
				if err != nil {
					fmt.Println("Cannot find Student")
				}
				//fmt.Println(studentFind)
				fmt.Print("\nID\tName\t\tMath\tPhysic\tChemical")
				fmt.Println()
				fmt.Printf("%d", studentFind.ID)
				fmt.Printf("\t%s\t", studentFind.Name)
				fmt.Printf("\t%.2f\t%.2f\t%.2f", studentFind.MathSco, studentFind.PhysicSco, studentFind.ChemiSco)
				fmt.Println()
				continue
			}
		case 4:
			{
				fmt.Println("Delete Student")
				fmt.Print("Type Student ID: ")
				id := keyboard.GetIntNumber()
				isDel := repo.Student.Delete(context.Background(), int64(id))
				if isDel != nil {
					fmt.Print("Deleted Success!")
				}
				continue
			}
		case 5:
			{
				fmt.Println("Update Student")
				fmt.Print("Type Student ID: ")
				id := keyboard.GetIntNumber()
				st, err := repo.Student.GetById(context.Background(), int64(id))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Print("Update student name: ")
				st.Name = namestand.Check(keyboard.GetText())
				fmt.Print("Update Math Score: ")
				st.MathSco = keyboard.GetFloatNumber()
				fmt.Print("Update Physic Score: ")
				st.PhysicSco = keyboard.GetFloatNumber()
				fmt.Print("Update Literature Score: ")
				st.ChemiSco = keyboard.GetFloatNumber()
				repo.Student.Update(context.Background(), st)
				continue
			}
		case 0:
			{
				fmt.Println("Saved Success!!")
				fmt.Println("Thank You for Using App ^^")
				break Menu
			}
		default:
			fmt.Println("NO options key !")
			break Menu
		}
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
