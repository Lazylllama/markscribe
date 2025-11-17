#### 📺 My latest anime additions
{{range slice (reverse animePlanetData.Anime.Entries) 0 4}}
- 🎬 **[{{.Name}}](https://www.anime-planet.com/anime/{{ regexReplaceAll "-+" (.Name | lower | replace " " "-" | replace ":" "" | replace "'" "") "-" }})**
    - Status: `{{title .Status}}` {{if eq .Status "watched"}}✅{{else if eq .Status "watching"}}▶️{{else}}📋{{end}}
    - Rating: {{if .Rating}}⭐ {{.Rating}}/5{{else}}Not rated{{end}}
    {{- if not (and (eq .Status "watched") (eq .Eps 1))}}
    - Episode: {{.Eps}} ep{{if gt .Eps 1}}s{{end}}
    {{- end}}
    - Times Watched: {{.Times}}x
{{- end}}

#### 📗 My latest manga additions
{{range slice (reverse animePlanetData.Manga.Entries) 0 4}}
- 📕 **[{{.Name}}](https://www.anime-planet.com/manga/{{ regexReplaceAll "-+" (.Name | lower | replace " " "-" | replace ":" "" | replace "'" "") "-" }})**
    - Status: `{{title .Status}}` {{if eq .Status "read"}}✅{{else if eq .Status "reading"}}▶️{{else}}📋{{end}}
    - Rating: {{if gt .Rating 0}}⭐ {{.Rating}}/5{{else}}Not rated{{end}}
    {{- if not (and (eq .Status "read") (eq .Vol 1))}}
    - Volumes: {{.Vol}} ep{{if gt .Vol 1}}s{{end}}
    {{- end}}
{{- end}}

#### 📗 Currently Reading
{{range $index, $element := animePlanetData.Manga.Entries}}
{{- if eq $element.Status "reading"}}
- 📕 **[{{.Name}}](https://www.anime-planet.com/manga/{{ regexReplaceAll "-+" (.Name | lower | replace " " "-" | replace ":" "" | replace "'" "") "-" }})**
    - Rating: {{if gt .Rating 0}}⭐ {{.Rating}}/5{{else}}Not rated{{end}}
    - Volumes: {{.Vol}} ep{{if gt .Vol 1}}s{{end}}
{{- end}}
{{- end}}

#### 📺 Currently Watching
{{range $index, $element := animePlanetData.Anime.Entries}}
{{- if eq $element.Status "watching"}}
- 🎬 **[{{.Name}}](https://www.anime-planet.com/anime/{{ regexReplaceAll "-+" (.Name | lower | replace " " "-" | replace ":" "" | replace "'" "") "-" }})**
    - Rating: {{if .Rating}}⭐ {{.Rating}}/5{{else}}Not rated{{end}}
    - Episode: {{.Eps}} ep{{if gt .Eps 1}}s{{end}}
{{- end}}
{{- end}}