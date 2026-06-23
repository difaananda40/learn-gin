serve:
	go run cmd/app/main.go

seed:
	go run . --seed

gen-secret:
	go run pkg/scripts/gen-secret.go
