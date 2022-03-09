# gocred - simple os based credential store
gocred allows unified access to the OS specific native credential store.
The credentials are stored with the service prefix `GOCRED`.

The following credential stores are used:

| OS | credential store          |
| --- |---------------------------|
| windows | Windows Credental Manager |
| macos | MacOS Keychain            |
| linux | GNOME Keyring             |