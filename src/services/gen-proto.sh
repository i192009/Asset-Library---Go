#!/bin/bash

pwd
files="UserService.proto OrgMagService.proto StorageService.proto AssetLibraryService.proto CommonBase.proto"
    
for file in $files; do
    echo $file
    protoc --proto_path=../grpc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional ../grpc/$file
done
