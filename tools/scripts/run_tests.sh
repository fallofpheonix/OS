#!/bin/bash
echo "Running go test ./..."
go test ./... > test_out1.txt 2>&1 || echo "failed"

echo "Running go test -race ./..."
go test -race ./... > test_out2.txt 2>&1 || echo "failed"

echo "Running go test -count=1000 ./phoenix_os/containment/rollback/..."
go test -count=1000 ./phoenix_os/containment/rollback/... > test_out3.txt 2>&1 || echo "failed"

echo "Running go test -shuffle=on -count=200 ./phoenix_os/containment/..."
go test -shuffle=on -count=200 ./phoenix_os/containment/... > test_out4.txt 2>&1 || echo "failed"

echo "Running go test -race -shuffle=on ./phoenix_os/containment/..."
go test -race -shuffle=on ./phoenix_os/containment/... > test_out5.txt 2>&1 || echo "failed"
