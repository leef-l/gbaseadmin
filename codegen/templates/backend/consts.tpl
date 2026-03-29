package consts
{{range .Fields}}{{if .IsEnum}}
// {{$.ModelName}}{{.NameCamel}} {{.Label}}
const (
{{- range .EnumValues}}
	{{$.ModelName}}{{$.NameCamel}}{{.Label}} = {{.Value}}
{{- end}}
)
{{end}}{{end}}
