package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc{
	return func(c *gin.Context){

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listening order item"})
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
	}
}

func GetOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = contet.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("order_id")
		var food models.Order 

		err := foodCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()
		if err!= nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching the orders"})
		}
		c.JSON(http.StatusOK, order)
	}	
}

func CreateOrder() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func UpdateOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var table models.Table
		var order models.Order

		var updateObj primitive.D

		orderId := c.Param("order_id")
		if err := c.BindJSON(&order); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return 
		}

		if order.Table_id!= nil{
			err: menuCollection.FindOne(ctx, bson.M{"table_id":food.Table_id}):Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				return 
			}
			updateObj = append(updateObj, bson.E{"menu", order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"update_at", food.Updated_at})

		upsert : true

		filter := bson.M{"order_id":orderId}
		opt := options.UpdateOptions{
			Upsert:&upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$st", updateObj},
			},
			&opt,
		)

		if err != nil{
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}