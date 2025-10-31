package jobs

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"time"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"gorm.io/gorm"
)

func SendEmail(to, subject, htmlBody string) error {
	auth := smtp.PlainAuth("", config.Env.FromEmail, config.Env.FromEmailPassword, config.Env.FromEmailSMTP)

	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", config.Env.FromEmail, to, subject, htmlBody)

	return smtp.SendMail(config.Env.SMTPAddr, auth, config.Env.FromEmail, []string{to}, []byte(msg))
}

func createEmailTemplate(clientName, invoiceNumber string, dueDate time.Time, total, amountPaid float64, daysLeft int) string {
	remainingAmount := total - amountPaid
	
	var html strings.Builder
	
	html.WriteString(
		`<!DOCTYPE html>
		<html lang="id">
		<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Invoice Reminder</title>
		<style>
			body {
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				background-color: #f5f5f5;
				margin: 0;
				padding: 20px;
			}

			.container {
				max-width: 600px;
				margin: 0 auto;
				background-color: white;
				border-radius: 10px;
				box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
				overflow: hidden;
			}

			.header {
				background: #019bd3;
				color: white;
				padding: 30px;
				text-align: center;
			}

			.header h1 {
				margin: 0;
				font-size: 24px;
				font-weight: 600;
			}

			.content {
				padding: 40px 30px;
			}

			.greeting {
				font-size: 18px;
				color: #333;
				margin-bottom: 20px;
			}

			.invoice-info {
				background-color: #f8f9fa;
				border-left: 4px solid #019bd3;
				padding: 20px;
				margin: 25px 0;
				border-radius: 5px;
			}

			.invoice-row {
				display: flex;
				justify-content: space-between;
				margin-bottom: 10px;
				padding: 5px 0;
			}

			.invoice-row:last-child {
				margin-bottom: 0;
			}

			.label {
				font-weight: 600;
				color: #555;
			}

			.value {
				color: #333;
			}

			.urgent {
				background-color: #fff3cd;
				border: 1px solid #ffeaa7;
				padding: 15px;
				border-radius: 5px;
				margin: 20px 0;
				text-align: center;
			}

			.urgent.critical {
				background-color: #f8d7da;
				border-color: #f5c6cb;
			}

			.footer {
				background-color: #f8f9fa;
				padding: 25px 30px;
				text-align: center;
				color: #666;
				font-size: 14px;
				border-top: 1px solid #e9ecef;
			}

			.highlight {
				color: #e74c3c;
				font-weight: bold;
			}

			.amount {
				font-weight: bold;
				color: #2c3e50;
			}
    	</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>Pembayaran Invoice</h1>
			</div>
        
        	<div class="content">
            	<div class="greeting">
					Halo 
					<strong>`) 
						html.WriteString(clientName)
						html.WriteString(`
					</strong>,
				</div>
            	<p>Kami ingin mengingatkan Anda mengenai invoice yang akan segera jatuh tempo:</p>
            
				<div class="invoice-info">
					<div class="invoice-row">
						<span class="label">Nomor Invoice:</span>
						<span class="value"><strong>`)
							html.WriteString(invoiceNumber)
							html.WriteString(`</strong></span>
					</div>
                	<div class="invoice-row">
						<span class="label">Tanggal Jatuh Tempo:</span>
						<span class="value"><strong>`)
							html.WriteString(dueDate.Format("02 January 2006"))
							html.WriteString(`</strong></span>
						</div>
                	<div class="invoice-row">
						<span class="label">Total Invoice:</span>
						<span class="value amount">Rp `)
							html.WriteString(fmt.Sprintf("%.2f", total))
							html.WriteString(`</span>
                	</div>
                	<div class="invoice-row">
						<span class="label">Sudah Dibayar:</span>
						<span class="value">Rp `)
							html.WriteString(fmt.Sprintf("%.2f", amountPaid))
							html.WriteString(`</span>
                		</div>
                	<div class="invoice-row">
						<span class="label">Sisa Pembayaran:</span>
						<span class="value amount highlight">Rp `)
						html.WriteString(fmt.Sprintf("%.2f", remainingAmount))
						html.WriteString(`</span>
					</div>
            	</div>
            `)
		
			html.WriteString(`
            <p>Mohon segera lakukan pembayaran untuk menghindari keterlambatan. Jika Anda sudah melakukan pembayaran, mohon abaikan email ini.</p>
        
        </div>
        
        <div class="footer">
            Email ini dikirim secara otomatis, mohon tidak membalas email ini.<br>
            Untuk pertanyaan, silakan hubungi customer service kami.</p>
        </div>
    </div>
</body>
</html>`)

	return html.String()
}

func EmailCron(db *gorm.DB) {

	var invoices []models.Invoice
	var client models.Client
	
	// Ambil invoice yang belum lunas (unpaid dan partial_paid)
	db.Where("payment_status IN ?", []string{"unpaid", "partial_paid"}).Find(&invoices)
	
	now := time.Now()
	for _, inv := range invoices {
		// Hitung sisa hari sampai due date
		daysLeft := int(inv.DueDate.Sub(now).Hours() / 24)
		fmt.Println("Testing CRon")
		// Ambil data client
		if err := db.Where("client_id = ?", inv.ClientID).First(&client).Error; err != nil {
			log.Printf("Client tidak ditemukan untuk invoice %s: %v", inv.InvoiceID, err)
			continue
		}


		// if true {
		if daysLeft == 5 || daysLeft == 3 || daysLeft == 1 {
			subject := fmt.Sprintf("Pengingat Pembayaran Invoice %s - %d Hari Lagi", 
				inv.InvoiceNumber, daysLeft)
			
			htmlBody := createEmailTemplate(
				client.ClientName,
				inv.InvoiceNumber,
				inv.DueDate,
				inv.Total,
				inv.AmountPaid,
				daysLeft,
			)

	
			if err := SendEmail("tio.michael2004@gmail.com", subject, htmlBody); err != nil {
				log.Printf("❌ Gagal kirim email ke %s untuk invoice %s: %v", 
					client.ClientEmail, inv.InvoiceNumber, err)
			} else {
				log.Printf("✅ Email berhasil dikirim ke %s untuk invoice %s (%d hari tersisa)", 
					client.ClientEmail, inv.InvoiceNumber, daysLeft)
			}
		}
	}
}