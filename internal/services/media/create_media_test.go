package media

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateMedia(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании mock соединения: %v", err)
	}
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("ошибка при открытии GORM подключения: %v", err)
	}

	// Генерируем тестовые данные
	testFile := []byte{1, 2, 3, 4, 5}
	testContentType := "image/jpeg"
	expectedID := generateUniqueId()

	// Настраиваем ожидания для SQL запроса
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO media \([^)]+\) VALUES \(\testFile,\testContentType,\expectedID,\)`).
		WillReturnResult(sqlmock.NewResult(expectedID, 1))
	mock.ExpectCommit()

	// Выполняем тестируемый метод
	id, err := CreateMedia(testFile, testContentType, gormDB)

	// Проверяем результаты
	if err != nil {
		t.Errorf("неожиданная ошибка: %v", err)
	}

	if id != expectedID {
		t.Errorf("ожидаемый ID = %d, полученный ID = %d", expectedID, id)
	}

	// Проверяем выполнение всех ожиданий
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все ожидания были выполнены: %v", err)
	}
}

func TestCreateMedia_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании mock соединения: %v", err)
	}
	defer db.Close()

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("ошибка при открытии GORM подключения: %v", err)
	}

	// Настраиваем ожидания для ошибочного сценария
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO media \([^)]+\)`).
		WillReturnError(sql.ErrConnDone)
	mock.ExpectRollback()

	// Выполняем тестируемый метод
	_, err = CreateMedia([]byte{1, 2, 3}, "image/jpeg", gormDB)

	// Проверяем, что ошибка не nil
	if err == nil {
		t.Errorf("ожидалась ошибка, но она не произошла")
	}

	// Проверяем выполнение всех ожиданий
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все ожидания были выполнены: %v", err)
	}
}
