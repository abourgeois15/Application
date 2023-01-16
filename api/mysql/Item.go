package mysqloperations

import (
	"api/entities"
	"database/sql"
)

type ItemModel struct {
	Db *sql.DB
}

// func (itemModel ItemModel) Delete(id int64) (int64, error) {

// 	result, err := itemModel.Db.Exec("DELETE FROM items WHERE id=?", id)
// 	if err != nil {
// 		return 0, err
// 	} else {
// 		return result.RowsAffected()
// 	}

// }

// func (itemModel ItemModel) Update(item entities.Item) (int64, error) {

// 	result, err := itemModel.Db.Exec("UPDATE author SET name=?, email=?, WHERE id=?", item.Name, item.Email, item.Id)
// 	if err != nil {
// 		return 0, err
// 	} else {
// 		return result.RowsAffected()
// 	}

// }

// func (authorModel ItemModel) Create(item *entities.Item) error {

// 	result, err := authorModel.Db.Exec("INSERT INTO author(name, email) VALUES (?,?)", item.Name, item.Email)
// 	if err != nil {
// 		return err
// 	} else {
// 		item.Id, _ = result.LastInsertId()
// 		return nil
// 	}

// }

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
				if number != 0 {
					recipe = append(recipe, entities.Ingredient{Number: number, Item: ingredients[i]})
				} else {
					break
				}
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
			if number != 0 {
				recipe = append(recipe, entities.Ingredient{Number: number, Item: ingredients[i]})
			} else {
				break
			}
		}
		item := entities.Item{Name: name, Time: time, Recipe: recipe, MachineType: machineType, Result: result}

		items = append(items, item)

	}
	return items, nil

}
