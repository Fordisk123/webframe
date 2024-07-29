package generator

type Option func(gen *Generator)

// UseJSONNamesForFields. if disabled, the original proto name will be used for generating OpenAPI definitions
func UseJSONNamesForFields(b bool) Option {
	return func(gen *Generator) {
		gen.reg.SetUseJSONNamesForFields(b)
	}
}

// RecursiveDepth. maximum recursion count allowed for a field type
func RecursiveDepth(depth int) Option {
	return func(gen *Generator) {
		gen.reg.SetRecursiveDepth(depth)
	}
}

// EnumsAsInts. whether to render enum values as integers, as opposed to string values
func EnumsAsInts(b bool) Option {
	return func(gen *Generator) {
		gen.reg.SetEnumsAsInts(b)
	}
}

// MergeFileName. target OpenAPI file name prefix after merge
func MergeFileName(name string) Option {
	return func(gen *Generator) {
		gen.reg.SetMergeFileName(name)
	}
}

// DisableDefaultErrors. if set, disables generation of default errors. This is useful if you have defined custom error handling
func DisableDefaultErrors(b bool) Option {
	return func(gen *Generator) {
		gen.reg.SetDisableDefaultErrors(b)
	}
}
