package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// App holds the application configuration
type App struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
