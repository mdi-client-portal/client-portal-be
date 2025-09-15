package seeders

import (
	"log"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

func UserSeeder(db *gorm.DB){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Gagal hash password:", err)
	}
	
	query := `
		INSERT INTO users (user_id, username, password, role, created_at, updated_at)
		VALUES (gen_random_uuid(), ?, ?, ?, NOW(), NOW()),
		       (gen_random_uuid(), ?, ?, ?, NOW(), NOW())
	`

	if err := db.Exec(query,
		"bejo", string(hashedPassword), "finance",
		"jobe", string(hashedPassword), "management",
	).Error; err != nil {
		log.Fatal("Gagal insert user seeder:", err)
	}

	log.Println("Seeder: User berhasil ditambahkan (RAW SQL)")
}