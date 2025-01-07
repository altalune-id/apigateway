format:
	gofmt -s -w .

deploy-local:
	cdklocal bootstrap
	cdklocal deploy "Local/*" --require-approval never --force

deploy-dev-id:
	cdk bootstrap
	cdk deploy "Dev-ID/*" --require-approval never --force
