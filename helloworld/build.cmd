
del goRm_helloworld
set GOOS=linux

go build -i -ldflags "-s -w -X 'main.buildTime=%date% %time%'" -o goRm_helloworld
"C:\Users\Admin\Desktop\putty+PSCP\pscp.exe" -pw 1 goRm_helloworld lhm@192.168.44.139:/home/lhm/Gorm