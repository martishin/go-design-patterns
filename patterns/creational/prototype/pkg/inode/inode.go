package inode

import (
	"io"
)

// Inode is the Prototype interface.
//
// The Prototype pattern allows cloning objects without coupling code to their concrete types.
type Inode interface {
	// Print writes a hierarchical representation of the inode to w.
	Print(w io.Writer, indentation string)

	// Clone returns a copy of the receiver.
	Clone() Inode
}
