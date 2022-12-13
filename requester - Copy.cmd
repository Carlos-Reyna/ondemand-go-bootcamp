@echo off

for /l %%i in (1,1,905) do (
  curl localhost:8080/read/%%i
)