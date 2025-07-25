// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.906
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Login() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<form method=\"POST\" action=\"/admin/login\" class=\"max-w-sm mx-auto mt-20 p-6 border border-gray-300 rounded shadow\"><h2 class=\"text-white font-bold mb-4 text-center\">Admin Login</h2><div class=\"mb-4 text-white\"><label for=\"username\" class=\"block text-white-700 mb-1\">Username</label> <input type=\"text\" id=\"username\" name=\"username\" required class=\"text-white w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500\"></div><div class=\"mb-6 text-white\"><label for=\"password\" class=\"block text-white-700 mb-1\">Password</label> <input type=\"password\" id=\"password\" name=\"password\" required class=\"text-white w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500\"></div><button type=\"submit\" class=\"cursor-pointer w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700\">Login</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
