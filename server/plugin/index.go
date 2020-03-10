package plugin

import (
	_ "github.com/icoma89/filestash/server/plugin/plg_starter_tunnel"
	_ "github.com/icoma89/filestash/server/plugin/plg_starter_tor"
	_ "github.com/icoma89/filestash/server/plugin/plg_video_transcoder"
	_ "github.com/icoma89/filestash/server/plugin/plg_editor_onlyoffice"
	_ "github.com/icoma89/filestash/server/plugin/plg_handler_syncthing"
	//_ "github.com/icoma89/filestash/server/plugin/plg_image_light"
	_ "github.com/icoma89/filestash/server/plugin/plg_backend_backblaze"
	_ "github.com/icoma89/filestash/server/plugin/plg_backend_dav"
	_ "github.com/icoma89/filestash/server/plugin/plg_backend_mysql"
	_ "github.com/icoma89/filestash/server/plugin/plg_backend_ldap"
	_ "github.com/icoma89/filestash/server/plugin/plg_backend_dropbox"
	_ "github.com/icoma89/filestash/server/plugin/plg_security_scanner"
	_ "github.com/icoma89/filestash/server/plugin/plg_security_svg"
	_ "github.com/icoma89/filestash/server/plugin/plg_handler_console"
	. "github.com/icoma89/filestash/server/common"
)

func init() {
	Log.Debug("Plugin loader")
}
