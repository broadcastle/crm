# Changelog

All changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

- Additional functionality for notes.
- Pretty output for contacts.
- Pretty output for notes.
- Tests.
- Email address validation.

## [0.5.0] - 2018-08-10

### Added

- 'email' command.
- Ability to view existing emails.
- Slugs for contacts.
- Ability to use slugs in addition to id's with contact commands.

## [0.4.0] - 2018-08-09

### Added

- Alias 'contacts' for 'contact'.
- Alias 'notes' for 'note'.
- 'fast' flag to note creation.
- Untested email functionality.

### Changed

- Moved cobra flag input functionality to utils.
- 'note' now asks for input like 'contact'.

### Fixed

- Removed extra colon from note creation.
- Mentions of client in 'cmd/contact.go'

## [0.3.0] - 2018-08-05

### Added

- Note functionality 
- Minimum arguments to 'remove' and 'edit' commands.
- Prevent arguments from being added to 'add' commands..
- Ability to view all contacts.

### Changed

- Moved all 'crm contact' commands to a single file.
- Renamed 'contact' sub-command functions.
- Moved 'Note{}' to different file to prevent a single large file.

## [0.2.0] - 2018-08-04

### Added

- Delete relationships from deleted contacts.
- View contacts in database.
- JSON fields for structs.

## [0.1.1] - 2018-08-04

### Added

- Enforcement that all subscribers are customers.
- Prevent contacts from being added without a name or email.

### Changed

- Moved the 'manage.go' contact functions into seperate files.

### Fixed

- Duplication issue when updating contact relationships.
- Relationship table not being added when empty.

## [0.1.0] - 2018-07-31

### Added

- This changelog file was added.
- Logrus debug entries to more files.
- Additional fields in notes.
- README file.
- Relationships for contacts.
- Boolean input.
- Additional questions in contact creation.

### Changed

- 'client' commands and files were renamed to 'contact'.
- Using 'logrus.Fatal' instead of 'logrus.Fatalln' where possible.
- Joined 'contact add' and 'contact edit' commands.
- Contact fields.
- Changed maximum stars in hidden input hint.

### Fixed

- Debug command.
