package mysqloperations

import (
	"api/entities"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
	rows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery("SELECT name FROM items").
		WillReturnRows(rows)

	if _, err = itemModel.FindAll(); err != nil {
		t.Errorf("error was not expected while getting all items: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectQuery("SELECT name FROM items").
		WillReturnError(fmt.Errorf("some error"))

	if _, err = itemModel.FindAll(); err == nil {
		t.Errorf("error was not expected while getting all items: %s", err)
	}

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
	rows := sqlmock.NewRows([]string{"name", "time", "number1", "ingredient1", "number1", "ingredient1", "number1", "ingredient1", "result", "machineType"})
	name := "Advanced circuit"

	mock.ExpectQuery("SELECT * FROM items WHERE name=?").
		WithArgs(name).
		WillReturnRows(rows)

	if _, err = itemModel.Find(name); err != nil {
		t.Errorf("error was not expected while getting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectQuery("SELECT * FROM items WHERE name=?").
		WithArgs(name).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = itemModel.Find(name); err == nil {
		t.Errorf("error was expected while getting item: %s", err)
	}

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
	res := sqlmock.NewResult(0, 1)
	item := entities.Item{}

	mock.ExpectExec("INSERT INTO items(name, time, number1, ingredient1, number2, ingredient2, number3, ingredient3, result, machineType) VALUES (?,?,?,?,?,?,?,?,?,?)").
		WithArgs(item.Name, item.Time, item.Recipe[0].Number, item.Recipe[0].Item, item.Recipe[1].Number, item.Recipe[1].Item, item.Recipe[2].Number, item.Recipe[2].Item, item.Result, item.MachineType).
		WillReturnResult(res)

	if _, err = itemModel.Create(&item); err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("INSERT INTO items(name, time, number1, ingredient1, number2, ingredient2, number3, ingredient3, result, machineType) VALUES (?,?,?,?,?,?,?,?,?,?)").
		WithArgs(item.Name, item.Time, item.Recipe[0].Number, item.Recipe[0].Item, item.Recipe[1].Number, item.Recipe[1].Item, item.Recipe[2].Number, item.Recipe[2].Item, item.Result, item.MachineType).
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
	res := sqlmock.NewResult(0, 1)
	item := entities.Item{}

	mock.ExpectExec("UPDATE items SET time=?, number1=?, ingredient1=?, number2=?, ingredient2=?, number3=?, ingredient3=?, result=?, machineType=? WHERE name=?").
		WithArgs(item.Time, item.Recipe[0].Number, item.Recipe[0].Item, item.Recipe[1].Number, item.Recipe[1].Item, item.Recipe[2].Number, item.Recipe[2].Item, item.Result, item.MachineType, item.Name).
		WillReturnResult(res)

	if _, err = itemModel.Update(item); err != nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("UPDATE items SET time=?, number1=?, ingredient1=?, number2=?, ingredient2=?, number3=?, ingredient3=?, result=?, machineType=? WHERE name=?").
		WithArgs(item.Time, item.Recipe[0].Number, item.Recipe[0].Item, item.Recipe[1].Number, item.Recipe[1].Item, item.Recipe[2].Number, item.Recipe[2].Item, item.Result, item.MachineType, item.Name).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = itemModel.Update(item); err == nil {
		t.Errorf("error was not expected while deleting item: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
