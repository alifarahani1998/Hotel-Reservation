package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/alifarahani1998/bookings/controllers/models"
)

type AppConfig struct {
	UsedCache     bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
