#!/usr/bin/env bash
set -e
# all application service l4
go run -tags generator main.go -p MIBS -m ALTEON-CHEETAH-LAYER4-MIB.mib .1.3.6.1.4.1.1872.2.5.4.1.1
# all oper
go run -tags generator main.go -p MIBS -m ALTEON-CHEETAH-LAYER4-MIB.mib .1.3.6.1.4.1.1872.2.5.4.4
# all stat
go run -tags generator main.go -p MIBS -m ALTEON-CHEETAH-LAYER4-MIB.mib .1.3.6.1.4.1.1872.2.5.4.2
# AgSaveTable
go run -tags generator main.go -p MIBS -m ALTEON-CHEETAH-SWITCH-MIB.mib .1.3.6.1.4.1.1872.2.5.1.1.17.3
# AgApplyTable
go run -tags generator main.go -p MIBS -m ALTEON-CHEETAH-SWITCH-MIB.mib  .1.3.6.1.4.1.1872.2.5.1.1.8.5
