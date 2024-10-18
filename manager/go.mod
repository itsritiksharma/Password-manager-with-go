module password-manager/manager

go 1.23.2

replace password-manager/vaultCreation => ../vaultCreation

require password-manager/vaultCreation v0.0.0-00010101000000-000000000000

require password-manager/fileOperations v0.0.0-00010101000000-000000000000 // indirect

replace password-manager/fileOperations => ../fileOperations
