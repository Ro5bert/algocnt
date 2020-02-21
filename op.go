package algocnt

// OpType represents a type of primitive operation.
// Users can create their own OpTypes or use the ones provided by this package.
// OpTypes must not begin with an underscore.
type OpType string

const (
	All        OpType = "_All"
	enterScope OpType = "_enterScope"
	exitScope  OpType = "_exitScope"

	// Assignment primitive operation.
	A OpType = "Assignment"

	// Comparison primitive operation
	C OpType = "Comparison"

	// Boolean expression primitive operation
	B OpType = "BoolExpr"

	// Array index primitive operation
	I OpType = "ArrayIndex"

	// Record selection primitive operation
	R OpType = "RecordSelect"

	// Addition/subtraction primitive operation
	S OpType = "Add/Sub"

	// Multiplication/division primitive operation
	D OpType = "Mult/Div"

	// Trigonometric function primitive operation
	T OpType = "Trig"

	// Method/function/procedure/routine call primitive operation
	M OpType = "Call"
)

// op represents an primitive operation of a particular OpType along with a comment describing the operation.
type op struct {
	opType  OpType
	comment string
}
