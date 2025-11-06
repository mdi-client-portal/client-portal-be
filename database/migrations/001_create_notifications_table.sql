-- Create notifications table
CREATE TABLE IF NOT EXISTS notifications (
    notification_id SERIAL PRIMARY KEY,
    client_id VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES clients(client_id) ON DELETE CASCADE
);

-- Create index on client_id and created_at for better query performance
CREATE INDEX IF NOT EXISTS idx_notifications_client_created 
ON notifications(client_id, created_at DESC);

-- Create index on read status for filtering unread notifications
CREATE INDEX IF NOT EXISTS idx_notifications_read 
ON notifications(client_id, read);
