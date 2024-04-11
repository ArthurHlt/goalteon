# How Alteon Api is Working ?

If you want to create a new client for the Alteon API you need to understand few things

## Alteon API is an SNMP bridge

The Alteon Rest api is a bridge to make call in snmp so if you can take MIBS (included in this repo) and generate
automatically the code by using table in mibs.

Target all table inside an oids and indexes will be index in the url path and others columns are json parameters.

All columns name (including indexes) have the common prefix between table **entry** and column name removed, example:

For `slbNewCfgGroupTable` the entry name is `slbNewCfgGroupEntry` and a column looks like `slbNewCfgGroupAddServer` so 
parameter key should be `AddServer`.

But for `slbNewCfgEnhVirtServerTable` the entry name is `slbNewCfgEnhVirtualServerEntry` 
(note the difference between VirtServ and VirtualServer) and a column looks like `slbNewCfgEnhVirtServerIpAddress` and 
the prefix in common between entry and colum is `slbNewCfgEnh` due to `VirtualServer` in entry so the parameter key 
should be `VirtServerIpAddress`.

