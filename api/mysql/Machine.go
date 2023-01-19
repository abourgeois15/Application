package mysqloperations

import (
	"api/entities"
	"database/sql"
	"fmt"
)

type MachineModel struct {
	Db *sql.DB
}

func (machineModel MachineModel) Delete(name string) (int64, error) {

	result, err := machineModel.Db.Exec("DELETE FROM machines WHERE name=?", name)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (machineModel MachineModel) Update(machine entities.Machine) (int64, error) {
	result, err := machineModel.Db.Exec("UPDATE machines SET time=?, number1=?, ingredient1=?, number2=?, ingredient2=?, number3=?, ingredient3=?, type=?, speed=? WHERE name=?", machine.Time, machine.Recipe[0].Number, machine.Recipe[0].Item, machine.Recipe[1].Number, machine.Recipe[1].Item, machine.Recipe[2].Number, machine.Recipe[2].Item, machine.Type, machine.Speed, machine.Name)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (machineModel MachineModel) Create(machine *entities.Machine) (int64, error) {
	fmt.Println(*machine)
	result, err := machineModel.Db.Exec("INSERT INTO machines(name, time, number1, ingredient1, number2, ingredient2, number3, ingredient3, type, speed) VALUES (?,?,?,?,?,?,?,?,?,?)", machine.Name, machine.Time, machine.Recipe[0].Number, machine.Recipe[0].Item, machine.Recipe[1].Number, machine.Recipe[1].Item, machine.Recipe[2].Number, machine.Recipe[2].Item, machine.Type, machine.Speed)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}

}

func (machineModel MachineModel) FindName(name string) (entities.Machine, error) {

	rows, err := machineModel.Db.Query("SELECT * FROM `machines` WHERE name=?", name)
	if err != nil {
		return entities.Machine{}, err
	} else {
		machine := entities.Machine{}
		for rows.Next() {
			var name string
			var mtype string
			var numbers [3]int
			var ingredients [3]string
			var time float32
			var speed float32
			err := rows.Scan(&name, &mtype, &numbers[0], &ingredients[0], &numbers[1], &ingredients[1], &numbers[2], &ingredients[2], &time, &speed)
			if err != nil {
				return entities.Machine{}, err
			}
			recipe := []entities.Ingredient{}
			for i, number := range numbers {
				if number != 0 {
					recipe = append(recipe, entities.Ingredient{Number: number, Item: ingredients[i]})
				} else {
					break
				}
			}
			machine = entities.Machine{Name: name, Type: mtype, Recipe: recipe, Time: time, Speed: speed}
		}
		return machine, nil
	}
}

func (machineModel MachineModel) FindType(mtype string) ([]entities.Machine, error) {

	rows, err := machineModel.Db.Query("SELECT * FROM `machines` WHERE type=?", mtype)
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
