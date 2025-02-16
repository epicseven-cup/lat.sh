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

	createTableQuery := fmt.Sprintf("CREATE TABLE redirects (id SERIAL PRIMARY KEY NOT NULL, source VARCHAR (%s) UNIQUE NOT NULL, destination TEXT NOT NULL, visitors INTEGER NOT NULL)", length)

	rows, err := dbh.connpool.Query(timeoutContext, createTableQuery)
	defer rows.Close()
	if err != nil {
		dbh.logger.Error(err.Error())
		return
	}
}

func (dbh PostgresDatabase) InsertUrl(source string, destination string) {
	timeoutContext, cancelFunc := context.WithTimeout(context.TODO(), time.Second*time.Duration(dbh.timeout))

	defer cancelFunc()

	//NOTE: When Using `Query` need to rememebr to close the rows connection
	rows, err := dbh.connpool.Query(timeoutContext, "INSERT INTO redirects (source, destination, visitors) VALUES ($1, $2, $3)", source, destination, 0)
	defer rows.Close()

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

	//NOTE: All QueryRow need to be closed after used, this is done so that it can free up the connection
	row := dbh.connpool.QueryRow(timeoutContext, "SELECT EXISTS(SELECT 1 FROM redirects WHERE source = $1)", source)
	// The connection is return when Scan is called
	err := row.Scan(&notDuplicateUrl)
	if err != nil {
		dbh.logger.Error(err.Error())
		return true
	}

	if notDuplicateUrl {
		dbh.logger.Info(fmt.Sprintf("Duplicate URL detected: %s", source))
		return true
	}

	dbh.logger.Info(fmt.Sprintf("Duplicate Status: %s, %v", source, notDuplicateUrl))
	return notDuplicateUrl
}

func (dbh PostgresDatabase) SelectDestination(source string) (string, string, int, error) {
	timeoutContext, cancelFunc := context.WithTimeout(context.TODO(), time.Second*time.Duration(dbh.timeout))
	defer cancelFunc()
	var tbDestination string
	var tbSource string
	var visitors int

	row := dbh.connpool.QueryRow(timeoutContext, "SELECT source, destination, visitors FROM redirects WHERE source=$1 LIMIT 1", source)
	err := row.Scan(&tbSource, &tbDestination, &visitors)

	if err != nil {
		dbh.logger.Error(err.Error())
		return "", "", 0, err
	}

	dbh.logger.Info(fmt.Sprintf("Current visitor count: %d", visitors))

	return tbSource, tbDestination, visitors, nil
}

func (dbh PostgresDatabase) AddVisitor(source string) bool {
	timeoutContext, cancelFunc := context.WithTimeout(context.TODO(), time.Second*time.Duration(dbh.timeout))
	defer cancelFunc()
	rows, err := dbh.connpool.Query(timeoutContext, "UPDATE redirects SET visitors = visitors + 1 WHERE source=$1", source)
	defer rows.Close()

	if err != nil {
		dbh.logger.Error("I'm looking at error bro")
		dbh.logger.Error(err.Error())
		return false
	}
	return true
}
