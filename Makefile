CURRENT_DIR=$(pwd)

run:
	go run cmd/main.go

run_script:
	./script/gen-proto.sh

swag:
	swag init -g ./api/router.go -o api/docs

proto-gen:
	bash ${CURRENT_DIR}/script/gen-proto.sh
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
