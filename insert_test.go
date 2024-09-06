package main

import (
	"context"
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

var dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

func insertRawRecord(b *testing.B, db *gorm.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.WithContext(ctx).Exec("INSERT INTO temp VALUES('rVUDTwPNOC')").Error; err != nil {
		b.Fatal(err)
	}

}

func BenchmarkMaxOpenConnectionEq1(b *testing.B) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	sqlDb.SetMaxOpenConns(1)
	defer func(sqlDb *sql.DB) {
		_ = sqlDb.Close()
	}(sqlDb)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRawRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConnectionEq5(b *testing.B) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	sqlDb.SetMaxOpenConns(5)
	defer func(sqlDb *sql.DB) {
		_ = sqlDb.Close()
	}(sqlDb)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRawRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConnectionEq10(b *testing.B) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	sqlDb.SetMaxOpenConns(10)
	defer func(sqlDb *sql.DB) {
		_ = sqlDb.Close()
	}(sqlDb)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRawRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConnectionEq100(b *testing.B) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		b.Fatal(err)
	}
	sqlDb.SetMaxOpenConns(100)
	defer func(sqlDb *sql.DB) {
		_ = sqlDb.Close()
	}(sqlDb)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRawRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConnectionUnlimited(b *testing.B) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRawRecord(b, db)
		}
	})
}
