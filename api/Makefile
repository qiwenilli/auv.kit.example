all:
	@rm -rfv *.go
	cd ..
	protoc --twirp_out=. --go_out=. --twirp_swagger_out=. -I. ./user.proto
	protoc --twirp_out=. --go_out=. --twirp_swagger_out=. -I. ./acl.proto
	mv api/* ./
