package main

import (
	"bytes"
	"strings"

	google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// ProtoFile reprpesents a parsed proto file
type ProtoFile struct {
	PackageName string
	CsNamespace string
	Messages    []*ProtoMessage
	Services    []*ProtoService
}

// ProtoMessage represents a parsed message in a proto file
type ProtoMessage struct {
	Name       string
	PascalName string
}

// ProtoService represents a parsed service in a proto file
type ProtoService struct {
	Name       string
	PascalName string
	Methods    []*ProtoMethod
}

// ProtoMethod represents a parsed method in a proto service
type ProtoMethod struct {
	Name         string
	PascalName   string
	InputStream  bool
	Input        *ProtoMessage
	OutputStream bool
	Output       *ProtoMessage
}

//ProtoAst transforms a FileDescriptor to an AST that can be used for code generation
func ProtoAst(file *google_protobuf.FileDescriptorProto) *ProtoFile {

	pkg := &ProtoFile{}
	pkg.PackageName = file.GetPackage()
	pkg.CsNamespace = file.Options.GetCsharpNamespace()
	messages := make(map[string]*ProtoMessage)
	for _, message := range file.GetMessageType() {
		m := &ProtoMessage{}
		m.Name = message.GetName()
		m.PascalName = MakeFirstLowerCase(m.Name)
		pkg.Messages = append(pkg.Messages, m)
		messages[m.Name] = m
	}

	for _, service := range file.GetService() {
		s := &ProtoService{}
		s.Name = service.GetName()
		s.PascalName = MakeFirstLowerCase(s.Name)
		pkg.Services = append(pkg.Services, s)

		for _, method := range service.GetMethod() {
			m := &ProtoMethod{}
			m.Name = method.GetName()
			m.PascalName = MakeFirstLowerCase(m.Name)
			//		m.InputStream = *method.ClientStreaming
			//		m.OutputStream = *method.ServerStreaming
			input := removePackagePrefix(method.GetInputType(), pkg.PackageName)
			output := removePackagePrefix(method.GetOutputType(), pkg.PackageName)
			m.Input = messages[input]
			m.Output = messages[output]
			s.Methods = append(s.Methods, m)
		}
	}
	return pkg
}

// MakeFirstLowerCase makes the first character in a string lower case
func MakeFirstLowerCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
