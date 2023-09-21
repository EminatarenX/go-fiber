package models

import (

)

type Cliente struct {
	ID int `json:"id"`
	Nombre string `json:"nome"`
	Apellidos string `json:"email"`
	Email string `json:"telefone"`
	CreatedAt string `json:"created_at"`
	
}