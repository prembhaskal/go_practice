package fs_test

import (
	"fmt"
	"io/fs"
	"strings"
	"testing"

	cfs "github.com/prembhaskal/go_practice/pkg/util/fsclean/fs"
)

func TestPathUtil(t *testing.T) {
	t.Logf("Testing path util")

	pvisit := printvisit{}
	cfs.WalkDir("/tmp/prem", &pvisit)
}

type printvisit struct{
	pad int
}

func (p *printvisit) Visit(path string, d fs.DirEntry, err error) error {
	p.pad++
	var sb strings.Builder
	for i := 0; i < p.pad; i++ {
		sb.WriteString(" ")
	}

	fmt.Printf("%s Visit: %s\n", sb.String(), path)
	return nil
}

func (p *printvisit) Postvisit(path string, d fs.DirEntry, err error) error {
	var sb strings.Builder
	for i := 0; i < p.pad; i++ {
		sb.WriteString(" ")
	}
	
	fmt.Printf("%s Post Visit: %s\n", sb.String(), path)
	p.pad--
	return nil
}
