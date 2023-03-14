package mysql

import (
	"api/entities"
)

func (model Model) DeleteItem(id int) (int64, error) {

	result, err := model.Db.Exec("DELETE FROM items WHERE id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (model Model) UpdateItem(item entities.Item) (int64, error) {

	result, err := model.Db.Exec("UPDATE items SET name=?, time=?, result=?, machineType=? WHERE id=?", item.Name, item.Time, item.Result, item.MachineType, item.Id)
	if err != nil {
		return 0, err
	}

	tx, err := model.Db.Begin()
	if err != nil {
		return 0, err
	}
	for _, ingredient := range item.Recipe {
		if ingredient.Id == -1 {
			_, err = tx.Exec("INSERT INTO recipes(itemId, number, ingredientId) VALUES (?,?,(SELECT id FROM items WHERE name=?))", item.Id, ingredient.Number, ingredient.Item)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		} else if ingredient.Number != -1 {
			_, err = tx.Exec("UPDATE recipes SET itemId=?, number=?, ingredientId=(SELECT id FROM items WHERE name=?) WHERE id=?", item.Id, ingredient.Number, ingredient.Item, ingredient.Id)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		} else {
			_, err = tx.Exec("DELETE FROM recipes name WHERE id=?", ingredient.Id)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}
	tx.Commit()
	return result.RowsAffected()
}

func (model Model) CreateItem(item *entities.Item) (int64, error) {

	result, err := model.Db.Exec("INSERT INTO items(name, time, result, machineType) VALUES (?,?,?,?)", item.Name, item.Time, item.Result, item.MachineType)
	if err != nil {
		return 0, err
	}

	tx, err := model.Db.Begin()
	if err != nil {
		return 0, err
	}
	for _, ingredient := range item.Recipe {
		_, err = tx.Exec("INSERT INTO recipes(itemId, number, ingredientId) VALUES ((SELECT id FROM items WHERE name=?),?,(SELECT id FROM items WHERE name=?))", item.Name, ingredient.Number, ingredient.Item)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return result.RowsAffected()
}

func (model Model) FindItem(id int) (entities.Item, error) {

	rows, err := model.Db.Query("SELECT * FROM items WHERE id=?", id)
	if err != nil {
		return entities.Item{}, err
	}
	var item entities.Item
	for rows.Next() {
		err := rows.Scan(&item.Id, &item.Name, &item.Time, &item.Result, &item.MachineType)
		if err != nil {
			return entities.Item{}, err
		}
	}
	rows, err = model.Db.Query("SELECT recipes.id, recipes.number, items.name FROM recipes INNER JOIN items ON recipes.ingredientId=items.id WHERE recipes.itemId=?", id)
	if err != nil {
		return entities.Item{}, err
	}
	recipe := []entities.Ingredient{}
	for rows.Next() {
		var ingredient entities.Ingredient
		err := rows.Scan(&ingredient.Id, &ingredient.Number, &ingredient.Item)
		if err != nil {
			return entities.Item{}, err
		}
		recipe = append(recipe, ingredient)
	}
	item.Recipe = recipe
	return item, nil

}

func (model Model) FindAllItems() ([]entities.Item, error) {

	rows, err := model.Db.Query("SELECT id, name FROM items  ORDER BY name ASC")

	if err != nil {
		return []entities.Item{}, err
	}
	items := []entities.Item{}
	for rows.Next() {
		var item entities.Item
		err := rows.Scan(&item.Id, &item.Name)
		if err != nil {
			return []entities.Item{}, err
		}
		items = append(items, item)
	}
	return items, nil
}
