package config

/*
# YAML Example

log:
    level: debug

queryTemplate: &queryTemplate
    # query option

datasource:
    AWSRoute53:
        enabled: true
        query:
            <<: *queryTemplate
        auth:
            # auth option
        option:
            TTL: 300 # second
            # more options...
    Kubernetes:
        enabled: true
        query:
            <<: *queryTemplate
        auth:
            # auth option
    ArgoCD:
        enabled: true
        query:
            <<: *queryTemplate
        auth:
            # auth option
    static:
        enabled: true
        values:
            - name: site-a.example.com
              tags:
                region: seoul
*/

type Config struct {
    Log Log `yaml:"log"`
}

type Log struct {
    Level string `yaml:"level"`
}
