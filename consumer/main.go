package main

import (
	"consumer/repositories"
	"events"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	db := initDatabase()
	accountRepo := repositories.NewAccountRepository(db)
}

func main() {
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// repositories.NewAccountRepository()
	// services.NewConsumerHandler()
	consumer.Consume(context.Background(), events.Topics)
}

// func main() {
// 	servers := []string{"localhost:9092"}

// 	consumer, err := sarama.NewConsumer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer consumer.Close()

// 	partitionConsumer, err := consumer.ConsumePartition("bondhello", 0, sarama.OffsetNewest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer partitionConsumer.Close()

// 	fmt.Println("Consumer start.")
// 	for {
// 		select {
// 		case err := <-partitionConsumer.Errors():
// 			fmt.Println(err)
// 		case msg := <-partitionConsumer.Messages():
// 			fmt.Println(string(msg.Value))
// 		}
// 	}
// }
