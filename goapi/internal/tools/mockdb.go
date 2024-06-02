package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"marco": {
		AuthToken: "123ABC",
		Username:  "marco",
	},
	"martina": {
		AuthToken: "456DEF",
		Username:  "martina",
	},
	"massimo": {
		AuthToken: "789GHI",
		Username:  "massimo",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"marco": {
		Coins:    100,
		Username: "marco",
	},
	"martina": {
		Coins:    200,
		Username: "martina",
	},
	"massimo": {
		Coins:    300,
		Username: "massimo",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetUpDatabase() error {
	return nil
}
