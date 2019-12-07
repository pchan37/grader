#!/usr/bin/env bash

echo 'Generating binary for linux...'
gox -osarch="linux/amd64"
mv grader_linux_amd64 grader
tar -czf grader-$1-python-x86_64-linux.tar.gz grader

echo 'Generating binary for darwin...'
gox -osarch="darwin/amd64"
mv grader_darwin_amd64 grader
tar -czf grader-$1-python-x86_64-darwin.tar.gz grader

echo 'Generating binary for windows...'
gox -osarch="windows/amd64"
mv grader_windows_amd64.exe grader.exe
zip grader-$1-python-x86_64-windows.zip grader.exe

echo 'Cleaning up...'
rm grader grader.exe
