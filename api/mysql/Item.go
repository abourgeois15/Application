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

	result, err := itemModel.Db.Exec("UPDATE items SET name=?, time=?, result=?, machineType=? WHERE id=?", item.Name, item.Time, item.Result, item.MachineType, item.Id)
	fmt.Println(item)
	if err != nil {
		return 0, err
	}
	for _, ingredient := range item.Recipe {
		if ingredient.Id == -1 {
			_, err = itemModel.Db.Exec("INSERT INTO recipes(item, number, ingredient) VALUES (?,?,?)", item.Name, ingredient.Number, ingredient.Item)
			if err != nil {
				return 0, err
			}
		} else if ingredient.Number != -1 {
			_, err = itemModel.Db.Exec("UPDATE recipes SET item=?, number=?, ingredient=? WHERE id=?", item.Name, ingredient.Number, ingredient.Item, ingredient.Id)
			if err != nil {
				return 0, err
			}
		} else {
			_, err = itemModel.Db.Exec("DELETE FROM recipes name WHERE id=?", ingredient.Id)
			if err != nil {
				return 0, err
			}
		}
	}
	return result.RowsAffected()

}

func (itemModel ItemModel) Create(item *entities.Item) (int64, error) {

	result, err := itemModel.Db.Exec("INSERT INTO items(name, time, result, machineType) VALUES (?,?,?,?)", item.Name, item.Time, item.Result, item.MachineType)
	if err != nil {
		return 0, err
	}
	for _, ingredient := range item.Recipe {
		_, err = itemModel.Db.Exec("INSERT INTO recipes(item, number, ingredient) VALUES (?,?,?)", item.Name, ingredient.Number, ingredient.Item)
		if err != nil {
			return 0, err
		}
	}
	return result.RowsAffected()
}

func (itemModel ItemModel) Find(name string) (entities.Item, error) {

	rows, err := itemModel.Db.Query("SELECT * FROM items WHERE name=?", name)
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
	rows, err = itemModel.Db.Query("SELECT id, number, ingredient FROM recipes WHERE item=?", name)
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

func (itemModel ItemModel) FindAll() ([]string, error) {

	rows, err := itemModel.Db.Query("SELECT name FROM items  ORDER BY name ASC")

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
