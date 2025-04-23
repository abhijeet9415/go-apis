
package main
import (
	"sms-microservices/database" // Import centralized database package
	"go.mongodb.org/mongo-driver/mongo"
)


var transactionDB *mongo.Collection
// GetStoreCollection returns the "messages" collection from MongoDB
func InitTransactionDB (){
	transactionDB = database.GetCollection("transactions")
}
