#
# The include should be a single file that contains:
# export APIKEY := {APIKEY}
# export SECRET := {SECRET}
#
include env

$(info $$APIKEY is [${APIKEY}])
$(info $$SECRET is [${SECRET}])

all:
	go run *.go

build: ## Build
	go build *.go

.DEFAULT_GOAL := all