all: go py

go: const.pb.go

py: const_pb2.py

%.pb.go: %.proto
	protoc --go_out=plugins=grpc:. $<

%_pb2.py: %.proto
	python3 -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. $<