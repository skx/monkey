#!/bin/sh


# I don't even ..
go env -w GOFLAGS="-buildvcs=false"

# Install the lint-tool, and the shadow-tool
go install golang.org/x/lint/golint@latest
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

# At this point failures cause aborts
set -e

# Run the linter
echo "Launching linter .."

#
# We have a bunch of errors which we need to mask
#
#  opcode/opcode.go:8:2: don't use ALL_CAPS in Go names; use CamelCase
#  opcode/opcode.go:11:2: don't use ALL_CAPS in Go names; use CamelCase
#  opcode/opcode.go:14:2: don't use ALL_CAPS in Go names; use CamelCase
#  opcode/opcode.go:17:2: don't use ALL_CAPS in Go names; use CamelCase
#  opcode/opcode.go:20:2: don't use ALL_CAPS in Go names; use CamelCase
#
( golint  ./...  | grep -v ALL_CAPS > lint.out ) || true
if [ -s lint.out ]; then
    echo "Linter errors: "
    cat lint.out
    exit 1
else
    rm lint.out
fi

echo "Completed linter .."

# Run the shadow-checker
echo "Launching shadowed-variable check .."
go vet -vettool=$(which shadow) ./...
echo "Completed shadowed-variable check .."

# Build a binary and run it.
echo "Launching monkey-test .."
go build -ldflags "-X main.version=test-code" -o "monkey"
./monkey -eval 'puts("ok - stdlib\n"); exit(0);'
echo "Completed monkey-test .."
