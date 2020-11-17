{{- define "hazelcastName" -}} hazelcast {{- end -}}
{{- define "hazelcastEnvironmentConfig" -}} hz-env-config {{- end -}}
{{- define "hazelcastConfig" -}} hz-config {{- end -}}


{{- define "isHazelcastEnabled" }}
 {{- if .Values.enabled -}}
  true
 {{- else -}}
  false
 {{- end -}}
{{- end }}

{{- define "waitForHazelcast" -}}
- name: wait-for-hazelcast
  image: busybox:1.31.0
  # Init container for waiting for hazelcast service to initialize.
  args:
  - sh
  - -c
  - >
    set -e
    counter=0;
    while [ $(wget -q -O - "http://hazelcast-service.{{ .Release.Namespace }}:5701/hazelcast/health/cluster-size" /dev/null) -ne {{ .Values.hazelcast.replicas }} ] || [ $(wget -S "http://hazelcast-service.{{ .Release.Namespace }}:5701/hazelcast/health/cluster-safe" 2>&1 | grep "HTTP/" | awk '{print $2}') -ne 200 ]; do
    echo "waiting for hazelcast pods to start and join the cluster..." ;
    counter=$(($counter+5));
    sleep 5;
    if [ $counter -gt 300 ]; then
    echo "Timeout Reached. Hazelcast pods failed to join the cluster within 5 minutes";
    exit 1;
    fi
    done;
    echo "Hazelcast cluster is up now"
    exit 0
{{- end }}


