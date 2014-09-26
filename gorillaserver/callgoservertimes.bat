@echo off
for /L %%n in (1,1,1000) do (
cscript /nologo wget.js http://localhost:3000/GetAllCompanyData
)
exit
