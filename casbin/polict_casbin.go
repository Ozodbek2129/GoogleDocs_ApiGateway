package casbin

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

const (
	host     = "postgres3"
	port     = "5432"
	dbname   = "casbin"
	username = "postgres"
	password = "1234"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to database", "error", err.Error())
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS casbin")
	if err != nil {
		logger.Error("Error dropping Casbin database", "error", err.Error())
		return nil, err
	}

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
		{"admin", "api/user/getbyuser/:email", "GET"},
		{"admin", "api/user/update_user", "PUT"},
		{"admin", "api/user/delete_user/:id", "DELETE"},

		{"user", "api/user/getbyuser/:email", "GET"},
		{"user", "api/user/update_user", "PUT"},
		{"user", "api/user/delete_user/:id", "DELETE"},


		// docs
		{"admin", "api/user/docs/createDocument", "POST"},
		{"admin", "api/user/docs/SearchDocument", "GET"},
		{"admin", "api/user/docs/GetAllDocuments", "GET"},
		{"admin", "api/user/docs/UpdateDocument", "PUT"},
		{"admin", "api/user/docs/DeleteDocument", "DELETE"},
		{"admin", "api/user/docs/ShareDocument", "POST"},

		{"user", "api/user/docs/createDocument", "POST"},
		{"user", "api/user/docs/SearchDocument", "GET"},
		{"user", "api/user/docs/GetAllDocuments", "GET"},
		{"user", "api/user/docs/UpdateDocument", "PUT"},
		{"user", "api/user/docs/DeleteDocument", "DELETE"},
		{"user", "api/user/docs/ShareDocument", "POST"},


		// version
		{"admin", "api/user/version/GetAllVersions", "GET"},
		{"admin", "api/user/version/RestoreVersion", "PUT"},

		{"user", "api/user/version/GetAllVersions", "GET"},
		{"user", "api/user/version/RestoreVersion", "PUT"},
		
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