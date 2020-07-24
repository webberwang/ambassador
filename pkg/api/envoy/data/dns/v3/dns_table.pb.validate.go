// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/dns/v3/dns_table.proto

package envoy_data_dns_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _dns_table_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DnsTable with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DnsTable) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ExternalRetryCount

	if len(m.GetVirtualDomains()) < 1 {
		return DnsTableValidationError{
			field:  "VirtualDomains",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetVirtualDomains() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTableValidationError{
					field:  fmt.Sprintf("VirtualDomains[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetKnownSuffixes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTableValidationError{
					field:  fmt.Sprintf("KnownSuffixes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DnsTableValidationError is the validation error returned by
// DnsTable.Validate if the designated constraints aren't met.
type DnsTableValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTableValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTableValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTableValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTableValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTableValidationError) ErrorName() string { return "DnsTableValidationError" }

// Error satisfies the builtin error interface
func (e DnsTableValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTableValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTableValidationError{}

// Validate checks the field values on DnsTable_AddressList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_AddressList) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetAddress()) < 1 {
		return DnsTable_AddressListValidationError{
			field:  "Address",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetAddress() {
		_, _ = idx, item

		if utf8.RuneCountInString(item) < 3 {
			return DnsTable_AddressListValidationError{
				field:  fmt.Sprintf("Address[%v]", idx),
				reason: "value length must be at least 3 runes",
			}
		}

	}

	return nil
}

// DnsTable_AddressListValidationError is the validation error returned by
// DnsTable_AddressList.Validate if the designated constraints aren't met.
type DnsTable_AddressListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_AddressListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_AddressListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_AddressListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_AddressListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_AddressListValidationError) ErrorName() string {
	return "DnsTable_AddressListValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_AddressListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_AddressList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_AddressListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_AddressListValidationError{}

// Validate checks the field values on DnsTable_DnsEndpoint with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_DnsEndpoint) Validate() error {
	if m == nil {
		return nil
	}

	switch m.EndpointConfig.(type) {

	case *DnsTable_DnsEndpoint_AddressList:

		if v, ok := interface{}(m.GetAddressList()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsTable_DnsEndpointValidationError{
					field:  "AddressList",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return DnsTable_DnsEndpointValidationError{
			field:  "EndpointConfig",
			reason: "value is required",
		}

	}

	return nil
}

// DnsTable_DnsEndpointValidationError is the validation error returned by
// DnsTable_DnsEndpoint.Validate if the designated constraints aren't met.
type DnsTable_DnsEndpointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsEndpointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsEndpointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsEndpointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsEndpointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsEndpointValidationError) ErrorName() string {
	return "DnsTable_DnsEndpointValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsEndpointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsEndpoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsEndpointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsEndpointValidationError{}

// Validate checks the field values on DnsTable_DnsVirtualDomain with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsTable_DnsVirtualDomain) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 2 {
		return DnsTable_DnsVirtualDomainValidationError{
			field:  "Name",
			reason: "value length must be at least 2 runes",
		}
	}

	if !_DnsTable_DnsVirtualDomain_Name_Pattern.MatchString(m.GetName()) {
		return DnsTable_DnsVirtualDomainValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
		}
	}

	if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsTable_DnsVirtualDomainValidationError{
				field:  "Endpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if d := m.GetAnswerTtl(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return DnsTable_DnsVirtualDomainValidationError{
				field:  "AnswerTtl",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return DnsTable_DnsVirtualDomainValidationError{
				field:  "AnswerTtl",
				reason: "value must be greater than 0s",
			}
		}

	}

	return nil
}

// DnsTable_DnsVirtualDomainValidationError is the validation error returned by
// DnsTable_DnsVirtualDomain.Validate if the designated constraints aren't met.
type DnsTable_DnsVirtualDomainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsTable_DnsVirtualDomainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsTable_DnsVirtualDomainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsTable_DnsVirtualDomainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsTable_DnsVirtualDomainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsTable_DnsVirtualDomainValidationError) ErrorName() string {
	return "DnsTable_DnsVirtualDomainValidationError"
}

// Error satisfies the builtin error interface
func (e DnsTable_DnsVirtualDomainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsTable_DnsVirtualDomain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsTable_DnsVirtualDomainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsTable_DnsVirtualDomainValidationError{}

var _DnsTable_DnsVirtualDomain_Name_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")
