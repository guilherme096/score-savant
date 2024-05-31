// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package Search

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "guilherme096/score-savant/templates/Layout"

var leagues = []map[string]interface{}{
	{
		"LeagueID":   1,
		"LeagueName": "Premier League",
		"Nation":     "England",
		"ValueTotal": 7000000000.00,
	},
	{
		"LeagueID":   2,
		"LeagueName": "LaLiga",
		"Nation":     "Spain",
		"ValueTotal": 5000000000.00,
	},
	{
		"LeagueID":   3,
		"LeagueName": "Serie A",
		"Nation":     "Italy",
		"ValueTotal": 4500000000.00,
	},
	{
		"LeagueID":   4,
		"LeagueName": "Bundesliga",
		"Nation":     "Germany",
		"ValueTotal": 4200000000.00,
	},
	{
		"LeagueID":   5,
		"LeagueName": "Ligue 1",
		"Nation":     "France",
		"ValueTotal": 3500000000.00,
	},
	{
		"LeagueID":   6,
		"LeagueName": "Eredivisie",
		"Nation":     "Netherlands",
		"ValueTotal": 1000000000.00,
	},
	{
		"LeagueID":   7,
		"LeagueName": "Primeira Liga",
		"Nation":     "Portugal",
		"ValueTotal": 1200000000.00,
	},
	{
		"LeagueID":   8,
		"LeagueName": "Major League Soccer",
		"Nation":     "USA",
		"ValueTotal": 1500000000.00,
	},
	{
		"LeagueID":   9,
		"LeagueName": "Brasileirão",
		"Nation":     "Brazil",
		"ValueTotal": 1300000000.00,
	},
	{
		"LeagueID":   10,
		"LeagueName": "J1 League",
		"Nation":     "Japan",
		"ValueTotal": 800000000.00,
	},
}

func LeagueSearchPage() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-row h-full mr-6 p-4\"><div class=\"w-80 bg-gray-200 p-4 ml-10 mr-5 rounded-lg h-fit\"><h2 class=\"text-xl font-bold mb-4\">Filters</h2><form hx-get=\"/api/list-leagues\" hx-target=\"#table-body\" class=\"space-y-4\"><div class=\"flex flex-rol justify-between space-x-2\"><div><label for=\"sort\" class=\"block text-sm font-medium\">Sort By</label> <select name=\"sort\" id=\"sort\" class=\"select select-bordered w-full max-w-1/2\"><option value=\"\">Default</option> <option value=\"League\">League</option> <option value=\"Nation\">Nation</option> <option value=\"Value\">Value</option></select></div><div class=\"min-w-1/2\"><label for=\"direction\" class=\"block text-sm font-medium\">Direction</label> <select name=\"direction\" id=\"direction\" class=\"select select-bordered w-full max-w-xs\"><option value=\"DESC\">Desc</option> <option value=\"ASC\">Asc</option></select></div></div><div><label for=\"leagueName\" class=\"block text-sm font-medium\">League Name</label> <input type=\"text\" name=\"leagueName\" id=\"leagueName\" placeholder=\"League Name\" class=\"w-full p-2 border border-gray-300 rounded\"></div><div><label for=\"nationName\" class=\"block text-sm font-medium\">Nation</label> <input type=\"text\" name=\"nationName\" id=\"nationName\" placeholder=\"Nation\" class=\"w-full p-2 border border-gray-300 rounded\"></div><div><label for=\"valueTotalRange\" class=\"block text-sm font-medium\">Value Total Range</label><div class=\"flex space-x-2\"><input type=\"number\" name=\"minValue\" id=\"minValue\" placeholder=\"Min\" class=\"w-full p-2 border border-gray-300 rounded\"> <input type=\"number\" name=\"maxValue\" id=\"maxValue\" placeholder=\"Max\" class=\"w-full p-2 border border-gray-300 rounded\"></div></div><div><input type=\"submit\" value=\"Apply Filters\" class=\"w-full p-2 btn btn-primary text-white font-bold rounded cursor-pointer\"></div></form></div><div class=\"w-full\"><div class=\"club-list\"><div class=\"mx-auto\"><table class=\"table table-zebra w-full rounded-lg overflow-clip\"><thead class=\"bg-gray-300\"><tr class=\"text-md\"><th class=\"py-2 px-6 border-b cursor-pointer\">Star</th><th class=\"py-2 px-6 border-b cursor-pointer\">League</th><th class=\"py-2 px-6 border-b cursor-pointer\">Nation</th><th class=\"py-2 px-6 border-b cursor-pointer\">Value Total</th></tr></thead> <tbody id=\"table-body\" hx-get=\"/api/list-leagues?page=1\" hx-trigger=\"load\" hx-swap=\"innerHTML\"></tbody></table><button id=\"load-more\" class=\"w-full p-2 btn btn-primary text-white rounded mt-8\">Load More</button></div></div></div></div><script>\n        let currentPage = 1; // Initialize the current page\n\n        document.getElementById('load-more').addEventListener('click', function () {\n                currentPage++; // Increment the current page\n                fetchPage(currentPage); // Fetch the new page\n                });\n\n        function fetchPage(page) {\n            htmx.ajax('GET', `/api/list-leagues?page=${page}`, {\n            swap: 'beforeend',\n            target: '#table-body',\n        });\n        }\n        </script>")
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
