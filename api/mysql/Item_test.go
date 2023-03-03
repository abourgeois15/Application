package mysqloperations

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
	itemModel := ItemModel{
		Db: db,
	}
	defer db.Close()

	nameList := []string{"Electronic circuit", "Copper cable", "Iron plate"}
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nameList[0]).
		AddRow(nameList[1]).
		AddRow(nameList[2])

	mock.ExpectQuery("SELECT name FROM items  ORDER BY name ASC").
		WillReturnRows(rows)

	var nameList_resp []string
	if nameList_resp, err = itemModel.FindAll(); err != nil {
		t.Errorf("error was not expected while getting all items: %s", err)
	}
	assert.Equal(t, nameList, nameList_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectQuery("SELECT name FROM items  ORDER BY name ASC").
		WillReturnError(fmt.Errorf("some error"))

	if nameList_resp, err = itemModel.FindAll(); err == nil {
		t.Errorf("error was not expected while getting all items: %s", err)
	}
	assert.Equal(t, []string{}, nameList_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestItemFind(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	itemModel := ItemModel{
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

	mock.ExpectQuery("SELECT * FROM items WHERE name=?").
		WithArgs(item.Name).
		WillReturnRows(itemRows)

	mock.ExpectQuery("SELECT id, number, ingredient FROM recipes WHERE item=?").
		WithArgs(item.Name).
		WillReturnRows(recipeRows)

	var item_resp entities.Item
	if item_resp, err = itemModel.Find(item.Name); err != nil {
		t.Errorf("error was not expected while getting item: %s", err)
	}
	assert.Equal(t, item, item_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectQuery("SELECT * FROM items WHERE name=?").
		WithArgs(item.Name).
		WillReturnError(fmt.Errorf("some error"))

	if item_resp, err = itemModel.Find(item.Name); err == nil {
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
	itemModel := ItemModel{
		Db: db,
	}
	defer db.Close()
	res := sqlmock.NewResult(0, 1)
	name := "Advanced circuit"

	mock.ExpectExec("DELETE FROM items WHERE name=?").
		WithArgs(name).
		WillReturnResult(res)

	if _, err = itemModel.Delete(name); err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("DELETE FROM items WHERE name=?").
		WithArgs(name).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = itemModel.Delete(name); err == nil {
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
	itemModel := ItemModel{
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
		mock.ExpectExec("INSERT INTO recipes(item, number, ingredient) VALUES (?,?,?)").
			WithArgs(item.Name, ingredient.Number, ingredient.Item).
			WillReturnResult(res)
	}
	item.Recipe = recipe

	if _, err = itemModel.Create(&item); err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("INSERT INTO items(name, time, result, machineType) VALUES (?,?,?,?)").
		WithArgs(item.Name, item.Time, item.Result, item.MachineType).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = itemModel.Create(&item); err == nil {
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
	itemModel := ItemModel{
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

	mock.ExpectExec("INSERT INTO recipes(item, number, ingredient) VALUES (?,?,?)").
		WithArgs(item.Name, recipe[0].Number, recipe[0].Item).
		WillReturnResult(res)

	mock.ExpectExec("UPDATE recipes SET item=?, number=?, ingredient=? WHERE id=?").
		WithArgs(item.Name, recipe[1].Number, recipe[1].Item, recipe[1].Id).
		WillReturnResult(res)

	mock.ExpectExec("DELETE FROM recipes name WHERE id=?").
		WithArgs(recipe[2].Id).
		WillReturnResult(res)
	item.Recipe = recipe

	if _, err = itemModel.Update(item); err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("UPDATE items SET name=?, time=?, result=?, machineType=? WHERE id=?").
		WithArgs(item.Name, item.Time, item.Result, item.MachineType, item.Id).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = itemModel.Update(item); err == nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
