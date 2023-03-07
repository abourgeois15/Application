# API

The API is working at http://localhost:8080 and is using go(lang).

It uses the http web framework GIN to handle the requests from the client.

Each request is handled with a router which trigger the execution of a service function. This function will call other functions from the mysqloperations package with sql queries to the database.

## Data structures

To make the data handling more efficient, there are several struct declared in the API:

```go
type Ingredient struct {
	Id     int     `json:"id"`
	Number float32 `json:"number"`
	Item   string  `json:"item"`
}

type Item struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Recipe      []Ingredient `json:"recipe"`
	Time        float32      `json:"time"`
	Result      int          `json:"result"`
	MachineType string       `json:"machineType"`
}

type Machine struct {
	Id     int          `json:"id"`
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	Recipe []Ingredient `json:"recipe"`
	Time   float32      `json:"time"`
	Speed  float32      `json:"speed"`
}
```

## Database

The API is connected to a MySQL database server. This database is named newdockerdb with the following Tables:

*Item table*
| Field       | Type        | Null | Key | Default | Extra          |
|-------------|-------------|------|-----|---------|----------------|
| id          | int         | NO   | PRI | NULL    | auto_increment |
| name        | varchar(50) | NO   |     | NULL    |                |
| time        | float       | YES  |     | 0       |                |
| result      | int         | NO   |     | NULL    |                |
| machineType | varchar(30) | NO   |     | NULL    |                |

*Machine table*
| Field | Type        | Null | Key | Default | Extra          |
|-------|-------------|------|-----|---------|----------------|
| id    | int         | NO   | PRI | NULL    | auto_increment |
| name  | varchar(50) | NO   |     | NULL    |                |
| type  | varchar(50) | NO   |     | NULL    |                |
| time  | float       | YES  |     | 0       |                |
| speed | float       | NO   |     | NULL    |                |

Each field correspond to one element of the struct defined in the API except for the recipes, which are stored in the same table.

*Recipe table*
| Field      | Type        | Null | Key | Default | Extra          |
|------------|-------------|------|-----|---------|----------------|
| id         | int         | NO   | PRI | NULL    | auto_increment |
| item       | varchar(50) | NO   |     | NULL    |                |
| number     | int         | NO   |     | NULL    |                |
| ingredient | varchar(50) | NO   |     | NULL    |                |

The item field correspond either to an item or a machine to craft and the number and ingredient name represents the items to use to craft.
