//=============================================================================
/*
Copyright © 2024 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package service

import (
	"github.com/algotiqa/core/auth"
	"github.com/algotiqa/system-adapter/pkg/adapter"
	"github.com/algotiqa/system-adapter/pkg/business"
)

//=============================================================================

func getAdapters(c *auth.Context) {
	list := business.GetAdapters()
	_ = c.ReturnList(list, 0, 1000, len(*list))
}

//=============================================================================

func getAdapter(c *auth.Context) {
	code := c.GetCodeFromUrl()
	a, err := business.GetAdapter(code)

	if err == nil {
		_ = c.ReturnObject(a)
	}

	c.ReturnError(err)
}

//=============================================================================

func getConnectionParams(c *auth.Context) {
	code := c.GetCodeFromUrl()
	config := map[string]any{}
	err := c.BindParamsFromBody(&config)

	if err == nil {
		var params []*adapter.ParamDef
		params, err = business.GetConnectionParams(code, config)
		if err == nil {
			_ = c.ReturnObject(params)
		}
	}

	c.ReturnError(err)
}

//=============================================================================
