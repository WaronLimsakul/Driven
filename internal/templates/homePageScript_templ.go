// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func HomePageScript() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<script type=\"text/javascript\">\n\t\tdocument.addEventListener(\"DOMContentLoaded\", (e) => {\n\t\t    // force client machine to send UTC time.\n  \t\t    document.querySelector(\"#task-form\").addEventListener(\"htmx:configRequest\", (event) => {\n              let formData = event.detail.parameters; // This is the form data HTMX will send\n              // console.log(formData);\n\n              if (formData[\"task-date\"]) {\n                  let localDate = new Date(formData[\"task-date\"] + \"T00:00:00\");\n                  let utcDate = new Date(Date.UTC(localDate.getFullYear(), localDate.getMonth(), localDate.getDate()));\n                  formData[\"task-date\"] = utcDate.toISOString().split(\"T\")[0]; // Convert to UTC format\n                  // console.log(\"task-date: \", formData[\"task-date\"]);\n              }\n            });\n\n\t\t    document.addEventListener(\"htmx:beforeSwap\", (event) => {\n\t\t\t\tif (event.detail.xhr.status == 403) {\n\t\t\t     event.detail.shouldSwap = true;\n\t\t\t\t} else if (event.detail.target == document.querySelector(\"#task-form-msg\")\n\t\t\t\t  && event.detail.xhr.status == 201) { // creat task success\n\t\t\t\t\tevent.detail.shouldSwap = false;\n\t\t\t\t}\n\t\t\t});\n\t\t});\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
