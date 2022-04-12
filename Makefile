test-coverage:
	@echo "integration testing..."
	go test -v -coverpkg=./... -coverprofile=tmp.cov ./...
	@EXIT_CODE=$$?
	@cat tmp.cov | grep -v '.pb.go' > nongenerated_code.cov
	@go tool cover -func nongenerated_code.cov -o coverage.out 
	# -@rm *.cov
	@exit $$EXIT_CODE