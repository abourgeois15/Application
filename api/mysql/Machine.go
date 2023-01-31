package mysqloperations

import (
	"api/entities"
	"database/sql"
	"fmt"
)

type MachineModel struct {
	Db *sql.DB
}

func (machineModel MachineModel) CreateTable() (int64, error) {
	result, err := machineModel.Db.Exec("CREATE TABLE `machines` (`id` int NOT NULL AUTO_INCREMENT,`name` varchar(50) NOT NULL,`type` varchar(50) NOT NULL,`number1` int DEFAULT '0',`ingredient1` varchar(50) DEFAULT '',`number2` int DEFAULT '0',`ingredient2` varchar(50) DEFAULT '',`number3` int DEFAULT '0',`ingredient3` varchar(50) DEFAULT '',`time` float DEFAULT '0',`speed` float NOT NULL,PRIMARY KEY (`id`))")
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (machineModel MachineModel) DeleteTable() (int64, error) {
	result, err := machineModel.Db.Exec("DROP TABLE `machines`")
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
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
	result, err := machineModel.Db.Exec("UPDATE machines SET name=?, time=?, number1=?, ingredient1=?, number2=?, ingredient2=?, number3=?, ingredient3=?, type=?, speed=? WHERE id=?", machine.Name, machine.Time, machine.Recipe[0].Number, machine.Recipe[0].Item, machine.Recipe[1].Number, machine.Recipe[1].Item, machine.Recipe[2].Number, machine.Recipe[2].Item, machine.Type, machine.Speed, machine.Id)
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

	rows, err := machineModel.Db.Query("SELECT * FROM machines WHERE name=?", name)
	if err != nil {
		return entities.Machine{}, err
	} else {
		machine := entities.Machine{}
		for rows.Next() {
			var id int
			var name string
			var mtype string
			var numbers [3]int
			var ingredients [3]string
			var time float32
			var speed float32
			err := rows.Scan(&id, &name, &mtype, &numbers[0], &ingredients[0], &numbers[1], &ingredients[1], &numbers[2], &ingredients[2], &time, &speed)
			if err != nil {
				return entities.Machine{}, err
			}
			recipe := [3]entities.Ingredient{}
			for i, number := range numbers {
				recipe[i] = entities.Ingredient{Number: number, Item: ingredients[i]}
			}
			machine = entities.Machine{Id: id, Name: name, Type: mtype, Recipe: recipe, Time: time, Speed: speed}
		}
		return machine, nil
	}
}

func (machineModel MachineModel) FindType(mtype string) ([]string, error) {

	rows, err := machineModel.Db.Query("SELECT name FROM machines WHERE type=?", mtype)
	if err != nil {
		return nil, err
	} else {
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
}

func (machineModel MachineModel) FindAll() ([]string, error) {

	rows, err := machineModel.Db.Query("SELECT name FROM machines")

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

func (machineModel MachineModel) FindAllTypes() ([]string, error) {

	rows, err := machineModel.Db.Query("SELECT DISTINCT type FROM machines")

	if err != nil {
		return nil, err
	}
	types := []string{}
	for rows.Next() {
		var mtype string
		err := rows.Scan(&mtype)
		if err != nil {
			return nil, err
		}
		types = append(types, mtype)

	}
	return types, nil

}
