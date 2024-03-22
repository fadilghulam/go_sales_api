package controllers

import (
	"encoding/json"
	"fmt"
	db "go_sales_api/config"
	"go_sales_api/helpers"
	"go_sales_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {

	bodyBytes := c.Body()

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(bodyBytes), &data); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	datas, err := helpers.InsertDataDynamic(c.UserContext(), data)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    datas,
	})
}

// func CreateCashier2(c *fiber.Ctx) error {
// 	var data map[string]string
// 	err := c.BodyParser(&data)
// 	if err != nil {
// 		log.Fatalf("registeration error in post request %v", err)
// 	}

// 	if data["name"] == "" || data["passcode"] == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Cashier Name is required",
// 			"error":   map[string]interface{}{},
// 		})
// 	}
// 	//passCode := strconv.Itoa(rand.Intn(1000000))
// 	//fmt.Println("passCode:::", passCode
// 	cashier := models.Cashier{
// 		Name:     data["name"],
// 		Passcode: data["passcode"],
// 	}
// 	db.DB.Create(&cashier)

// 	return c.Status(200).JSON(fiber.Map{
// 		"success": true,
// 		"message": "Success",
// 		"data":    cashier,
// 	})
// }

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")

	var cashier models.Cashier
	db.DB.Select("id ,name").Where("id=?", cashierId).First(&cashier)
	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashierData,
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier
	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}
	db.DB.Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
	})
}

func UpdateCashier(c *fiber.Ctx) error {

	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Find(&cashier, "id = ?", cashierId)
	if cashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}

	var updateCashierData models.Cashier
	c.BodyParser(&updateCashierData)
	if updateCashierData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
			"error":   map[string]interface{}{},
		})
	}

	cashier.Name = updateCashierData.Name
	db.DB.Save(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashier,
	})

}

// Cashiers struct with two values
type Cashiers struct {
	Id   uint   `json:"cashierId"`
	Name string `json:"name"`
}

func CashiersList(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	query := "SELECT id,name FROM cashiers"

	if limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}

	if skip != 0 {
		query += fmt.Sprintf(" OFFSET %d", skip)
	}

	rows, err := db.DB.Raw(query).Rows()
	if err != nil {
		return err
	}

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	defer rows.Close()

	results, err := helpers.SaveRowToDynamicStruct(rows, columns)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Error",
			"error":   err.Error(),
		})
	}

	metaMap := map[string]interface{}{
		"total": len(results),
		"limit": limit,
		"skip":  skip,
	}
	cashiersData := map[string]interface{}{
		"cashiers": results,
		"meta":     metaMap,
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    cashiersData,
	})

}

// func CashiersList2(c *fiber.Ctx) error {
// 	limit, _ := strconv.Atoi(c.Query("limit"))
// 	skip, _ := strconv.Atoi(c.Query("skip"))

// 	// fmt.Println("limit : ", limit)

// 	var count int64
// 	var cashier []Cashiers

// 	query := "SELECT * FROM cashiers"

// 	if limit != 0 {
// 		query += fmt.Sprintf(" LIMIT %d", limit)
// 	}

// 	db.DB.Raw(query).Find(&cashier).Count(&count)

// 	metaMap := map[string]interface{}{
// 		"total": count,
// 		"limit": limit,
// 		"skip":  skip,
// 	}
// 	cashiersData := map[string]interface{}{
// 		"cashiers": cashier,
// 		"meta":     metaMap,
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"success": true,
// 		"Message": "Success",
// 		"data":    cashiersData,
// 	})
// }
