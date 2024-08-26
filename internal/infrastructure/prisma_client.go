package infrastructure

import (
	"log"

	"github.com/BakingUp/BakingUp-Backend/prisma/db"
)


func InitializePrismaClient() *db.PrismaClient {
    client := db.NewClient()

    if err := client.Connect(); err != nil {
        log.Fatalf("Failed to connect to Prisma: %v", err)
    }

    return client
}