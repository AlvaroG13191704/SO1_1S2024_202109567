cmd_/home/alvarog/Modules/ram.mod := printf '%s\n'   ram.o | awk '!x[$$0]++ { print("/home/alvarog/Modules/"$$0) }' > /home/alvarog/Modules/ram.mod
