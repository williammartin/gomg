vet:
	go vet

test: vet
	ginkgo -p -r --randomizeAllSpecs --failOnPending --randomizeSuites --race

coverage: vet
	ginkgo -p -r --randomizeAllSpecs --failOnPending --randomizeSuites --race -cover -coverprofile=coverage.txt -outputdir=.

.PHONY: test
