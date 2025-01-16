module user-service

go 1.21.3

require (
	github.com/go-chi/chi/v5 v5.2.0
	shared v0.0.0
)

require github.com/lib/pq v1.10.9 // indirect

replace shared => ../shared
