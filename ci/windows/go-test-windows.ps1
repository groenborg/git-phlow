#SET GOPATH TO TMP DIRECTORY, BECAUSE WE WANT THE DEPENDENCIES LOCALLY
$env:GOPATH = "$PWD\"
$env:GOBIN = "$PWD\bin"

#NAVIGATE TO FOLDER
cd git-phlow

#GET DEPENDENCIES
go get -d -t -v ./...

#RUN TESTS
go test -p 1 ./...

# USES THE LATEST EXITCODE, WHICH IS FROM GO TEST, AND EXITS WITH THAT. THIS FIXIS 
# ISSUE #152
exit $lastexitcode