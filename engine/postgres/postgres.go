package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

const Name = "postgres"

func Query(ctx context.Context, dsn, statement string) (result any, err error) {
	db, err := sql.Open(Name, dsn)
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
