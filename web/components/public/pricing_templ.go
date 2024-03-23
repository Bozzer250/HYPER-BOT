// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package publicComponents

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Bots() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section><div class=\"container max-w-full mx-auto py-12 px-6\"><h1 class=\"text-center text-4xl text-black font-medium leading-snug tracking-wider\">Our bots</h1><p class=\"text-center text-lg text-gray-700 mt-2 px-6\"></p><div class=\"h-1 mx-auto bg-indigo-200 w-24 opacity-75 mt-4 rounded\"></div><div class=\"max-w-full md:max-w-6xl mx-auto my-3 md:px-8\"><div class=\"relative block flex flex-col md:flex-row items-center\"><div class=\"w-11/12 max-w-sm sm:w-3/5 lg:w-1/3 sm:my-5 my-8 relative z-0 rounded-lg shadow-lg md:-mr-4\"><div class=\"bg-white text-black rounded-lg shadow-inner shadow-lg overflow-hidden\"><div class=\"block text-left text-sm sm:text-md max-w-sm mx-auto mt-2 text-black px-8 lg:px-6\"><h1 class=\"text-lg font-medium uppercase p-3 pb-0 text-center tracking-wide\">Hobby</h1><h2 class=\"text-sm text-gray-500 text-center pb-6\">$10</h2></div><div class=\"flex flex-wrap mt-3 px-6\"><ul><li class=\"flex items-center\"><div class=\" rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">1% Daily Profit</span></li><li class=\"flex items-center\"><div class=\" rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">Weekly Withdrawals</span></li><li class=\"flex items-center\"><div class=\" rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">Speed</span></li></ul></div><div class=\"block flex items-center p-8  uppercase\"><button class=\"mt-3 text-lg font-semibold\n								bg-green w-full text-white rounded-lg\n								px-6 py-3 block shadow-xl hover:bg-gray-700\" onClick=\"window.location.href = &#39;/invest?plan=hobby&amp;price=10&#39;\">Invest</button></div></div></div><div class=\"w-full max-w-md sm:w-2/3 lg:w-1/3 sm:my-5 my-8 relative z-10 bg-white rounded-lg shadow-lg\"><div class=\"text-sm leading-none rounded-t-lg bg-gray-200 text-black font-semibold uppercase py-4 text-center tracking-wide\">Most Popular</div><div class=\"block text-left text-sm sm:text-md max-w-sm mx-auto mt-2 text-black px-8 lg:px-6\"><h1 class=\"text-lg font-medium uppercase p-3 pb-0 text-center tracking-wide\">ADVANCED</h1><h2 class=\"text-sm text-gray-500 text-center pb-6\"><span class=\"text-3xl\">$50</span></h2></div><div class=\"flex pl-12 justify-start sm:justify-start mt-3\"><ul><li class=\"flex items-center\"><div class=\"rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">6% Daily Profit</span></li><li class=\"flex items-center\"><div class=\" rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">Daily Withdrawals</span></li><li class=\"flex items-center\"><div class=\" rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">Referrals Income</span></li></ul></div><div class=\"block flex items-center p-8  uppercase\"><button class=\"mt-3 text-lg font-semibold\n	bg-green w-full text-white rounded-lg\n	px-6 py-3 block shadow-xl hover:bg-gray-700\">Invest</button></div></div><div class=\"w-11/12 max-w-sm sm:w-3/5 lg:w-1/3 sm:my-5 my-8 relative z-0 rounded-lg shadow-lg md:-ml-4\"><div class=\"bg-white text-black rounded-lg shadow-inner shadow-lg overflow-hidden\"><div class=\"block text-left text-sm sm:text-md max-w-sm mx-auto mt-2 text-black px-8 lg:px-6\"><h1 class=\"text-lg font-medium uppercase p-3 pb-0 text-center tracking-wide\">EXPERT</h1><h2 class=\"text-sm text-gray-500 text-center pb-6\">$300</h2></div><div class=\"flex flex-wrap mt-3 px-6\"><ul><li class=\"flex items-center\"><div class=\" rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">20% Daily profit</span></li><li class=\"flex items-center\"><div class=\" rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">Instant Withdrawals</span></li><li class=\"flex items-center\"><div class=\" rounded-full p-2 fill-current text-green-700\"><svg class=\"w-6 h-6 align-middle\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M22 11.08V12a10 10 0 1 1-5.93-9.14\"></path> <polyline points=\"22 4 12 14.01 9 11.01\"></polyline></svg></div><span class=\"text-gray-700 text-lg ml-3\">Referrals Income</span></li></ul></div><div class=\"block flex items-center p-8  uppercase\"><button class=\"mt-3 text-lg font-semibold\n	bg-green w-full text-white rounded-lg\n	px-6 py-3 block shadow-xl hover:bg-gray-700\">Invest</button></div></div></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
