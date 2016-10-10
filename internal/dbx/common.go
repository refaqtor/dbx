// Copyright (C) 2016 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dbx

import (
	"io"

	"github.com/spacemonkeygo/errors"
)

var Error = errors.NewClass("dbx")

type Language interface {
	Name() string
	RenderHeader(w io.Writer, schema *Schema) error
	RenderSelect(w io.Writer, sql string, params *SelectParams) error
	RenderCount(w io.Writer, sql string, params *SelectParams) error
	RenderDelete(w io.Writer, sql string, params *DeleteParams) error
	RenderInsert(w io.Writer, sql string, params *InsertParams) error
	RenderUpdate(w io.Writer, sql string, params *UpdateParams) error
	Format([]byte) ([]byte, error)
}

type Dialect interface {
	Name() string
	ColumnName(column *Column) string
	RenderSchema(schema *Schema) (string, error)
	RenderSelect(params *SelectParams) (string, error)
	RenderCount(params *SelectParams) (string, error)
	RenderDelete(params *DeleteParams) (string, error)
	RenderInsert(params *InsertParams) (string, error)
	RenderUpdate(params *UpdateParams) (string, error)
	SupportsReturning() bool
}