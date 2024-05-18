// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package Layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base(children ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>ScoreSavant</title><link href=\"/static/css/tailwind.css\" rel=\"stylesheet\"><link rel=\"icon\" href=\"/static/icons/logo.png\"></head><body><div class=\"h-screen w-screen m-0 p-0 left-0 top-0 absolute bg-gradient-to-tl from-primary to-base-100 overflow-clip overflow-y-scroll\"><div class=\"drawer drawer-start\"><input id=\"my-drawer-4\" type=\"checkbox\" class=\"drawer-toggle\"><div class=\"drawer-content\"><label for=\"my-drawer-4\" class=\"drawer-button btn btn-primary p-3 mt-3 ml-3\"><img class=\"h-5 w-5\" src=\"/static/icons/menu.png\"></label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"drawer-side\"><label for=\"my-drawer-4\" aria-label=\"close sidebar\" class=\"drawer-overlay\"></label><ul class=\"menu p-4 w-80 min-h-full bg-base-200 text-base-content\"><li><a>Search Player</a></li><li><a>Compare Players</a></li><li><a>Compare Clubs</a></li><li><a>Add Players</a></li></ul></div></div></div><!-- htmx --><script src=\"https://unpkg.com/htmx.org@1.9.12\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
