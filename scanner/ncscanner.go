package scanner

import (
	"fmt"

	"github.com/wozuo/droidalyzer"
)

var networkingKeywords = []string{
	"http",
	"json",
	"xml",
	"url",
	"request",
	"API",
	"java.net.URLConnection",
	"java.net.HttpURLConnection",
	"HttpsURLConnection",
	"java.net.Socket",
	"DefaultHttpClient",
	"webview",
	"websocket",
}

var okHTTPKeywords = []string{
	"OkHttpClient()",
	"Request.Builder()",
	"addQueryParameter",
}

var retrofitKeywords = []string{
	"ResponseBody",
	"@POST",
	"@GET",
	"@PUT",
	"@PATCH",
}

var glideKeywords = []string{
	"GlideApp",
	"Glide",
}

var apacheKeywords = []string{
	"DefaultHttpClient",
	"HttpGet",
	"HttpPost",
	"request",
	"http",
}

var volleyKeywords = []string{
	"JsonObjectRequest",
	"JsonArrayRequest",
	"Request.Method",
	"Volley.newRequestQueue",
}

var asyncHTTPClientKeywords = []string{
	"AsyncHttpClient",
	"prepareGet",
}

// FindNetworkingCodeInProject scans source files in a
// project for networking related code
func FindNetworkingCodeInProject(p *droidalyzer.Project) error {

	for _, apf := range p.APFiles {
		if apf.Extension == ".java" {
			err := apf.Scan(&networkingKeywords)
			if err != nil {
				return err
			}

			if len(apf.NCode) > 0 {
				fmt.Println("Keyword occurrences of file:", apf.Name,
					":", len(apf.NCode))
				fmt.Println("Path:", apf.Path)
				apf.PrintNCSResults()
				fmt.Println("=============================")
			}
		}
	}

	return nil
}
