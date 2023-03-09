package mysqloperations

import (
	"api/entities"
	"database/sql"
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

func (machineModel MachineModel) Delete(id int) (int64, error) {

	result, err := machineModel.Db.Exec("DELETE FROM machines WHERE id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (machineModel MachineModel) Update(machine entities.Machine) (int64, error) {
	result, err := machineModel.Db.Exec("UPDATE machines SET name=?, time=?, type=?, speed=? WHERE id=?", machine.Name, machine.Time, machine.Type, machine.Speed, machine.Id)
	if err != nil {
		return 0, err
	}
	for _, ingredient := range machine.Recipe {
		if ingredient.Id == -1 {
			result, err = machineModel.Db.Exec("INSERT INTO recipes(itemId, number, ingredientId) VALUES ((SELECT id FROM items WHERE name=?),?,(SELECT id FROM items WHERE name=?))", machine.Name, ingredient.Number, ingredient.Item)
			if err != nil {
				return 0, err
			}
		} else if ingredient.Number != -1 {
			result, err = machineModel.Db.Exec("UPDATE recipes SET itemId=(SELECT id FROM items WHERE name=?), number=?, ingredientId=(SELECT id FROM items WHERE name=?) WHERE id=?", machine.Name, ingredient.Number, ingredient.Item, ingredient.Id)
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
	result, err := machineModel.Db.Exec("INSERT INTO machines(name, time, type, speed) VALUES (?,?,?,?)", machine.Name, machine.Time, machine.Type, machine.Speed)
	if err != nil {
		return 0, err
	}
	result, err = machineModel.Db.Exec("INSERT INTO items(name, time, result, machineType) VALUES (?,?,?,?)", machine.Name, machine.Time, 1, machine.Type)
	if err != nil {
		return 0, err
	}
	for _, ingredient := range machine.Recipe {
		result, err = machineModel.Db.Exec("INSERT INTO recipes(itemId, number, ingredientId) VALUES ((SELECT id FROM items WHERE name=?),?,(SELECT id FROM items WHERE name=?))", machine.Name, ingredient.Number, ingredient.Item)
		if err != nil {
			return 0, err
		}
	}
	return result.RowsAffected()
}

func (machineModel MachineModel) FindId(id int) (entities.Machine, error) {

	rows, err := machineModel.Db.Query("SELECT * FROM machines WHERE id=?", id)
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
	rows, err = machineModel.Db.Query("SELECT recipes.id, recipes.number, items.name FROM recipes INNER JOIN items ON recipes.ingredientId=items.id WHERE recipes.itemId=(SELECT id FROM items WHERE name=?)", machine.Name)
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

func (machineModel MachineModel) FindType(mtype string) ([]entities.Machine, error) {

	rows, err := machineModel.Db.Query("SELECT id, name FROM machines WHERE type=? ORDER BY name ASC", mtype)
	if err != nil {
		return []entities.Machine{}, err
	} else {
		machines := []entities.Machine{}
		for rows.Next() {
			var machine entities.Machine
			err := rows.Scan(&machine.Id, &machine.Name)
			if err != nil {
				return []entities.Machine{}, err
			}
			machines = append(machines, machine)
		}
		return machines, nil
	}
}

func (machineModel MachineModel) FindAll() ([]entities.Machine, error) {

	rows, err := machineModel.Db.Query("SELECT id, name FROM machines ORDER BY name ASC")

	if err != nil {
		return []entities.Machine{}, err
	}
	machines := []entities.Machine{}
	for rows.Next() {
		var machine entities.Machine
		err := rows.Scan(&machine.Id, &machine.Name)
		if err != nil {
			return []entities.Machine{}, err
		}
		machines = append(machines, machine)

	}
	return machines, nil

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
