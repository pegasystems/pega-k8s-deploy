export CHART_VERSION=$(cat version.yaml | grep -A1 version | cut -d: -f2)
export PEGA_FILE_NAME=pega-${CHART_VERSION}.tgz
export ADDONS_FILE_NAME=addons-${CHART_VERSION}.tgz
cat descriptor-template.json | jq '.files[0].includePattern=env.PEGA_FILE_NAME' | jq '.files[0].uploadPattern=env.PEGA_FILE_NAME' | jq '.files[1].includePattern=env.ADDONS_FILE_NAME' | jq '.files[1].uploadPattern=env.ADDONS_FILE_NAME' > descriptor.json
curl -o index.yaml https://kishor.bintray.com/pega-helm-charts/index.yaml
helm package --version ${CHART_VERSION} ./pega/
helm package --version ${CHART_VERSION} ./addons/
ls -l
helm repo index --merge index.yaml --url https://kishor.bintray.com/pega-helm-charts/ .