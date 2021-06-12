# go postgres  and gRPC 

1. gen file proto:

   ```console
   $  protoc --go_out=plugins=grpc:../ helloworld.proto 
 