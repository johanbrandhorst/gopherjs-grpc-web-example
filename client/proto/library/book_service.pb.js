/**
 * @fileoverview gRPC Web JS generated client stub for library
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!


goog.provide('proto.library.BookServiceClient');

goog.require('grpc.web.GatewayClientBase');
goog.require('grpc.web.AbstractClientBase');
goog.require('grpc.web.ClientReadableStream');
goog.require('proto.library.Book');
goog.require('proto.library.GetBookRequest');
goog.require('proto.library.QueryBooksRequest');



goog.scope(function() {

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @constructor
 * @struct
 * @final
 */
proto.library.BookServiceClient =
    function(hostname, credentials, options) {
  /**
   * @private @const {!grpc.web.GatewayClientBase} The client
   */
  this.client_ = new grpc.web.GatewayClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.library.GetBookRequest,
 *   !proto.library.Book>}
 */
const methodInfo_GetBook = new grpc.web.AbstractClientBase.MethodInfo(
  proto.library.Book,
  /** @param {!proto.library.GetBookRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.library.Book.deserializeBinary
);


/**
 * @param {!proto.library.GetBookRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.library.Book)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.library.Book>|undefined}
 *     The XHR Node Readable Stream
 */
proto.library.BookServiceClient.prototype.getBook =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/library.BookService/GetBook',
      request,
      metadata,
      methodInfo_GetBook,
      callback);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.library.QueryBooksRequest,
 *   !proto.library.Book>}
 */
const methodInfo_QueryBooks = new grpc.web.AbstractClientBase.MethodInfo(
  proto.library.Book,
  /** @param {!proto.library.QueryBooksRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.library.Book.deserializeBinary
);


/**
 * @param {!proto.library.QueryBooksRequest} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.library.Book>}
 *     The XHR Node Readable Stream
 */
proto.library.BookServiceClient.prototype.queryBooks =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/library.BookService/QueryBooks',
      request,
      metadata,
      methodInfo_QueryBooks);
};


}); // goog.scope

