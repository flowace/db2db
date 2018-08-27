package main

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type AirQuality struct {
	// gorm.Model
	// ID      uint   `gorm:"column:id"`
	Index   string `gorm:"column:index"`
	BEN     string `gorm:"column:BEN"`
	CH4     string `gorm:"column:CH4"`
	CO      string `gorm:"column:CO"`
	EBE     string `gorm:"column:EBE"`
	MXY     string `gorm:"column:MXY"`
	NMHC    string `gorm:"column:NMHC"`
	NO      string `gorm:"column:NO"`
	NO2     string `gorm:"column:NO_2"`
	NOX     string `gorm:"column:NOx"`
	OXY     string `gorm:"column:OXY"`
	O3      string `gorm:"column:O_3"`
	PM10    string `gorm:"column:PM10"`
	PM25    string `gorm:"column:PM25"`
	PXY     string `gorm:"column:PXY"`
	SO2     string `gorm:"column:SO_2"`
	TCH     string `gorm:"column:TCH"`
	TOL     string `gorm:"column:TOL"`
	Time    string `gorm:"column:date; type:timestamp"`
	Station string `gorm:"column:station"`
}

func (AirQuality) TableName() string {
	return "AQ"
}

func main() {
	c := generateRowsConcurrent("boring!!")

	for row := range c {
		fmt.Println(row)
	}
	// for {
	// 	fmt.Println(<-c)
	// 	if c == nil {
	// 		fmt.Println("Bye")
	// 		break
	// 	}
	// }
}

func generateRowsConcurrent(msg string) <-chan []string {
	c := make(chan []string)
	go func() {
		db, err := gorm.Open("sqlite3", "./load_testing_7.6m.db")
		if err != nil {
			panic("failed to connect database")
		}
		defer db.Close()
		rows, err := db.Model(&AirQuality{}).Limit(20).Rows()
		defer rows.Close()
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var aq AirQuality
			db.ScanRows(rows, &aq)
			v := reflect.Indirect(reflect.ValueOf(aq))
			var buf []string
			for i := 0; i < v.NumField(); i++ {
				buf = append(buf, v.Field(i).String())
			}
			c <- buf
		}

		defer close(c)
	}()
	return c
}

