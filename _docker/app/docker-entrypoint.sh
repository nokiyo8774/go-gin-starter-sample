#!/bin/bash

set -e

cd $SRC_DIR
go install github.com/go-delve/delve/cmd/dlv@v1.6.0
go install golang.org/x/tools/gopls@v0.6.9
go install github.com/ramya-rao-a/go-outline@v0.0.0-20200117021646-2a048b4510eb

go install github.com/volatiletech/sqlboiler/v4@v4.5.0
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.5.0

exec "$@"