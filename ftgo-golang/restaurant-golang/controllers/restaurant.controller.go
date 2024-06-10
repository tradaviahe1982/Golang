package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"restaurant-golang/models"
	"strconv"
	"strings"
)

type RestaurantController struct {
	DB *gorm.DB
}

func NewRestaurantController(DB *gorm.DB) RestaurantController {
	return RestaurantController{DB}
}

func (pc *RestaurantController) CreateRestaurant(ctx *gin.Context) {
	var payload *models.CreateRestaurantRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newRestaurant := models.Restaurant{
		Name:    payload.Name,
		Address: payload.Address,
	}

	result := pc.DB.Create(&newRestaurant)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Cửa hàng đã tồn tại"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newRestaurant})
}

func (pc *RestaurantController) UpdateRestaurant(ctx *gin.Context) {
	restaurantId := ctx.Param("restaurantId")

	var payload *models.UpdateRestaurantRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedRestaurant models.Restaurant

	result := pc.DB.First(&updatedRestaurant, "id = ?", restaurantId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Cửa hàng không tồn tại trong hệ thống"})
		return
	}

	restaurantToUpdate := models.Restaurant{
		Name:    payload.Name,
		Address: payload.Address,
	}

	pc.DB.Model(&updatedRestaurant).Updates(restaurantToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedRestaurant})
}

func (pc *RestaurantController) FindRestaurantById(ctx *gin.Context) {
	restaurantId := ctx.Param("restaurantId")
	var restaurant models.Restaurant
	result := pc.DB.First(&restaurant, "id = ?", restaurantId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Cửa hàng không tồn tại trong hệ thống"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": restaurant})
}

func (pc *RestaurantController) FindRestaurant(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var restaurants []models.Restaurant
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&restaurants)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(restaurants), "data": restaurants})
}

func (pc *RestaurantController) DeleteRestaurant(ctx *gin.Context) {
	restaurantId := ctx.Param("restaurantId")

	result := pc.DB.Delete(&models.Restaurant{}, "id = ?", restaurantId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Cửa hàng không tồn tại trong hệ thống"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (pc *RestaurantController) CreateMenu(ctx *gin.Context) {
	//
	restaurantId := ctx.Param("restaurantId")
	var restaurant models.Restaurant
	resultRestaurant := pc.DB.First(&restaurant, "id = ?", restaurantId)
	if resultRestaurant.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Cửa hàng tồn tại trong hệ thống"})
		return
	}
	//
	var payload *models.CreateMenuRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//
	if restaurantId != payload.RestaurantId {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "RestaurantId không tồn tại"})
		return
	}
	//
	newMenu := models.Menu{
		Name:         payload.Name,
		RestaurantId: payload.RestaurantId,
	}
	//
	result := pc.DB.Create(&newMenu)
	//
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Menu không tồn tại trong hệ thống"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	//
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newMenu})
}

func (pc *RestaurantController) UpdateMenuAtId(ctx *gin.Context) {
	restaurantId := ctx.Param("restaurantId")
	menuId := ctx.Param("menuId")
	//
	var restaurant models.Restaurant
	resultRestaurant := pc.DB.First(&restaurant, "id = ?", restaurantId)
	if resultRestaurant.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Cửa hàng không tồn tại trong hệ thống"})
		return
	}
	//
	var payload *models.UpdateMenuRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedMenu models.Menu
	resultMenu := pc.DB.First(&updatedMenu, "id = ?", menuId)
	if resultMenu.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Menu không tồn tại trong hệ thống"})
		return
	}
	//
	if restaurantId != payload.RestaurantId {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "RestaurantId không tồn tại"})
		return
	}
	//
	if menuId != updatedMenu.ID.String() {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "MenuId không tồn tại"})
		return
	}
	//
	menuToUpdate := models.Menu{
		Name:         payload.Name,
		RestaurantId: payload.RestaurantId,
	}
	pc.DB.Model(&updatedMenu).Updates(menuToUpdate)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": menuToUpdate})
}
