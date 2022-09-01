package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

const JEL_CODE = `function jel(){var e=this===window?document.body:this;return e.jel.apply(e,arguments)}jel.settings={},jel.settings.allowScripts=!0,jel.settings.mapKeywords={innerHTML:"innerHTML",html:"html",children:"children",chi:"chi",properties:"properties",prop:"prop",jel:"jel"},jel._templates={},jel.SetTemplate=function(e,t){return"string"!=typeof e||"function"!=typeof t&&"object"!=typeof t&&"string"!=typeof t?null:(this._templates[e]=t,this)},jel.GetTemplate=function(e){return this._templates[e]},HTMLElement.prototype.jel=function(){function e(t,r){if(Array.isArray(r))for(var l=0;l<r.length;l++)e(t,r[l]);else if("function"==typeof r)e(t,r.call(t,t,t.style));else for(var l in r){var n;n="function"==typeof r[l]?r[l].call(t,t,t.style[l]):r[l],("string"!=typeof t.style[l]||-1==t.style[l].search("!important")||"string"==typeof n&&-1!=n.search("!important"))&&(t.style[l]=n)}}function t(e,r){if(Array.isArray(r))for(var l=0;l<r.length;l++)t(e,r[l]);else{var n=e.getAttribute("class");n||(n=""),n="function"==typeof r?n+" "+r.call(e,e,n.trim()):n+" "+r,e.setAttribute("class",n.trim())}}function r(e,t){return jel.settings.allowScripts||-1==t.search("<script")?e.appendChild(document.createRange().createContextualFragment(t)):null}function l(e,t,r){for(var l in t)void 0!==t[l].tagName?e.appendChild(t[l]):"object"==typeof t[l]||"string"==typeof t[l]?e.jel(t[l],{_appliedTemplates:r}):"function"==typeof t[l]&&t[l].call(e,e)}function n(l,n,o){for(var a in n)switch(a){case jel.settings.mapKeywords.innerHTML:case jel.settings.mapKeywords.html:r(l,n[a]);break;case"style":if("string"==typeof n[a])null===(i=l.getAttribute(a))&&(i=""),l.setAttribute(a,(i+" "+n[a]).trim());else"object"==typeof n[a]&&e(l,n[a]);break;case"class":var i;if("string"==typeof n[a])null===(i=l.getAttribute(a))&&(i=""),l.setAttribute(a,(i+" "+n[a]).trim());else"object"==typeof n[a]&&t(l,n[a]);break;case jel.settings.mapKeywords.children:case jel.settings.mapKeywords.chi:if("object"==typeof n[a])if(Array.isArray(n[a]))for(var s in n[a])switch(typeof n[a][s]){case"object":l.jel(n[a][s],{_appliedTemplates:o});break;case"string":r(l,n[a][s]);break;case"function":if(!1===(p=n[a][s].call(l,l)))return;p&&l.jel(p,{_appliedTemplates:o})}else l.jel(n[a],{_appliedTemplates:o});else if("function"==typeof n[a]){var p;if(!1===(p=n[a].call(l,l)))return;p&&l.jel(p,{_appliedTemplates:o})}break;case jel.settings.mapKeywords.jel:if("object"==typeof n[a])for(var s in n[a])switch(s){case"name":l.jelEx._namedParent!==l&&(l.jelEx._namedParent[n[a][s]]=l,l.jelEx._namedParentForChildren=l);break;case"root":l.jelEx._componentRoot=l;break;case"links":if("object"!=typeof n[a][s])break;for(var c in n[a][s])l.jelEx.AddPropertyLink(c,n[a][s][c])}break;case"_appliedTemplates":case jel.settings.mapKeywords.properties:case jel.settings.mapKeywords.prop:break;default:"function"==typeof n[a]?l.setAttribute(a,n[a].call(l,l,l.getAttribute(a))):l.setAttribute(a,n[a])}}function o(e){for(var t=new Array(e.length),r=0;r<t.length;r++)t[r]=e[r];return t}if(0!==arguments.length){var a="object"==typeof arguments[arguments.length-1]&&arguments[arguments.length-1]._appliedTemplates?function(e){var t={};for(var r in e)t[r]=e[r];return t}(arguments[arguments.length-1]._appliedTemplates):null;if("object"==typeof arguments[0]&&Array.isArray(arguments[0])){for(var i=[],s=0;s<arguments[0].length;s++){var p=o(arguments);p[0]=arguments[0][s];var c=this.jel.apply(this,p);"undefined"!=c&&i.push(c)}return i}if("object"==typeof arguments[0]){var f=null,j=null;for(var y in arguments[0]){f=y,j=arguments[0][y];break}return Array.prototype.unshift.call(arguments,f),arguments[1]=j,this.jel.apply(this,arguments)}if("function"==typeof arguments[0])return arguments[0].call(this);if("string"!=typeof arguments[0])throw"Format error";if("function"==typeof jel._templates[arguments[0]])return jel._templates[arguments[0]].call(this,this,o(arguments));if(!jel.settings.allowScripts&&"script"==arguments[0].trim())return null;var h={};if(!("object"!=typeof jel._templates[arguments[0]]||a&&a[arguments[0]])){a||(a={}),a[arguments[0]]=!0;var u=o(arguments);return u[0]=jel._templates[arguments[0]],u.push({_appliedTemplates:a}),h=this.jel.apply(this,u)}if(""==arguments[0]&&"string"==typeof arguments[1]&&(2==arguments.length||3==arguments.length&&void 0!==arguments[arguments.length-1]._appliedTemplates))return r(this,arguments[1]),null;if((h=document.createElement(arguments[0])).jelEx={},h.jelEx._ownerElement=h,h.jelEx._namedParentForChildren=h,h.jelEx._namedParent=h,h.jelEx._componentRoot=h,h.jelEx._topComponentRoot=h,h.jelEx.AddPropertyLink=function e(t,r){if("object"==typeof r&&Array.isArray(r))for(var l=0;l<r.length;l++)e.call(this,t,r[l]);else{var n=this._ownerElement,o=t.split("."),a=r.split("."),i=n.jelEx._namedParent;"root"===o[0]&&(i=n.jelEx._componentRoot);for(var s=1;s<o.length-1;s++)void 0===i[o[s]]&&(i[o[s]]={}),i=i[o[s]];for(var p=n,c=0;c<a.length-1;c++){if(void 0===p[a[c]])throw"Invalid local property";p=p[a[c]]}!function(e,t,r,l){var n=e[e.length-1],o=r[r.length-1];void 0===t[n]&&(Object.defineProperty(t,n,{get:function(){var e=[],r=t.jelExt[n]._properties;for(var l in r)e.push(r[l].localPropParent[r[l].localPropName]);return 1===e.length&&(e=e[0]),e},set:function(e){var r=t.jelExt[n]._properties;for(var l in r)r[l].localPropParent[r[l].localPropName]=e}}),t.jelExt=t.jelExt||{},t.jelExt[n]={},t.jelExt[n]._properties=[]),t.jelExt[n]._properties.push({localPropParent:l,localPropName:o})}(o,i,a,p)}},void 0!==this.jelEx&&(h.jelEx._namedParentForChildren=this.jelEx._namedParentForChildren,h.jelEx._namedParent=this.jelEx._namedParentForChildren,h.jelEx._componentRoot=this.jelEx._componentRoot),void 0===h)throw"Unknown error";for(var m=1;m<arguments.length;m++)if("object"==typeof arguments[m]&&("object"==typeof arguments[m][jel.settings.mapKeywords.properties]||"object"==typeof arguments[m][jel.settings.mapKeywords.prop])){var d=arguments[m][jel.settings.mapKeywords.properties]||arguments[m][jel.settings.mapKeywords.prop];for(var g in d){for(var v=g.split("."),b=h,_=0;_<v.length-1;_++)void 0===b[v[_]]&&(b[v[_]]={}),b=b[v[_]];b[v[v.length-1]]=d[g]}}for(m=1;m<arguments.length;m++)switch(typeof arguments[m]){case"string":r(h,arguments[m]);break;case"object":Array.isArray(arguments[m])?l(h,arguments[m],a):n(h,arguments[m],a);break;case"function":var E=arguments[m].call(h,h);if(!1===E)return null;E&&(arguments[m]=E,m-=1)}return this.appendChild(h),h}};`

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, `<html><head></head><body><script type="text/javascript">%s</script></body></html>`, JEL_CODE)
	fmt.Fprintf(ctx, `<script type="text/javascript">jel({"h1": "test"})</script>`)

	// fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	// fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	// fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	// fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	// fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	// fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	// fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	// fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	// fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	// fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	// fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	// ctx.SetContentType("text/plain; charset=utf8")
	ctx.SetContentType("text/html; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}
