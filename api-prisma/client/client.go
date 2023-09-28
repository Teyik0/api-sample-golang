package client

import "github.com/Teyik0/api-sample-golang/api-prisma/prisma/db"

var Prisma *db.PrismaClient

func Connect() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	Prisma = client

	return nil
}
