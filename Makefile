CURRENT_DIR=$(pwd)
pull_submodule:
	git submodule update --init --recursive

update_submodule:
	git submodule update --remote --merge

run:
	go run cmd/main.go

create:
	migrate create -ext sql -dir migrations -seq create_admin_table

up-version:
	migrate -source file:./migrations/ -database 'postgres://postgres:compos1995@locahost:5432/api?sslmode=disable' up


create_proto_submodule:
	git submodule add git@github.com:Asliddin3/Proto-Submodule-Product-servise.git

run_script:
	./script/gen-proto.sh

swag:
	swag init -g ./api/router.go -o api/docs

proto-gen:
	bash ${CURRENT_DIR}/script/gen-proto.sh
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
