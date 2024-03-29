// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package dashboardPages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"hyperbot/web/components/dashboard"
	"hyperbot/web/layouts"
)

func DashboardLandingPage(totalAssets float64) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mt-12\"><div class=\"mb-12 grid gap-y-10 gap-x-6 md:grid-cols-2 xl:grid-cols-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = dashboardComponents.AccountSummary(totalAssets).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"mb-4 w-full grid grid-cols-1 gap-6 xl:grid-cols-3\"><div class=\"relative flex flex-col bg-clip-border rounded-xl bg-white text-gray-700 shadow-md overflow-hidden xl:col-span-2\"><div class=\"relative bg-clip-border rounded-xl overflow-hidden bg-transparent text-gray-700 shadow-none m-0 flex items-center justify-between p-6\"><div><h6 class=\"block antialiased tracking-normal font-sans text-base font-semibold leading-relaxed text-blue-gray-900 mb-1\">Recent activities</h6></div><button aria-expanded=\"false\" aria-haspopup=\"menu\" id=\":r5:\" class=\"relative middle none font-sans font-medium text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none w-8 max-w-[32px] h-8 max-h-[32px] rounded-lg text-xs text-blue-gray-500 hover:bg-blue-gray-500/10 active:bg-blue-gray-500/30\" type=\"button\"><span class=\"absolute top-1/2 left-1/2 transform -translate-y-1/2 -translate-x-1/2\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currenColor\" viewBox=\"0 0 24 24\" stroke-width=\"3\" stroke=\"currentColor\" aria-hidden=\"true\" class=\"h-6 w-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z\"></path></svg></span></button></div><div class=\"p-6 overflow-x-scroll px-0 pt-0 pb-2\"><table class=\"w-full min-w-[640px] table-auto\"><thead><tr><th class=\"border-b border-blue-gray-50 py-3 px-6 text-left\"><p class=\"block antialiased font-sans text-[11px] font-medium uppercase text-blue-gray-400\">companies</p></th><th class=\"border-b border-blue-gray-50 py-3 px-6 text-left\"><p class=\"block antialiased font-sans text-[11px] font-medium uppercase text-blue-gray-400\">budget</p></th><th class=\"border-b border-blue-gray-50 py-3 px-6 text-left\"><p class=\"block antialiased font-sans text-[11px] font-medium uppercase text-blue-gray-400\">completion</p></th></tr></thead> <tbody><tr><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"flex items-center gap-4\"><p class=\"block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-bold\">Material XD Version</p></div></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><p class=\"block antialiased font-sans text-xs font-medium text-blue-gray-600\">$14,000</p></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"w-10/12\"><p class=\"antialiased font-sans mb-1 block text-xs font-medium text-blue-gray-600\">60%</p><div class=\"flex flex-start bg-blue-gray-50 overflow-hidden w-full rounded-sm font-sans text-xs font-medium h-1\"><div class=\"flex justify-center items-center h-full bg-gradient-to-tr from-blue-600 to-blue-400 text-white\" style=\"width: 60%;\"></div></div></div></td></tr><tr><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"flex items-center gap-4\"><p class=\"block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-bold\">Add Progress Track</p></div></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><p class=\"block antialiased font-sans text-xs font-medium text-blue-gray-600\">$3,000</p></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"w-10/12\"><p class=\"antialiased font-sans mb-1 block text-xs font-medium text-blue-gray-600\">10%</p><div class=\"flex flex-start bg-blue-gray-50 overflow-hidden w-full rounded-sm font-sans text-xs font-medium h-1\"><div class=\"flex justify-center items-center h-full bg-gradient-to-tr from-blue-600 to-blue-400 text-white\" style=\"width: 10%;\"></div></div></div></td></tr><tr><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"flex items-center gap-4\"><p class=\"block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-bold\">Fix Platform Errors</p></div></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><p class=\"block antialiased font-sans text-xs font-medium text-blue-gray-600\">Not set</p></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"w-10/12\"><p class=\"antialiased font-sans mb-1 block text-xs font-medium text-blue-gray-600\">100%</p><div class=\"flex flex-start bg-blue-gray-50 overflow-hidden w-full rounded-sm font-sans text-xs font-medium h-1\"><div class=\"flex justify-center items-center h-full bg-gradient-to-tr from-green-600 to-green-400 text-white\" style=\"width: 100%;\"></div></div></div></td></tr><tr><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"flex items-center gap-4\"><p class=\"block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-bold\">Launch our Mobile App</p></div></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><p class=\"block antialiased font-sans text-xs font-medium text-blue-gray-600\">$20,500</p></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"w-10/12\"><p class=\"antialiased font-sans mb-1 block text-xs font-medium text-blue-gray-600\">100%</p><div class=\"flex flex-start bg-blue-gray-50 overflow-hidden w-full rounded-sm font-sans text-xs font-medium h-1\"><div class=\"flex justify-center items-center h-full bg-gradient-to-tr from-green-600 to-green-400 text-white\" style=\"width: 100%;\"></div></div></div></td></tr><tr><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"flex items-center gap-4\"><p class=\"block antialiased font-sans text-sm leading-normal text-blue-gray-900 font-bold\">Add the New Pricing Page</p></div></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><p class=\"block antialiased font-sans text-xs font-medium text-blue-gray-600\">$500</p></td><td class=\"py-3 px-5 border-b border-blue-gray-50\"><div class=\"w-10/12\"><p class=\"antialiased font-sans mb-1 block text-xs font-medium text-blue-gray-600\">25%</p><div class=\"flex flex-start bg-blue-gray-50 overflow-hidden w-full rounded-sm font-sans text-xs font-medium h-1\"><div class=\"flex justify-center items-center h-full bg-gradient-to-tr from-blue-600 to-blue-400 text-white\" style=\"width: 25%;\"></div></div></div></td></tr></tbody></table></div></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = webLayouts.DashboardLaoyts().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
