# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import s3g_pb2 as s3g__pb2


class S3GatewayStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateBucket = channel.unary_unary(
                '/aveplen.s3g.S3Gateway/CreateBucket',
                request_serializer=s3g__pb2.CreateBucketRequest.SerializeToString,
                response_deserializer=s3g__pb2.CreateBucketResponse.FromString,
                )
        self.DeleteBucket = channel.unary_unary(
                '/aveplen.s3g.S3Gateway/DeleteBucket',
                request_serializer=s3g__pb2.DeleteBucketRequest.SerializeToString,
                response_deserializer=s3g__pb2.DeleteBucketResponse.FromString,
                )
        self.ListBuckets = channel.unary_unary(
                '/aveplen.s3g.S3Gateway/ListBuckets',
                request_serializer=s3g__pb2.ListBucketsRequest.SerializeToString,
                response_deserializer=s3g__pb2.ListBucketsResponse.FromString,
                )
        self.PutObject = channel.unary_unary(
                '/aveplen.s3g.S3Gateway/PutObject',
                request_serializer=s3g__pb2.PutObjectRequest.SerializeToString,
                response_deserializer=s3g__pb2.PutObjectResponse.FromString,
                )
        self.GetObject = channel.unary_unary(
                '/aveplen.s3g.S3Gateway/GetObject',
                request_serializer=s3g__pb2.GetObjectRequest.SerializeToString,
                response_deserializer=s3g__pb2.GetObjectResponse.FromString,
                )


class S3GatewayServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateBucket(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteBucket(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListBuckets(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PutObject(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetObject(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_S3GatewayServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateBucket': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateBucket,
                    request_deserializer=s3g__pb2.CreateBucketRequest.FromString,
                    response_serializer=s3g__pb2.CreateBucketResponse.SerializeToString,
            ),
            'DeleteBucket': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteBucket,
                    request_deserializer=s3g__pb2.DeleteBucketRequest.FromString,
                    response_serializer=s3g__pb2.DeleteBucketResponse.SerializeToString,
            ),
            'ListBuckets': grpc.unary_unary_rpc_method_handler(
                    servicer.ListBuckets,
                    request_deserializer=s3g__pb2.ListBucketsRequest.FromString,
                    response_serializer=s3g__pb2.ListBucketsResponse.SerializeToString,
            ),
            'PutObject': grpc.unary_unary_rpc_method_handler(
                    servicer.PutObject,
                    request_deserializer=s3g__pb2.PutObjectRequest.FromString,
                    response_serializer=s3g__pb2.PutObjectResponse.SerializeToString,
            ),
            'GetObject': grpc.unary_unary_rpc_method_handler(
                    servicer.GetObject,
                    request_deserializer=s3g__pb2.GetObjectRequest.FromString,
                    response_serializer=s3g__pb2.GetObjectResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'aveplen.s3g.S3Gateway', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class S3Gateway(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateBucket(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/aveplen.s3g.S3Gateway/CreateBucket',
            s3g__pb2.CreateBucketRequest.SerializeToString,
            s3g__pb2.CreateBucketResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteBucket(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/aveplen.s3g.S3Gateway/DeleteBucket',
            s3g__pb2.DeleteBucketRequest.SerializeToString,
            s3g__pb2.DeleteBucketResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListBuckets(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/aveplen.s3g.S3Gateway/ListBuckets',
            s3g__pb2.ListBucketsRequest.SerializeToString,
            s3g__pb2.ListBucketsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PutObject(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/aveplen.s3g.S3Gateway/PutObject',
            s3g__pb2.PutObjectRequest.SerializeToString,
            s3g__pb2.PutObjectResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetObject(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/aveplen.s3g.S3Gateway/GetObject',
            s3g__pb2.GetObjectRequest.SerializeToString,
            s3g__pb2.GetObjectResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)