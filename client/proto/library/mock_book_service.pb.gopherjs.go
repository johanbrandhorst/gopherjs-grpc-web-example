// This is what protoc-gen-gopherjs should generate
package library

import (
	"github.com/gopherjs/gopherjs/js"

	// gRPC-web Bindings
	grpcweb "github.com/johanbrandhorst/gopherjs-grpc-web"
)

type Book struct {
	*js.Object
	Isbn   int64  `js:"isbn"`
	Title  string `js:"title"`
	Author string `js:"author"`
}

func NewBook(isbn int64, title, author string) *Book {
	x := &Book{
		Object: js.Global.Get("Object").New(),
	}

	x.Isbn = isbn
	x.Title = title
	x.Author = author

	return x
}

func (x *Book) Serialize() (rawBytes []byte, err error) {
	// Recover any thrown JS errors
	defer func() {
		e := recover()
		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
			err = e
		} else {
			panic(e)
		}
	}()

	writer := js.Global.Get("jspb").Get("BinaryWriter").New()
	if x.Isbn != 0 {
		writer.Call("writeInt64", 1, x.Isbn)
	}
	if x.Title != "" {
		writer.Call("writeString", 2, x.Title)
	}
	if x.Author != "" {
		writer.Call("writeString", 3, x.Author)
	}
	rawBytes = writer.Call("getResultBuffer").Interface().([]byte)

	return rawBytes, err
}

func (x *Book) Deserialize(rawBytes []byte) (err error) {
	// Recover any thrown JS errors
	defer func() {
		e := recover()
		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
			err = e
		} else {
			panic(e)
		}
	}()

	reader := js.Global.Get("jspb").Get("BinaryReader").New(rawBytes)
	x = &Book{
		Object: js.Global.Get("Object").New(),
	}

	// Read contents
	for {
		reader.Call("nextField")
		if reader.Call("isEndGroup").Bool() {
			break
		}

		switch reader.Call("getFieldNumber").Int() {
		case 1:
			x.Isbn = reader.Call("readInt64").Int64()
		case 2:
			x.Title = reader.Call("readString").String()
		case 3:
			x.Author = reader.Call("readString").String()
		default:
			reader.Call("skipField")
		}
	}

	return err
}

type GetBookRequest struct {
	*js.Object
	Isbn int64 `js:"isbn"`
}

func NewGetBookRequest(isbn int64) *GetBookRequest {
	x := &GetBookRequest{
		Object: js.Global.Get("Object").New(),
	}

	x.Isbn = isbn

	return x
}

func (x *GetBookRequest) Serialize() (rawBytes []byte, err error) {
	// Recover any thrown JS errors
	defer func() {
		e := recover()
		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
			err = e
		} else {
			panic(e)
		}
	}()

	writer := js.Global.Get("jspb").Get("BinaryWriter").New()
	if x.Isbn != 0 {
		writer.Call("writeInt64", 1, x.Isbn)
	}
	rawBytes = writer.Call("getResultBuffer").Interface().([]byte)

	return rawBytes, err
}

func (x *GetBookRequest) Deserialize(rawBytes []byte) (err error) {
	// Recover any thrown JS errors
	defer func() {
		e := recover()
		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
			err = e
		} else {
			panic(e)
		}
	}()

	reader := js.Global.Get("jspb").Get("BinaryReader").New(rawBytes)
	x = &GetBookRequest{
		Object: js.Global.Get("Object").New(),
	}

	// Read contents
	for {
		reader.Call("nextField")
		if reader.Call("isEndGroup").Bool() {
			break
		}

		switch reader.Call("getFieldNumber").Int() {
		case 1:
			x.Isbn = reader.Call("readInt64").Int64()
		default:
			reader.Call("skipField")
		}
	}

	return err
}

type QueryBooksRequest struct {
	*js.Object
	AuthorPrefix string `js:"authorPrefix"`
}

func NewQueryBooksRequest(authorPrefix string) *QueryBooksRequest {
	x := &QueryBooksRequest{
		Object: js.Global.Get("Object").New(),
	}

	x.AuthorPrefix = authorPrefix

	return x
}

func (x *QueryBooksRequest) Serialize() (rawBytes []byte, err error) {
	// Recover any thrown JS errors
	defer func() {
		e := recover()
		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
			err = e
		} else {
			panic(e)
		}
	}()

	writer := js.Global.Get("jspb").Get("BinaryWriter").New()
	if x.AuthorPrefix != "" {
		writer.Call("writeString", 1, x.AuthorPrefix)
	}
	rawBytes = writer.Call("getResultBuffer").Interface().([]byte)

	return rawBytes, err
}

func (x *QueryBooksRequest) Deserialize(rawBytes []byte) (err error) {
	// Recover any thrown JS errors
	defer func() {
		e := recover()
		if e == nil {
			return
		}

		if e, ok := e.(*js.Error); ok {
			err = e
		} else {
			panic(e)
		}
	}()

	reader := js.Global.Get("jspb").Get("BinaryReader").New(rawBytes)
	x = &QueryBooksRequest{
		Object: js.Global.Get("Object").New(),
	}

	// Read contents
	for {
		reader.Call("nextField")
		if reader.Call("isEndGroup").Bool() {
			break
		}

		switch reader.Call("getFieldNumber").Int() {
		case 1:
			x.AuthorPrefix = reader.Call("readString").String()
		default:
			reader.Call("skipField")
		}
	}

	return err
}

type BookServiceClient interface {
	GetBook(*GetBookRequest, ...grpcweb.CallOption) (*Book, error)
	QueryBooks(*QueryBooksRequest, ...grpcweb.CallOption) (BookService_QueryBooksClient, error)
}

type bookServiceClient struct {
	*js.Object
	hostname string
	client   *grpcweb.GatewayClientBase
}

type BookService_QueryBooksClient interface {
	Recv() (*Book, error)
}

type bookServiceQueryBooksClient struct {
	stream *grpcweb.StreamReader
}

func (x *bookServiceQueryBooksClient) Recv() (*Book, error) {
	m := new(Book)
	resp, err := x.stream.Recv()
	if err != nil {
		// Could be EOF, on success
		return nil, err
	}

	err = m.Deserialize(resp)

	return m, err
}

func NewBookServiceClient(hostname string) BookServiceClient {
	bsc := &bookServiceClient{
		Object: js.Global.Get("Object").New(),
	}
	bsc.hostname = hostname
	bsc.client = grpcweb.NewGatewayClientBase()
	return bsc
}

func (x *bookServiceClient) GetBook(in *GetBookRequest, opts ...grpcweb.CallOption) (*Book, error) {
	println("Calling RPC")
	resp, err := x.client.RPCCall(x.hostname+"/library.BookService/GetBook", in, opts...)
	if err != nil {
		return nil, err
	}

	println("Got reply")
	b := new(Book)
	err = b.Deserialize(resp)

	return b, err
}

func (x *bookServiceClient) QueryBooks(in *QueryBooksRequest, opts ...grpcweb.CallOption) (BookService_QueryBooksClient, error) {
	srv, err := x.client.ServerStreaming(x.hostname+"/library.BookService/QueryBooks", in, opts...)
	if err != nil {
		return nil, err
	}

	return &bookServiceQueryBooksClient{
		stream: srv,
	}, nil
}
