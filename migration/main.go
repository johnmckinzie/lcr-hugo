package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	//	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/gosimple/slug"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	db, err := sql.Open("mysql", "creoto_lcr:ark135@tcp(mysql.creoto.com:3306)/creoto_lcr_dev")
	// db, err := sql.Open("mysql", "root:password@/lcr")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected")
	}

	fmt.Println("here")

	rows, err := db.Query("select id, name from lcr_bands")
	if err != nil {
		fmt.Println("select")
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		id   int
		name string
	)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		var bandSlug = slug.Make(name)
		log.Println(id, name, bandSlug)

		// ioutil.WriteFile("", d1, 0644)
	}

	defer db.Close()

	/*
		session, err := mgo.Dial("104.131.121.179")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		c := session.DB("bands").C("people")
		err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})

		if err != nil {
			log.Fatal(err)
		}

		result := Person{}
		err = c.Find(bson.M{"name": "Ale"}).One(&result)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Phone:", result.Phone)
	*/
}
