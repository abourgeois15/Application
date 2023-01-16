package mysqloperations

import (
	"api/entities"
	"database/sql"
)

type MachineModel struct {
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

func (machineModel MachineModel) Find(id string) ([]entities.Machine, error) {

	rows, err := machineModel.Db.Query("SELECT * FROM `machines` WHERE name=? OR type=?", id, id)
	if err != nil {
		return nil, err
	} else {
		machines := []entities.Machine{}
		for rows.Next() {
			var name string
			var mtype string
			var numbers [3]int
			var ingredients [3]string
			var time float32
			var speed float32
			err := rows.Scan(&name, &mtype, &numbers[0], &ingredients[0], &numbers[1], &ingredients[1], &numbers[2], &ingredients[2], &time, &speed)
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
			machine := entities.Machine{Name: name, Type: mtype, Recipe: recipe, Time: time, Speed: speed}

			machines = append(machines, machine)
		}
		return machines, nil
	}

}

func (machineModel MachineModel) FindAll() ([]entities.Machine, error) {

	rows, err := machineModel.Db.Query("SELECT * FROM machines")

	if err != nil {
		return nil, err
	}
	machines := []entities.Machine{}
	for rows.Next() {
		var name string
		var mtype string
		var numbers [3]int
		var ingredients [3]string
		var time float32
		var speed float32
		err := rows.Scan(&name, &mtype, &numbers[0], &ingredients[0], &numbers[1], &ingredients[1], &numbers[2], &ingredients[2], &time, &speed)
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
		machine := entities.Machine{Name: name, Type: mtype, Recipe: recipe, Time: time, Speed: speed}

		machines = append(machines, machine)

	}
	return machines, nil

}
