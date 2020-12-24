# File              : Makefile
# Author            : qiwen <yangdongyong@qianxin.com>
# Date              : 24.12.2020
# Last Modified Date: 24.12.2020
# Last Modified By  : qiwen <yangdongyong@qianxin.com>

all:
	@rm -rfv *.go
	cd ..
	protoc --twirp_out=. --go_out=. --twirp_swagger_out=. -I. ./user.proto
	protoc --twirp_out=. --go_out=. --twirp_swagger_out=. -I. ./acl.proto
	mv api/* ./
