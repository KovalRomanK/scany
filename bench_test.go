package scany_test

import (
	"github.com/georgysavva/scany/dbscan"
	"github.com/stretchr/testify/require"
	"testing"
)

type rows struct {
	count int
}

func (r *rows) Close() error {
	return nil
}

func (r *rows) Err() error {
	return nil
}

func (r *rows) Next() bool {
	if r.count == 0 {
		r.count++
		return true
	}
	return false
}

func (r *rows) Columns() ([]string, error) {
	return columns, nil
}

func (r *rows) Scan(dest ...interface{}) error {
	return nil
}

var err error

func BenchmarkStruct(b *testing.B) {
	dbscan.UseStructCache = false
	model := &Data{}
	r := &rows{}
	for i := 0; i < b.N; i++ {
		rs := dbscan.NewRowScanner(r)
		err = rs.Scan(model)
	}
}

func BenchmarkStructCache(b *testing.B) {
	dbscan.UseStructCache = true
	model := &Data{}
	r := &rows{}
	for i := 0; i < b.N; i++ {
		rs := dbscan.NewRowScanner(r)
		err = rs.Scan(model)
	}
}

func BenchmarkMap(b *testing.B) {
	model := map[string]interface{}{}
	for i := 0; i < b.N; i++ {
		rs := dbscan.NewRowScanner(&rows{})
		err = rs.Scan(&model)
	}
}

func TestStruct(t *testing.T) {
	model := &Data{}
	for i := 0; i < 100; i++ {
		rs := dbscan.NewRowScanner(&rows{})
		err := rs.Scan(model)
		require.NoError(t, err)
	}
}

func TestMap(t *testing.T) {
	model := map[string]interface{}{}
	for i := 0; i < 100; i++ {
		rs := dbscan.NewRowScanner(&rows{})
		err := rs.Scan(&model)
		require.NoError(t, err)
	}
}
