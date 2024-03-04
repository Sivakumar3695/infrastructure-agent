// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
package logs

var fbConfigFormat = `{{- range .Inputs }}
[INPUT]
    Name {{ .Name }}
    {{- if .Path }}
    Path {{ .Path }}
    {{- end }}
    {{- if .BufferChunkSize }}
    Buffer_Chunk_Size {{ .BufferChunkSize }}
    {{- end }}
    {{- if .BufferMaxSize }}
    Buffer_Max_Size {{ .BufferMaxSize }}
    {{- end }}
	{{- if .MemBufferLimit }}
    Mem_Buf_Limit {{ .MemBufferLimit }}
    {{- end }}
    {{- if .SkipLongLines }}
    Skip_Long_Lines {{ .SkipLongLines }}
    {{- end }}
    {{- if .PathKey }}
    Path_Key {{ .PathKey }}
    {{- end }}
    {{- if .Tag }}
    Tag  {{ .Tag }}
    {{- end }}
    {{- if .DB }}
    DB   {{ .DB }}
    {{- end }}
    {{- if .Systemd_Filter }}
    Systemd_Filter {{ .Systemd_Filter }}
    {{- end }}
    {{- if .Channels }}
    Channels {{ .Channels }}
    {{- end }}
    {{- if .SyslogMode }}
    Mode {{ .SyslogMode }}
    {{- end }}
    {{- if .SyslogListen }}
    Listen {{ .SyslogListen }}
    {{- end }}
    {{- if .SyslogPort }}
    Port {{ .SyslogPort }}
    {{- end }}
    {{- if .SyslogParser }}
    Parser {{ .SyslogParser }}
    {{- end }}
    {{- if .SyslogUnixPath }}
    Path {{ .SyslogUnixPath }}
    {{- end }}
    {{- if .SyslogUnixPermissions }}
    Unix_Perm {{ .SyslogUnixPermissions }}
    {{- end }}
    {{- if .TcpListen }}
    Listen {{ .TcpListen }}
    {{- end }}
    {{- if .TcpPort }}
    Port {{ .TcpPort }}
    {{- end }}
    {{- if .TcpFormat }}
    Format {{ .TcpFormat }}
    {{- end }}
    {{- if .TcpSeparator }}
    Separator {{ .TcpSeparator }}
    {{- end }}
    {{- if .TcpBufferSize }}
    Buffer_Size {{ .TcpBufferSize }}
    {{- end }}
 	{{- if .UseANSI }}
    Use_ANSI {{ .UseANSI }}
    {{- end }}
{{ end -}}

{{- range .Filters }}
[FILTER]
    {{- if .Name }}
    Name  {{ .Name }}
    {{- end }}
    {{- if .Match }}
    Match {{ .Match }}
    {{- end }}
    {{- if .Regex }}
    Regex {{ .Regex }}
    {{- end }}
    {{- if .Records }}
        {{- range $key, $value := .Records }}
    Record {{ $key }} {{ $value }}
        {{- end }}
    {{- end }}
    {{- if .Modifiers }}
        {{- range $key, $value := .Modifiers }}
    Rename {{ $key }} {{ $value }}
        {{- end }}
    {{- end }}
    {{- if .Script }}
    script {{ .Script }}
    {{- end }}
    {{- if .Call }}
    call {{ .Call }}
    {{- end }}
{{ end -}}

{{- if .Output }}
[OUTPUT]
    Name                {{ .Output.Name }}
    Match               {{ .Output.Match }}
    {{- if .Output.LicenseKey }}
    licenseKey          ${NR_LICENSE_KEY_ENV_VAR}
    {{- end }}
    {{- if .Output.Endpoint }}
    endpoint            {{ .Output.Endpoint }}
    {{- end }}
    {{- if .Output.Proxy }}
    proxy               {{ .Output.Proxy }}
    {{- end }}
	{{- if .Output.IgnoreSystemProxy }}
    ignoreSystemProxy   true
    {{- end }}
	{{- if .Output.CABundleFile }}
    caBundleFile        {{ .Output.CABundleFile }}
    {{- end }}
    {{- if .Output.CABundleDir }}
    caBundleDir         {{ .Output.CABundleDir }}
    {{- end }}
    {{- if not .Output.ValidateCerts }}
    validateProxyCerts  false
    {{- end }}
    {{- if .Output.Retry_Limit}}
    Retry_Limit         {{ .Output.Retry_Limit }}
    {{- end }}
    {{- if .Output.SendMetrics}}
    sendMetrics         {{ .Output.SendMetrics}}
    {{- end}}
{{ end -}}

{{- if .ExternalCfg.CfgFilePath }}
@INCLUDE {{ .ExternalCfg.CfgFilePath }}
{{ end -}}`

var fbLuaScriptFormat = `function {{ .FnName }}(tag, timestamp, record)
    eventId = record["EventID"]
    -- Discard log records matching any of these conditions
    if {{ .ExcludedEventIds }} then
        return -1, 0, 0
    end
    -- Include log records matching any of these conditions
    if {{ .IncludedEventIds }} then
        return 0, 0, 0
    end
    -- If there is not any matching conditions discard everything
    return -1, 0, 0
 end`

var tooManyFilesWarnMsg = `
The amount of open files targeted by your Log Forwarding configuration
files ({{ .TotalFiles }}) exceeds the recommended maximum ({{ .DefaultFileLimit }}). The Operating System
may kill the Log Forwarder process or not even allow it to start.
These are the amount of files targeted by each of your configuration blocks:{{"\n"}}

{{- range .LogsCfg }}
{{- if .File }}
- name: {{ .Name }}
  file: {{ .File }}
  targeted files: {{ .TargetFilesCnt }}
{{ end -}}
{{ end -}}

{{"\n"}}We recommend the following tips:
- Consider adjusting the wildcards used in the "file" configuration attributes
to target a smaller amount of files
- Consider using another log rotation strategy that uses a different naming
for the rotated logs (i.e. my_log.log.20240214 instead of my_log.20240214.log)
- You may also consider increasing the maximum amount of allowed file
descriptors and inotify watchers. See: https://docs.newrelic.com/docs/logs/forward-logs/forward-your-logs-using-infrastructure-agent/#too-many-files

Please note that this is a friendly warning message. If your operating system allows more than {{ .DefaultFileLimit }} file descriptors/inotify watchers 
or if you already increased their maximum amount by following the above link, you can safely ignore this message.
`
