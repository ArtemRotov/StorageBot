package commands

// import (
// 	"github.com/sadbard/StorageBot/internal/storage"
// )

// type MockRecordDB struct{}

// func (m *MockRecordDB) All() ([]storage.Record, error) {
// 	var r []storage.Record = []storage.Record{
// 		{
// 			UserID: 0,
// 			Name:   "testName0",
// 			Login:  "testLogin0",
// 			Passw:  "testPassw0",
// 		},
// 		{
// 			UserID: 1,
// 			Name:   "testName1",
// 			Login:  "testLogin1",
// 			Passw:  "testPassw1",
// 		},
// 		{
// 			UserID: 2,
// 			Name:   "testName2",
// 			Login:  "testLogin2",
// 			Passw:  "testPassw2",
// 		},
// 	}

// 	return r, nil
// }

// func TestCalculateSalesRate(t *testing.T) {
// 	// Инициализируем заглушку.
// 	m := &MockShopDB{}
// 	// Передаём заглушку в функцию calculateSalesRate().
// 	sr := calculateSalesRate(m)

// 	// Проверяем, соответствует ли возвращаемое значение ожиданиям на основе
// 	// фальшивых входных данных.
// 	exp := "0.33"
// 	if sr != exp {
// 		t.Fatalf("got %v; expected %v", sr, exp)
// 	}
// }
