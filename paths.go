package profitbricks

// slash returns "/<str>" great for making url paths
func slash(str string) string {
	return "/" + str
}

// dc_col_path	returns the string "/datacenters"
func dcColPath() string {
	return slash("datacenters")
}

// dc_path returns the string "/datacenters/<dcid>"
func dcPath(dcid string) string {
	return dcColPath() + slash(dcid)
}

// image_col_path returns the string" /images"
func imageColPath() string {
	return slash("images")
}

// image_path returns the string"/images/<imageid>"
func imagePath(imageid string) string {
	return imageColPath() + slash(imageid)
}

// ipblockColPath returns the string "/ipblocks"
func ipblockColPath() string {
	return slash("ipblocks")
}

//  ipblock_path returns the string "/ipblocks/<ipblockid>"
func ipblockPath(ipblockid string) string {
	return ipblockColPath() + slash(ipblockid)
}

// location_col_path returns the string  "/locations"
func locationColPath() string {
	return slash("locations")
}

// location_path returns the string   "/locations/<locid>"
func locationPath(locid string) string {
	return locationColPath() + slash(locid)
}

// location_path returns the string   "/locations/<regid>"
func locationRegPath(regid string) string {
	return locationColPath() + slash(regid)
}

// snapshot_col_path returns the string "/snapshots"
func snapshotColPath() string {
	return slash("snapshots")
}

// lan_col_path returns the string "/datacenters/<dcid>/lans"
func lanColPath(dcid string) string {
	return dcPath(dcid) + slash("lans")
}

// lan_path returns the string	"/datacenters/<dcid>/lans/<lanid>"
func lanPath(dcid, lanid string) string {
	return lanColPath(dcid) + slash(lanid)
}

//  lbal_col_path returns the string "/loadbalancers"
func lbalColPath(dcid string) string {
	return dcPath(dcid) + slash("loadbalancers")
}

// lbalpath returns the string "/loadbalancers/<lbalid>"
func lbalPath(dcid, lbalid string) string {
	return lbalColPath(dcid) + slash(lbalid)
}

// server_col_path returns the string	"/datacenters/<dcid>/servers"
func serverColPath(dcid string) string {
	return dcPath(dcid) + slash("servers")
}

// serverPath returns the string   "/datacenters/<dcid>/servers/<srvid>"
func serverPath(dcid, srvid string) string {
	return serverColPath(dcid) + slash(srvid)
}

// server_cmd_path returns the string   "/datacenters/<dcid>/servers/<srvid>/<cmd>"
func serverCommandPath(dcid, srvid, cmd string) string {
	return serverPath(dcid, srvid) + slash(cmd)
}

// volume_col_path returns the string "/volumes"
func volumeColPath(dcid string) string {
	return dcPath(dcid) + slash("volumes")
}

// volume_path returns the string "/volumes/<volid>"
func volumePath(dcid, volid string) string {
	return volumeColPath(dcid) + slash(volid)
}

//  balnic_col_path returns the string "/loadbalancers/<lbalid>/balancednics"
func balnicColPath(dcid, lbalid string) string {
	return lbalPath(dcid, lbalid) + slash("balancednics")
}

//  balnic_path returns the string "/loadbalancers/<lbalid>/balancednics<balnicid>"
func balnicPath(dcid, lbalid, balnicid string) string {
	return balnicColPath(dcid, lbalid) + slash(balnicid)
}

// server_cdrom_col_path returns the string   "/datacenters/<dcid>/servers/<srvid>/cdroms"
func serverCdromColPath(dcid, srvid string) string {
	return serverPath(dcid, srvid) + slash("cdroms")
}

// server_cdrom_path returns the string   "/datacenters/<dcid>/servers/<srvid>/cdroms/<cdid>"
func serverCdromPath(dcid, srvid, cdid string) string {
	return serverCdromColPath(dcid, srvid) + slash(cdid)
}

// server_volume_col_path returns the string   "/datacenters/<dcid>/servers/<srvid>/volumes"
func serverVolumeColPath(dcid, srvid string) string {
	return serverPath(dcid, srvid) + slash("volumes")
}

// server_volume_path returns the string   "/datacenters/<dcid>/servers/<srvid>/volumes/<volid>"
func serverVolumePath(dcid, srvid, volid string) string {
	return serverVolumeColPath(dcid, srvid) + slash(volid)
}

// nic_path returns the string   "/datacenters/<dcid>/servers/<srvid>/nics"
func nicColPath(dcid, srvid string) string {
	return serverPath(dcid, srvid) + slash("nics")
}

// nic_path returns the string   "/datacenters/<dcid>/servers/<srvid>/nics/<nicid>"
func nicPath(dcid, srvid, nicid string) string {
	return nicColPath(dcid, srvid) + slash(nicid)
}

// fwrule_col_path returns the string   "/datacenters/<dcid>/servers/<srvid>/nics/<nicid>/firewallrules"
func fwruleColPath(dcid, srvid, nicid string) string {
	return nicPath(dcid, srvid, nicid) + slash("firewallrules")
}

// fwrule_path returns the string
//  "/datacenters/<dcid>/servers/<srvid>/nics/<nicid>/firewallrules/<fwruleid>"
func fwrulePath(dcid, srvid, nicid, fwruleid string) string {
	return fwruleColPath(dcid, srvid, nicid) + slash(fwruleid)
}

// contract_resource_path returns the string "/contracts"
func contractResourcePath() string {
	return slash("contracts")
}

func um() string {
	return slash("um")
}

// um_groups	returns the string "/groups"
func umGroups() string {
	return um() + slash("groups")
}

// um_group_path	returns the string "/groups/groupid"
func umGroupPath(grpid string) string {
	return umGroups() + slash(grpid)
}

// um_group_shares	returns the string "groups/{groupId}/shares"
func umGroupShares(grpid string) string {
	return um() + slash("groups") + slash(grpid) + slash("shares")
}

// um_group_share_path	returns the string "groups/{groupId}/shares/{resourceId}"
func umGroupSharePath(grpid string, resourceid string) string {
	return um() + slash("groups") + slash(grpid) + slash("shares") + slash(resourceid)
}

// um_group_users	returns the string "/groups/groupid/users"
func umGroupUsers(grpid string) string {
	return um() + slash("groups") + slash(grpid) + slash("users")
}

// um_group_users_path	returns the string "/groups/groupid/users/userid"
func umGroupUsersPath(grpid string, usrid string) string {
	return um() + slash("groups") + slash(grpid) + slash("users") + slash(usrid)
}

// um_users returns the string "/users"
func umUsers() string {
	return um() + slash("users")
}

// um_users returns the string "/users/usrid"
func umUsersPath(usrid string) string {
	return um() + slash("users") + slash(usrid)
}

// um_resources returns the string "/resources"
func umResources() string {
	return um() + slash("resources")
}

// um_resources_type returns the string "/resources/resourceType"
func umResourcesType(restype string) string {
	return um() + slash("resources") + slash(restype)
}

// um_resources_type_path returns the string "resources/{resourceType}/{resourceId}"
func umResourcesTypePath(restype string, resourceid string) string {
	return um() + slash("resources") + slash(restype) + slash(resourceid)
}

// tokenColPath returns the string "/tokens"
func tokenColPath() string {
	return slash("tokens")
}

// tokenPath returns the string "/tokens/<tokenid>"
func tokenPath(tokenid string) string {
	return tokenColPath() + slash(tokenid)
}
