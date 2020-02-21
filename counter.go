package algocnt

import (
	"fmt"
	"regexp"
	"strings"
)

type stringStack []string

func (ps *stringStack) push(p string) {
	*ps = append(*ps, p)
}

func (ps *stringStack) pop() string {
	// This will panic if the stack is empty.
	p := (*ps)[len(*ps)-1]
	*ps = (*ps)[:len(*ps)-1]
	return p
}

func (ps stringStack) isEmpty() bool {
	return len(ps) == 0
}

// Counter counts primitive operations.
// The zero value of a Counter is ready to be used.
type Counter struct {
	ops    []op
	sstack stringStack
}

func (c *Counter) append(op op) {
	c.ops = append(c.ops, op)
}

// Addc adds a primitive operation to the Counter of the specified type with the given comment.
func (c *Counter) Addc(ot OpType, comment string) {
	if strings.HasPrefix(string(ot), "_") {
		panic("user-defined OpType must not start with '_'")
	} else if ot == "" {
		panic("user-defined OpType must not be the empty string")
	}
	c.append(op{ot, comment})
}

// Add adds a primitive operation to the Counter of the specified type with no comment.
func (c *Counter) Add(ot OpType) {
	c.Addc(ot, "")
}

// EnterScope enters a scope with the given name.
// Scopes can be used to group primitive operations in algorithms with distinct steps.
// After counting the primitive operations in an algorithm with a Counter, one can sum primitive operations based on
// what scope they are in (see CountPath).
// Each call to EnterScope must be matched with a call to ExitScope.
// Scope names must not contain a forward slash ("/").
func (c *Counter) EnterScope(scope string) {
	if strings.ContainsRune(scope, '/') {
		panic("EnterScope: scope name cannot contain '/'")
	}
	c.sstack.push(scope)
	c.append(op{enterScope, scope})
}

// ExitScope exits the most-recently entered scope (panics if not in any scope).
func (c *Counter) ExitScope() {
	if c.sstack.isEmpty() {
		panic("ExitScope: not in any scope")
	}
	c.sstack.pop()
	c.append(op{exitScope, ""})
}

// CountPath counts the primitive operations added to the Counter which have paths matching pathRegex and which have
// types in opTypes.
// An operation's path is formed by taking its containing scopes and separating them by forward slashes ("/").
// For example, a primitive operation with a path "A/B/C" means it is in a scope "C", which is in turn inside a scope
// "B", which is in turn inside a scope "A".
// Operations not in any scopes have the empty string as their path.
// If the special All OpType is in opTypes, primitive operations will be counted regardless of their type.
func (c *Counter) CountPath(pathRegex string, opTypes ...OpType) int {
	if !c.sstack.isEmpty() {
		panic("Count: not all scopes closed")
	}

	r := regexp.MustCompile(pathRegex)
	cnt := 0
	scope := ""
	doCount := []bool{r.MatchString("")}

	all := false
	for _, ot := range opTypes {
		if ot == All {
			all = true
			break
		}
	}

	for _, op := range c.ops {
		switch op.opType {
		case enterScope:
			scope += "/" + op.comment
			doCount = append(doCount, r.MatchString(scope)) // Push.
		case exitScope:
			scope = scope[:strings.LastIndex(scope, "/")]
			doCount = doCount[:len(doCount)-1] // Pop.
		default:
			if !doCount[len(doCount)-1] {
				break // Do not count it.
			}
			if all {
				// Count it.
				cnt++
				break
			}
			for _, ot := range opTypes {
				if ot == op.opType {
					// Count it.
					cnt++
					break
				}
			}
		}
	}

	return cnt
}

// Count counts the primitive operations added to the Counter similarly to CountPath, except there is no restriction on
// the operations' paths.
// (This is equivalent to CountPath("", opTypes...).)
func (c *Counter) Count(opTypes ...OpType) int {
	return c.CountPath("", opTypes...)
}

// Print out the operations added to the Counter to stdout.
func (c *Counter) Print() {
	if !c.sstack.isEmpty() {
		panic("Print: not all scopes closed")
	}
	indent := ""
	for _, op := range c.ops {
		switch op.opType {
		case enterScope:
			fmt.Printf("%s%s:\n", indent, op.comment)
			indent += "\t"
		case exitScope:
			indent = indent[:len(indent)-1]
		default:
			fmt.Printf("%s%s", indent, op.opType)
			if op.comment != "" {
				fmt.Println(":", op.comment)
			} else {
				fmt.Println()
			}
		}
	}
}
