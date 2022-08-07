package tests

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	"gitlab.com/nurmanhabib/go-oauth2/util"
)

func loadEnvTesting() {
	err := godotenv.Load(filepath.Join(util.RootDir(), ".env.testing"))
	if err != nil {
		log.Fatalf("no .env.testing file provided: %v", err)
	}
}
