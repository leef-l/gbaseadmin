package consts
{{range .Fields}}{{if .IsEnum}}
// {{$.ModelName}}{{.NameCamel}} {{.Label}}
const (
{{- $fieldCamel := .NameCamel -}}
{{- range .EnumValues}}
	{{$.ModelName}}{{$fieldCamel}}{{.Label}} = {{.Value}}
{{- end}}
)
{{end}}{{end}}
