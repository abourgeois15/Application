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
	result, err := machineModel.Db.Exec("UPDATE machines SET name=?, time=?, type=?, speed=? WHERE id=?", machine.Name, machine.Time, machine.Type, machine.Speed, machine.Id)
	fmt.Println(machine)
	if err != nil {
		return 0, err
	}
	for _, ingredient := range machine.Recipe {
		if ingredient.Id == -1 {
			result, err = machineModel.Db.Exec("INSERT INTO recipes(item, number, ingredient) VALUES (?,?,?)", machine.Name, ingredient.Number, ingredient.Item)
			if err != nil {
				return 0, err
			}
		} else if ingredient.Number != -1 {
			result, err = machineModel.Db.Exec("UPDATE recipes SET item=?, number=?, ingredient=? WHERE id=?", machine.Name, ingredient.Number, ingredient.Item, ingredient.Id)
			if err != nil {
				return 0, err
			}
		} else {
			result, err = machineModel.Db.Exec("DELETE FROM recipes name WHERE id=?", ingredient.Id)
			if err != nil {
				return 0, err
			}
		}
	}
	return result.RowsAffected()

}

func (machineModel MachineModel) Create(machine *entities.Machine) (int64, error) {
	fmt.Println(*machine)
	result, err := machineModel.Db.Exec("INSERT INTO machines(name, time, type, speed) VALUES (?,?,?,?)", machine.Name, machine.Time, machine.Type, machine.Speed)
	if err != nil {
		return 0, err
	}
	for _, ingredient := range machine.Recipe {
		result, err = machineModel.Db.Exec("INSERT INTO recipes(item, number, ingredient) VALUES (?,?,?)", machine.Name, ingredient.Number, ingredient.Item)
		if err != nil {
			return 0, err
		}
	}
	return result.RowsAffected()
}

func (machineModel MachineModel) FindName(name string) (entities.Machine, error) {

	rows, err := machineModel.Db.Query("SELECT * FROM machines WHERE name=?", name)
	if err != nil {
		return entities.Machine{}, err
	}
	var machine entities.Machine
	for rows.Next() {
		err := rows.Scan(&machine.Id, &machine.Name, &machine.Type, &machine.Time, &machine.Speed)
		if err != nil {
			return entities.Machine{}, err
		}
	}
	rows, err = machineModel.Db.Query("SELECT id, number, ingredient FROM recipes WHERE item=?", name)
	if err != nil {
		return entities.Machine{}, err
	}
	recipe := []entities.Ingredient{}
	for rows.Next() {
		var ingredient entities.Ingredient
		err := rows.Scan(&ingredient.Id, &ingredient.Number, &ingredient.Item)
		if err != nil {
			return entities.Machine{}, err
		}
		recipe = append(recipe, ingredient)
	}
	machine.Recipe = recipe
	return machine, nil
}

func (machineModel MachineModel) FindType(mtype string) ([]string, error) {

	rows, err := machineModel.Db.Query("SELECT name FROM machines WHERE type=? ORDER BY name ASC", mtype)
	if err != nil {
		return []string{}, err
	} else {
		names := []string{}
		for rows.Next() {
			var name string
			err := rows.Scan(&name)
			if err != nil {
				return []string{}, err
			}
			names = append(names, name)
		}
		return names, nil
	}
}

func (machineModel MachineModel) FindAll() ([]string, error) {

	rows, err := machineModel.Db.Query("SELECT name FROM machines ORDER BY name ASC")

	if err != nil {
		return []string{}, err
	}
	names := []string{}
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return []string{}, err
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
