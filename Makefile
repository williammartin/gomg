vet:
	go vet

test: vet
	./scripts/test.sh

coverage: vet
	./scripts/generate-coverage.sh

release: test
	./scripts/release.sh

.PHONY: test
