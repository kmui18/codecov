test-coverage:
	@echo "integration testing..."
	go test -v -coverpkg=./... -coverprofile=tmp.cov ./...
	@EXIT_CODE=$$?
	@cat tmp.cov | grep -v '.pb.go' > nongenerated_code.cov
	@exit $$EXIT_CODE
