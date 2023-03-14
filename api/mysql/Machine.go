package mysql

import (
	"api/entities"
)

func (model Model) DeleteMachine(id int) (int64, error) {

	result, err := model.Db.Exec("DELETE FROM machines WHERE id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (model Model) UpdateMachine(machine entities.Machine) (int64, error) {
	result, err := model.Db.Exec("UPDATE machines SET name=?, time=?, type=?, speed=? WHERE id=?", machine.Name, machine.Time, machine.Type, machine.Speed, machine.Id)
	if err != nil {
		return 0, err
	}

	tx, err := model.Db.Begin()
	if err != nil {
		return 0, err
	}
	for _, ingredient := range machine.Recipe {
		if ingredient.Id == -1 {
			result, err = tx.Exec("INSERT INTO recipes(itemId, number, ingredientId) VALUES ((SELECT id FROM items WHERE name=?),?,(SELECT id FROM items WHERE name=?))", machine.Name, ingredient.Number, ingredient.Item)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		} else if ingredient.Number != -1 {
			result, err = model.Db.Exec("UPDATE recipes SET itemId=(SELECT id FROM items WHERE name=?), number=?, ingredientId=(SELECT id FROM items WHERE name=?) WHERE id=?", machine.Name, ingredient.Number, ingredient.Item, ingredient.Id)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		} else {
			result, err = model.Db.Exec("DELETE FROM recipes name WHERE id=?", ingredient.Id)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}
	tx.Commit()
	return result.RowsAffected()

}

func (model Model) CreateMachine(machine *entities.Machine) (int64, error) {
	result, err := model.Db.Exec("INSERT INTO machines(name, time, type, speed) VALUES (?,?,?,?)", machine.Name, machine.Time, machine.Type, machine.Speed)
	if err != nil {
		return 0, err
	}

	result, err = model.Db.Exec("INSERT INTO items(name, time, result, machineType) VALUES (?,?,?,?)", machine.Name, machine.Time, 1, machine.Type)
	if err != nil {
		return 0, err
	}

	tx, err := model.Db.Begin()
	if err != nil {
		return 0, err
	}
	for _, ingredient := range machine.Recipe {
		result, err = tx.Exec("INSERT INTO recipes(itemId, number, ingredientId) VALUES ((SELECT id FROM items WHERE name=?),?,(SELECT id FROM items WHERE name=?))", machine.Name, ingredient.Number, ingredient.Item)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return result.RowsAffected()
}

func (model Model) FindMachineById(id int) (entities.Machine, error) {

	rows, err := model.Db.Query("SELECT * FROM machines WHERE id=?", id)
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
	rows, err = model.Db.Query("SELECT recipes.id, recipes.number, items.name FROM recipes INNER JOIN items ON recipes.ingredientId=items.id WHERE recipes.itemId=(SELECT id FROM items WHERE name=?)", machine.Name)
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

func (model Model) FindMachinesByType(mtype string) ([]entities.Machine, error) {

	rows, err := model.Db.Query("SELECT id, name FROM machines WHERE type=? ORDER BY name ASC", mtype)
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

func (model Model) FindAllMachines() ([]entities.Machine, error) {

	rows, err := model.Db.Query("SELECT id, name FROM machines ORDER BY name ASC")

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

func (model Model) FindAllTypes() ([]string, error) {

	rows, err := model.Db.Query("SELECT DISTINCT type FROM machines")

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
