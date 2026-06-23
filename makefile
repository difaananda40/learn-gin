serve:
	go run main.go

seed:
	go run . --seed

gen-secret:
	go run tools/gen-secret.go
