package database

type DB struct {
	Connection interface{}
}

func FactoryConnection(tipo string)*DB{
	return nil
}
