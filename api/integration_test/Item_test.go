package integrationtest

import (
	"api/config"
	"api/entities"
	"api/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ItemTestSuite struct {
	suite.Suite
	Router *gin.Engine
}

func TestItemTestSuite(t *testing.T) {
	suite.Run(t, new(ItemTestSuite))
}

func (s *ItemTestSuite) SetupSuite() {
	config.ConnectMySQLDBTest()
	router := gin.Default()
	router.POST("/tables", service.CreateTables)
	router.GET("/items", service.GetAllItems)
	router.GET("/item/:item_name", service.GetItemByName)
	router.DELETE("/item/:item_name", service.DeleteItem)
	router.POST("/item", service.CreateItem)
	router.PUT("/item", service.UpdateItem)
	go router.Run("localhost:8080")
	s.Router = router
}

func (s *ItemTestSuite) TestCreateUpdateGetDeleteItem() {
	var (
		item         entities.Item
		itemReceived entities.Item
		nameList     []string
		itemJson     []byte
		err          error
		req          *http.Request
		w            *httptest.ResponseRecorder
	)

	// Create tables
	req, err = http.NewRequest("POST", "/tables", nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)

	// Create item
	item = entities.Item{Id: 0, Name: "Advanced circuit", Time: 6, Result: 1, MachineType: "Assembling"}
	item.Recipe = []entities.Ingredient{
		{Id: 0, Number: 4, Item: "Copper cable"},
		{Id: 1, Number: 2, Item: "Electronic circuit"},
		{Id: 2, Number: 2, Item: "Plastic bar"},
	}
	itemJson, err = json.Marshal(item)
	s.Nil(err)
	req, err = http.NewRequest("POST", "/item", strings.NewReader(string(itemJson)))
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)

	// Get item
	req, err = http.NewRequest("GET", "/item/"+item.Name, nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Nil(json.Unmarshal(w.Body.Bytes(), &itemReceived))
	s.Equal(http.StatusOK, w.Code)
	s.Equal(item.Name, itemReceived.Name)
	s.Equal(item.Time, itemReceived.Time)
	s.Equal(item.Recipe, itemReceived.Recipe)
	s.Equal(item.Result, itemReceived.Result)
	s.Equal(item.MachineType, itemReceived.MachineType)

	// Update item
	item = entities.Item{Id: 0, Name: "Electronic circuit", Time: 0.5, Result: 1, MachineType: "Assembling"}
	item.Recipe = []entities.Ingredient{
		{Id: 0, Number: 3, Item: "Copper cable"},
		{Id: 1, Number: -1, Item: "Electronic circuit"},
		{Id: 2, Number: -1, Item: "Plastic bar"},
		{Id: -1, Number: 1, Item: "Iron ore"},
	}
	itemJson, err = json.Marshal(item)
	s.Nil(err)
	req, err = http.NewRequest("PUT", "/item", strings.NewReader(string(itemJson)))
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)

	// Get all items
	req, err = http.NewRequest("GET", "/items", nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Nil(json.Unmarshal(w.Body.Bytes(), &nameList))
	s.Equal(http.StatusOK, w.Code)

	// Get item
	req, err = http.NewRequest("GET", "/item/"+item.Name, nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Nil(json.Unmarshal(w.Body.Bytes(), &itemReceived))
	s.Equal(http.StatusOK, w.Code)
	s.Equal(item.Name, itemReceived.Name)
	s.Equal(item.Time, itemReceived.Time)
	s.Equal(item.Recipe, itemReceived.Recipe)
	s.Equal(item.Result, itemReceived.Result)
	s.Equal(item.MachineType, itemReceived.MachineType)

	// Delete item
	req, err = http.NewRequest("DELETE", "/item/"+item.Name, nil)
	s.Nil(err)
	w = httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
}
