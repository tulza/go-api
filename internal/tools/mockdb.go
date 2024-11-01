package tools

import (
	"time"
)

type mockDB struct {}

var  mockLoginDetails = map[string]LoginDetails{
	"alex":{
		Username: "alex",
		AuthToken: "123",
	},
	"jason":{
		Username: "jason",
		AuthToken: "456",
	},
	"marine":{
		Username: "marine",
		AuthToken: "789",
	},
}

var  mockCoinDetails = map[string]CoinDetails{
	"alex":{
		Username: "alex",
		Coins: 100,
	},
	"jason":{
		Username: "jason",
		Coins: 200,
	},
	"marine":{
		Username: "marine",
		Coins: 300,
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}