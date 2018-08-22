// Copyright 2015 Garrett D'Amore
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encoding

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUTF8(t *testing.T) {
	Convey("ASCII UTF8 identity transforms", t, func() {
		for i := 0; i < 128; i++ {
			verifyMap(UTF8, byte(i), rune(i))
		}
	})

	// We need to add tests for ErrSrcShort, ErrDstShort, and
	// larger rune values.
}
