// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package club

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"guilherme096/score-savant/templates/Layout"
	Utils "guilherme096/score-savant/utils"
	"strconv"
)

var league map[string]interface{} = map[string]interface{}{
	"league_id":     1,
	"league_name":   "Premier League",
	"nation_name":   "England",
	"total_clubs":   20,
	"total_players": 500,
	"total_value":   1000000000.00,
	"total_wage":    100000000.00,
	"avg_att":       12,
}

func LeaguePage() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full max-h-full m-8 flex flex-row\"><div class=\"w-96\"><div class=\"w-full bg-gray-200 rounded-lg p-4\"><h1 class=\"text-2xl font-bold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(league["league_name"].(string))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/League/LeaguePage.templ`, Line: 27, Col: 81}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><div class=\"w-full h-1/2 flex flex-col mt-6\"><div class=\"w-full h-1/2 flex flex-col\"><div class=\"flex flex-row\"><h1 class=\"text-lg font-bold\">Nation:</h1><h1 class=\"text-lg ml-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(league["nation_name"].(string))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/League/LeaguePage.templ`, Line: 32, Col: 88}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div><div class=\"flex flex-row\"><h1 class=\"text-lg font-bold\">Total Clubs:</h1><h1 class=\"text-lg ml-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(league["total_clubs"].(int)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/League/LeaguePage.templ`, Line: 36, Col: 99}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div><div class=\"flex flex-row\"><h1 class=\"text-lg font-bold\">Player Count:</h1><h1 class=\"text-lg ml-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(league["total_players"].(int)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/League/LeaguePage.templ`, Line: 40, Col: 101}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div><div class=\"flex flex-row\"><h1 class=\"text-lg font-bold\">Average Attribute Rating:</h1><h1 class=\"text-lg ml-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(league["avg_att"].(int)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/League/LeaguePage.templ`, Line: 44, Col: 95}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div><div class=\"flex flex-row\"><h1 class=\"text-lg font-bold\">Total Wage:</h1><h1 class=\"text-lg ml-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(Utils.FormatNumber(league["total_wage"].(float64)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/League/LeaguePage.templ`, Line: 48, Col: 108}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div><div class=\"flex flex-row\"><h1 class=\"text-lg font-bold\">Total Value:</h1><h1 class=\"text-lg ml-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(Utils.FormatNumber(league["total_value"].(float64)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/League/LeaguePage.templ`, Line: 52, Col: 109}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div></div></div></div></div><div class=\"mx-auto px-12 flex-1\"><div class=\"search-bar w-full mb-5\"><form action=\"/search\" method=\"get\" class=\"flex\"><input type=\"text\" name=\"search\" class=\"w-full p-2 rounded-l-lg border border-gray-300\" placeholder=\"Search for a player\"> <input type=\"submit\" value=\"Search\" class=\"p-2 bg-blue-500 text-white font-bold rounded-r-lg cursor-pointer\"></form></div><div class=\"w-full\"><div class=\"club-list\"><div class=\"mx-auto\"><table class=\"table table-zebra w-full rounded-lg overflow-clip\"><thead class=\"bg-gray-300\"><tr class=\"text-md\"><th class=\"py-2 px-6 border-b\">Star</th><th class=\"py-2 px-6 border-b\">Name</th><th class=\"py-2 px-6 border-b\">Age</th><th class=\"py-2 px-6 border-b\">Position</th><th class=\"py-2 px-6 border-b\">Club</th><th class=\"py-2 px-6 border-b\">Nation</th><th class=\"py-2 px-6 border-b\">League</th><th class=\"py-2 px-6 border-b\">Wage</th><th class=\"py-2 px-6 border-b\">Value</th></tr></thead> <tbody id=\"table-body\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/api/list-players?page=1&leagueName=%s", league["league_name"].(string)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/League/LeaguePage.templ`, Line: 82, Col: 143}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"load\" hx-swap=\"innerHTML\"></tbody></table><button id=\"load-more\" class=\"w-full p-2 btn btn-primary text-white rounded mt-8\">Load More</button></div></div></div></div></div>")
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