# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: sms.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='sms.proto',
  package='',
  syntax='proto3',
  serialized_options=b'Z\027services/smsqueue/proto',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\tsms.proto\"_\n\x0bSendRequest\x12\x0f\n\x07\x61\x63\x63ount\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\r\n\x05phone\x18\x03 \x01(\t\x12\x0f\n\x07\x63ontent\x18\x04 \x01(\t\x12\r\n\x05\x65xtra\x18\x05 \x01(\t\"C\n\x0cSendResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\x12\n\nrawContent\x18\x03 \x01(\t2?\n\x11UnicomMessagePush\x12*\n\x0bSendMessage\x12\x0c.SendRequest\x1a\r.SendResponseB\x19Z\x17services/smsqueue/protob\x06proto3'
)




_SENDREQUEST = _descriptor.Descriptor(
  name='SendRequest',
  full_name='SendRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account', full_name='SendRequest.account', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='password', full_name='SendRequest.password', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='phone', full_name='SendRequest.phone', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='content', full_name='SendRequest.content', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='extra', full_name='SendRequest.extra', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=13,
  serialized_end=108,
)


_SENDRESPONSE = _descriptor.Descriptor(
  name='SendResponse',
  full_name='SendResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='SendResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='SendResponse.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='rawContent', full_name='SendResponse.rawContent', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=110,
  serialized_end=177,
)

DESCRIPTOR.message_types_by_name['SendRequest'] = _SENDREQUEST
DESCRIPTOR.message_types_by_name['SendResponse'] = _SENDRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SendRequest = _reflection.GeneratedProtocolMessageType('SendRequest', (_message.Message,), {
  'DESCRIPTOR' : _SENDREQUEST,
  '__module__' : 'sms_pb2'
  # @@protoc_insertion_point(class_scope:SendRequest)
  })
_sym_db.RegisterMessage(SendRequest)

SendResponse = _reflection.GeneratedProtocolMessageType('SendResponse', (_message.Message,), {
  'DESCRIPTOR' : _SENDRESPONSE,
  '__module__' : 'sms_pb2'
  # @@protoc_insertion_point(class_scope:SendResponse)
  })
_sym_db.RegisterMessage(SendResponse)


DESCRIPTOR._options = None

_UNICOMMESSAGEPUSH = _descriptor.ServiceDescriptor(
  name='UnicomMessagePush',
  full_name='UnicomMessagePush',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=179,
  serialized_end=242,
  methods=[
  _descriptor.MethodDescriptor(
    name='SendMessage',
    full_name='UnicomMessagePush.SendMessage',
    index=0,
    containing_service=None,
    input_type=_SENDREQUEST,
    output_type=_SENDRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_UNICOMMESSAGEPUSH)

DESCRIPTOR.services_by_name['UnicomMessagePush'] = _UNICOMMESSAGEPUSH

# @@protoc_insertion_point(module_scope)
