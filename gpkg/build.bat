@echo %Proto%\protoc.exe --proto_path=grpc --go_out=. data.proto model.proto

for /f "tokens=*" %%i in ('go env GOPATH') do set GOPATH=%%i
echo GOPATH is: %GOPATH%


@echo %Proto%\protoc.exe --proto_path=grpc --proto_path=E:\proto-git\protobuf\src --gogofaster_out=. data.proto model.proto

%Proto%\protoc.exe --proto_path=grpc -I=E:\proto-git\protobuf\src -I=%GOPATH%\pkg\mod\github.com\gogo\protobuf@v1.3.2 --gogofaster_out=. data.proto model.proto