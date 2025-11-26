package seeders

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func PaymentSeeder(db *gorm.DB) {
	type InvoiceData struct {
		InvoiceID     string
		InvoiceNumber string
		Total         float64
		AmountPaid    float64
		ClientID      string
		DueDate       time.Time
	}

	var invoices []InvoiceData
	if err := db.Raw(`
		SELECT invoice_id, invoice_number, total, amount_paid, client_id, due_date 
		FROM invoices 
		WHERE payment_status IN ('partial', 'paid')
		LIMIT 10
	`).Scan(&invoices).Error; err != nil {
		log.Fatal("Gagal mengambil invoices:", err)
	}

	if len(invoices) == 0 {
		log.Println("⚠️ Tidak ada invoice dengan status partial/paid. Jalankan InvoiceSeeder terlebih dahulu.")
		return
	}

	now := time.Now()

	for idx, inv := range invoices {

		remainingAmount := inv.Total - inv.AmountPaid

		if remainingAmount > 0 {
			paymentDate := now.AddDate(0, 0, -(idx % 5))
			proofOfTransfer := fmt.Sprintf("proof-%d-%d.pdf", idx+1, time.Now().Unix())

			paymentQuery := `
				INSERT INTO payments (
				payment_id, payment_date, amount_paid, proof_of_transfer, invoice_id, invoice_number,
				created_at, updated_at)
				VALUES (gen_random_uuid(), ?, ?, ?, ?, ?, NOW(), NOW())
			`

			if err := db.Exec(paymentQuery,
				paymentDate, remainingAmount, proofOfTransfer, inv.InvoiceID, inv.InvoiceNumber,
			).Error; err != nil {
				log.Printf("⚠️ Gagal insert payment untuk invoice %s: %v", inv.InvoiceID, err)
				continue
			}
		}
	}

	log.Println("✅ Seeder: Payment berhasil ditambahkan")
}
