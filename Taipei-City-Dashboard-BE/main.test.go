package main

import (
	"testing"
)

var connStr string

func TestGetDBConnection_EmptyConnectionString(t *testing.T) {
	// Temporarily replace the connection string with an empty string
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = ""

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when connection string is empty, got nil")
	}
}

func TestGetDBConnection_MissingHost(t *testing.T) {
	// Temporarily replace the connection string with one missing the host
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "user=postgres password=Sammy0928 dbname=violations_db"

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when database host is not provided, got nil")
	}
}

func TestGetDBConnection_MissingUser(t *testing.T) {
	// Temporarily replace the connection string with one missing the user
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "host=postgres-data password=Sammy0928 dbname=violations_db"

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when database user is not provided, got nil")
	}
}

func TestGetDBConnection_MissingPassword(t *testing.T) {
	// Temporarily replace the connection string with one missing the password
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "host=postgres-data user=postgres dbname=violations_db"

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when database password is not provided, got nil")
	}
}

func TestGetDBConnection_MissingDBName(t *testing.T) {
	// Temporarily replace the connection string with one missing the database name
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "host=postgres-data user=postgres password=Sammy0928"

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when database name is not provided, got nil")
	}
}

func TestGetDBConnection_InvalidDriver(t *testing.T) {
	// Temporarily replace the connection string with an invalid driver
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "invalid_driver://host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when database driver is not 'postgres', got nil")
	}
}

func TestGetDBConnection_SuccessfulConnection(t *testing.T) {
	// Temporarily replace the connection string with a valid one
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"

	db, err := getDBConnection()
	if err != nil {
		t.Errorf("Expected successful connection, got error: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Errorf("Expected successful ping, got error: %v", err)
	}
}

func TestGetDBConnection_ServerUnreachable(t *testing.T) {
	// Temporarily replace the connection string with one pointing to an unreachable server
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "host=unreachable-server user=postgres password=Sammy0928 dbname=violations_db"

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when database server is unreachable, got nil")
	}
}

func TestGetDBConnection_IncorrectCredentials(t *testing.T) {
	// Temporarily replace the connection string with incorrect credentials
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "host=postgres-data user=wronguser password=wrongpassword dbname=violations_db"

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when database credentials are incorrect, got nil")
	}
}

func TestGetDBConnection_DatabaseNotRunning(t *testing.T) {
	// Temporarily replace the connection string with one pointing to a non-running database
	originalConnStr := "host=postgres-data user=postgres password=Sammy0928 dbname=violations_db"
	defer func() { connStr = originalConnStr }()
	connStr = "host=localhost user=postgres password=Sammy0928 dbname=violations_db"

	_, err := getDBConnection()
	if err == nil {
		t.Errorf("Expected error when database is not running, got nil")
	}
}
