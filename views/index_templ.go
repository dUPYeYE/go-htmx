// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/dUPYeYE/go-htmx/internal/models"
	"github.com/dUPYeYE/go-htmx/views/components"
)

func Index(userData *models.User) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"min-h-screen bg-background text-foreground\"><div class=\"relative flex flex-col items-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Navbar(userData).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mt-6 w-[65%] flex justify-center items-center flex-col\"><h1 class=\"text-4xl font-bold\">GO + HTMX Fullstack app</h1><p class=\"text-m mt-4\">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam et commodo lorem. In pulvinar nunc id purus accumsan, vitae mollis sapien gravida. Duis ligula justo, malesuada ut nibh et, consequat cursus augue. Pellentesque vitae condimentum odio. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec imperdiet eu quam id rhoncus. Sed orci orci, aliquam sit amet mattis eget, sagittis ac arcu. Donec venenatis scelerisque turpis, id pulvinar felis bibendum eget. Aenean mollis risus sit amet dolor fringilla hendrerit. Donec sit amet lectus eu nisl faucibus pellentesque in eu risus. Etiam vel varius sem. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Sed consectetur eros sit amet felis lacinia, vitae semper diam bibendum. Cras cursus ultricies elit nec lobortis. Praesent mattis tincidunt faucibus. Integer eleifend molestie lorem, ut pulvinar elit gravida et.</p><p class=\"text-m mt-4\">In porttitor lacus ante, ut facilisis ipsum pellentesque a. Donec ac erat nec nibh tincidunt viverra sed a lectus. Sed interdum sed ante vitae interdum. Nam quis enim venenatis velit vestibulum tempus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam maximus, libero a commodo finibus, enim urna commodo nulla, eu tincidunt nisi erat volutpat neque. Praesent nisi nibh, bibendum dignissim ultricies sagittis, laoreet vel lacus. Donec eu tortor et arcu fermentum mattis eget quis diam. Sed lectus enim, volutpat bibendum aliquam nec, porttitor eu augue. Cras diam est, hendrerit eu felis non, pretium semper quam. Nulla auctor quis nisl quis sodales. Ut eu erat iaculis, fringilla tellus id, dapibus lorem. Ut nec leo lacus.</p><img src=\"/images/go.png\" height=\"200\" width=\"200\" class=\"mt-10\"></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Footer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
