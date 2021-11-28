# VPT
SSH Tunnel

# Install
`go install github.com/simba-fs/vpt@latest`

# Usage
## Host
`vpt connect host <local port>:<server ip>[:server port=22] [-p <password>]`

## Client
`vpt connect client <local port>:<server ip>[:server port=22] [-p <password> ]`

# Woking progress
- [x] `connect`  
- [x] `key`  
- [x] `key renew`  
- [x] ~~`key add <key>`~~  
- [x] ~~`key remove <keySHA256>`~~  
