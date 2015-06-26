package profitbricks

// ListIpBlocks
func ListIpBlocks() Collection {
	path := ipblock_col_path()
	return is_list(path)
}

func ReserveIpBlock(jason []byte) Instance {
	path := ipblock_col_path()
	return is_post(path, jason)

}
func GetIpBlock(ipblockid string) Instance {
	path := ipblock_path(ipblockid)
	return is_get(path)
}

func ReleaseIpBlock(ipblockid string) Resp {
	path := ipblock_path(ipblockid)
	return is_delete(path)
}
