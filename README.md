# Remotgo

Send commands over ssh to AWS EC2 instances

### Example
Execute "df -H" command in all instances with core as role and test as environment
```go
go run main.go -role core -environment test -command "df -H"
```

