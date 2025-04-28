package gopsql_test

import (
	"encoding/json"
	"testing"

	go_test "github.com/karim-w/go-test"
	"github.com/karim-w/gopsql"
	"github.com/stretchr/testify/assert"
)

func TestSQLScan(t *testing.T) {
	db, cleanup := go_test.InitDockerPostgresSQLDBTest(t)
	defer cleanup()

	const migration = `
	CREATE TABLE test (
		id SERIAL PRIMARY KEY,
		data JSONB NOT NULL
	);`

	_, err := db.Exec(migration)
	assert.Nil(t, err)

	sample := map[string]any{
		"foo": "bar",
	}

	byts, err := json.Marshal(sample)
	assert.Nil(t, err)

	_, err = db.Exec("INSERT INTO test (data) VALUES ($1)", byts)
	assert.Nil(t, err)

	var res gopsql.JSONB

	err = db.QueryRow("SELECT data FROM test").Scan(&res)
	assert.Nil(t, err)

	assert.Equal(t, sample["foo"], res["foo"])
}

func TestSQLScanNull(t *testing.T) {
	db, cleanup := go_test.InitDockerPostgresSQLDBTest(t)
	defer cleanup()

	const migration = `
	CREATE TABLE test (
		id SERIAL PRIMARY KEY,
		data JSONB NULL
	);

	INSERT INTO test (data) VALUES (NULL)
	`

	_, err := db.Exec(migration)
	assert.Nil(t, err)

	var res gopsql.JSONB

	err = db.QueryRow("SELECT data FROM test").Scan(&res)
	assert.Nil(t, err)

	assert.Equal(t, nil, res["foo"])
}

func TestSQLExec(t *testing.T) {
	db, cleanup := go_test.InitDockerPostgresSQLDBTest(t)
	defer cleanup()

	const migration = `
	CREATE TABLE test (
		id SERIAL PRIMARY KEY,
		data JSONB NOT NULL
	);`

	_, err := db.Exec(migration)
	assert.Nil(t, err)

	sample := gopsql.JSONB{
		"foo": "bar",
	}

	_, err = db.Exec("INSERT INTO test (data) VALUES ($1)", sample)
	assert.Nil(t, err)

	var result string

	err = db.QueryRow("SELECT data->>'foo' FROM test").Scan(&result)
	assert.Nil(t, err)

	assert.Equal(t, result, "bar")
}

func TestSQLExecNull(t *testing.T) {
	db, cleanup := go_test.InitDockerPostgresSQLDBTest(t)
	defer cleanup()

	const migration = `
	CREATE TABLE test (
		id SERIAL PRIMARY KEY,
		data JSONB NULL
	);`

	_, err := db.Exec(migration)
	assert.Nil(t, err)

	var sample gopsql.JSONB

	_, err = db.Exec("INSERT INTO test (data) VALUES ($1)", sample)
	assert.Nil(t, err)

	var res gopsql.JSONB

	err = db.QueryRow("SELECT data FROM test").Scan(&res)
	assert.Nil(t, err)

	t.Log(res)

	assert.Equal(t, map[string]any{}, map[string]any(res))

	res["foo"] = 2
}
