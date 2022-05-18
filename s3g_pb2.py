# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: s3g.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\ts3g.proto\x12\x0b\x61veplen.s3g\"=\n\x13\x43reateBucketRequest\x12\x12\n\nbucketname\x18\x02 \x01(\t\x12\x12\n\npublicread\x18\x04 \x01(\t\"?\n\x14\x43reateBucketResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\t\"8\n\x13\x44\x65leteBucketRequest\x12\r\n\x05store\x18\x01 \x01(\t\x12\x12\n\nbucketname\x18\x02 \x01(\t\"?\n\x14\x44\x65leteBucketResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\t\"#\n\x12ListBucketsRequest\x12\r\n\x05store\x18\x01 \x01(\t\"X\n\x13ListBucketsResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12&\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x18.aveplen.s3g.ListBuckets\"\x1d\n\x0bListBuckets\x12\x0e\n\x06\x62ucket\x18\x01 \x03(\t\"\x9f\x01\n\x10PutObjectRequest\x12\x12\n\nbucketname\x18\x02 \x01(\t\x12\x0b\n\x03key\x18\x03 \x01(\t\x12\x12\n\npublicread\x18\x04 \x01(\t\x12\x13\n\x0b\x63ontenttype\x18\x05 \x01(\t\x12\x18\n\x10\x63ontentmaxlength\x18\x06 \x01(\x03\x12\x13\n\x0b\x66ilecontent\x18\x07 \x01(\t\x12\x12\n\nexpiretime\x18\x08 \x01(\x03\"<\n\x11PutObjectResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\t\"H\n\x10GetObjectRequest\x12\x12\n\nbucketname\x18\x02 \x01(\t\x12\x0b\n\x03key\x18\x03 \x01(\t\x12\x13\n\x0b\x63ontenttype\x18\x04 \x01(\t\"<\n\x11GetObjectResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\t2\xa9\x03\n\tS3Gateway\x12U\n\x0c\x43reateBucket\x12 .aveplen.s3g.CreateBucketRequest\x1a!.aveplen.s3g.CreateBucketResponse\"\x00\x12U\n\x0c\x44\x65leteBucket\x12 .aveplen.s3g.DeleteBucketRequest\x1a!.aveplen.s3g.DeleteBucketResponse\"\x00\x12R\n\x0bListBuckets\x12\x1f.aveplen.s3g.ListBucketsRequest\x1a .aveplen.s3g.ListBucketsResponse\"\x00\x12L\n\tPutObject\x12\x1d.aveplen.s3g.PutObjectRequest\x1a\x1e.aveplen.s3g.PutObjectResponse\"\x00\x12L\n\tGetObject\x12\x1d.aveplen.s3g.GetObjectRequest\x1a\x1e.aveplen.s3g.GetObjectResponse\"\x00\x42\x39Z7github.com/aveplen-bach/s3-grpc-gateway;s3_grpc_gatewayb\x06proto3')



_CREATEBUCKETREQUEST = DESCRIPTOR.message_types_by_name['CreateBucketRequest']
_CREATEBUCKETRESPONSE = DESCRIPTOR.message_types_by_name['CreateBucketResponse']
_DELETEBUCKETREQUEST = DESCRIPTOR.message_types_by_name['DeleteBucketRequest']
_DELETEBUCKETRESPONSE = DESCRIPTOR.message_types_by_name['DeleteBucketResponse']
_LISTBUCKETSREQUEST = DESCRIPTOR.message_types_by_name['ListBucketsRequest']
_LISTBUCKETSRESPONSE = DESCRIPTOR.message_types_by_name['ListBucketsResponse']
_LISTBUCKETS = DESCRIPTOR.message_types_by_name['ListBuckets']
_PUTOBJECTREQUEST = DESCRIPTOR.message_types_by_name['PutObjectRequest']
_PUTOBJECTRESPONSE = DESCRIPTOR.message_types_by_name['PutObjectResponse']
_GETOBJECTREQUEST = DESCRIPTOR.message_types_by_name['GetObjectRequest']
_GETOBJECTRESPONSE = DESCRIPTOR.message_types_by_name['GetObjectResponse']
CreateBucketRequest = _reflection.GeneratedProtocolMessageType('CreateBucketRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEBUCKETREQUEST,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.CreateBucketRequest)
  })
_sym_db.RegisterMessage(CreateBucketRequest)

CreateBucketResponse = _reflection.GeneratedProtocolMessageType('CreateBucketResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEBUCKETRESPONSE,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.CreateBucketResponse)
  })
_sym_db.RegisterMessage(CreateBucketResponse)

DeleteBucketRequest = _reflection.GeneratedProtocolMessageType('DeleteBucketRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEBUCKETREQUEST,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.DeleteBucketRequest)
  })
_sym_db.RegisterMessage(DeleteBucketRequest)

DeleteBucketResponse = _reflection.GeneratedProtocolMessageType('DeleteBucketResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETEBUCKETRESPONSE,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.DeleteBucketResponse)
  })
_sym_db.RegisterMessage(DeleteBucketResponse)

ListBucketsRequest = _reflection.GeneratedProtocolMessageType('ListBucketsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTBUCKETSREQUEST,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.ListBucketsRequest)
  })
_sym_db.RegisterMessage(ListBucketsRequest)

ListBucketsResponse = _reflection.GeneratedProtocolMessageType('ListBucketsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTBUCKETSRESPONSE,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.ListBucketsResponse)
  })
_sym_db.RegisterMessage(ListBucketsResponse)

ListBuckets = _reflection.GeneratedProtocolMessageType('ListBuckets', (_message.Message,), {
  'DESCRIPTOR' : _LISTBUCKETS,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.ListBuckets)
  })
_sym_db.RegisterMessage(ListBuckets)

PutObjectRequest = _reflection.GeneratedProtocolMessageType('PutObjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _PUTOBJECTREQUEST,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.PutObjectRequest)
  })
_sym_db.RegisterMessage(PutObjectRequest)

PutObjectResponse = _reflection.GeneratedProtocolMessageType('PutObjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _PUTOBJECTRESPONSE,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.PutObjectResponse)
  })
_sym_db.RegisterMessage(PutObjectResponse)

GetObjectRequest = _reflection.GeneratedProtocolMessageType('GetObjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETOBJECTREQUEST,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.GetObjectRequest)
  })
_sym_db.RegisterMessage(GetObjectRequest)

GetObjectResponse = _reflection.GeneratedProtocolMessageType('GetObjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETOBJECTRESPONSE,
  '__module__' : 's3g_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3g.GetObjectResponse)
  })
_sym_db.RegisterMessage(GetObjectResponse)

_S3GATEWAY = DESCRIPTOR.services_by_name['S3Gateway']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z7github.com/aveplen-bach/s3-grpc-gateway;s3_grpc_gateway'
  _CREATEBUCKETREQUEST._serialized_start=26
  _CREATEBUCKETREQUEST._serialized_end=87
  _CREATEBUCKETRESPONSE._serialized_start=89
  _CREATEBUCKETRESPONSE._serialized_end=152
  _DELETEBUCKETREQUEST._serialized_start=154
  _DELETEBUCKETREQUEST._serialized_end=210
  _DELETEBUCKETRESPONSE._serialized_start=212
  _DELETEBUCKETRESPONSE._serialized_end=275
  _LISTBUCKETSREQUEST._serialized_start=277
  _LISTBUCKETSREQUEST._serialized_end=312
  _LISTBUCKETSRESPONSE._serialized_start=314
  _LISTBUCKETSRESPONSE._serialized_end=402
  _LISTBUCKETS._serialized_start=404
  _LISTBUCKETS._serialized_end=433
  _PUTOBJECTREQUEST._serialized_start=436
  _PUTOBJECTREQUEST._serialized_end=595
  _PUTOBJECTRESPONSE._serialized_start=597
  _PUTOBJECTRESPONSE._serialized_end=657
  _GETOBJECTREQUEST._serialized_start=659
  _GETOBJECTREQUEST._serialized_end=731
  _GETOBJECTRESPONSE._serialized_start=733
  _GETOBJECTRESPONSE._serialized_end=793
  _S3GATEWAY._serialized_start=796
  _S3GATEWAY._serialized_end=1221
# @@protoc_insertion_point(module_scope)
