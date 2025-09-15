package seeders

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ClientSeeder(db *gorm.DB){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Gagal hash password:", err)
	}

	query := `
		INSERT INTO clients (client_id, client_email, client_password, client_name, currency, country, client_address, postal_code, client_phone, created_at, updated_at)
		VALUES (gen_random_uuid(), ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()),
		(gen_random_uuid(), ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()),
		(gen_random_uuid(), ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()),
		(gen_random_uuid(), ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()),
		(gen_random_uuid(), ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
		
	`

	if err := db.Exec(query,
		"clientA@gmail.com", string(hashedPassword), "clientA", "USD", "United State", "Wall Street", "55555", "123456789",
		"clientB@gmail.com", string(hashedPassword), "clientB", "IDR", "Jakarta", "Monas", "22222", "22222222222",
		"clientC@gmail.com", string(hashedPassword), "clientC", "USD", "USA", "Wall Street", "33333", "333333333333",
		"clientD@gmail.com", string(hashedPassword), "clientD", "IDR", "Jakut", "Kebon Ancol", "1111", "22222222",
		"clientE@gmail.com", string(hashedPassword), "clientE", "USD", "USA", "Wall Street", "66666", "111111111111",
		
	).Error; err != nil {
		log.Fatal("Gagal insert client seeder:", err)
	}

	log.Println("Seeder: CLient berhasil ditambahkan (RAW SQL)")
}