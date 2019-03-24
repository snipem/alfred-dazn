install:
	go build -o "${HOME}/bin/alfred-dazn"
run:
	alfred_workflow_data=workflow alfred_workflow_cache=/tmp/alfred alfred_workflow_bundleid=mk_testing go run alfred-dazn.go



