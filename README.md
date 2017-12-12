These tools uses https://github.com/rkrd/sn.go to interface with https://simplenote.com

note_create.go - Creates new note. Tags of note can be set with -t or -tag, accepts multiple tags. Last argument can be a file or - to read note body from stdin.
note_init_directory.go - Create note structure in current directory or path in first argument.
note_user_handler.go - Common handler for login and settings up sn.User.
notes_sync.go - Syncs all notes on server in current directory or path in first argument.
