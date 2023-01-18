package mysqloperations

import (
	"api/entities"
	"database/sql"
	"fmt"
)

type ItemModel struct {
	Db *sql.DB
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
	result, err := itemModel.Db.Exec("UPDATE items SET time=?, number1=?, ingredient1=?, number2=?, ingredient2=?, number3=?, ingredient3=?, result=?, machineType=? WHERE name=?", item.Time, item.Recipe[0].Number, item.Recipe[0].Item, item.Recipe[1].Number, item.Recipe[1].Item, item.Recipe[2].Number, item.Recipe[2].Item, item.Result, item.MachineType, item.Name)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (itemModel ItemModel) Create(item *entities.Item) (int64, error) {
	fmt.Println(*item)
	result, err := itemModel.Db.Exec("INSERT INTO items(name, time, number1, ingredient1, number2, ingredient2, number3, ingredient3, result, machineType) VALUES (?,?,?,?,?,?,?,?,?,?)", item.Name, item.Time, item.Recipe[0].Number, item.Recipe[0].Item, item.Recipe[1].Number, item.Recipe[1].Item, item.Recipe[2].Number, item.Recipe[2].Item, item.Result, item.MachineType)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (itemModel ItemModel) Find(name string) (entities.Item, error) {

	rows, err := itemModel.Db.Query("SELECT * FROM `items` WHERE name=?", name)
	if err != nil {
		return entities.Item{}, err
	} else {
		item := entities.Item{}
		for rows.Next() {
			var name string
			var time float32
			var numbers [3]int
			var ingredients [3]string
			var machineType string
			var result int
			err := rows.Scan(&name, &time, &numbers[0], &ingredients[0], &numbers[1], &ingredients[1], &numbers[2], &ingredients[2], &result, &machineType)
			if err != nil {
				return entities.Item{}, err
			}
			recipe := []entities.Ingredient{}
			for i, number := range numbers {
				recipe = append(recipe, entities.Ingredient{Number: number, Item: ingredients[i]})
			}
			item = entities.Item{Name: name, Time: time, Recipe: recipe, MachineType: machineType, Result: result}
		}
		return item, nil
	}

}

func (itemModel ItemModel) FindAll() ([]entities.Item, error) {

	rows, err := itemModel.Db.Query("SELECT * FROM items")

	if err != nil {
		return nil, err
	}
	items := []entities.Item{}
	for rows.Next() {
		var name string
		var time float32
		var numbers [3]int
		var ingredients [3]string
		var machineType string
		var result int
		err := rows.Scan(&name, &time, &numbers[0], &ingredients[0], &numbers[1], &ingredients[1], &numbers[2], &ingredients[2], &result, &machineType)
		if err != nil {
			return nil, err
		}
		recipe := []entities.Ingredient{}
		for i, number := range numbers {
			recipe = append(recipe, entities.Ingredient{Number: number, Item: ingredients[i]})
		}
		item := entities.Item{Name: name, Time: time, Recipe: recipe, MachineType: machineType, Result: result}

		items = append(items, item)

	}
	return items, nil

}
