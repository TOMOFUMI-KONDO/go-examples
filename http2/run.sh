#!/bin/bash

echo go run .
go run .

echo GODEBUG=http2client=0 go run .
GODEBUG=http2client=0 go run .

echo GODEBUG=http2debug=1 go run .
GODEBUG=http2debug=1 go run .

echo GODEBUG=http2debug=2 go run .
GODEBUG=http2debug=2 go run .
