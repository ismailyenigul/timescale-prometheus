package pgmodel

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	labelCount      = 5
	metricName      = "benchmark_metric"
	otherMetricName = "other_benchmark_metric"
)

func BenchmarkGetSeriesIDForKeyValueArrayExistingSeries(b *testing.B) {
	b.StopTimer()
	withDB(b, "bench_1", func(db *pgxpool.Pool, t testing.TB) {
		err := createMetricTableName(db, metricName)
		if err != nil {
			t.Fatal(err)
		}

		keys, values := generateKeysAndValues(b.N+labelCount, "label")
		for n := 0; n < b.N; n++ {
			err = getSeriesIDForKeyValueArray(db, metricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}

		var bench *testing.B
		var ok bool
		if bench, ok = t.(*testing.B); !ok {
			t.Fatal("Not a benchmarking instance, stopping benchmark")
		}

		bench.ResetTimer()
		bench.StartTimer()
		for n := 0; n < b.N; n++ {
			err = getSeriesIDForKeyValueArray(db, metricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}
		bench.StopTimer()
	})
}

func BenchmarkGetSeriesIDForKeyValueArrayNewSeriesExistingLabels(b *testing.B) {
	b.StopTimer()
	withDB(b, "bench_2", func(db *pgxpool.Pool, t testing.TB) {
		err := createMetricTableName(db, metricName)
		if err != nil {
			t.Fatal(err)
		}

		keys, values := generateKeysAndValues(b.N+labelCount, "label")
		for n := 0; n < b.N; n++ {
			err = getSeriesIDForKeyValueArray(db, otherMetricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}

		var bench *testing.B
		var ok bool
		if bench, ok = t.(*testing.B); !ok {
			t.Fatal("Not a benchmarking instance, stopping benchmark")
		}

		bench.ResetTimer()
		bench.StartTimer()
		for n := 0; n < b.N; n++ {
			err = getSeriesIDForKeyValueArray(db, metricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}
		bench.StopTimer()
	})
}

func BenchmarkGetSeriesIDForKeyValueArrayNewSeriesNewLabels(b *testing.B) {
	b.StopTimer()
	withDB(b, "bench_3", func(db *pgxpool.Pool, t testing.TB) {
		err := createMetricTableName(db, metricName)
		if err != nil {
			t.Fatal(err)
		}

		keys, values := generateKeysAndValues(b.N+labelCount, "label")

		var bench *testing.B
		var ok bool
		if bench, ok = t.(*testing.B); !ok {
			t.Fatal("Not a benchmarking instance, stopping benchmark")
		}

		bench.ResetTimer()
		bench.StartTimer()
		for n := 0; n < b.N; n++ {
			err = getSeriesIDForKeyValueArray(db, metricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}
		bench.StopTimer()
	})
}

func BenchmarkKeyValueArrayToLabelArrayCreateNewLabels(b *testing.B) {
	b.StopTimer()
	withDB(b, "bench_4", func(db *pgxpool.Pool, t testing.TB) {
		err := createMetricTableName(db, metricName)
		if err != nil {
			t.Fatal(err)
		}

		keys, values := generateKeysAndValues(b.N+labelCount, "label")

		var bench *testing.B
		var ok bool
		if bench, ok = t.(*testing.B); !ok {
			t.Fatal("Not a benchmarking instance, stopping benchmark")
		}

		bench.ResetTimer()
		bench.StartTimer()
		for n := 0; n < b.N; n++ {
			err = keyValueArrayToLabelArray(db, metricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}
		bench.StopTimer()
	})
}

func BenchmarkKeyValueArrayToLabelArrayExistingLabels(b *testing.B) {
	b.StopTimer()
	withDB(b, "bench_5", func(db *pgxpool.Pool, t testing.TB) {
		err := createMetricTableName(db, metricName)
		if err != nil {
			t.Fatal(err)
		}

		keys, values := generateKeysAndValues(b.N+labelCount, "label")
		for n := 0; n < b.N; n++ {
			err = getSeriesIDForKeyValueArray(db, metricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}

		var bench *testing.B
		var ok bool
		if bench, ok = t.(*testing.B); !ok {
			t.Fatal("Not a benchmarking instance, stopping benchmark")
		}

		bench.ResetTimer()
		bench.StartTimer()
		for n := 0; n < b.N; n++ {
			err = keyValueArrayToLabelArray(db, metricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}
		bench.StopTimer()
	})
}

func BenchmarkKeyValueArrayToLabelArrayCreateNewLabelKeys(b *testing.B) {
	b.StopTimer()
	withDB(b, "bench_6", func(db *pgxpool.Pool, t testing.TB) {

		err := createMetricTableName(db, metricName)
		if err != nil {
			t.Fatal(err)
		}

		err = createMetricTableName(db, otherMetricName)
		if err != nil {
			t.Fatal(err)
		}

		keys, values := generateKeysAndValues(b.N+labelCount, "label")
		for n := 0; n < b.N; n++ {
			err = getSeriesIDForKeyValueArray(db, otherMetricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}

		keys, values = generateKeysAndValues(b.N+labelCount, "label_new")

		var bench *testing.B
		var ok bool
		if bench, ok = t.(*testing.B); !ok {
			t.Fatal("Not a benchmarking instance, stopping benchmark")
		}

		bench.ResetTimer()
		bench.StartTimer()
		for n := 0; n < b.N; n++ {
			err = keyValueArrayToLabelArray(db, metricName, keys[n:n+labelCount], values[n:n+labelCount])
			if err != nil {
				t.Fatal(err)
			}
		}
		bench.StopTimer()
	})
}

func BenchmarkGetOrCreateMetricTableName(b *testing.B) {
	b.StopTimer()
	withDB(b, "bench_7", func(db *pgxpool.Pool, t testing.TB) {
		metricNames, _ := generateKeysAndValues(b.N, "metric")

		var bench *testing.B
		var ok bool
		if bench, ok = t.(*testing.B); !ok {
			t.Fatal("Not a benchmarking instance, stopping benchmark")
		}

		bench.ResetTimer()
		bench.StartTimer()
		var err error
		for n := 0; n < b.N; n++ {
			err = createMetricTableName(db, metricNames[n])
			if err != nil {
				t.Fatal(err)
			}
		}
		bench.StopTimer()
	})
}

func keyValueArrayToLabelArray(db *pgxpool.Pool, metricName string, keys []string, values []string) error {
	var labelArray []int
	return db.QueryRow(context.Background(), "SELECT label_array($1, $2, $3)", metricName, keys, values).Scan(&labelArray)
}

func createMetricTableName(db *pgxpool.Pool, name string) error {
	var metricID int
	var tableName string
	return db.QueryRow(context.Background(), "SELECT * FROM _prom_catalog.get_or_create_metric_table_name($1)", name).Scan(&metricID, &tableName)
}

func getSeriesIDForKeyValueArray(db *pgxpool.Pool, metricName string, keys []string, values []string) error {
	var seriesIDKeyVal int
	return db.QueryRow(context.Background(), "SELECT _prom_catalog.get_series_id_for_key_value_array($1, $2, $3)", metricName, keys, values).Scan(&seriesIDKeyVal)
}

func generateKeysAndValues(count int, prefix string) ([]string, []string) {
	keys, values := make([]string, count), make([]string, count)

	for i := 0; i < count; i++ {
		keys[i] = fmt.Sprintf("%s_key_%d", prefix, i%labelCount)
		values[i] = fmt.Sprintf("%s_value_%d", prefix, i)
	}

	return keys, values
}
