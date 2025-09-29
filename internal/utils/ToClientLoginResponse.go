package utils

import "github.com/mdi-client-portal/client-portal-be/database/models"

func ToClientLoginResponse(c *models.Client) models.ClientLoginResponse {
	return models.ClientLoginResponse{
		ClientID:    c.ClientID,
		ClientName:  c.ClientName,
		ClientEmail: c.ClientEmail,
	}
}
