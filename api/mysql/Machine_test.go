package mysqloperations

import (
	"api/entities"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMachineFindAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	machineModel := MachineModel{
		Db: db,
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery("SELECT name FROM machines").
		WillReturnRows(rows)

	if _, err = machineModel.FindAll(); err != nil {
		t.Errorf("error was not expected while getting all machines: %s", err)
	}

	mock.ExpectQuery("SELECT name FROM machines").
		WillReturnError(fmt.Errorf("some error"))

	if _, err = machineModel.FindAll(); err == nil {
		t.Errorf("error was not expected while getting all machines: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMachineFindName(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	machineModel := MachineModel{
		Db: db,
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"name", "type", "number1", "ingredient1", "number1", "ingredient1", "number1", "ingredient1", "time", "speed"})
	name := "Advanced circuit"

	mock.ExpectQuery("SELECT * FROM machines WHERE name=?").
		WithArgs(name).
		WillReturnRows(rows)

	if _, err = machineModel.FindName(name); err != nil {
		t.Errorf("error was not expected while getting machine: %s", err)
	}

	mock.ExpectQuery("SELECT * FROM machines WHERE name=?").
		WithArgs(name).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = machineModel.FindName(name); err == nil {
		t.Errorf("error was expected while getting machine: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMachineFindType(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	machineModel := MachineModel{
		Db: db,
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"name", "type", "number1", "ingredient1", "number1", "ingredient1", "number1", "ingredient1", "time", "speed"})
	mtype := "Assembling"

	mock.ExpectQuery("SELECT name FROM machines WHERE type=?").
		WithArgs(mtype).
		WillReturnRows(rows)

	if _, err = machineModel.FindType(mtype); err != nil {
		t.Errorf("error was not expected while getting machine: %s", err)
	}

	mock.ExpectQuery("SELECT name FROM machines WHERE type=?").
		WithArgs(mtype).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = machineModel.FindType(mtype); err == nil {
		t.Errorf("error was expected while getting machine: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestMachineDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	machineModel := MachineModel{
		Db: db,
	}
	defer db.Close()
	res := sqlmock.NewResult(0, 1)
	name := "Advanced circuit"

	mock.ExpectExec("DELETE FROM machines WHERE name=?").
		WithArgs(name).
		WillReturnResult(res)

	if _, err = machineModel.Delete(name); err != nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	mock.ExpectExec("DELETE FROM machines WHERE name=?").
		WithArgs(name).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = machineModel.Delete(name); err == nil {
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
	machineModel := MachineModel{
		Db: db,
	}
	defer db.Close()
	res := sqlmock.NewResult(0, 1)
	machine := entities.Machine{}

	mock.ExpectExec("INSERT INTO machines(name, time, number1, ingredient1, number2, ingredient2, number3, ingredient3, type, speed) VALUES (?,?,?,?,?,?,?,?,?,?)").
		WithArgs(machine.Name, machine.Time, machine.Recipe[0].Number, machine.Recipe[0].Item, machine.Recipe[1].Number, machine.Recipe[1].Item, machine.Recipe[2].Number, machine.Recipe[2].Item, machine.Type, machine.Speed).
		WillReturnResult(res)

	if _, err = machineModel.Create(&machine); err != nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	mock.ExpectExec("INSERT INTO machines(name, time, number1, ingredient1, number2, ingredient2, number3, ingredient3, type, speed) VALUES (?,?,?,?,?,?,?,?,?,?)").
		WithArgs(machine.Name, machine.Time, machine.Recipe[0].Number, machine.Recipe[0].Item, machine.Recipe[1].Number, machine.Recipe[1].Item, machine.Recipe[2].Number, machine.Recipe[2].Item, machine.Type, machine.Speed).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = machineModel.Create(&machine); err == nil {
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
	machineModel := MachineModel{
		Db: db,
	}
	defer db.Close()
	res := sqlmock.NewResult(0, 1)
	machine := entities.Machine{}

	mock.ExpectExec("UPDATE machines SET name=?, time=?, number1=?, ingredient1=?, number2=?, ingredient2=?, number3=?, ingredient3=?, type=?, speed=? WHERE id=?").
		WithArgs(machine.Name, machine.Time, machine.Recipe[0].Number, machine.Recipe[0].Item, machine.Recipe[1].Number, machine.Recipe[1].Item, machine.Recipe[2].Number, machine.Recipe[2].Item, machine.Type, machine.Speed, machine.Id).
		WillReturnResult(res)

	if _, err = machineModel.Update(machine); err != nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	mock.ExpectExec("UPDATE machines SET name=?, time=?, number1=?, ingredient1=?, number2=?, ingredient2=?, number3=?, ingredient3=?, type=?, speed=? WHERE id=?").
		WithArgs(machine.Name, machine.Time, machine.Recipe[0].Number, machine.Recipe[0].Item, machine.Recipe[1].Number, machine.Recipe[1].Item, machine.Recipe[2].Number, machine.Recipe[2].Item, machine.Type, machine.Speed, machine.Id).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = machineModel.Update(machine); err == nil {
		t.Errorf("error was not expected while deleting machine: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
