package mysql

import (
	"api/entities"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMachineFindAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	machineList := []entities.Machine{{Id: 0, Name: "Assembling machine 1"}, {Id: 1, Name: "Electric furnace"}, {Id: 2, Name: "Electric mining drill"}}
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(machineList[0].Id, machineList[0].Name).
		AddRow(machineList[1].Id, machineList[1].Name).
		AddRow(machineList[2].Id, machineList[2].Name)

	mock.ExpectQuery("SELECT id, name FROM machines ORDER BY name ASC").
		WillReturnRows(rows)

	var machineList_resp []entities.Machine
	if machineList_resp, err = model.FindAllMachines(); err != nil {
		t.Errorf("error was not expected while getting all machines: %s", err)
	}
	assert.Equal(t, machineList, machineList_resp)

	mock.ExpectQuery("SELECT id, name FROM machines ORDER BY name ASC").
		WillReturnError(fmt.Errorf("some error"))

	if machineList_resp, err = model.FindAllMachines(); err == nil {
		t.Errorf("error was not expected while getting all machines: %s", err)
	}
	assert.Equal(t, []entities.Machine{}, machineList_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMachineFindId(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	machine := entities.Machine{Id: 0, Name: "Assembling machine 1", Time: 0.5, Speed: 0.5, Type: "Assembling"}
	machineRows := sqlmock.NewRows([]string{"id", "name", "type", "time", "speed"}).
		AddRow(machine.Id, machine.Name, machine.Type, machine.Time, machine.Speed)

	recipe := []entities.Ingredient{
		{Id: 0, Number: 3, Item: "Electronic circuit"},
		{Id: 1, Number: 5, Item: "Iron gear wheel"},
		{Id: 2, Number: 9, Item: "Iron plate"},
	}
	recipeRows := sqlmock.NewRows([]string{"ingrement_id", "number", "ingredient"}).
		AddRow(recipe[0].Id, recipe[0].Number, recipe[0].Item).
		AddRow(recipe[1].Id, recipe[1].Number, recipe[1].Item).
		AddRow(recipe[2].Id, recipe[2].Number, recipe[2].Item)
	machine.Recipe = recipe

	mock.ExpectQuery("SELECT * FROM machines WHERE id=?").
		WithArgs(machine.Id).
		WillReturnRows(machineRows)

	mock.ExpectQuery("SELECT recipes.id, recipes.number, items.name FROM recipes INNER JOIN items ON recipes.ingredientId=items.id WHERE recipes.itemId=(SELECT id FROM items WHERE name=?)").
		WithArgs(machine.Name).
		WillReturnRows(recipeRows)

	var machine_resp entities.Machine
	if machine_resp, err = model.FindMachineById(machine.Id); err != nil {
		t.Errorf("error was not expected while getting machine: %s", err)
	}
	assert.Equal(t, machine, machine_resp)

	mock.ExpectQuery("SELECT * FROM machines WHERE id=?").
		WithArgs(machine.Id).
		WillReturnError(fmt.Errorf("some error"))

	if machine_resp, err = model.FindMachineById(machine.Id); err == nil {
		t.Errorf("error was expected while getting machine: %s", err)
	}
	assert.Equal(t, entities.Machine{}, machine_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMachineFindType(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	mtype := "Assembling"
	machineList := []entities.Machine{{Id: 0, Name: "Assembling machine 1"}, {Id: 1, Name: "Assembling machine 2"}}
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(machineList[0].Id, machineList[0].Name).
		AddRow(machineList[1].Id, machineList[1].Name)

	mock.ExpectQuery("SELECT id, name FROM machines WHERE type=? ORDER BY name ASC").
		WithArgs(mtype).
		WillReturnRows(rows)

	var machineList_resp []entities.Machine
	if machineList_resp, err = model.FindMachinesByType(mtype); err != nil {
		t.Errorf("error was not expected while getting machine: %s", err)
	}
	assert.Equal(t, machineList, machineList_resp)

	mock.ExpectQuery("SELECT id, name FROM machines WHERE type=? ORDER BY name ASC").
		WithArgs(mtype).
		WillReturnError(fmt.Errorf("some error"))

	if machineList_resp, err = model.FindMachinesByType(mtype); err == nil {
		t.Errorf("error was expected while getting machine: %s", err)
	}
	assert.Equal(t, []entities.Machine{}, machineList_resp)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMachineDelete(t *testing.T) {
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

	mock.ExpectExec("DELETE FROM machines WHERE id=?").
		WithArgs(id).
		WillReturnResult(res)

	if _, err = model.DeleteMachine(id); err != nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	mock.ExpectExec("DELETE FROM machines WHERE id=?").
		WithArgs(id).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = model.DeleteMachine(id); err == nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMachineCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	machine := entities.Machine{Id: 0, Name: "Assembling machine 1", Time: 0.5, Type: "Assembling", Speed: 0.5}
	res := sqlmock.NewResult(0, 1)

	mock.ExpectExec("INSERT INTO machines(name, time, type, speed) VALUES (?,?,?,?)").
		WithArgs(machine.Name, machine.Time, machine.Type, machine.Speed).
		WillReturnResult(res)

	mock.ExpectExec("INSERT INTO items(name, time, result, machineType) VALUES (?,?,?,?)").
		WithArgs(machine.Name, machine.Time, 1, machine.Type).
		WillReturnResult(res)

	recipe := []entities.Ingredient{
		{Id: 0, Number: 3, Item: "Electronic circuit"},
		{Id: 1, Number: 5, Item: "Iron gear wheel"},
		{Id: 2, Number: 9, Item: "Iron plate"},
	}
	res = sqlmock.NewResult(0, 3)

	for _, ingredient := range recipe {
		mock.ExpectExec("INSERT INTO recipes(itemId, number, ingredientId) VALUES ((SELECT id FROM items WHERE name=?),?,(SELECT id FROM items WHERE name=?))").
			WithArgs(machine.Name, ingredient.Number, ingredient.Item).
			WillReturnResult(res)
	}
	machine.Recipe = recipe

	if _, err = model.CreateMachine(&machine); err != nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("INSERT INTO machines(name, time, type, speed) VALUES (?,?,?,?)").
		WithArgs(machine.Name, machine.Time, machine.Type, machine.Speed).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = model.CreateMachine(&machine); err == nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMachineUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	model := Model{
		Db: db,
	}
	defer db.Close()

	machine := entities.Machine{Id: 0, Name: "Assembling machine 1", Time: 0.5, Type: "Assembling", Speed: 0.5}
	res := sqlmock.NewResult(0, 1)

	mock.ExpectExec("UPDATE machines SET name=?, time=?, type=?, speed=? WHERE id=?").
		WithArgs(machine.Name, machine.Time, machine.Type, machine.Speed, machine.Id).
		WillReturnResult(res)

	recipe := []entities.Ingredient{
		{Id: -1, Number: 3, Item: "Electronic circuit"},
		{Id: 1, Number: 5, Item: "Iron gear wheel"},
		{Id: 2, Number: -1, Item: "Iron plate"},
	}
	res = sqlmock.NewResult(0, 3)

	mock.ExpectExec("INSERT INTO recipes(itemId, number, ingredientId) VALUES ((SELECT id FROM items WHERE name=?),?,(SELECT id FROM items WHERE name=?))").
		WithArgs(machine.Name, recipe[0].Number, recipe[0].Item).
		WillReturnResult(res)

	mock.ExpectExec("UPDATE recipes SET itemId=(SELECT id FROM items WHERE name=?), number=?, ingredientId=(SELECT id FROM items WHERE name=?) WHERE id=?").
		WithArgs(machine.Name, recipe[1].Number, recipe[1].Item, recipe[1].Id).
		WillReturnResult(res)

	mock.ExpectExec("DELETE FROM recipes name WHERE id=?").
		WithArgs(recipe[2].Id).
		WillReturnResult(res)
	machine.Recipe = recipe

	if _, err = model.UpdateMachine(machine); err != nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("UPDATE machines SET name=?, time=?, type=?, speed=? WHERE id=?").
		WithArgs(machine.Name, machine.Time, machine.Type, machine.Speed, machine.Id).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = model.UpdateMachine(machine); err == nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
