module 1

go 1.23.2

replace example.com/node => ../node

replace local/node => ../node

require local/node v0.0.0-00010101000000-000000000000 // indirect
