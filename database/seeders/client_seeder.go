package seeders

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ClientSeeder(db *gorm.DB) {
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
		"contact@nusantaratech.co.id", string(hashedPassword), "nusantaratech", "IDR", "Indonesia", "Jl. Gatot Subroto No.15, Jakarta Selatan", "12710", "081234567890",
		"info@majuyaja.co.id", string(hashedPassword), "majuyaja", "IDR", "Indonesia", "Jl. Asia Afrika No.20, Bandung", "40111", "085711223344",
		"admin@cahayasejahtera.co.id", string(hashedPassword), "cahayasejahtera", "IDR", "Indonesia", "Jl. Diponegoro No.8, Surabaya", "33333", "082199887766",
		"hello@sinarabadi.co.id", string(hashedPassword), "sinarabadi", "IDR", "Indonesia", "Jl. Malioboro No.33, Yogyakarta", "55213", "087865432109",
		"support@mitrabersama.co.id", string(hashedPassword), "mitrabersama", "IDR", "Indonesia", "Jl. Gatot Subroto No.99, Medan", "20217", "081322445566",
	).Error; err != nil {
		log.Fatal("Gagal insert client seeder:", err)
	}

	log.Println("Seeder: CLient berhasil ditambahkan (RAW SQL)")
}
