// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Logo() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<svg width=\"400\" height=\"100\" viewBox=\"0 0 400 100\" xmlns=\"http://www.w3.org/2000/svg\"><g fill=\"#00A8FF\"><!-- D --><polygon points=\"10,20 30,50 10,80 0,80 20,50 0,20\"></polygon> <polygon points=\"30,20 50,50 30,80 40,80 60,50 40,20\"></polygon><!-- R --><polygon points=\"70,20 90,20 105,50 90,50 110,80 90,80 75,55 70,55\"></polygon><!-- I --><polygon points=\"120,20 140,20 130,40 140,60 120,60 130,40\"></polygon><!-- V --><polygon points=\"150,20 165,50 180,20 195,50 180,80 165,80 150,50\"></polygon><!-- E --><polygon points=\"200,20 220,20 240,35 220,50 240,65 220,80 200,80 220,50\"></polygon><!-- N --><polygon points=\"250,80 270,45 290,80 300,80 270,30 240,80\"></polygon></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
