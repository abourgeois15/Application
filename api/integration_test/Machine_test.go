package integrationtest

import (
	"api/config"
	"api/entities"
	mysqloperations "api/mysql"
	"api/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type MachineTestSuite struct {
	suite.Suite
	Router *gin.Engine
}

func TestMachineTestSuite(t *testing.T) {
	suite.Run(t, new(MachineTestSuite))
}

func (s *MachineTestSuite) SetupSuite() {
	config.ConnectMySQLDBTest()
	db, err := config.GetMySQLDB()
	s.Nil(err)
	machineModel := mysqloperations.MachineModel{Db: db}
	_, err = machineModel.CreateTable()
	s.Nil(err)
	router := gin.Default()
	router.GET("/machines", service.GetAllMachines)
	router.GET("/machines/type", service.GetAllTypes)
	router.GET("/machine/name/:machine_name", service.GetMachineByName)
	router.GET("/machine/type/:machine_type", service.GetMachineByType)
	router.DELETE("/machine/:machine_name", service.DeleteMachine)
	router.POST("/machine", service.CreateMachine)
	router.PUT("/machine", service.UpdateMachine)
	go router.Run("localhost:8080")
	s.Router = router
}

func (s *MachineTestSuite) TestCreateUpdateGetDeleteMachine() {
	var (
		machine         entities.Machine
		machineReceived entities.Machine
		nameList        []string
		machineJson     []byte
		err             error
		req             *http.Request
		w               *httptest.ResponseRecorder
	)
	// Create machine
	machine = entities.Machine{
		Name:   "machine_name",
		Time:   1,
		Recipe: [3]entities.Ingredient{},
		Speed:  1,
		Type:   "Assembling",
	}
	machine.Recipe[0] = entities.Ingredient{Number: 1, Item: "ingredient"}
	machine.Recipe[1] = entities.Ingredient{Number: 1, Item: "ingredient"}
	machine.Recipe[2] = entities.Ingredient{Number: 1, Item: "ingredient"}
	machineJson, err = json.Marshal(machine)
	s.Nil(err)
	req, err = http.NewRequest("POST", "/machine", strings.NewReader(string(machineJson)))
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)

	// Update machine
	machine = entities.Machine{
		Name:  "machine_name",
		Time:  0.5,
		Speed: 2,
		Type:  "Furnace",
	}
	machine.Recipe[0] = entities.Ingredient{Number: 2, Item: "new_ingredient"}
	machine.Recipe[1] = entities.Ingredient{Number: 2, Item: "new_ingredient"}
	machine.Recipe[2] = entities.Ingredient{Number: 2, Item: "new_ingredient"}
	machineJson, err = json.Marshal(machine)
	s.Nil(err)
	req, err = http.NewRequest("PUT", "/machine", strings.NewReader(string(machineJson)))
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)

	// Get all machines
	req, err = http.NewRequest("GET", "/machines", nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Nil(json.Unmarshal(w.Body.Bytes(), &nameList))
	s.Equal(http.StatusOK, w.Code)
	s.Equal(machine.Name, nameList[0])

	// Get all types
	req, err = http.NewRequest("GET", "/machines/type", nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Nil(json.Unmarshal(w.Body.Bytes(), &nameList))
	s.Equal(http.StatusOK, w.Code)
	s.Equal(machine.Type, nameList[0])

	// Get machine by name
	req, err = http.NewRequest("GET", "/machine/name/"+machine.Name, nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Nil(json.Unmarshal(w.Body.Bytes(), &machineReceived))
	s.Equal(http.StatusOK, w.Code)
	s.Equal(machine, machineReceived)

	// Get machine by type
	req, err = http.NewRequest("GET", "/machine/type/"+machine.Type, nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Nil(json.Unmarshal(w.Body.Bytes(), &nameList))
	s.Equal(http.StatusOK, w.Code)
	s.Equal(machine.Name, nameList[0])

	// Delete machine
	req, err = http.NewRequest("DELETE", "/machine/"+machine.Name, nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}

func (s *MachineTestSuite) TearDownSuite() {
	db, err := config.GetMySQLDB()
	s.Nil(err)
	machineModel := mysqloperations.MachineModel{Db: db}
	_, err = machineModel.DeleteTable()
	s.Nil(err)
}
