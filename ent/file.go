// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-file/ent/file"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// 创建人id
	CreateId string `json:"createId,omitempty"`
	// 部门id
	DepartmentId string `json:"departmentId,omitempty"`
	// 分类id
	CategoryID int `json:"category_id,omitempty"`
	// File's name | 文件名称
	Name string `json:"name,omitempty"`
	// File's type | 文件类型
	FileType uint8 `json:"file_type,omitempty"`
	// File's size | 文件大小
	Size uint64 `json:"size,omitempty"`
	// File's path | 文件路径
	Path string `json:"path,omitempty"`
	// User's UUID | 用户的 UUID
	UserID string `json:"user_id,omitempty"`
	// The md5 of the file | 文件的 md5
	Md5 string `json:"md5,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges        FileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Tags holds the value of the tags edge.
	Tags []*FileTag `json:"tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) TagsOrErr() ([]*FileTag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldStatus, file.FieldCategoryID, file.FieldFileType, file.FieldSize:
			values[i] = new(sql.NullInt64)
		case file.FieldCreateId, file.FieldDepartmentId, file.FieldName, file.FieldPath, file.FieldUserID, file.FieldMd5:
			values[i] = new(sql.NullString)
		case file.FieldCreatedAt, file.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case file.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case file.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case file.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case file.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = uint8(value.Int64)
			}
		case file.FieldCreateId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field createId", values[i])
			} else if value.Valid {
				f.CreateId = value.String
			}
		case file.FieldDepartmentId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field departmentId", values[i])
			} else if value.Valid {
				f.DepartmentId = value.String
			}
		case file.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				f.CategoryID = int(value.Int64)
			}
		case file.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case file.FieldFileType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field file_type", values[i])
			} else if value.Valid {
				f.FileType = uint8(value.Int64)
			}
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = uint64(value.Int64)
			}
		case file.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				f.Path = value.String
			}
		case file.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				f.UserID = value.String
			}
		case file.FieldMd5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field md5", values[i])
			} else if value.Valid {
				f.Md5 = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the File.
// This includes values selected through modifiers, order, etc.
func (f *File) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryTags queries the "tags" edge of the File entity.
func (f *File) QueryTags() *FileTagQuery {
	return NewFileClient(f.config).QueryTags(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return NewFileClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", f.Status))
	builder.WriteString(", ")
	builder.WriteString("createId=")
	builder.WriteString(f.CreateId)
	builder.WriteString(", ")
	builder.WriteString("departmentId=")
	builder.WriteString(f.DepartmentId)
	builder.WriteString(", ")
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", f.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("file_type=")
	builder.WriteString(fmt.Sprintf("%v", f.FileType))
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(f.Path)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(f.UserID)
	builder.WriteString(", ")
	builder.WriteString("md5=")
	builder.WriteString(f.Md5)
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File
