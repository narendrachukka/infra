key=$1

# create serviceaccount
kubectl create serviceaccount infra
kubectl create clusterrolebinding infra --clusterrole=cluster-admin --serviceaccount=default:infra

# get ca & token
# TODO: make me work without jq
secret=$(kubectl get sa infra -o json | jq -r '.secrets[].name')
token=$(kubectl get secret $secret -o json | jq -r '.data["token"]' | base64 --decode)
ca=$(kubectl get secret $secret -o json | jq -r '.data["ca.crt"]' | base64 --decode | jq -sR .)

# get endpoint
context=`kubectl config current-context`
name=`kubectl config get-contexts $context | awk '{print $3}' | tail -n 1`
url=`kubectl config view -o jsonpath="{.clusters[?(@.name == \"$name\")].cluster.server}"`

curl -X POST http://localhost/v1/destinations \
    -H "Authorization: Bearer $key" \
    -H 'Content-Type: application/json' \
    --data-binary @- << EOF
{
    "name": "kubernetes.$name",
    "connection": {
        "ca": $ca,
        "url": "$url",
        "credential": "$token"
    }
}
EOF
