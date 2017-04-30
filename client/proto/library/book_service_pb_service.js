// package: library
// file: proto/library/book_service.proto

var jspb = require("google-protobuf");
var proto_library_book_service_pb = require("../../proto/library/book_service_pb");
var BookService = {
  serviceName: "library.BookService"
};
BookService.GetBook = {
  methodName: "GetBook",
  service: BookService,
  requestStream: false,
  responseStream: false,
  requestType: proto_library_book_service_pb.GetBookRequest,
  responseType: proto_library_book_service_pb.Book
};
BookService.QueryBooks = {
  methodName: "QueryBooks",
  service: BookService,
  requestStream: false,
  responseStream: true,
  requestType: proto_library_book_service_pb.QueryBooksRequest,
  responseType: proto_library_book_service_pb.Book
};
module.exports = {
  BookService: BookService,
};

