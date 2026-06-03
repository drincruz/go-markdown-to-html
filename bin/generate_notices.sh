#!/usr/bin/env bash
set -euo pipefail

PACKAGE="github.com/drincruz/go-markdown-to-html/..."
NOTICES_DIR="./third_party"
TEMPLATE="./bin/notices.tmpl"

# Locate or install go-licenses
if command -v go-licenses &> /dev/null; then
  GO_LICENSES="go-licenses"
elif [ -f "$(go env GOPATH)/bin/go-licenses" ]; then
  GO_LICENSES="$(go env GOPATH)/bin/go-licenses"
else
  echo "Installing go-licenses..."
  go install github.com/google/go-licenses@latest
  GO_LICENSES="$(go env GOPATH)/bin/go-licenses"
fi

# Ignore the project's own package to only list third-party modules
IGNORE_PKG="github.com/drincruz/go-markdown-to-html"

# 1. Check for any forbidden or unknown licenses — fail fast
echo "Checking licenses..."
"$GO_LICENSES" check "$PACKAGE" --ignore "$IGNORE_PKG"

# 2. Save all license files into third_party/
echo "Saving license files..."
rm -rf "$NOTICES_DIR"
"$GO_LICENSES" save "$PACKAGE" --save_path="$NOTICES_DIR" --force --ignore "$IGNORE_PKG"

# 3. Generate a human-readable THIRD_PARTY_NOTICES.md
echo "Generating THIRD_PARTY_NOTICES.md..."
"$GO_LICENSES" report "$PACKAGE" --template="$TEMPLATE" --ignore "$IGNORE_PKG" > THIRD_PARTY_NOTICES.md

echo "Done. Review THIRD_PARTY_NOTICES.md and $NOTICES_DIR/"
