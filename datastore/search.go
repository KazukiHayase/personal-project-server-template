package datastore

import "cloud.google.com/go/datastore"

const utf8LastChar = "\xef\xbf\xbd"

func SetSearchKeyword(q *datastore.Query, f string, k string) *datastore.Query {
	// 前方一致
	return q.
		FilterField(f, ">=", k).
		FilterField(f, "<=", k+utf8LastChar).
		Order(f)
}
