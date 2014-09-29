@echo off
for /L %%n in (1,1,100) do (
cscript /nologo wget.js http://localhost:3000/Profile
)
exit
