package Models

import (
	"../Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Joke struct {
	ID     int     `json:"id" binding:"required"`
	Likes  int     `json:"likes"`
	Joke   string  `json:"joke" binding:"required"`
}

func (b *Joke) TableName() string {
	return "joke"
}

func GetAllJoke(b *[]Joke) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewJoke(b *Joke) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneJoke(b *Joke, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneJoke(b *Joke, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteJoke(b *Joke, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}