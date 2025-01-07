format:
	gofmt -s -w .

deploy-local:
	cdklocal bootstrap
	cdklocal deploy "Local/*" --require-approval never --force

deploy:
	cdk deploy "${STAGE}/*" --require-approval never
