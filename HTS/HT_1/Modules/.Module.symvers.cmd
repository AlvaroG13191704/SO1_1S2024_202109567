cmd_/home/alvarog/Modules/Module.symvers :=  sed 's/ko$$/o/'  /home/alvarog/Modules/modules.order | scripts/mod/modpost  -a   -N -o /home/alvarog/Modules/Module.symvers -e -i Module.symvers -T - 
