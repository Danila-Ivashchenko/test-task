run:

mock_gen:
	@mockgen -source=./internal/domain/ports/enricher/enricher.go -destination=./pkg/moks/enricher/enricher.go
	@mockgen -source=./internal/domain/ports/storage/user.go -destination=./pkg/moks/storage/user.go