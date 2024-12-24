package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type PostgresDatabase struct {
	connpool *pgxpool.Pool
	timeout  int
	logger   slog.Logger
}

func NewPostgresDatabase(logger slog.Logger) PostgresDatabase {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_DATABASE_URL"))

	if err != nil {
		logger.Error(err.Error())
		return PostgresDatabase{}
	}

	dbTimeout, err := strconv.Atoi(os.Getenv("DATABASE_TIMEOUT"))

	if err != nil {
		logger.Error(err.Error())
	}

	return PostgresDatabase{
		connpool: dbpool,
		timeout:  dbTimeout,
		logger:   logger,
	}
}

func (dbh PostgresDatabase) Setup() {
	timeoutContext, cancelFunc := context.WithTimeout(context.TODO(), time.Second*time.Duration(dbh.timeout))
	defer cancelFunc()

	length, ok := os.LookupEnv("URL_LENGTH")

	if !ok {
		length = "1024"
		dbh.logger.Info("`URL_LENGTH` ENV not found")
	}

	createTableQuery := fmt.Sprintf("CREATE TABLE redirects (id SERIAL PRIMARY KEY NOT NULL, source VARCHAR (%s) UNIQUE NOT NULL, destination TEXT NOT NULL, vistors INTEGER NOT NULL)", length)

	_, err := dbh.connpool.Query(timeoutContext, createTableQuery)
	if err != nil {
		dbh.logger.Error(err.Error())
		return
	}
}

func (dbh PostgresDatabase) InsertUrl(source string, destination string) {
	timeoutContext, cancelFunc := context.WithTimeout(context.TODO(), time.Second*time.Duration(dbh.timeout))

	defer cancelFunc()

	_, err := dbh.connpool.Query(timeoutContext, "INSERT INTO redirects (source, destination, vistors) VALUES ($1, $2, $3)", source, destination, 0)

	if err != nil {
		dbh.logger.Error(err.Error())
		return
	}

	dbh.logger.Info(fmt.Sprintf("Insert Source: %s, Destination: %s", source, destination))
}

func (dbh PostgresDatabase) CheckDuplicateUrl(source string) bool {
	timeoutContext, cancelFunc := context.WithTimeout(context.TODO(), time.Second*time.Duration(dbh.timeout))

	defer cancelFunc()

	var notDuplicateUrl bool
	err := dbh.connpool.QueryRow(timeoutContext, "SELECT EXISTS(SELECT 1 FROM redirects WHERE source = $1)", source).Scan(&notDuplicateUrl)

	if err != nil {
		dbh.logger.Error(err.Error())
		return true
	}

	if notDuplicateUrl {
		dbh.logger.Info(fmt.Sprintf("Duplicate URL detected: %s", source))
		return true
	}

	return notDuplicateUrl
}

func (dbh PostgresDatabase) SelectDestination(source string) (string, string, int, error) {
	timeoutContext, cancelFunc := context.WithTimeout(context.TODO(), time.Second*time.Duration(dbh.timeout))
	defer cancelFunc()
	var tbDestination string
	var tbSource string
	var vistors int
	err := dbh.connpool.QueryRow(timeoutContext, "SELECT source, destination, vistors FROM redirects WHERE source=$1", source).Scan(&tbSource, &tbDestination, &vistors)

	if err != nil {
		dbh.logger.Error(err.Error())
		return "", "", 0, err
	}
	return tbSource, tbDestination, vistors, nil
}

func (dbh PostgresDatabase) AddVisitor(source string) bool {
	timeoutContext, cancelFunc := context.WithTimeout(context.TODO(), time.Second*time.Duration(dbh.timeout))
	defer cancelFunc()
	_, err := dbh.connpool.Query(timeoutContext, "UPDATE redirects SET vistors = vistors + 1 WHERE source=$1", source)

	if err != nil {
		dbh.logger.Error(err.Error())
		return false
	}

	return true
}
