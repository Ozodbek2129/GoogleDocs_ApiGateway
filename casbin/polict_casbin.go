package casbin

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres_container"
	port     = "5432"
	dbname   = "casbin"
	username = "macbookpro"
	password = "1111"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to database", "error", err.Error())
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error("Error pinging the database", "error", err.Error())
		return nil, err
	}
	query := `DROP TABLE IF EXISTS "casbin_rule";`
	db.Exec(query)

	adapter, err := xormadapter.NewAdapter("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password))
	if err != nil {
		logger.Error("Error creating Casbin adapter", "error", err.Error())
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("casbin/model.conf", adapter)
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}

	
	policies := [][]string{
		//user
		{"admin", "api/user/profile/:id", "GET"},
		{"admin", "api/user/updateUser/:id", "PUT"},
		{"admin", "api/user/email/:email", "GET"},

		{"patient", "api/user/profile/:id", "GET"},
		{"patient", "api/user/updateUser/:id", "PUT"},
		{"patient", "api/user/email/:email", "GET"},

		{"doctor", "api/user/profile/:id", "GET"},
		{"doctor", "api/user/updateUser/:id", "PUT"},
		{"doctor", "api/user/email/:email", "GET"},

		
	}

	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		logger.Error("Error adding Casbin policy", "error", err.Error())
		return nil, err
	}

	err = enforcer.SavePolicy()
	if err != nil {
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}

	return enforcer, nil
}