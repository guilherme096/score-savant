// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package TopPlayers

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"guilherme096/score-savant/templates/Layout"
)

func TopPlayersPage() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-row h-full mr-6 p-4\"><div class=\"w-full\"><div class=\"player-list\"><div class=\"container mx-auto\"><table class=\"w-full table-zebra rounded-lg shadow-lg overflow-clip table\"><thead class=\"bg-gray-300\"><tr class=\"text-md\"><th class=\"py-2 px-6 border-b cursor-pointer\">Star</th><th class=\"py-2 px-6 border-b cursor-pointer text-left\">Name</th><th class=\"py-2 px-6 border-b cursor-pointer\">Age</th><th class=\"py-2 px-6 border-b cursor-pointer\">Position</th><th class=\"py-2 px-6 border-b cursor-pointer\">Club</th><th class=\"py-2 px-6 border-b cursor-pointer\">Nation</th><th class=\"py-2 px-6 border-b cursor-pointer\">League</th><th class=\"py-2 px-6 border-b cursor-pointer\">Wage</th><th class=\"py-2 px-6 border-b cursor-pointer\">Value</th></tr></thead> <tbody id=\"table-body\" hx-get=\"api/list-players?page=1\" hx-trigger=\"load\" hx-swap=\"innerHTML\"></tbody></table><button id=\"load-more\" class=\"w-full p-2 btn btn-primary text-white rounded mt-8\">Load More</button></div></div></div></div><script>\n        let currentPage = 1; // Initialize the current page\n\n        document.getElementById('load-more').addEventListener('click', function () {\n                currentPage++; // Increment the current page\n                fetchPage(currentPage); // Fetch the new page\n                });\n\n        function fetchPage(page) {\n            htmx.ajax('GET', `/api/list-players?page=${page}`, {\n            swap: 'beforeend',\n            target: '#table-body',\n        });\n        }\n\n// Initially load the first page\nfetchPage(currentPage);\n</script>")
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
