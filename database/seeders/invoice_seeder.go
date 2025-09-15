package seeders

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func InvoiceSeeder(db *gorm.DB){
	pricePerDelivery := 100
	taxRate := 0.02
	deliveryCount := 2000
	amount := pricePerDelivery * deliveryCount
	subTotal := amount
	tax_amount := taxRate * float64(subTotal)
	total := tax_amount + float64(subTotal)


	err := db.Transaction(func(tx *gorm.DB) error {
		invoiceQuery := `
			INSERT INTO invoices (
			invoice_id, invoice_number, issue_date, due_date, tax_rate, tax_amount, 
			sub_total, total, tax_invoice_number, amount_paid, payment_status, client_id,
			created_at,updated_at)
			VALUES (gen_random_uuid(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
			RETURNING invoice_id
		`

		var invoiceID string
		if err := tx.Raw(invoiceQuery,
			"1", time.Now(), time.Now().AddDate(0, 0, 30), 0.02, tax_amount, subTotal, total, "tax-12345", 0, "unpaid", "d444aa6a-dbc3-47f5-bbca-639ea61c834b",
		).Scan(&invoiceID).Error; err != nil {
			return err 
		}

		invoiceDetailQuery := `
			INSERT INTO invoice_details
			(invoice_detail_id, invoice_id, amount, created_at, price_per_delivery, transaction_note, updated_at, delivery_count)
			VALUES (gen_random_uuid(), ?, ?, NOW(), ?, ?, NOW(), ?)
		`

		if err := tx.Exec(invoiceDetailQuery,
			invoiceID, amount, pricePerDelivery, "Testing Client", deliveryCount,
		).Error; err != nil {
			return err 
		}

		return nil 
	})

	if err != nil {
		log.Fatal("Gagal insert invoice + invoice_detail (TRANSACTION):", err)
	}

	log.Println("Seeder: Client + Profile berhasil ditambahkan (TRANSACTION)")
}