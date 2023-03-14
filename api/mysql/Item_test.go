package mysql

import (
	"api/entities"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestItemFindAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	itemList := []entities.Item{{Id: 0, Name: "Electronic circuit"}, {Id: 1, Name: "Copper cable"}, {Id: 2, Name: "Iron plate"}}
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(itemList[0].Id, itemList[0].Name).
		AddRow(itemList[1].Id, itemList[1].Name).
		AddRow(itemList[2].Id, itemList[2].Name)

	mock.ExpectQuery("SELECT id, name FROM items  ORDER BY name ASC").
		WillReturnRows(rows)

	var itemList_resp []entities.Item
	if itemList_resp, err = model.FindAllItems(); err != nil {
		t.Errorf("error was not expected while getting all items: %s", err)
	}
	assert.Equal(t, itemList, itemList_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectQuery("SELECT id, name FROM items  ORDER BY name ASC").
		WillReturnError(fmt.Errorf("some error"))

	if itemList_resp, err = model.FindAllItems(); err == nil {
		t.Errorf("error was not expected while getting all items: %s", err)
	}
	assert.Equal(t, []entities.Item{}, itemList_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestItemFind(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	item := entities.Item{Id: 0, Name: "Advanced circuit", Time: 6, Result: 1, MachineType: "Assembling"}
	itemRows := sqlmock.NewRows([]string{"id", "name", "time", "result", "machineType"}).
		AddRow(item.Id, item.Name, item.Time, item.Result, item.MachineType)

	recipe := []entities.Ingredient{
		{Id: 0, Number: 4, Item: "Copper cable"},
		{Id: 1, Number: 2, Item: "Electronic circuit"},
		{Id: 2, Number: 2, Item: "Plastic bar"},
	}
	recipeRows := sqlmock.NewRows([]string{"ingrement_id", "number", "ingredient"}).
		AddRow(recipe[0].Id, recipe[0].Number, recipe[0].Item).
		AddRow(recipe[1].Id, recipe[1].Number, recipe[1].Item).
		AddRow(recipe[2].Id, recipe[2].Number, recipe[2].Item)
	item.Recipe = recipe

	mock.ExpectQuery("SELECT * FROM items WHERE id=?").
		WithArgs(item.Id).
		WillReturnRows(itemRows)

	mock.ExpectQuery("SELECT recipes.id, recipes.number, items.name FROM recipes INNER JOIN items ON recipes.ingredientId=items.id WHERE recipes.itemId=?").
		WithArgs(item.Id).
		WillReturnRows(recipeRows)

	var item_resp entities.Item
	if item_resp, err = model.FindItem(item.Id); err != nil {
		t.Errorf("error was not expected while getting item: %s", err)
	}
	assert.Equal(t, item, item_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectQuery("SELECT * FROM items WHERE id=?").
		WithArgs(item.Id).
		WillReturnError(fmt.Errorf("some error"))

	if item_resp, err = model.FindItem(item.Id); err == nil {
		t.Errorf("error was expected while getting item: %s", err)
	}
	assert.Equal(t, entities.Item{}, item_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestItemDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()
	res := sqlmock.NewResult(0, 1)
	id := 1

	mock.ExpectExec("DELETE FROM items WHERE id=?").
		WithArgs(id).
		WillReturnResult(res)

	if _, err = model.DeleteItem(id); err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("DELETE FROM items WHERE id=?").
		WithArgs(id).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = model.DeleteItem(id); err == nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestItemCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	item := entities.Item{Id: 0, Name: "Advanced circuit", Time: 6, Result: 1, MachineType: "Assembling"}
	res := sqlmock.NewResult(0, 1)

	mock.ExpectExec("INSERT INTO items(name, time, result, machineType) VALUES (?,?,?,?)").
		WithArgs(item.Name, item.Time, item.Result, item.MachineType).
		WillReturnResult(res)

	recipe := []entities.Ingredient{
		{Id: 0, Number: 4, Item: "Copper cable"},
		{Id: 1, Number: 2, Item: "Electronic circuit"},
		{Id: 2, Number: 2, Item: "Plastic bar"},
	}
	res = sqlmock.NewResult(0, 3)

	for _, ingredient := range recipe {
		mock.ExpectExec("INSERT INTO recipes(itemId, number, ingredientId) VALUES ((SELECT id FROM items WHERE name=?),?,(SELECT id FROM items WHERE name=?))").
			WithArgs(item.Name, ingredient.Number, ingredient.Item).
			WillReturnResult(res)
	}
	item.Recipe = recipe

	if _, err = model.CreateItem(&item); err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("INSERT INTO items(name, time, result, machineType) VALUES (?,?,?,?)").
		WithArgs(item.Name, item.Time, item.Result, item.MachineType).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = model.CreateItem(&item); err == nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestItemUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	item := entities.Item{Id: 0, Name: "Advanced circuit", Time: 6, Result: 1, MachineType: "Assembling"}
	res := sqlmock.NewResult(0, 1)

	mock.ExpectExec("UPDATE items SET name=?, time=?, result=?, machineType=? WHERE id=?").
		WithArgs(item.Name, item.Time, item.Result, item.MachineType, item.Id).
		WillReturnResult(res)

	recipe := []entities.Ingredient{
		{Id: -1, Number: 4, Item: "Copper cable"},
		{Id: 1, Number: 2, Item: "Electronic circuit"},
		{Id: 2, Number: -1, Item: "Plastic bar"},
	}

	mock.ExpectExec("INSERT INTO recipes(itemId, number, ingredientId) VALUES (?,?,(SELECT id FROM items WHERE name=?))").
		WithArgs(item.Id, recipe[0].Number, recipe[0].Item).
		WillReturnResult(res)

	mock.ExpectExec("UPDATE recipes SET itemId=?, number=?, ingredientId=(SELECT id FROM items WHERE name=?) WHERE id=?").
		WithArgs(item.Id, recipe[1].Number, recipe[1].Item, recipe[1].Id).
		WillReturnResult(res)

	mock.ExpectExec("DELETE FROM recipes name WHERE id=?").
		WithArgs(recipe[2].Id).
		WillReturnResult(res)
	item.Recipe = recipe

	if _, err = model.UpdateItem(item); err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("UPDATE items SET name=?, time=?, result=?, machineType=? WHERE id=?").
		WithArgs(item.Name, item.Time, item.Result, item.MachineType, item.Id).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = model.UpdateItem(item); err == nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
