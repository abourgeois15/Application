package mysqloperations

import (
	"api/entities"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

type ItemModel struct {
	Db *sql.DB
}

func (itemModel ItemModel) CreateTables() (int64, error) {
	file, err := os.Open("c:\\Users\\bor6rt\\go\\Application\\database\\database.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result sql.Result
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		query := scanner.Text()
		if scanner.Text() != "" {
			result, err = itemModel.Db.Exec(query)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (itemModel ItemModel) DeleteTable() (int64, error) {
	result, err := itemModel.Db.Exec("DROP TABLE `items`")
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (itemModel ItemModel) Delete(name string) (int64, error) {

	result, err := itemModel.Db.Exec("DELETE FROM items WHERE name=?", name)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (itemModel ItemModel) Update(item entities.Item) (int64, error) {

	result, err := itemModel.Db.Exec("UPDATE items SET name=?, time=?, number1=?, ingredient1=?, number2=?, ingredient2=?, number3=?, ingredient3=?, result=?, machineType=? WHERE id=?", item.Name, item.Time, item.Recipe[0].Number, item.Recipe[0].Item, item.Recipe[1].Number, item.Recipe[1].Item, item.Recipe[2].Number, item.Recipe[2].Item, item.Result, item.MachineType, item.Id)
	fmt.Println(item)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (itemModel ItemModel) Create(item *entities.Item) (int64, error) {

	result, err := itemModel.Db.Exec("INSERT INTO items(name, time, number1, ingredient1, number2, ingredient2, number3, ingredient3, result, machineType) VALUES (?,?,?,?,?,?,?,?,?,?)", item.Name, item.Time, item.Recipe[0].Number, item.Recipe[0].Item, item.Recipe[1].Number, item.Recipe[1].Item, item.Recipe[2].Number, item.Recipe[2].Item, item.Result, item.MachineType)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (itemModel ItemModel) Find(name string) (entities.Item, error) {

	rows, err := itemModel.Db.Query("SELECT * FROM items WHERE name=?", name)
	if err != nil {
		return entities.Item{}, err
	} else {
		item := entities.Item{}
		for rows.Next() {
			var id int
			var name string
			var time float32
			var numbers [3]int
			var ingredients [3]string
			var machineType string
			var result int
			err := rows.Scan(&id, &name, &time, &numbers[0], &ingredients[0], &numbers[1], &ingredients[1], &numbers[2], &ingredients[2], &result, &machineType)
			if err != nil {
				return entities.Item{}, err
			}
			recipe := [3]entities.Ingredient{}
			for i, number := range numbers {
				recipe[i] = entities.Ingredient{Number: number, Item: ingredients[i]}
			}
			item = entities.Item{Id: id, Name: name, Time: time, Recipe: recipe, MachineType: machineType, Result: result}
		}
		return item, nil
	}
}

func (itemModel ItemModel) FindAll() ([]string, error) {

	rows, err := itemModel.Db.Query("SELECT name FROM items")

	if err != nil {
		return nil, err
	}
	names := []string{}
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}
