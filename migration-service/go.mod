module migration-service

go 1.22.11

require (
	github.com/golang-migrate/migrate/v4 v4.18.1
	shared v0.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-chi/chi/v5 v5.2.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/redis/go-redis/v9 v9.7.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace shared => ../shared
