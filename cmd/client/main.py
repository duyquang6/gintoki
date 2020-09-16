import grpc
import product_inventory_pb2_grpc
import product_inventory_pb2

channel = grpc.insecure_channel('0.0.0.0:8080')
stub = product_inventory_pb2_grpc.ProductInventoryServiceStub(channel)

req = product_inventory_pb2.GetMultiProductInventoryRequest(api="v1", product_ids=[106,107])
res = stub.GetMultiProductInventory(req)
print(len(res.data))
