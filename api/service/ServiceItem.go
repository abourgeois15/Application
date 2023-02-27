package service

import (
	"api/config"
	"api/entities"
	mysqloperations "api/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTables(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	itemModel := mysqloperations.ItemModel{Db: db}
	rows, err := itemModel.CreateTables()
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusCreated, rows)
	}
}

func DeleteTableItem(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	itemModel := mysqloperations.ItemModel{Db: db}
	rows, err := itemModel.DeleteTable()
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusOK, rows)
	}
}

func CreateItem(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdItem entities.Item
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err := c.BindJSON(&createdItem); err != nil {
		fmt.Println(err)
		return
	}

	itemModel := mysqloperations.ItemModel{Db: db}
	item := entities.Item{
		Name:        createdItem.Name,
		Time:        createdItem.Time,
		Recipe:      createdItem.Recipe,
		Result:      createdItem.Result,
		MachineType: createdItem.MachineType,
	}
	rows, err := itemModel.Create(&item)
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusCreated, rows)
	}
}

func UpdateItem(c *gin.Context) {
	db, _ := config.GetMySQLDB()
	var createdItem entities.Item
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err := c.BindJSON(&createdItem); err != nil {
		return
	}

	itemModel := mysqloperations.ItemModel{Db: db}
	item := entities.Item{
		Id:          createdItem.Id,
		Name:        createdItem.Name,
		Time:        createdItem.Time,
		Recipe:      createdItem.Recipe,
		Result:      createdItem.Result,
		MachineType: createdItem.MachineType,
	}
	rows, err := itemModel.Update(item)
	if err != nil {
		fmt.Println(err)
	} else {
		if rows > 0 {
			fmt.Println("done")
		}
		c.IndentedJSON(http.StatusOK, rows)
	}
}

func DeleteItem(c *gin.Context) {
	db, err := config.GetMySQLDB()
	name := c.Param("item_name")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		itemModel := mysqloperations.ItemModel{Db: db}
		rows, err := itemModel.Delete(name)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("done")
			}
		}
		c.IndentedJSON(http.StatusOK, rows)
	}

}

func GetAllItems(c *gin.Context) {
	db, err := config.GetMySQLDB()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		itemModel := mysqloperations.ItemModel{Db: db}
		names, err := itemModel.FindAll()

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, names)
	}
}

func GetItemByName(c *gin.Context) {
	db, err := config.GetMySQLDB()
	name := c.Param("item_name")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		fmt.Println(err)
	} else {
		itemModel := mysqloperations.ItemModel{Db: db}
		items, err := itemModel.Find(name)

		if err != nil {
			fmt.Println(err)
		}
		c.IndentedJSON(http.StatusOK, items)
	}
}

func GetCraftPlan(c *gin.Context) {
	db, err := config.GetMySQLDB()
	var craftPlans []entities.CraftPlan
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if err := c.BindJSON(&craftPlans); err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	for i, craftPlan := range craftPlans {
		if craftPlan.Item != "" {
			itemModel := mysqloperations.ItemModel{Db: db}
			item, err := itemModel.Find(craftPlan.Item)
			if err != nil {
				fmt.Println(err)
			}
			machineModel := mysqloperations.MachineModel{Db: db}
			machines, err := machineModel.FindType(item.MachineType)
			if err != nil {
				fmt.Println(err)
			}
			craftPlans[i].Machines = machines
			if craftPlan.Machine == "" {
				craftPlans[i].Machine = machines[0]
			}
			machine, err := machineModel.FindName(craftPlans[i].Machine)
			if err != nil {
				fmt.Println(err)
			}

			if craftPlan.ParentId != -1 {
				fmt.Println(craftPlan.ParentId)
				for _, ingredient := range craftPlans[craftPlan.ParentId].Recipe {
					if ingredient.Item == craftPlan.Item {
						craftPlans[i].Number = ingredient.Number
					}
				}
			}

			var timeMult float32
			switch craftPlan.Time {
			case "s":
				timeMult = 1
			case "min":
				timeMult = 60
			case "h":
				timeMult = 3600
			default:
				timeMult = 1
			}
			craftPlans[i].NumberMachine = (float32(craftPlans[i].Number) / timeMult * item.Time) / (machine.Speed * float32(item.Result))
			craftPlans[i].Recipe[0].Number = (float32(craftPlans[i].Number) * float32(item.Recipe[0].Number)) / float32(item.Result)
			craftPlans[i].Recipe[1].Number = (float32(craftPlans[i].Number) * float32(item.Recipe[1].Number)) / float32(item.Result)
			craftPlans[i].Recipe[2].Number = (float32(craftPlans[i].Number) * float32(item.Recipe[2].Number)) / float32(item.Result)
			craftPlans[i].Recipe[0].Item = item.Recipe[0].Item
			craftPlans[i].Recipe[1].Item = item.Recipe[1].Item
			craftPlans[i].Recipe[2].Item = item.Recipe[2].Item
		}
		fmt.Println(craftPlans)
	}
	c.IndentedJSON(http.StatusOK, craftPlans)
}
