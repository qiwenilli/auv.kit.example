
all:
	@rm -rfv *.go
	protoc --twirp_out=. --go_out=. --twirp_swagger_out=. -I. ./user.proto
