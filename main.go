package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	dsn := "mysql://appuser:password@localhost:3306/appdb"

	cmd := exec.Command("atlas", "migrate", "apply", "--url", dsn, "--dir", "file://migrations")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running DB migrations with Atls...")
	if err := cmd.Run(); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("Migration applied successfully!")
}
