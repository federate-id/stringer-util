# stringer-util
A Go package for taking Stringer (has `func (*T) String() string` declared) structs and allowing you to annotate fields for inclusion, exclusion, masked value ("" or "*****") or the type (useful for any or interface{}) when printing/logging.
