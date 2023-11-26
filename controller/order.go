package controller

import (
	"fmt"
	"rest-api/config"
	"rest-api/model"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func ReadOrder(c echo.Context) error {
	db := config.GetDB()

	Order := []model.Order{}
	err := db.Model(&model.Order{}).Preload("Items").Find(&Order).Error

	if err != nil {
		fmt.Println(err)
		return SetErrorResponse(c, err.Error())
	}

	return SetSuccessResponse(c, "Get Data Success", Order)
}

func CreateOrder(c echo.Context) error {
	db := config.GetDB()

	Order := model.Order{}

	if err := c.Bind(&Order); err != nil {
		fmt.Println(err)
		return SetErrorResponse(c, "Format data salah")
	}

	err := db.Create(&Order).Error

	if err != nil {
		fmt.Println(err)
		return SetErrorResponse(c, err.Error())
	}

	fmt.Println("Insert Data Success")
	return SetSuccessResponse(c, "Insert Data Success", nil)
}

func UpdateOrder(c echo.Context) error {
	db := config.GetDB()

	id := c.Param("id")
	if id == "" {
		return SetErrorResponse(c, "Data ID Tidak Ditemukan")
	}

	paramID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return SetErrorResponse(c, err.Error())
	}

	Orders := model.Order{}

	if err := c.Bind(&Orders); err != nil {
		fmt.Println(err)
		return SetErrorResponse(c, "Format Data Salah")
	}

	Orders.ID = uint(paramID)

	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&Orders).Error
	if err != nil {
		fmt.Println(err)
		return SetErrorResponse(c, err.Error())
	}

	return SetSuccessResponse(c, "Update Data Success", Orders)
}

func DeleteOrder(c echo.Context) error {
	db := config.GetDB()

	id := c.Param("id")
	if id == "" {
		return SetErrorResponse(c, "Data ID Tidak Ditemukan")
	}

	paramID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return SetErrorResponse(c, err.Error())
	}

	Orders := model.Order{}

	if err := c.Bind(&Orders); err != nil {
		fmt.Println(err)
		return SetErrorResponse(c, "Format Data Salah")
	}

	Orders.ID = uint(paramID)

	tx := db.Begin()
	err = tx.Where("order_id = ?", uint(paramID)).Unscoped().Delete(&model.Item{}).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return SetErrorResponse(c, err.Error())
	}
	err = tx.Unscoped().Delete(&Orders).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return SetErrorResponse(c, err.Error())
	}

	tx.Commit()

	return SetSuccessResponse(c, "Deleta Data Success", nil)
}
