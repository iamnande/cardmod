// Command cardmodweb serves the real cardmod refinement calculator over plain
// HTTP instead of gRPC. It imports the same refinement repository the gRPC
// CalculateAPI uses (internal/repositories/refinements) and runs the same
// math api.go's Calculate handler does — this is the actual calculator, just
// with a browser-friendly face instead of a protobuf one.
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/iamnande/cardmod/internal/repositories/refinements"
)

var repo = refinements.NewRepository()

// calcLine is one row of a calculation result: how many of Source you need,
// and (if Source is itself refined from something) its own breakdown.
type calcLine struct {
	Source   string
	Count    int32
	Children []calcLine
}

// uniqueTargets returns every distinct refinement target the data knows
// about, sorted, for the dropdown.
func uniqueTargets() []string {
	seen := map[string]bool{}
	out := make([]string, 0, 64)
	for _, r := range repo.ListRefinements(nil) {
		if !seen[r.Target()] {
			seen[r.Target()] = true
			out = append(out, r.Target())
		}
	}
	sort.Strings(out)
	return out
}

// calculate mirrors internal/api/calculatev1/api.go's Calculate handler:
// for a target and a desired count, find every source that refines into it,
// and one level of that source's own sources.
func calculate(target string, count int32) []calcLine {
	matches := repo.ListRefinements(refinements.NewFilter("", target))
	out := make([]calcLine, 0, len(matches))
	for _, r := range matches {
		c := count / r.Denominator() * r.Numerator()
		line := calcLine{Source: r.Source(), Count: c}
		for _, cr := range repo.ListRefinements(refinements.NewFilter("", r.Source())) {
			line.Children = append(line.Children, calcLine{
				Source: cr.Source(),
				Count:  c * cr.Numerator(),
			})
		}
		out = append(out, line)
	}
	return out
}

type pageData struct {
	Targets  []string
	Target   string
	Count    int32
	Results  []calcLine
	NotFound bool
}

var page = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>cardmod</title>
<style>
  :root { color-scheme: dark; }
  body { background:#101310; color:#dde8dd; font-family:ui-monospace,"SF Mono",Consolas,monospace;
         max-width:620px; margin:64px auto; padding:0 20px; line-height:1.6; }
  h1 { font-size:1.4rem; margin:0 0 4px; }
  .sub { color:#8a9a8a; font-size:0.9rem; margin:0 0 28px; }
  form { display:flex; gap:10px; flex-wrap:wrap; margin-bottom:28px; }
  select, input, button {
    font-family:inherit; font-size:0.95rem; background:#161c16; color:#dde8dd;
    border:1px solid #2a332a; border-radius:6px; padding:8px 10px;
  }
  input[type=number] { width:90px; }
  button { background:#7a9e7e; color:#101310; font-weight:700; border:none; cursor:pointer; }
  button:hover { opacity:0.9; }
  .card { background:#161c16; border:1px solid #2a332a; border-radius:8px; padding:14px 16px; margin-bottom:10px; }
  .card .head { font-weight:700; }
  .card .count { color:#7a9e7e; }
  .children { list-style:none; margin:8px 0 0; padding:10px 0 0 16px; border-top:1px solid #2a332a; }
  .children li { color:#8a9a8a; font-size:0.9rem; margin-bottom:4px; }
  .empty { color:#8a9a8a; }
  footer { margin-top:40px; font-size:0.78rem; color:#5a6a56; border-top:1px solid #2a332a; padding-top:14px; }
  a { color:#7a9e7e; }
</style>
</head>
<body>
<h1>⚔️ cardmod</h1>
<p class="sub">the real ff8 refinement calculator — pick what you want, how many, see what it costs.</p>

<form method="get" action="/">
  <select name="target" required>
    <option value="" {{if not .Target}}selected{{end}} disabled>choose a target…</option>
    {{range .Targets}}<option value="{{.}}" {{if eq . $.Target}}selected{{end}}>{{.}}</option>{{end}}
  </select>
  <input type="number" name="count" min="1" value="{{.Count}}">
  <button type="submit">calculate</button>
</form>

{{if .NotFound}}
  <p class="empty">nothing refines into "{{.Target}}" in this dataset.</p>
{{end}}

{{range .Results}}
  <div class="card">
    <div class="head">{{.Source}} <span class="count">× {{.Count}}</span></div>
    {{if .Children}}
    <ul class="children">
      {{range .Children}}<li>from: {{.Source}} × {{.Count}}</li>{{end}}
    </ul>
    {{end}}
  </div>
{{end}}

<footer>
  same data + math as the real grpc <code>CalculateAPI</code> in
  <a href="https://github.com/iamnande/cardmod">iamnande/cardmod</a> —
  just served over http instead of protobuf.
</footer>
</body>
</html>`))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target := r.URL.Query().Get("target")
		count := int32(1)
		if v, err := strconv.Atoi(r.URL.Query().Get("count")); err == nil && v > 0 {
			count = int32(v)
		}

		data := pageData{Targets: uniqueTargets(), Target: target, Count: count}
		if target != "" {
			data.Results = calculate(target, count)
			data.NotFound = len(data.Results) == 0
		}

		if err := page.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Printf("cardmodweb listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
