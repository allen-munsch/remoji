#!/usr/bin/env bash
set -euo pipefail

REMOJI="./remoji"   # path to your binary
TMP_DIR="$(mktemp -d)"
PASS=0
FAIL=0

cleanup() {
    rm -rf "$TMP_DIR"
}
trap cleanup EXIT

# --- helper functions ---------------------------------------------------------

assert_eq() {
    local name="$1"
    local got="$2"
    local expected="$3"

    if [[ "$got" == "$expected" ]]; then
        echo "✔ $name"
        PASS=$((PASS+1))
    else
        echo "❌ $name"
        echo "   expected: '$expected'"
        echo "   got     : '$got'"
        FAIL=$((FAIL+1))
    fi
}

run_case() {
    local name="$1"
    local cmd="$2"
    local expected="$3"

    local output
    output="$(eval "$cmd")"

    assert_eq "$name" "$output" "$expected"
}

# --- test cases ---------------------------------------------------------------

# 1. File → stdout
echo -n "Hello 🌍 World 🚀" > "$TMP_DIR/a.txt"
run_case "File → stdout" \
"$REMOJI $TMP_DIR/a.txt" \
"Hello  World "

# 2. Pipe input
run_case "Pipe input → stdout" \
'printf "Hi 😄 there" | '"$REMOJI" \
"Hi  there"

# 3. In-place edit (-i)
echo -n "Test 🔥🔥🔥 Wow" > "$TMP_DIR/b.txt"
$REMOJI -i "$TMP_DIR/b.txt"
output="$(cat "$TMP_DIR/b.txt")"
assert_eq "In-place edit" "$output" "Test  Wow"

# 4. No emoji input
echo -n "no emoji here" > "$TMP_DIR/c.txt"
run_case "No emoji untouched" \
"$REMOJI $TMP_DIR/c.txt" \
"no emoji here"

# 5. Mixed content
echo -n "123 ❤️ 456 😀 789" > "$TMP_DIR/d.txt"
run_case "Mixed content" \
"$REMOJI $TMP_DIR/d.txt" \
"123  456  789"

# --- summary ------------------------------------------------------------------

echo
echo "Summary:"
echo "  ✔ Passed: $PASS"
echo "  ❌ Failed: $FAIL"

if [[ $FAIL -gt 0 ]]; then
    exit 1
fi
