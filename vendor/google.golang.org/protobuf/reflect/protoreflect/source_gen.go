// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-protos. DO NOT EDIT.

package protoreflect

func (p *SourcePath) appendFileDescriptorProto(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name", nil)
	case 2:
		b = p.appendSingularField(b, "package", nil)
	case 3:
		b = p.appendRepeatedField(b, "dependency", nil)
	case 10:
		b = p.appendRepeatedField(b, "public_dependency", nil)
	case 11:
		b = p.appendRepeatedField(b, "weak_dependency", nil)
	case 4:
		b = p.appendRepeatedField(b, "message_type", (*SourcePath).appendDescriptorProto)
	case 5:
		b = p.appendRepeatedField(b, "enum_type", (*SourcePath).appendEnumDescriptorProto)
	case 6:
		b = p.appendRepeatedField(b, "service", (*SourcePath).appendServiceDescriptorProto)
	case 7:
		b = p.appendRepeatedField(b, "extension", (*SourcePath).appendFieldDescriptorProto)
	case 8:
		b = p.appendSingularField(b, "options", (*SourcePath).appendFileOptions)
	case 9:
		b = p.appendSingularField(b, "source_code_info", (*SourcePath).appendSourceCodeInfo)
	case 12:
		b = p.appendSingularField(b, "syntax", nil)
	case 14:
		b = p.appendSingularField(b, "edition", nil)
	}
	return b
}

func (p *SourcePath) appendDescriptorProto(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name", nil)
	case 2:
		b = p.appendRepeatedField(b, "field", (*SourcePath).appendFieldDescriptorProto)
	case 6:
		b = p.appendRepeatedField(b, "extension", (*SourcePath).appendFieldDescriptorProto)
	case 3:
		b = p.appendRepeatedField(b, "nested_type", (*SourcePath).appendDescriptorProto)
	case 4:
		b = p.appendRepeatedField(b, "enum_type", (*SourcePath).appendEnumDescriptorProto)
	case 5:
		b = p.appendRepeatedField(b, "extension_range", (*SourcePath).appendDescriptorProto_ExtensionRange)
	case 8:
		b = p.appendRepeatedField(b, "oneof_decl", (*SourcePath).appendOneofDescriptorProto)
	case 7:
		b = p.appendSingularField(b, "options", (*SourcePath).appendMessageOptions)
	case 9:
		b = p.appendRepeatedField(b, "reserved_range", (*SourcePath).appendDescriptorProto_ReservedRange)
	case 10:
		b = p.appendRepeatedField(b, "reserved_name", nil)
	}
	return b
}

func (p *SourcePath) appendEnumDescriptorProto(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name", nil)
	case 2:
		b = p.appendRepeatedField(b, "value", (*SourcePath).appendEnumValueDescriptorProto)
	case 3:
		b = p.appendSingularField(b, "options", (*SourcePath).appendEnumOptions)
	case 4:
		b = p.appendRepeatedField(b, "reserved_range", (*SourcePath).appendEnumDescriptorProto_EnumReservedRange)
	case 5:
		b = p.appendRepeatedField(b, "reserved_name", nil)
	}
	return b
}

func (p *SourcePath) appendServiceDescriptorProto(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name", nil)
	case 2:
		b = p.appendRepeatedField(b, "method", (*SourcePath).appendMethodDescriptorProto)
	case 3:
		b = p.appendSingularField(b, "options", (*SourcePath).appendServiceOptions)
	}
	return b
}

func (p *SourcePath) appendFieldDescriptorProto(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name", nil)
	case 3:
		b = p.appendSingularField(b, "number", nil)
	case 4:
		b = p.appendSingularField(b, "label", nil)
	case 5:
		b = p.appendSingularField(b, "type", nil)
	case 6:
		b = p.appendSingularField(b, "type_name", nil)
	case 2:
		b = p.appendSingularField(b, "extendee", nil)
	case 7:
		b = p.appendSingularField(b, "default_value", nil)
	case 9:
		b = p.appendSingularField(b, "oneof_index", nil)
	case 10:
		b = p.appendSingularField(b, "json_name", nil)
	case 8:
		b = p.appendSingularField(b, "options", (*SourcePath).appendFieldOptions)
	case 17:
		b = p.appendSingularField(b, "proto3_optional", nil)
	}
	return b
}

func (p *SourcePath) appendFileOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "java_package", nil)
	case 8:
		b = p.appendSingularField(b, "java_outer_classname", nil)
	case 10:
		b = p.appendSingularField(b, "java_multiple_files", nil)
	case 20:
		b = p.appendSingularField(b, "java_generate_equals_and_hash", nil)
	case 27:
		b = p.appendSingularField(b, "java_string_check_utf8", nil)
	case 9:
		b = p.appendSingularField(b, "optimize_for", nil)
	case 11:
		b = p.appendSingularField(b, "go_package", nil)
	case 16:
		b = p.appendSingularField(b, "cc_generic_services", nil)
	case 17:
		b = p.appendSingularField(b, "java_generic_services", nil)
	case 18:
		b = p.appendSingularField(b, "py_generic_services", nil)
	case 23:
		b = p.appendSingularField(b, "deprecated", nil)
	case 31:
		b = p.appendSingularField(b, "cc_enable_arenas", nil)
	case 36:
		b = p.appendSingularField(b, "objc_class_prefix", nil)
	case 37:
		b = p.appendSingularField(b, "csharp_namespace", nil)
	case 39:
		b = p.appendSingularField(b, "swift_prefix", nil)
	case 40:
		b = p.appendSingularField(b, "php_class_prefix", nil)
	case 41:
		b = p.appendSingularField(b, "php_namespace", nil)
	case 44:
		b = p.appendSingularField(b, "php_metadata_namespace", nil)
	case 45:
		b = p.appendSingularField(b, "ruby_package", nil)
	case 50:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	}
	return b
}

func (p *SourcePath) appendSourceCodeInfo(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendRepeatedField(b, "location", (*SourcePath).appendSourceCodeInfo_Location)
	}
	return b
}

func (p *SourcePath) appendDescriptorProto_ExtensionRange(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "start", nil)
	case 2:
		b = p.appendSingularField(b, "end", nil)
	case 3:
		b = p.appendSingularField(b, "options", (*SourcePath).appendExtensionRangeOptions)
	}
	return b
}

func (p *SourcePath) appendOneofDescriptorProto(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name", nil)
	case 2:
		b = p.appendSingularField(b, "options", (*SourcePath).appendOneofOptions)
	}
	return b
}

func (p *SourcePath) appendMessageOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "message_set_wire_format", nil)
	case 2:
		b = p.appendSingularField(b, "no_standard_descriptor_accessor", nil)
	case 3:
		b = p.appendSingularField(b, "deprecated", nil)
	case 7:
		b = p.appendSingularField(b, "map_entry", nil)
	case 11:
		b = p.appendSingularField(b, "deprecated_legacy_json_field_conflicts", nil)
	case 12:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	}
	return b
}

func (p *SourcePath) appendDescriptorProto_ReservedRange(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "start", nil)
	case 2:
		b = p.appendSingularField(b, "end", nil)
	}
	return b
}

func (p *SourcePath) appendEnumValueDescriptorProto(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name", nil)
	case 2:
		b = p.appendSingularField(b, "number", nil)
	case 3:
		b = p.appendSingularField(b, "options", (*SourcePath).appendEnumValueOptions)
	}
	return b
}

func (p *SourcePath) appendEnumOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 2:
		b = p.appendSingularField(b, "allow_alias", nil)
	case 3:
		b = p.appendSingularField(b, "deprecated", nil)
	case 6:
		b = p.appendSingularField(b, "deprecated_legacy_json_field_conflicts", nil)
	case 7:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	}
	return b
}

func (p *SourcePath) appendEnumDescriptorProto_EnumReservedRange(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "start", nil)
	case 2:
		b = p.appendSingularField(b, "end", nil)
	}
	return b
}

func (p *SourcePath) appendMethodDescriptorProto(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name", nil)
	case 2:
		b = p.appendSingularField(b, "input_type", nil)
	case 3:
		b = p.appendSingularField(b, "output_type", nil)
	case 4:
		b = p.appendSingularField(b, "options", (*SourcePath).appendMethodOptions)
	case 5:
		b = p.appendSingularField(b, "client_streaming", nil)
	case 6:
		b = p.appendSingularField(b, "server_streaming", nil)
	}
	return b
}

func (p *SourcePath) appendServiceOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 34:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 33:
		b = p.appendSingularField(b, "deprecated", nil)
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	}
	return b
}

func (p *SourcePath) appendFieldOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "ctype", nil)
	case 2:
		b = p.appendSingularField(b, "packed", nil)
	case 6:
		b = p.appendSingularField(b, "jstype", nil)
	case 5:
		b = p.appendSingularField(b, "lazy", nil)
	case 15:
		b = p.appendSingularField(b, "unverified_lazy", nil)
	case 3:
		b = p.appendSingularField(b, "deprecated", nil)
	case 10:
		b = p.appendSingularField(b, "weak", nil)
	case 16:
		b = p.appendSingularField(b, "debug_redact", nil)
	case 17:
		b = p.appendSingularField(b, "retention", nil)
	case 19:
		b = p.appendRepeatedField(b, "targets", nil)
	case 20:
		b = p.appendRepeatedField(b, "edition_defaults", (*SourcePath).appendFieldOptions_EditionDefault)
	case 21:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 22:
		b = p.appendSingularField(b, "feature_support", (*SourcePath).appendFieldOptions_FeatureSupport)
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	}
	return b
}

func (p *SourcePath) appendFeatureSet(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "field_presence", nil)
	case 2:
		b = p.appendSingularField(b, "enum_type", nil)
	case 3:
		b = p.appendSingularField(b, "repeated_field_encoding", nil)
	case 4:
		b = p.appendSingularField(b, "utf8_validation", nil)
	case 5:
		b = p.appendSingularField(b, "message_encoding", nil)
	case 6:
		b = p.appendSingularField(b, "json_format", nil)
	case 7:
		b = p.appendSingularField(b, "enforce_naming_style", nil)
	}
	return b
}

func (p *SourcePath) appendUninterpretedOption(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 2:
		b = p.appendRepeatedField(b, "name", (*SourcePath).appendUninterpretedOption_NamePart)
	case 3:
		b = p.appendSingularField(b, "identifier_value", nil)
	case 4:
		b = p.appendSingularField(b, "positive_int_value", nil)
	case 5:
		b = p.appendSingularField(b, "negative_int_value", nil)
	case 6:
		b = p.appendSingularField(b, "double_value", nil)
	case 7:
		b = p.appendSingularField(b, "string_value", nil)
	case 8:
		b = p.appendSingularField(b, "aggregate_value", nil)
	}
	return b
}

func (p *SourcePath) appendSourceCodeInfo_Location(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendRepeatedField(b, "path", nil)
	case 2:
		b = p.appendRepeatedField(b, "span", nil)
	case 3:
		b = p.appendSingularField(b, "leading_comments", nil)
	case 4:
		b = p.appendSingularField(b, "trailing_comments", nil)
	case 6:
		b = p.appendRepeatedField(b, "leading_detached_comments", nil)
	}
	return b
}

func (p *SourcePath) appendExtensionRangeOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	case 2:
		b = p.appendRepeatedField(b, "declaration", (*SourcePath).appendExtensionRangeOptions_Declaration)
	case 50:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 3:
		b = p.appendSingularField(b, "verification", nil)
	}
	return b
}

func (p *SourcePath) appendOneofOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	}
	return b
}

func (p *SourcePath) appendEnumValueOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "deprecated", nil)
	case 2:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 3:
		b = p.appendSingularField(b, "debug_redact", nil)
	case 4:
		b = p.appendSingularField(b, "feature_support", (*SourcePath).appendFieldOptions_FeatureSupport)
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	}
	return b
}

func (p *SourcePath) appendMethodOptions(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 33:
		b = p.appendSingularField(b, "deprecated", nil)
	case 34:
		b = p.appendSingularField(b, "idempotency_level", nil)
	case 35:
		b = p.appendSingularField(b, "features", (*SourcePath).appendFeatureSet)
	case 999:
		b = p.appendRepeatedField(b, "uninterpreted_option", (*SourcePath).appendUninterpretedOption)
	}
	return b
}

func (p *SourcePath) appendFieldOptions_EditionDefault(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 3:
		b = p.appendSingularField(b, "edition", nil)
	case 2:
		b = p.appendSingularField(b, "value", nil)
	}
	return b
}

func (p *SourcePath) appendFieldOptions_FeatureSupport(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "edition_introduced", nil)
	case 2:
		b = p.appendSingularField(b, "edition_deprecated", nil)
	case 3:
		b = p.appendSingularField(b, "deprecation_warning", nil)
	case 4:
		b = p.appendSingularField(b, "edition_removed", nil)
	}
	return b
}

func (p *SourcePath) appendUninterpretedOption_NamePart(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "name_part", nil)
	case 2:
		b = p.appendSingularField(b, "is_extension", nil)
	}
	return b
}

func (p *SourcePath) appendExtensionRangeOptions_Declaration(b []byte) []byte {
	if len(*p) == 0 {
		return b
	}
	switch (*p)[0] {
	case 1:
		b = p.appendSingularField(b, "number", nil)
	case 2:
		b = p.appendSingularField(b, "full_name", nil)
	case 3:
		b = p.appendSingularField(b, "type", nil)
	case 5:
		b = p.appendSingularField(b, "reserved", nil)
	case 6:
		b = p.appendSingularField(b, "repeated", nil)
	}
	return b
}
