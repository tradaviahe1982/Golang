package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-golang/controllers"
)

type RestaurantRouteController struct {
	restaurantController controllers.RestaurantController
}

func NewRouteRestaurantController(restaurantController controllers.RestaurantController) RestaurantRouteController {
	return RestaurantRouteController{restaurantController}
}

func (pc *RestaurantRouteController) RestaurantRoute(rg *gin.RouterGroup) {
	router := rg.Group("restaurants")
	//
	router.POST("/", pc.restaurantController.CreateRestaurant)
	router.PUT("/:restaurantId", pc.restaurantController.UpdateRestaurant)
	router.DELETE("/:restaurantId", pc.restaurantController.DeleteRestaurant)
	router.GET("/", pc.restaurantController.FindRestaurant)
	router.GET("/:restaurantId", pc.restaurantController.FindRestaurantById)
	//
	router.POST("/:restaurantId/menu", pc.restaurantController.CreateMenu)
	router.PUT("/:restaurantId/menu/:menuId", pc.restaurantController.UpdateMenuAtId)
	/*router.DELETE("/:restaurantId/menu/:menuId", pc.restaurantController.DeleteMenuAtId)
	router.GET("/:restaurantId/menu/:menuId", pc.restaurantController.FindMenuAtMenuId)
	//
	router.POST("/:restaurantId/menu/:menuId/menuitem", pc.restaurantController.CreateMenuItem)
	router.PUT("/:restaurantId/menu/:menuId/menuitem/:menuitemId", pc.restaurantController.UpdateMenuItemAtId)
	router.DELETE("/:restaurantId/menu/:menuId/menuitem/:menuitemId", pc.restaurantController.DeleteMenuItemAtId)
	router.GET("/:restaurantId/menu/:menuId/menuitem/:menuitemId", pc.restaurantController.FindMenuItemAtMenuId)
	//*/
}
