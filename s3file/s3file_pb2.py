# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: s3file.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cs3file.proto\x12\x0e\x61veplen.s3file\x1a\x1bgoogle/protobuf/empty.proto\"+\n\x0bImageObject\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x10\n\x08\x63ontents\x18\x02 \x01(\x0c\"#\n\x15GetImageObjectRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x32\xa8\x01\n\tS3Gateway\x12\x45\n\x0ePutImageObject\x12\x1b.aveplen.s3file.ImageObject\x1a\x16.google.protobuf.Empty\x12T\n\x0eGetImageObject\x12%.aveplen.s3file.GetImageObjectRequest\x1a\x1b.aveplen.s3file.ImageObjectB Z\x1egithub.com/aveplen-back/s3fileb\x06proto3')



_IMAGEOBJECT = DESCRIPTOR.message_types_by_name['ImageObject']
_GETIMAGEOBJECTREQUEST = DESCRIPTOR.message_types_by_name['GetImageObjectRequest']
ImageObject = _reflection.GeneratedProtocolMessageType('ImageObject', (_message.Message,), {
  'DESCRIPTOR' : _IMAGEOBJECT,
  '__module__' : 's3file_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3file.ImageObject)
  })
_sym_db.RegisterMessage(ImageObject)

GetImageObjectRequest = _reflection.GeneratedProtocolMessageType('GetImageObjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETIMAGEOBJECTREQUEST,
  '__module__' : 's3file_pb2'
  # @@protoc_insertion_point(class_scope:aveplen.s3file.GetImageObjectRequest)
  })
_sym_db.RegisterMessage(GetImageObjectRequest)

_S3GATEWAY = DESCRIPTOR.services_by_name['S3Gateway']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\036github.com/aveplen-back/s3file'
  _IMAGEOBJECT._serialized_start=61
  _IMAGEOBJECT._serialized_end=104
  _GETIMAGEOBJECTREQUEST._serialized_start=106
  _GETIMAGEOBJECTREQUEST._serialized_end=141
  _S3GATEWAY._serialized_start=144
  _S3GATEWAY._serialized_end=312
# @@protoc_insertion_point(module_scope)