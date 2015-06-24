package profitbricks

// ListImages returns an Collection struct
func ListImages() Collection {
	path := image_col_path()
	return is_list(path)
}

// GetImage returns an Instance struct where id ==imageid
func GetImage(imageid string) Instance {
	path := image_path(imageid)
	return is_get(path)
}

// UpdateImage updates all image properties from values in jason
//returns an Instance struct where id ==imageid
func UpdateImage(imageid string, jason []byte) Instance {
	path := image_path(imageid)
	return is_put(path, jason)
}

// PatchImage replaces any image properties from values in jason
//returns an Instance struct where id ==imageid
func PatchImage(imageid string, jason []byte) Instance {
	path := image_path(imageid)
	return is_patch(path, jason)
}

// Deletes an image where id==imageid
func DeleteImage(imageid string) Resp {
	path := image_path(imageid)
	return is_delete(path)
}
