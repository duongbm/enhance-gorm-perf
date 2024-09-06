# Benchmark gorm

### setup database first

```shell
docker compose -f docker-compose-pg.yaml up
```

### Benchmark with SetMaxOpenConns

- code example

```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// set the maximum number of concurrently open connections (in-use + idle) to 10
// if this less than or equal 0 it means that there is no limit (it is also default setting).
sqlDb.SetMaxOpenConns(10)

```

- In this code, the pool now has a maximum of 10 concurrently open connections. \
  If all of 10 these connections are in-use and another new connection is needed, then the application will be forced to
  wait until one of the 10 connections is freed up and becomes idle.
- Run benchmark test with the maximum open connections set to 1, 5, 10, 100 and unlimited.\
  The benchmark executes parallel `INSERT` statements on a PostgreSQL database. And here is result

| Func                                  | Times | ns/op        |
|---------------------------------------|-------|--------------|
| BenchmarkMaxOpenConnectionEq1-8       | 1620  | 735514 ns/op |
| BenchmarkMaxOpenConnectionEq5-8       | 4060  | 287342 ns/op |
| BenchmarkMaxOpenConnectionEq10-8      | 4281  | 262860 ns/op |
| BenchmarkMaxOpenConnectionEq100-8     | 4314  | 280659 ns/op |
| BenchmarkMaxOpenConnectionUnlimited-8 | 4424  | 274939 ns/op |

**Summary**
- The more open connections that are allowed, the less time is taken to perform operation. \
- The more open connections that are permitted, the more queries can be performed concurrently.
