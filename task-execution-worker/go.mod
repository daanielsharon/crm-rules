module task-execution-service

go 1.22.11

require github.com/redis/go-redis/v9 v9.7.0

require (
	github.com/go-chi/chi/v5 v5.2.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	shared v0.0.0
)

replace shared => ../shared