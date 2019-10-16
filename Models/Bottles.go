package Models

import (
	"../Config"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type Bottle struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}


func (b *Bottle) TableName() string {
	return "bottle"
}


func GetAllBottle(b *[]Bottle) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}


func GetOneBottle(b *Bottle, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}