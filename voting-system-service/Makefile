CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}


exp:
	export DBURL='postgres://postgres:1234@localhost:5432/votingsystem?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://postgres:1234@localhost:5432/votingsystem?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table

swag-init:
	swag init -g api/docs/docs.go -o api/docs
# mig-delete:
#   rm -r db/migrations

