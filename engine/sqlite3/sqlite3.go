package sqlite3

import (
	"context"
	"database/sql"

	"github.com/go-zoox/core-utils/regexp"
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/fs"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// Name is the name of the engine
const Name = "sqlite3"

// Query executes a statement and returns the result
func Query(ctx context.Context, dsn, statement string) (result any, err error) {
	filepath := dsn
	// if dsn is an url, download the file to tmp
	if regexp.Match(`^https?://`, dsn) {
		filepath = fs.TmpFilePath()
		response, err := fetch.Download(dsn, filepath)
		if err != nil {
			return nil, err
		}
		if !response.Ok() {
			return nil, response.Error()
		}
		defer fs.Remove(filepath)
	}

	db, err := sql.Open(Name, filepath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, statement) // 使用 * 选择所有列
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	data := []map[string]any{}
	for rows.Next() {
		vals := make([]any, len(columns))
		for i := range vals {
			vals[i] = new(any)
		}

		err := rows.Scan(vals...)
		if err != nil {
			return nil, err
		}

		one := make(map[string]any)
		for i, colName := range columns {
			one[colName] = vals[i]
		}

		data = append(data, one)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
