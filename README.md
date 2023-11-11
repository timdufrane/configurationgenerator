# configurationgenerator
Simple configuration generator based on YAML templating. This is dirt simple code that I needed for quick templating uses when configuring network equipment.

# Example configuration
ip address 192.168.{{thirdoctet}}.5/24

# Example YAML template
thirdoctet:
  - 28
  - 29
  - 30
  - 31

# Example run
go build
configurationgenerator.exe -c thirdoctetconfiguration.txt -s thirdoctettemplate.yml -d thirdoctetout

# Example results
*thirdoctetout/out.0*: `ip address 192.168.28.5/24`  
*thirdoctetout/out.1*: `ip address 192.168.29.5/24`  
*thirdoctetout/out.2*: `ip address 192.168.30.5/24`  
*thirdoctetout/out.3*: `ip address 192.168.31.5/24`  
