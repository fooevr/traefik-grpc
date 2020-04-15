package main

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc/protoparse"
	"os"
	"path/filepath"
)


func main()  {
	files := []string{}
	paths := []string{}

	filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error)error{
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	x := protoparse.Parser{ImportPaths: paths}
	fs, _ := x.ParseFiles(files...)
	set := new(descriptor.FileDescriptorSet)
	for _, item := range fs{
		set.File = append(set.File, item.AsFileDescriptorProto())
	}

	bts, _ := proto.Marshal(set)
	fmt.Println(base64.StdEncoding.EncodeToString(bts))
}
