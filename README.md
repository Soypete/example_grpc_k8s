
# RUM
```bash
prototool lint
prototool config init
prototool generate
```

```
protoc -I=proto/ --go_out=gen/go/ proto/order.proto
protoc -I proto/ proto/real_estate.proto --go_out=plugins=grpc:gengo
```
