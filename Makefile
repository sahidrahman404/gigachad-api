# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the cmd/api application
.PHONY: build
build:
	go build -o=/tmp/bin/api ./cmd/api


## build/worker: build the cmd/worker application
.PHONY: build/worker
build/worker:
	go build -o=/tmp/bin/worker ./cmd/worker
	
## run: run the cmd/api application
.PHONY: run
run: build
	/tmp/bin/api

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/api" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"


## run/worker/live: run the worker with reloading on file changes
.PHONY: run/worker/live
run/worker/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build/worker" --build.bin "/tmp/bin/worker" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

## ent/init name=$: initialize a new ent schema
.PHONY: ent/init
ent/init:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}

## ent/gen: generate ent assets
.PHONY: ent/gen
ent/gen:
		go generate .

## atlas/migrate/diff name=$1
.PHONY: atlas/migrate/diff
atlas/migrate/diff:
	atlas migrate diff ${name} \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "sqlite://file?mode=memory&_fk=1"


## atlas/migrate/apply name=$1
.PHONY: atlas/migrate/apply
atlas/migrate/apply:
		atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "sqlite://file.db?_fk=1"

## atlas/migrate/inspect
.PHONY: atlas/migrate/inspect
atlas/migrate/inspect:
		atlas schema inspect --url "sqlite://file.db?_fk=1" > schema.hcl
