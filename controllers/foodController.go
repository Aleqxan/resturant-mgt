package controller 

import (
	"github.com/gin-gonic/gin"
	"context"
	"time"
	"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/bluesuncorp/validator.v5"
	"golang-resturant-management/database"
	"golang-resturant-management/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

var validate = validator.New()

func GetFoods() gin.HandlerFunc{
	return func(c *gin.Context){

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strvconv.Atoi(c.Query("recordPerPage"))
		if err!=nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strvconv.Atoi(c.Query("page"))
		if err != nil || page < 1{
			page = 1
		}

		startIndex := (page-1) + recordPerPage
		startIndex, err = strvconv.Atoi(c.Query("startIndex"))

		matchStage : bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id",bson.D {{"null"}}}, {"total_count", bson.D{{"$sum, 1"}}}, }}}
		projectStage := bson.D{
			{
				"$project",bson.D{
					{"_id", 0},
					{"total_count", 1}
					{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}
				}
			}
		}
	}
}

func GetFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = contet.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food 

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err!= nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return 
		}

		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err!=nil{
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil{
			msg := fmt.Sprintf("Food item was not craeted")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {
	
}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

