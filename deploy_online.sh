CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service  main.go

project_name=service
project_env=service
project_host=xxx

echo 'send server'
scp  *.yml *.sh root@$project_host:/root/app/$project_name
scp  $project_name root@$project_host:/root/app/$project_name/tmp_$project_name

echo 'backup server exe'
ssh  root@$project_host "mv -f /root/app/$project_name/$project_name /root/app/$project_name/$project_name-bak"
ssh  root@$project_host "mv -f /root/app/$project_name/tmp_$project_name /root/app//$project_name/$project_name"


echo 'start server'
ssh  root@$project_host "export IS_PROD=1 && /root/app/$project_name/start.sh >> /root/app/$project_name/nohup.log"