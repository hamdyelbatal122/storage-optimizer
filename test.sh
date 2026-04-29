#!/bin/bash

# Storage Optimizer - Test Suite

set -e

echo "Running Storage Optimizer test suite..."
echo "========================================="

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Counters
TESTS_RUN=0
TESTS_PASSED=0
TESTS_FAILED=0

# Test function
run_test() {
  local test_name=$1
  local command=$2
  
  TESTS_RUN=$((TESTS_RUN + 1))
  echo -e "\n${YELLOW}Test $TESTS_RUN: $test_name${NC}"
  echo "Command: $command"
  
  if eval "$command" > /tmp/test_output.log 2>&1; then
    TESTS_PASSED=$((TESTS_PASSED + 1))
    echo -e "${GREEN}PASSED${NC}"
  else
    TESTS_FAILED=$((TESTS_FAILED + 1))
    echo -e "${RED}FAILED${NC}"
    echo "Error:"
    cat /tmp/test_output.log | head -20
  fi
}

# Tests
echo -e "\nBuild Tests..."
run_test "Build application" "go build -o ./storage-optimizer main.go"
run_test "Verify executable exists" "test -f ./storage-optimizer"

echo -e "\nHelp Tests..."
run_test "Show main help" "./storage-optimizer --help"
run_test "Help for analyze command" "./storage-optimizer analyze --help"
run_test "Help for duplicates command" "./storage-optimizer duplicates --help"
run_test "Help for large command" "./storage-optimizer large --help"
run_test "Help for cleanup command" "./storage-optimizer cleanup --help"

echo -e "\nFunctionality Tests..."
run_test "Analyze current directory" "./storage-optimizer analyze ."
run_test "Find large files" "./storage-optimizer large . --limit 5"
run_test "Export large files as JSON" "./storage-optimizer large . --limit 5 -f json"

echo -e "\nReport Tests..."
run_test "Export analysis to JSON" "./storage-optimizer analyze . -f json -o /tmp/test-report.json"
run_test "Verify report file exists" "test -f /tmp/test-report.json && test -s /tmp/test-report.json"

echo -e "\n========================================="
echo -e "Test Results:"
echo -e "   Tests run: ${YELLOW}$TESTS_RUN${NC}"
echo -e "   Passed: ${GREEN}$TESTS_PASSED${NC}"
echo -e "   Failed: ${RED}$TESTS_FAILED${NC}"

if [ $TESTS_FAILED -eq 0 ]; then
  echo -e "\n${GREEN}All tests passed!${NC}"
  exit 0
else
  echo -e "\n${RED}Some tests failed${NC}"
  exit 1
fi
