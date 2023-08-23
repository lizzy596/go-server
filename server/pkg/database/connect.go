package main
 
import (
    "context"
    "fmt"
    "time"
 
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)
 
func connectToDB(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	defer close(client, ctx)

	if err := ping(client, ctx); err != nil {
		return err
	}

	return nil
}

func close(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connected successfully")
	return nil
}
