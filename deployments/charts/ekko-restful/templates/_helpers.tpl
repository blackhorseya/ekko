{{/*
Expand the name of the chart.
*/}}
{{- define "ekko-restful.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "ekko-restful.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "ekko-restful.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "ekko-restful.labels" -}}
helm.sh/chart: {{ include "ekko-restful.chart" . }}
{{ include "ekko-restful.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
release: {{ $.Release.Name | quote }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "ekko-restful.selectorLabels" -}}
app.kubernetes.io/name: {{ include "ekko-restful.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
release: {{ $.Release.Name | quote }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "ekko-restful.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "ekko-restful.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
