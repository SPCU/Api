protoc -I device/ --go_out=device/ device/*.proto
protoc -I metrics/ --go_out=metrics/ metrics/*.proto
