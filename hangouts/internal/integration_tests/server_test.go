// Full integration tests for hangouts server.
package server_test

import (
	"context"
	"hangouts/internal/database"
	"hangouts/internal/handler"
	"hangouts/internal/server"
	"hangouts/internal/services"
	"hangouts/internal/transactions"
	"hangouts/internal/utils"
	"log"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var (
	PORT   = "8008"
	CLIENT = utils.CreateTestClient(PORT)
	LOGGER = slog.New(slog.Default().Handler())
)

func runServer(port string) {
	ctx := context.Background()

	dbName := "users"
	dbUser := "user"
	dbPassword := "password"
	dbPort := "5432"
	LOGGER.Info("Creating postgres container")
	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}
	dbHostFn := func() (string, error) { return postgresContainer.Host(ctx) }
	dbHost := utils.SafeCall(dbHostFn)

	dbHostPortFn := func() (nat.Port, error) { return postgresContainer.MappedPort(ctx, nat.Port(dbPort)) }
	dbPort = utils.SafeCall(dbHostPortFn).Port()

	envConfig := &utils.EnvConfig{
		DB_HOST:     dbHost,
		DB_PORT:     dbPort,
		DB_USER:     dbUser,
		DB_PASSWORD: dbPassword,
		DB_NAME:     dbName,
	}
	db := database.CreateDatabase(*envConfig, LOGGER)

	database.AutoMigrate(db)

	transactions := transactions.CreateTransactions(LOGGER, db)

	services := services.CreateServices(LOGGER, transactions)

	h := handler.NewHandler(LOGGER, services)
	server.RunServer(h, ":"+port)
}

func TestMain(m *testing.M) {
	LOGGER.Info("Starting test server in a seperate go routine..")
	go runServer(PORT)
	if !CLIENT.CheckServer(time.Second * 30) {
		os.Exit(1)
	}
	LOGGER.Info("Finished setting up mock postgres container and server...")
	LOGGER.Info("Running tests...")
	code := m.Run()
	os.Exit(code)
}

func TestHealthCheck(t *testing.T) {
	expectedBody := map[string]interface{}{
		"message": "OK",
	}

	testVerify := CLIENT.GET("/healthcheck")
	testVerify.AssertStatusCode(200, t).AssertBody(expectedBody, t)
}
