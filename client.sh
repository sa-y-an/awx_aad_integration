echo "Creating Azure AD Application..."

appDisplayName="azure_awx_sm7"
appReplyUrl="http://localhost:10445/sso/complete/azuread-oauth2/"


az ad app create --display-name "$appDisplayName" --web-redirect-uris "$appReplyUrl"
appId=$(az ad app list --display-name "$appDisplayName" --query '[0].appId' --output tsv)
echo "Application ID (Client ID): $appId"
client=$(az ad app credential reset --id $appId --append)

echo "add permision"
az ad app permission add --id $appId --api 00000003-0000-0000-c000-000000000000 --api-permissions e1fe6dd8-ba31-4d61-89e7-88639da4683d=Scope
echo "hello"

export client

echo "\n"
echo ""
echo ""


echo "here is : client : $client"

go run .