package seeders

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func InvoiceSeeder(db *gorm.DB) {
	var clientIDs []string
	if err := db.Raw("SELECT client_id FROM clients LIMIT 5").Scan(&clientIDs).Error; err != nil {
		log.Fatal("Gagal mengambil client IDs:", err)
	}

	if len(clientIDs) == 0 {
		log.Println("⚠️ Tidak ada client. Jalankan ClientSeeder terlebih dahulu.")
		return
	}

	now := time.Now()
	invoiceNumber := 1001

	for idx, clientID := range clientIDs {
		for i := 0; i < 3; i++ {
			invoiceNum := fmt.Sprintf("INV-2024-%04d", invoiceNumber)
			taxInvNum := fmt.Sprintf("TAX-INV-2024-%04d", invoiceNumber)
			invoiceNumber++

			issueDateOffset := -(30 + i*10)
			dueDateOffset := (i + 1) * 10
			issueDate := now.AddDate(0, 0, issueDateOffset)
			dueDate := now.AddDate(0, 0, dueDateOffset)

			subTotal := 1000000 + float64(idx*100000+i*50000)
			taxAmount := subTotal * 0.10
			total := subTotal + taxAmount

			var paymentStatus string
			var amountPaid float64
			switch i {
			case 0:
				paymentStatus = "unpaid"
				amountPaid = 0
			case 1:
				paymentStatus = "partial"
				amountPaid = total / 2
			case 2:
				paymentStatus = "paid"
				amountPaid = total
			}

			err := db.Transaction(func(tx *gorm.DB) error {
				invoiceQuery := `
					INSERT INTO invoices (
					invoice_id, invoice_number, issue_date, due_date, tax_rate, tax_amount, 
					sub_total, total, tax_invoice_number, amount_paid, payment_status, client_id,
					created_at, updated_at)
					VALUES (gen_random_uuid(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
					RETURNING invoice_id
				`

				var invoiceID string
				if err := tx.Raw(invoiceQuery,
					invoiceNum, issueDate, dueDate, 0.10, taxAmount, subTotal, total, taxInvNum, amountPaid, paymentStatus, clientID,
				).Scan(&invoiceID).Error; err != nil {
					return err
				}

				invoiceDetailQuery := `
					INSERT INTO invoice_details
					(invoice_detail_id, invoice_id, amount, created_at, price_per_delivery, transaction_note, updated_at, delivery_count)
					VALUES (gen_random_uuid(), ?, ?, NOW(), ?, ?, NOW(), ?)
				`

				deliveryCount := 1000 + i*500
				pricePerDelivery := int(subTotal / float64(deliveryCount))
				transactionNote := fmt.Sprintf("Delivery %d units for client %d", deliveryCount, idx+1)

				if err := tx.Exec(invoiceDetailQuery,
					invoiceID, subTotal, pricePerDelivery, transactionNote, deliveryCount,
				).Error; err != nil {
					return err
				}

				return nil
			})

			if err != nil {
				log.Fatal("Gagal insert invoice + invoice_detail:", err)
			}
		}
	}

	log.Println("✅ Seeder: Invoice + Invoice Details berhasil ditambahkan")
}
