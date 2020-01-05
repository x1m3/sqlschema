package operations

import (
	"fmt"
	"io"

	"github.com/sqlbunny/sqlschema/schema"
)

type RenameTable struct {
	OldName string
	NewName string
}

func (o RenameTable) GetSQL() string {
	return fmt.Sprintf("ALTER TABLE \"%s\" RENAME TO \"%s\"", o.OldName, o.NewName)
}

func (o RenameTable) Dump(w io.Writer) {
	fmt.Fprint(w, "operations.RenameTable {\n")
	fmt.Fprint(w, "OldName: "+esc(o.OldName)+",\n")
	fmt.Fprint(w, "NewName: "+esc(o.NewName)+",\n")
	fmt.Fprint(w, "}")
}

func (o RenameTable) Apply(s *schema.Schema) error {
	t, ok := s.Tables[o.OldName]
	if !ok {
		return fmt.Errorf("no such table: %s", o.OldName)
	}
	if _, ok := s.Tables[o.NewName]; ok {
		return fmt.Errorf("destination table already exists: %s", o.NewName)
	}

	delete(s.Tables, o.OldName)
	s.Tables[o.NewName] = t

	for _, t2 := range s.Tables {
		for _, fk := range t2.ForeignKeys {
			if fk.ForeignTable == o.OldName {
				fk.ForeignTable = o.NewName
			}
		}
	}
	return nil
}
