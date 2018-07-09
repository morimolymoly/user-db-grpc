#user-db-grpc

## protoc
```
export PATH=$PATH:$GOPATH/bin
protoc --go_out=plugins=grpc:./ pb/user.proto
```
## grpcc
```
grpcc --proto pb/user.proto --address localhost:8100 -i

Connecting to pb.UserService on localhost:8100. Available globals:

  client - the client connection to UserService
    getAllUser (GetAllUserRequest, callback) returns Users

  printReply - function to easily print a unary call reply (alias: pr)
  streamReply - function to easily print stream call replies (alias: sr)
  createMetadata - convert JS objects into grpc metadata instances (alias: cm)
  printMetadata - function to easily print a unary call's metadata (alias: pm)

UserService@localhost:8100> client.getAllUser({}, pr)
EventEmitter {}
UserService@localhost:8100>
{
  "u": [
    {
      "userID": "0",
      "username": "もり"
    }
  ]
}
```
