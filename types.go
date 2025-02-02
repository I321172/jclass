package class

import (
	"io"
)

// ClassFile represents a single class file as specified in:
// http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
type ClassFile struct {
	// Magic number found in all valid Java class files.
	// This will always equal 0xCAFEBABE
	Magic uint32

	// Major.Minor denotes the class file version, that
	// has to be supported by the executing JVM.
	MinorVersion uint16
	MajorVersion uint16

	// The constant pool is a table of structures
	// representing various string constants, class,
	// interface & field names and other constants that
	// are referred to in the class file structure.
	// The reason we keep the ConstPoolSize instead of
	// just discarding it like with Fields and Methods,
	// is one of the most annoying bugs I've come across:
	// In the Java class file format, 64-bit constants
	// use up two spaces everywhere. So finding the size
	// of the constant pool is not just len(ConstantPool).
	ConstPoolSize uint16
	ConstantPool

	// AccessFlags is a mask of flags used to denote
	// access permissions and properties of this class
	// or interface.
	AccessFlags

	// Index into the constant pool, where you should
	// find a CONSTANT_Class_info struct that describes
	// this class.
	ThisClass ConstPoolIndex

	// Index into the constant pool or zero, where you
	// should find a CONSTANT_Class_info struct that
	// describes this class' super class.
	// If SuperClass is zero, then this class must
	// represent the Object class.
	// For an interface, the corresponding value in the
	// constant pool, must represent the Object class.
	SuperClass ConstPoolIndex

	// Interfaces contains indexes into the constant pool,
	// where every referenced entry describes a
	// CONSTANT_Class_info struct representing a direct
	// super-interface of this class or interface.
	Interfaces []ConstPoolIndex

	// Fields contains indexes into the constant pool,
	// referencing field_info structs, giving a complete
	// description of a field in this class or interface.
	// The Fields table only contains fields declared in
	// this class or interface, not any inherited ones.
	Fields []*Field

	// Methods contains method_info structs describing
	// a method of this class or interface.
	// If neither METHOD_ACC_NATIVE or METHOD_ACC_ABSTRACT
	// flags are set, the corresponding code for the method
	// will also be supplied.
	Methods []*Method

	// Attributes describes properties of this class or
	// interface through attribute_info structs.
	Attributes
}

// Dumper All Attributes and Constants, plus the actual class file
// have to fulfill this interface. As you can guess, it's
// used when writing the class file back to its original
// (binary) format.
type Dumper interface {
	Dump(io.Writer) error
}

// Attribute Attributes add extra/meta info to ClassFile, Field,
// Method and Code structs. Any JVM implementation or
// Java compiler, may create its own/new attribute(s).
// Though these should not effect the sematics of the program.
// http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7
type Attribute interface {
	Dumper

	Read(io.Reader, ConstantPool) error

	// GetTag Think of an Attribute value as a discriminated union.
	GetTag() AttributeType

	// UnknownAttr In order to actually access the fields of an attribute
	// you would need a type assertion in your code. But since
	// the Java spec is quite precise on when you can expect
	// what type of attribute (in a valid class file), we can
	// provide "safe" implementations of methods for casting
	// the values, that do not require type assertions.
	// You shouldn't call any of the following functions if you
	// aren't sure about what type an Attribute actually has,
	// since if you are wrong, the function will panic.
	UnknownAttr() *UnknownAttr
	ConstantValue() *ConstantValue
	Code() *Code
	// Exceptions StackMapTable() *StackMapTable
	Exceptions() *Exceptions
	InnerClasses() *InnerClasses
	EnclosingMethod() *EnclosingMethod
	Synthetic() *Synthetic
	Signature() *Signature
	SourceFile() *SourceFile
	SourceDebugExtension() *SourceDebugExtension
	LineNumberTable() *LineNumberTable
	LocalVariableTable() *LocalVariableTable
	LocalVariableTypeTable() *LocalVariableTypeTable
	Deprecated() *Deprecated
	// BootstrapMethods : RuntimeVisibleAnnotations() *RuntimeVisibleAnnotations
	// RuntimeInvisibleAnnotations() *RuntimeInvisibleAnnotations
	// RuntimeVisibleParameterAnnotations() *RuntimeVisibleParameterAnnotations
	// RuntimeInvisibleParameterAnnotations() *RuntimeInvisibleParameterAnnotations
	// AnnotationDefault() *AnnotationDefault
	BootstrapMethods() *BootstrapMethods
}

// Constant Constants reside in a class files constant pool and
// are used in various places in a class file. They can
// describe variable or method type signatures names of
// variables or other classes. The pool also contains all
// integer and string constants that can be found in the
// code (besides when the instruction lconst_1 or the like
// is used).
type Constant interface {
	Dumper

	Read(io.Reader) error

	GetTag() ConstantType

	// Class In order to actually access the fields of a constant
	// you would need a type assertion in your code. But since
	// the Java spec is quite precise on when you can expect
	// what type of constant (in a valid class file), we can
	// provide "safe" implementations of methods for casting
	// the values, that do not require type assertions.
	// You shouldn't call any of the following functions if you
	// aren't sure about what type a Constant actually has,
	// since if you are wrong, the function will panic.
	Class() *ClassRef
	Field() *FieldRef
	Method() *MethodRef
	InterfaceMethod() *InterfaceMethodRef
	StringRef() *StringRef
	Integer() *IntegerRef
	Float() *FloatRef
	Long() *LongRef
	Double() *DoubleRef
	NameAndType() *NameAndTypeRef
	UTF8() *UTF8Ref
	MethodHandle() *MethodHandleRef
	MethodType() *MethodTypeRef
	InvokeDynamic() *InvokeDynamicRef
}

// Attributes Describes a set of attributes as you would find them in a
// ClassFile, Method, Field or Code struct.
type Attributes []Attribute

// ConstantPool Every (valid) class will have a constant pool in which it
// stores constants referenced in the rest of the class file.
type ConstantPool []Constant

// ConstPoolIndex Index into the/a constant pool. This type is used so that
// you know, when a uint16 is actually an index.
type ConstPoolIndex uint16

// AccessFlags are a mask of flags that can determine things
// like access privileges, cache options or implementation
// details. There is no separate type for each access flag
// (i.e. for nested classes, methods, fields, etc.), but
// each of those separate types there are different constants
// describing those access flags. See METHOD_ACC_*, FIELD_ACC_*
// CLASS_ACC_*, and NESTED_CLASS_ACC_*.
type AccessFlags uint16
