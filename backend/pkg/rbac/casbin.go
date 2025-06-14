package rbac

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Mhirii/TaskHub/backend/pkg/config"
	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
)

func NewCasbin(cfg *config.Config) {
	DBcfg := cfg.RBAC.DB
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DBcfg.Host, DBcfg.Port, DBcfg.User, DBcfg.Password, DBcfg.Name, DBcfg.SSLMode)
	adapter, err := pgadapter.NewAdapter(dsn)
	if err != nil {
		log.Fatalf("Casbin: Failed to connect to database: %v", err)
	}

	enforcer, err := casbin.NewEnforcer(cfg.RBAC.ModelFile, adapter)
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}

	err = loadPolicyFromCSV(enforcer, "./policy.csv")
	if err != nil {
		log.Fatalf("Failed to load policy from file: %v", err)
	}

	err = enforcer.SavePolicy()
	if err != nil {
		log.Fatalf("Failed to save policy to database: %v", err)
	}
}

func loadPolicyFromCSV(e *casbin.Enforcer, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) < 3 {
			continue
		}
		switch parts[0] {
		case "p":
			e.AddPolicy(parts[1:])
		case "g":
			e.AddGroupingPolicy(parts[1:])
		}
	}
	return nil
}
