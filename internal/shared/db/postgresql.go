package db

import (
	"context"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var dbPool *pgxpool.Pool

func NewPostgresqlConn() *pgxpool.Pool {
	if dbPool != nil {
		return dbPool
	}

	//	# Example Keyword/Value
	//	user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10
	//
	//	# Example URL
	//	postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
	databaseURL := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v\n", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Minute
	config.ConnConfig.Tracer = &queryTracer{
		log: zerolog.New(os.Stdout).With().
			Str("foo", "bar").
			Logger(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	dbPool = pool
	return pool
}

type queryTracer struct {
	log zerolog.Logger
}

type contextKey string

func (tracer *queryTracer) TraceQueryStart(
	ctx context.Context,
	_ *pgx.Conn,
	data pgx.TraceQueryStartData) context.Context {

	ctx = context.WithValue(ctx, contextKey("query_start_time"), time.Now())
	ctx = context.WithValue(ctx, contextKey("query"), cleanSQL(data.SQL))
	ctx = context.WithValue(ctx, contextKey("args"), data.Args)
	return ctx
}

func (tracer *queryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	startTime := ctx.Value(contextKey("query_start_time")).(time.Time)
	duration := time.Since(startTime)

	logEvent := tracer.log.Info().Ctx(ctx)
	sql := ctx.Value(contextKey("query")).(string)
	args := ctx.Value(contextKey("args")).([]interface{})
	if data.Err != nil {
		logEvent = tracer.log.Error().Err(data.Err).Ctx(ctx)
	}
	logEvent.
		Str("sql", sql).
		Interface("args", args).
		Dur("duration", duration).
		Msg("Command executed")
}

func cleanSQL(sql string) string {
	re := regexp.MustCompile(`\s+`)
	cleanedSQL := re.ReplaceAllString(strings.TrimSpace(sql), " ")
	return cleanedSQL
}
