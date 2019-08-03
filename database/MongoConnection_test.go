package database

import "testing"

func TestConnectMongo(t *testing.T){
	err,_ := ConnectDB()
	if err != nil {
		t.Error("Error al hacer test con mongo",err)
	}
}

