install:
	go build -o "${HOME}/cloud/data/alfred/Alfred.alfredpreferences/workflows/user.workflow.A0C8AD8B-99A2-4B42-B7E0-B61021685ED9/alfred-dazn"
run:
	alfred_workflow_data=workflow alfred_workflow_cache=/tmp/alfred alfred_workflow_bundleid=mk_testing go run alfred-dazn.go



