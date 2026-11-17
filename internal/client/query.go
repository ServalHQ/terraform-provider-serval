package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// QueryFromMessage encodes a public API list/search request as query
// parameters keyed by proto JSON name. Only populated fields are sent and
// repeated fields are comma-joined.
func QueryFromMessage(m proto.Message) (url.Values, error) {
	if e, ok := m.(explicitFields); ok {
		return e.query()
	}
	q := url.Values{}
	var err error
	m.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.IsMap() {
			err = fmt.Errorf("%s: map fields cannot be sent as query parameters", fd.FullName())
			return false
		}
		if !fd.IsList() {
			var s string
			s, err = queryScalar(fd, v)
			q.Set(fd.JSONName(), s)
			return err == nil
		}
		list := v.List()
		parts := make([]string, list.Len())
		for i := range parts {
			if parts[i], err = queryScalar(fd, list.Get(i)); err != nil {
				return false
			}
		}
		q.Set(fd.JSONName(), strings.Join(parts, ","))
		return true
	})
	if err != nil {
		return nil, err
	}
	return q, nil
}

func queryScalar(fd protoreflect.FieldDescriptor, v protoreflect.Value) (string, error) {
	switch fd.Kind() {
	case protoreflect.StringKind:
		return v.String(), nil
	case protoreflect.BoolKind:
		return strconv.FormatBool(v.Bool()), nil
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind,
		protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return strconv.FormatInt(v.Int(), 10), nil
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return strconv.FormatUint(v.Uint(), 10), nil
	default:
		return "", fmt.Errorf("%s: %s fields cannot be sent as query parameters", fd.FullName(), fd.Kind())
	}
}

type explicitFields struct {
	proto.Message
	jsonNames []string
}

// WithExplicit sends the named scalar fields even when they hold their zero
// value, the way an optional parameter that was set is sent.
func WithExplicit(m proto.Message, jsonNames ...string) proto.Message {
	if len(jsonNames) == 0 {
		return m
	}
	return explicitFields{Message: m, jsonNames: jsonNames}
}

func (e explicitFields) query() (url.Values, error) {
	q, err := QueryFromMessage(e.Message)
	if err != nil {
		return nil, err
	}
	msg := e.ProtoReflect()
	for _, name := range e.jsonNames {
		fd := msg.Descriptor().Fields().ByJSONName(name)
		if fd == nil || fd.IsList() || fd.IsMap() {
			return nil, fmt.Errorf("%s has no scalar field %q", msg.Descriptor().FullName(), name)
		}
		if q.Has(name) {
			continue
		}
		s, err := queryScalar(fd, msg.Get(fd))
		if err != nil {
			return nil, err
		}
		q.Set(name, s)
	}
	return q, nil
}
