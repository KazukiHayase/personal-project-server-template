package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
)

const defaultLimit = 21

func SetPagination(q *datastore.Query, after *string, first *int, before *string, last *int) (*datastore.Query, error) {
	if after != nil && first != nil {
		limit := *first
		cursor, err := datastore.DecodeCursor(*after)
		if err != nil {
			return q, err
		}

		return q.Start(cursor).Limit(limit), nil
	}

	if before != nil && last != nil {
		limit := *last + 1
		cursor, err := datastore.DecodeCursor(*before)
		if err != nil {
			return q, err
		}

		return q.End(cursor).Limit(limit), nil
	}

	return q.Limit(defaultLimit), nil
}

func HasNextPage(dc datastore.Client, ctx context.Context, q *datastore.Query, c datastore.Cursor) (bool, error) {
	if c.String() == "" {
		return false, nil
	}

	keys, err := dc.GetAll(ctx, q.Limit(1).Start(c).End(datastore.Cursor{}).KeysOnly(), nil)
	if err != nil {
		return false, err
	}
	return len(keys) > 0, nil
}

func HasPreviusPage(dc datastore.Client, ctx context.Context, q *datastore.Query, c datastore.Cursor) (bool, error) {
	if c.String() == "" {
		return false, nil
	}

	// Endの場合はカーソル本体も含まれるのでLimit=2で取得
	keys, err := dc.GetAll(ctx, q.Limit(2).Start(datastore.Cursor{}).End(c).KeysOnly(), nil)
	if err != nil {
		return false, err
	}
	// そもそもカーソル自体が含まれない条件の場合lenghtは0
	return len(keys) > 1, nil
}
