These tools uses https://github.com/rkrd/sn.go to interface with https://simplenote.com

create_new.go - Creates new note. Tags of note can be set with -t or -tag, accepts multiple tags. Last argument can be a file or - to read note body from stdin.
init.go - Create note structure in current directory or path in first argument.
syncnotes.go - Syncs all notes on server in current directory or path in first argument.
user.go - Common handler for login and settings up sn.User.
