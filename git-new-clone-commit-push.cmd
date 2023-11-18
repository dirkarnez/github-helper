@echo off
set DOWNLOADS_DIR=%USERPROFILE%\Downloads

SET PATH=^
%DOWNLOADS_DIR%\PortableGit\bin;

@echo off
set USER_NAME=dirkarnez
set PASSWORD=
set REPO=

github-helper.exe --repo %REPO% &&^
git clone https://%USER_NAME%:%PASSWORD%@github.com/dirkarnez/%REPO%.git &&^
pause

cd %REPO% &&^
git add * && git commit -m "- upload files" && git branch -M main && git push -u origin main &&^
cd ..

pause

