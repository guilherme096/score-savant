// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package Home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "guilherme096/score-savant/templates/Layout"

func HomePage() templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-row h-[300px] w-fit items-center mx-auto mt-20\"><div class=\"w-[420px] h-full\"><h1 class=\"text-center text-5xl font-bold\">Score-Savant</h1><div class=\"w-full\"><ul class=\"w-full text-center p-4 w-80 min-h-full bg-base-200 text-base-content rounded-xl mt-8\"><li class=\"mt-2 h-12 hover:bg-base-300 rounded-xl flex flex-col justify-center items-center align-middle\"><a href=\"/search-player\">Search Player</a></li><li class=\"mt-2 h-12 hover:bg-base-300 rounded-xl flex flex-col justify-center items-center align-middle\"><a href=\"/search-club\">Search Club</a></li><li class=\"mt-2 h-12 hover:bg-base-300 rounded-xl flex flex-col justify-center items-center align-middle\"><a href=\"/player-insertion\">Add Players</a></li></ul></div></div><div class=\"max-w-xs h-fit w-full ml-16 mt-8\"><div class=\"mb-2\"><h1 class=\"text-left text-2xl font-semibold inline\">Random Player Pick</h1><button class=\"btn btn-primary ml-4\" hx-get=\"api/get-random-player\" hx-trigger=\"click\" hx-target=\"#random\">Pick</button></div><div id=\"random\" hx-get=\"api/get-random-player\" hx-swap=\"innerHTML\" hx-trigger=\"load\"></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
