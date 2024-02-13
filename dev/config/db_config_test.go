package config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// func TestCloseDatabaseConnection(t *testing.T) {
// 	// Membuat klien MongoDB untuk digunakan dalam pengujian
// 	fakeClient := &mongo.Client{}

// 	// Memanggil fungsi CloseDatabaseConnection untuk menutup koneksi
// 	CloseDatabaseConnection(fakeClient)

// 	assert.True(t, true, "expected database connection to be closed successfully")
// }

func TestNewMongoDBClient(t *testing.T) {
	client, err := NewMongoDBClient()

	assert.NoError(t, err, "expected no error when creating MongoDB client")

	assert.NotNil(t, client, "expected MongoDB client to be not nil")

	var result bson.M
	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result)

	// Periksa apakah ping ke database berhasil
	assert.NoError(t, err, "expected no error when pinging MongoDB database")
}
