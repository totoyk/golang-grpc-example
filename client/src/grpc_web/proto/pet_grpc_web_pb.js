/**
 * @fileoverview gRPC-Web generated client stub for pet
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.1
// 	protoc              v3.21.6
// source: proto/pet.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.pet = require('./pet_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pet.PetServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pet.PetServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.pet.GetPetReply>}
 */
const methodDescriptor_PetService_GetMyPet = new grpc.web.MethodDescriptor(
  '/pet.PetService/GetMyPet',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.pet.GetPetReply,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pet.GetPetReply.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pet.GetPetReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pet.GetPetReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pet.PetServiceClient.prototype.getMyPet =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pet.PetService/GetMyPet',
      request,
      metadata || {},
      methodDescriptor_PetService_GetMyPet,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pet.GetPetReply>}
 *     Promise that resolves to the response
 */
proto.pet.PetServicePromiseClient.prototype.getMyPet =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pet.PetService/GetMyPet',
      request,
      metadata || {},
      methodDescriptor_PetService_GetMyPet);
};


module.exports = proto.pet;
