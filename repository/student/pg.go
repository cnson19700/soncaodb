package student

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/soncaodb/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewPGRepository(getClient func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getClient}
}

func (r *pgRepository) GetById(ctx context.Context, ID int64) (*model.Student, error) {
	db := r.getClient(ctx)
	student := &model.Student{}

	err := db.Where("id = ?", ID).
		First(student).Error

	if err != nil {
		return nil, errors.Wrap(err, "get student by id")
	}

	return student, nil
}

func (r *pgRepository) GetAll(ctx context.Context) ([]model.Student, error) {
	db := r.getClient(ctx)
	listStu := []model.Student{}

	db.Find(&listStu)

	return listStu, nil
}

func (r *pgRepository) Delete(ctx context.Context, ID int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", ID).Delete(&model.Student{}).Error
	


	if err != nil {
		log.Fatal(err)
	}

	return errors.Wrap(err, "delete fail")
}

func (r *pgRepository) Create(ctx context.Context, student *model.Student) (*model.Student, error) {
	db := r.getClient(ctx)
	err := db.Create(student).Error

	return student, errors.Wrap(err, "create user")
}

func (r *pgRepository) Update(ctx context.Context, student *model.Student) (*model.Student, error) {
	db := r.getClient(ctx)

	err := db.Model(&student).Updates(&student).Error

	return student, errors.Wrap(err, "update user")
}

func (r *pgRepository) ShowStudent(ctx context.Context, st []model.Student) {
	fmt.Print("\nSTT\tID\tName\t\tMath\tPhysic\tChemical\t")
	for i := 0; i < len(st); i++ {
		fmt.Printf("\n%d", i+1)
		fmt.Printf("\t%d", st[i].ID)
		fmt.Printf("\t%s\t", st[i].Name)
		fmt.Printf("\t%.2f\t%.2f\t%.2f", st[i].MathSco, st[i].PhysicSco, st[i].ChemiSco)
	}
	fmt.Print("\n--------------------------------------------------------\n")
}
