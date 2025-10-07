env "dev" {
    url = "mysql://appuser:password@localhost:3306/appdb"
    dev = "docker://mysql/8"
    src = "file://schema.hcl"
    migration {
        dir = "file://migrations"
        format = atlas
    }
}
