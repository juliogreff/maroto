// Package core contains all core interfaces and basic implementations.
package core

import (
	"encoding/base64"
	"os"
)

type Pdf struct {
	bytes []byte
}

// NewPDF is responsible to create a new instance of PDF.
func NewPDF(bytes []byte) Document {
	return &Pdf{
		bytes: bytes,
	}
}

// GetBytes returns the PDF bytes.
func (p *Pdf) GetBytes() []byte {
	return p.bytes
}

// GetBase64 returns the PDF bytes in base64.
func (p *Pdf) GetBase64() string {
	return base64.StdEncoding.EncodeToString(p.bytes)
}

// Save saves the PDF in a file.
func (p *Pdf) Save(file string) error {
	return os.WriteFile(file, p.bytes, os.ModePerm)
}
