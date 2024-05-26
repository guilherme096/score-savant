// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package Search

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "guilherme096/score-savant/templates/Layout"

var clubs = []map[string]interface{}{
	{
		"name":         "Manchester United",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "Manchester City",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "38",
		"wage_total":   "£1,200,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "Liverpool",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "Chelsea",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "Arsenal",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "Tottenham Hotspur",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "Leicester City",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "West Ham United",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "Everton",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
	{
		"name":         "Aston Villa",
		"nation":       "England",
		"league":       "Premier League",
		"player_count": "30",
		"wage_total":   "£1,000,000",
		"value_total":  "£1,000,000,000",
	},
}

func ClubSearchPage() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex h-full\"><div class=\"w-1/5 bg-gray-200 p-4 ml-10 mt-3 mr-5 rounded-lg\"><h2 class=\"text-xl font-bold mb-4\">Filters</h2><form action=\"/search\" method=\"get\" class=\"space-y-4\"><div><label for=\"clubName\" class=\"block text-sm font-medium\">Club Name</label> <input type=\"text\" name=\"clubName\" id=\"clubName\" placeholder=\"Club name\" class=\"w-full p-2 border border-gray-300 rounded\"></div><div><label for=\"leagueName\" class=\"block text-sm font-medium\">League</label> <input type=\"text\" name=\"leagueName\" id=\"leagueName\" placeholder=\"League\" class=\"w-full p-2 border border-gray-300 rounded\"></div><div><label for=\"nationName\" class=\"block text-sm font-medium\">Nation</label> <input type=\"text\" name=\"nationName\" id=\"nationName\" placeholder=\"Nation\" class=\"w-full p-2 border border-gray-300 rounded\"></div><div><label for=\"playerCountRange\" class=\"block text-sm font-medium\">Player Count Range</label><div class=\"flex space-x-2\"><input type=\"number\" name=\"minPlayer\" id=\"minPlayer\" placeholder=\"Min\" class=\"w-full p-2 border border-gray-300 rounded\"> <input type=\"number\" name=\"maxPlayer\" id=\"maxPlayer\" placeholder=\"Max\" class=\"w-full p-2 border border-gray-300 rounded\"></div></div><div><label for=\"wageTotalRange\" class=\"block text-sm font-medium\">Wage Total Range</label><div class=\"flex space-x-2\"><input type=\"number\" name=\"minWage\" id=\"minWage\" placeholder=\"Min\" class=\"w-full p-2 border border-gray-300 rounded\"> <input type=\"number\" name=\"maxWage\" id=\"maxWage\" placeholder=\"Max\" class=\"w-full p-2 border border-gray-300 rounded\"></div></div><div><label for=\"valueTotalRange\" class=\"block text-sm font-medium\">Value Total Range</label><div class=\"flex space-x-2\"><input type=\"number\" name=\"minValue\" id=\"minValue\" placeholder=\"Min\" class=\"w-full p-2 border border-gray-300 rounded\"> <input type=\"number\" name=\"maxValue\" id=\"maxValue\" placeholder=\"Max\" class=\"w-full p-2 border border-gray-300 rounded\"></div></div><div><input type=\"submit\" value=\"Apply Filters\" class=\"w-full p-2 bg-blue-500 text-white font-bold rounded cursor-pointer\"></div></form></div><div class=\"w-full pr-32 pt-7\"><div class=\"search-bar mb-4\"><form action=\"/search\" method=\"get\" class=\"flex\"><input type=\"text\" name=\"club\" placeholder=\"Search for a club\" class=\"flex-grow p-2 border border-gray-300 rounded-l\"> <input type=\"submit\" value=\"Search\" class=\"p-2 bg-blue-500 text-white font-bold rounded-r cursor-pointer\"></form></div><div class=\"club-list\"><div class=\"container mx-auto p-5\"><table class=\"w-full bg-gray-100 rounded-lg\"><thead class=\"bg-gray-300\"><tr><th class=\"py-2 px-6 border-b cursor-pointer\">Star</th><th class=\"py-2 px-6 border-b cursor-pointer\">Club</th><th class=\"py-2 px-6 border-b cursor-pointer\">Nation</th><th class=\"py-2 px-6 border-b cursor-pointer\">League</th><th class=\"py-2 px-6 border-b cursor-pointer\">Player Count</th><th class=\"py-2 px-6 border-b cursor-pointer\">Wage Total</th><th class=\"py-2 px-6 border-b cursor-pointer\">Value Total</th></tr></thead> <tbody>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, club := range clubs {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr><th class=\"py-2 px-6 border-b\"><button id=\"starButton\" class=\"w-10 h-10 flex items-center justify-end\"><svg id=\"starIcon1\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 576 512\" class=\"w-5 h-5\"><path d=\"M287.9 0c9.2 0 17.6 5.2 21.6 13.5l68.6 141.3 153.2 22.6c9 1.3 16.5 7.6 19.3 16.3s.5 18.1-5.9 24.5L433.6 328.4l26.2 155.6c1.5 9-2.2 18.1-9.7 23.5s-17.3 6-25.3 1.7l-137-73.2L151 509.1c-8.1 4.3-17.9 3.7-25.3-1.7s-11.2-14.5-9.7-23.5l26.2-155.6L31.1 218.2c-6.5-6.4-8.7-15.9-5.9-24.5s10.3-14.9 19.3-16.3l153.2-22.6L266.3 13.5C270.4 5.2 278.7 0 287.9 0zm0 79L235.4 187.2c-3.5 7.1-10.2 12.1-18.1 13.3L99 217.9 184.9 303c5.5 5.5 8.1 13.3 6.8 21L171.4 443.7l105.2-56.2c7.1-3.8 15.6-3.8 22.6 0l105.2 56.2L384.2 324.1c-1.3-7.7 1.2-15.5 6.8-21l85.9-85.1L358.6 200.5c-7.8-1.2-14.6-6.1-18.1-13.3L287.9 79z\"></path></svg></button></th><th class=\"py-2 px-6 border-b\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(club["name"].(string))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/Search/ClubSearchPage.templ`, Line: 163, Col: 93}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th><th class=\"py-2 px-6 border-b\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(club["nation"].(string))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/Search/ClubSearchPage.templ`, Line: 164, Col: 95}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th><th class=\"py-2 px-6 border-b\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(club["league"].(string))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/Search/ClubSearchPage.templ`, Line: 165, Col: 95}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th><th class=\"py-2 px-6 border-b\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(club["player_count"].(string))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/Search/ClubSearchPage.templ`, Line: 166, Col: 101}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th><th class=\"py-2 px-6 border-b\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(club["wage_total"].(string))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/Search/ClubSearchPage.templ`, Line: 167, Col: 99}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th><th class=\"py-2 px-6 border-b\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(club["value_total"].(string))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/Search/ClubSearchPage.templ`, Line: 168, Col: 100}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th></tr>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div></div></div></div>")
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
